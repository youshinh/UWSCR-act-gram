package capture

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	goruntime "runtime"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Windows API constants
const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14
	WM_KEYDOWN     = 0x0100
	WM_LBUTTONDOWN = 0x0201
	WM_RBUTTONDOWN = 0x0204

	captureTimeout = 15 * time.Second
)

// UI Automation GUIDs
var (
	CLSID_CUIAutomation = ole.NewGUID("{ff48dba4-60ef-4201-aa87-54103eef594e}")
	IID_IUIAutomation   = ole.NewGUID("{30cbe57d-d9d0-452a-ab13-7ac5ac4825ee}")
)

// Win32 Structures
type POINT struct {
	X int32
	Y int32
}

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type MSLLHOOKSTRUCT struct {
	Pt          POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

type MSG struct {
	Hwnd    uintptr
	Message uint32
	Wparam  uintptr
	Lparam  uintptr
	Time    uint32
	Pt      POINT
}

// RecordEvent represents a structured user action
type RecordEvent struct {
	Timestamp  int64  `json:"timestamp"` // Epoch ms
	Type       string `json:"type"`      // "click", "keydown", "window_change"
	X          int    `json:"x,omitempty"`
	Y          int    `json:"y,omitempty"`
	RelX       int    `json:"rel_x,omitempty"`
	RelY       int    `json:"rel_y,omitempty"`
	Key        string `json:"key,omitempty"`
	Window     string `json:"window"`
	Control    string `json:"control,omitempty"`    // e.g. "Button", "ComboBox"
	ControlID  string `json:"control_id,omitempty"` // AutomationId
	Value      string `json:"value,omitempty"`      // Value or state
	ImagePath  string `json:"image_path,omitempty"` // Screenshot path
	WindowRect *Rect  `json:"-"`                    // Active window rect for evidence output
}

// Recorder manages the Win32 hooks and UI Automation
type Recorder struct {
	ctx context.Context

	stopMu        sync.Mutex
	stateMu       sync.Mutex
	isRecording   bool
	channelClosed bool

	dataMu       sync.Mutex
	logDir       string
	events       []RecordEvent
	lastWindow   string
	prevX, prevY int // For distance calculation later
	startedAt    int64
	eventsFile   *os.File
	jsonlEncoder *json.Encoder

	// Hook handles
	mouseHook    uintptr
	kbdHook      uintptr
	hookThreadId uint32

	// Async capture channel
	captureChan chan RecordEvent
	senderWG    sync.WaitGroup
	workerWG    sync.WaitGroup
	hookWG      sync.WaitGroup
	captureSem  chan struct{}
}

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	callNextHookEx      = user32.NewProc("CallNextHookEx")
	unhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	getMessage          = user32.NewProc("GetMessageW")
	postThreadMessage   = user32.NewProc("PostThreadMessageW")
	getForegroundWindow = user32.NewProc("GetForegroundWindow")
	getWindowText       = user32.NewProc("GetWindowTextW")
	getWindowRect       = user32.NewProc("GetWindowRect")
	getCurrentThreadId  = kernel32.NewProc("GetCurrentThreadId")
	kernel32            = syscall.NewLazyDLL("kernel32.dll")

	// Global recorder instance pointer for callback access
	globalRecorderMu sync.RWMutex
	globalRecorder   *Recorder
)

func setGlobalRecorder(r *Recorder) {
	globalRecorderMu.Lock()
	globalRecorder = r
	globalRecorderMu.Unlock()
}

func clearGlobalRecorder(r *Recorder) {
	globalRecorderMu.Lock()
	if globalRecorder == r {
		globalRecorder = nil
	}
	globalRecorderMu.Unlock()
}

func currentRecorder() *Recorder {
	globalRecorderMu.RLock()
	defer globalRecorderMu.RUnlock()
	return globalRecorder
}

func NewRecorder(ctx context.Context, logDir string) *Recorder {
	return &Recorder{
		ctx:         ctx,
		logDir:      logDir,
		events:      make([]RecordEvent, 0),
		captureChan: make(chan RecordEvent, 100),
		captureSem:  make(chan struct{}, 1),
	}
}

func (r *Recorder) writeSessionJSON() error {
	session := EvidenceSession{
		SchemaVersion: "evidence/v1",
		StartedAt:     r.startedAt,
		EventsPath:    "events.jsonl",
		CapturesDir:   "captures",
		TemplatesDir:  "templates",
	}
	data, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return fmt.Errorf("session JSONの作成に失敗: %v", err)
	}
	return os.WriteFile(filepath.Join(r.logDir, "session.json"), data, 0644)
}

func (r *Recorder) stopAcceptingEvents() bool {
	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	if !r.isRecording {
		return false
	}
	r.isRecording = false
	clearGlobalRecorder(r)
	return true
}

func (r *Recorder) closeCaptureChannel() {
	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	if r.channelClosed {
		return
	}
	close(r.captureChan)
	r.channelClosed = true
}

func (r *Recorder) closeEventsFile() {
	if r.eventsFile == nil {
		return
	}
	if err := r.eventsFile.Sync(); err != nil {
		log.Printf("[Recorder] events.jsonl sync failed: %v", err)
	}
	if err := r.eventsFile.Close(); err != nil {
		log.Printf("[Recorder] events.jsonl close failed: %v", err)
	}
	r.eventsFile = nil
	r.jsonlEncoder = nil
}

// Start hooks keyboard and mouse events
func (r *Recorder) Start(captureFunc func(string) error) error {
	r.stopMu.Lock()
	defer r.stopMu.Unlock()

	r.stateMu.Lock()
	if r.isRecording {
		r.stateMu.Unlock()
		return fmt.Errorf("すでに記録中です。")
	}
	r.captureChan = make(chan RecordEvent, 100)
	r.isRecording = true
	r.channelClosed = false
	r.mouseHook = 0
	r.kbdHook = 0
	r.hookThreadId = 0
	setGlobalRecorder(r)
	r.stateMu.Unlock()

	r.dataMu.Lock()
	r.events = r.events[:0]
	r.prevX, r.prevY = 0, 0
	r.dataMu.Unlock()

	// Ensure evidence directories exist
	if err := os.MkdirAll(filepath.Join(r.logDir, "captures"), 0755); err != nil {
		r.stopAcceptingEvents()
		return err
	}
	if err := os.MkdirAll(filepath.Join(r.logDir, "templates"), 0755); err != nil {
		r.stopAcceptingEvents()
		return err
	}
	r.startedAt = time.Now().UnixNano() / int64(time.Millisecond)
	if err := r.writeSessionJSON(); err != nil {
		r.stopAcceptingEvents()
		return err
	}
	eventsFile, err := os.OpenFile(filepath.Join(r.logDir, "events.jsonl"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		r.stopAcceptingEvents()
		return err
	}
	r.eventsFile = eventsFile
	r.jsonlEncoder = json.NewEncoder(eventsFile)

	// Start background Goroutine to process captures and write JSON logs
	r.workerWG.Add(1)
	go r.worker(captureFunc)

	// Start Windows message loop on a dedicated OS thread
	errChan := make(chan error, 1)
	r.hookWG.Add(1)
	go func() {
		defer r.hookWG.Done()
		runtimeLockOSThread()
		defer runtimeUnlockOSThread()

		r.hookThreadId = getThreadId()

		// Keyboard hook
		r.kbdHook = setHook(WH_KEYBOARD_LL, syscall.NewCallback(keyboardCallback))
		if r.kbdHook == 0 {
			errChan <- fmt.Errorf("キーボードフックの作成に失敗しました。")
			return
		}

		// Mouse hook
		r.mouseHook = setHook(WH_MOUSE_LL, syscall.NewCallback(mouseCallback))
		if r.mouseHook == 0 {
			unhookWindowsHookEx.Call(r.kbdHook)
			errChan <- fmt.Errorf("マウスフックの作成に失敗しました。")
			return
		}

		errChan <- nil // Signal success setting hooks

		// Windows Message Loop
		var msg MSG
		for {
			ret, _, _ := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
			if int32(ret) <= 0 {
				break
			}
		}

		// Cleanup hooks
		unhookWindowsHookEx.Call(r.kbdHook)
		unhookWindowsHookEx.Call(r.mouseHook)
	}()

	err = <-errChan
	if err != nil {
		r.stopAcceptingEvents()
		if cleanupErr := r.finishRecording(); cleanupErr != nil {
			log.Printf("[Recorder] cleanup after hook startup failure failed: %v", cleanupErr)
		}
		return err
	}

	log.Println("[Recorder] Low-level keyboard and mouse hooks registered successfully.")
	return nil
}

// Stop unhooks events and closes channels

func (r *Recorder) finishRecording() error {
	// Stop hooks first so callbacks can no longer enqueue new senders.
	if r.hookThreadId != 0 {
		postThreadMessage.Call(uintptr(r.hookThreadId), 0x0012, 0, 0) // WM_QUIT = 0x0012
	}
	r.hookWG.Wait()

	// Wait for in-flight senders before closing the channel.
	r.senderWG.Wait()
	r.closeCaptureChannel()

	// Wait for the worker to drain the closed capture channel.
	r.workerWG.Wait()

	r.closeEventsFile()
	return nil
}

func (r *Recorder) Stop() (string, error) {
	r.stopMu.Lock()
	defer r.stopMu.Unlock()

	if !r.stopAcceptingEvents() {
		r.stateMu.Lock()
		alreadyClosed := r.channelClosed
		r.stateMu.Unlock()
		if alreadyClosed {
			return r.logDir, nil
		}
		return "", fmt.Errorf("記録が開始されていません。")
	}

	if err := r.finishRecording(); err != nil {
		return "", err
	}

	// Save final log.json
	logPath := filepath.Join(r.logDir, "log.json")
	r.dataMu.Lock()
	legacyEvents := append([]RecordEvent(nil), r.events...)
	r.dataMu.Unlock()
	data, err := json.MarshalIndent(legacyEvents, "", "  ")
	if err != nil {
		return "", fmt.Errorf("ログJSONの作成に失敗: %v", err)
	}

	err = os.WriteFile(logPath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("ログファイルの書き込みに失敗: %v", err)
	}

	log.Printf("[Recorder] Recording stopped. Events saved to %s", logPath)

	// Emit notification to Svelte frontend
	if r.ctx != nil {
		wailsruntime.EventsEmit(r.ctx, "recording_stopped", r.logDir)
	}

	return r.logDir, nil
}

// OS Thread helper locks
func runtimeLockOSThread() {
	goruntime.LockOSThread()
}
func runtimeUnlockOSThread() {
	goruntime.UnlockOSThread()
}

func getThreadId() uint32 {
	r, _, _ := getCurrentThreadId.Call()
	return uint32(r)
}

func setHook(hookId int, callback uintptr) uintptr {
	h, _, _ := setWindowsHookEx.Call(uintptr(hookId), callback, 0, 0)
	return h
}

// Low-Level Callbacks
func keyboardCallback(code int32, wparam uintptr, lparam uintptr) uintptr {
	if code >= 0 && wparam == WM_KEYDOWN {
		kbd := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))

		// If F8 was pressed, stop recording
		if kbd.VkCode == 0x77 { // VK_F8 = 0x77
			log.Println("[Recorder] F8 detected. Stopping recording...")
			if rec := currentRecorder(); rec != nil {
				go func() {
					rec.Stop()
				}()
			}
			return 1 // Block event propagation
		}

		keyName := getVkKeyName(kbd.VkCode)
		if keyName != "" {
			if rec := currentRecorder(); rec != nil {
				rec.pushEvent(RecordEvent{
					Type: "keydown",
					Key:  keyName,
				})
			}
		}
	}
	ret, _, _ := callNextHookEx.Call(0, uintptr(code), wparam, lparam)
	return ret
}

func mouseCallback(code int32, wparam uintptr, lparam uintptr) uintptr {
	if code >= 0 {
		mouse := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lparam))
		if wparam == WM_LBUTTONDOWN {
			if rec := currentRecorder(); rec != nil {
				rec.pushEvent(RecordEvent{
					Type: "click",
					X:    int(mouse.Pt.X),
					Y:    int(mouse.Pt.Y),
				})
			}
		}
	}
	ret, _, _ := callNextHookEx.Call(0, uintptr(code), wparam, lparam)
	return ret
}

func (r *Recorder) pushEvent(ev RecordEvent) {
	ev.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	ev.Window = getActiveWindowTitle()

	hwnd, _, _ := getForegroundWindow.Call()
	if hwnd != 0 {
		var rect RECT
		ret, _, _ := getWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&rect)))
		if ret != 0 {
			ev.WindowRect = &Rect{
				Left:   int(rect.Left),
				Top:    int(rect.Top),
				Right:  int(rect.Right),
				Bottom: int(rect.Bottom),
			}
			if ev.Type == "click" {
				ev.RelX = ev.X - int(rect.Left)
				ev.RelY = ev.Y - int(rect.Top)
			}
		}
	}

	r.stateMu.Lock()
	if !r.isRecording || r.channelClosed {
		r.stateMu.Unlock()
		return
	}
	r.senderWG.Add(1)
	ch := r.captureChan
	r.stateMu.Unlock()

	go func() {
		defer r.senderWG.Done()
		ch <- ev
	}()
}

// Background Worker processes captures and UIA scans
func (r *Recorder) worker(captureFunc func(string) error) {
	defer r.workerWG.Done()

	for ev := range r.captureChan {
		r.stateMu.Lock()
		recording := r.isRecording
		r.stateMu.Unlock()
		r.processEvent(ev, captureFunc, recording)
	}
}

func (r *Recorder) processEvent(ev RecordEvent, captureFunc func(string) error, allowCapture bool) {
	evidence := EvidenceEvent{
		ID:        fmt.Sprintf("event_%d", ev.Timestamp),
		Timestamp: ev.Timestamp,
		Type:      ev.Type,
		X:         ev.X,
		Y:         ev.Y,
		RelX:      ev.RelX,
		RelY:      ev.RelY,
		Key:       ev.Key,
		Window: WindowInfo{
			Title: ev.Window,
			Rect:  ev.WindowRect,
		},
		CaptureStatus: CaptureStatus{OK: true},
	}

	if ev.Type == "click" {
		ctrl, cid, val := inspectElementAtPoint(ev.X, ev.Y)
		ev.Control = ctrl
		ev.ControlID = cid
		ev.Value = val
		evidence.UIAElement = &UIAElementInfo{ControlType: ctrl, AutomationID: cid, Value: val}

		if allowCapture && captureFunc != nil {
			images := &ImageEvidence{}
			beforeRel := filepath.Join("captures", fmt.Sprintf("event_%d_before.png", ev.Timestamp))
			beforeAbs := filepath.Join(r.logDir, beforeRel)
			if err := r.callCapture(captureFunc, beforeAbs); err != nil {
				evidence.CaptureStatus.OK = false
				evidence.CaptureStatus.Errors = append(evidence.CaptureStatus.Errors, fmt.Sprintf("before screenshot: %v", err))
				log.Printf("[Recorder Worker] Before screen capture failed: %v", err)
			} else {
				images.BeforePath = beforeRel
				markedRel := filepath.Join("captures", fmt.Sprintf("event_%d_marked.png", ev.Timestamp))
				markedAbs := filepath.Join(r.logDir, markedRel)
				var px, py int
				r.dataMu.Lock()
				px, py = r.prevX, r.prevY
				r.prevX, r.prevY = ev.X, ev.Y
				r.dataMu.Unlock()
				if err := DrawMeasurementMarker(beforeAbs, markedAbs, ev.X, ev.Y, px, py); err != nil {
					evidence.CaptureStatus.OK = false
					evidence.CaptureStatus.Errors = append(evidence.CaptureStatus.Errors, fmt.Sprintf("marker: %v", err))
					log.Printf("[Recorder Worker] Marker drawing failed: %v", err)
				} else {
					images.MarkedPath = markedRel
					ev.ImagePath = markedRel
				}
			}

			time.Sleep(150 * time.Millisecond)
			afterRel := filepath.Join("captures", fmt.Sprintf("event_%d_after.png", ev.Timestamp))
			afterAbs := filepath.Join(r.logDir, afterRel)
			if err := r.callCapture(captureFunc, afterAbs); err != nil {
				evidence.CaptureStatus.OK = false
				evidence.CaptureStatus.Errors = append(evidence.CaptureStatus.Errors, fmt.Sprintf("after screenshot: %v", err))
				log.Printf("[Recorder Worker] After screen capture failed: %v", err)
			} else {
				images.AfterPath = afterRel
			}
			if images.BeforePath != "" || images.AfterPath != "" || images.MarkedPath != "" {
				evidence.Images = images
			}
		}
	}

	r.dataMu.Lock()
	r.events = append(r.events, ev)
	r.dataMu.Unlock()

	if r.jsonlEncoder != nil {
		if err := r.jsonlEncoder.Encode(evidence); err != nil {
			log.Printf("[Recorder Worker] events.jsonl write failed: %v", err)
		}
	}
}

func (r *Recorder) callCapture(captureFunc func(string) error, outputPath string) (err error) {
	if captureFunc == nil {
		return fmt.Errorf("capture function is nil")
	}

	semTimer := time.NewTimer(captureTimeout)
	select {
	case r.captureSem <- struct{}{}:
		if !semTimer.Stop() {
			select {
			case <-semTimer.C:
			default:
			}
		}
		// Continue with this single capture. The semaphore bounds leaked capture
		// goroutines if a timed-out capture function never returns.
	case <-semTimer.C:
		return fmt.Errorf("capture skipped because another capture is still running after %s", captureTimeout)
	}

	done := make(chan error, 1)
	go func() {
		defer func() {
			<-r.captureSem
			if recovered := recover(); recovered != nil {
				done <- fmt.Errorf("capture panic: %v", recovered)
			}
		}()
		done <- captureFunc(outputPath)
	}()

	runTimer := time.NewTimer(captureTimeout)
	defer runTimer.Stop()
	select {
	case err := <-done:
		return err
	case <-runTimer.C:
		return fmt.Errorf("capture timed out after %s", captureTimeout)
	}
}

// Active Window helper
func getActiveWindowTitle() string {
	hwnd, _, _ := getForegroundWindow.Call()
	if hwnd == 0 {
		return "Unknown Window"
	}
	var buf [256]uint16
	getWindowText.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return syscall.UTF16ToString(buf[:])
}

// UI Automation OLE inspection
func inspectElementAtPoint(x, y int) (controlType string, automationId string, value string) {
	err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
	if err != nil {
		return "", "", ""
	}
	defer ole.CoUninitialize()

	unknown, err := ole.CreateInstance(CLSID_CUIAutomation, IID_IUIAutomation)
	if err != nil {
		return "", "", ""
	}
	defer unknown.Release()

	uiaDisp, err := unknown.QueryInterface(IID_IUIAutomation)
	if err != nil {
		return "", "", ""
	}
	defer uiaDisp.Release()

	pointVal := (int64(y) << 32) | (int64(x) & 0xFFFFFFFF)

	var elem *ole.IDispatch
	res, err := oleutil.CallMethod(uiaDisp, "ElementFromPoint", pointVal)
	if err != nil || res.Val == 0 {
		return "Unknown", "", ""
	}
	elem = res.ToIDispatch()
	defer elem.Release()

	// Get Control Type name
	typeIdVar, _ := oleutil.GetProperty(elem, "CurrentControlType")
	typeId := typeIdVar.Val
	controlType = translateControlType(int(typeId))

	// Get Name
	nameVar, _ := oleutil.GetProperty(elem, "CurrentName")
	name := nameVar.ToString()

	// Get Automation ID
	autoIdVar, _ := oleutil.GetProperty(elem, "CurrentAutomationId")
	automationId = autoIdVar.ToString()

	// Get ClassName
	classVar, _ := oleutil.GetProperty(elem, "CurrentClassName")
	className := classVar.ToString()

	if controlType == "Unknown" && className != "" {
		controlType = className
	}

	// Read Toggle/Value patterns if CheckBox/ComboBox
	if controlType == "CheckBox" || controlType == "RadioButton" {
		resPat, errPat := oleutil.CallMethod(elem, "GetCurrentPattern", 10003)
		if errPat == nil && resPat.Val != 0 {
			togglePat := resPat.ToIDispatch()
			defer togglePat.Release()

			stateVar, errState := oleutil.GetProperty(togglePat, "CurrentToggleState")
			if errState == nil {
				switch stateVar.Val {
				case 0:
					value = "Unchecked"
				case 1:
					value = "Checked"
				}
			}
		}
	} else {
		resPat, errPat := oleutil.CallMethod(elem, "GetCurrentPattern", 10002)
		if errPat == nil && resPat.Val != 0 {
			valPat := resPat.ToIDispatch()
			defer valPat.Release()
			valVar, errVal := oleutil.GetProperty(valPat, "CurrentValue")
			if errVal == nil {
				value = valVar.ToString()
			}
		}
	}

	if value == "" && name != "" {
		value = name
	}

	return controlType, automationId, value
}

func translateControlType(id int) string {
	switch id {
	case 50002:
		return "Button"
	case 50004:
		return "CheckBox"
	case 50005:
		return "ComboBox"
	case 50006:
		return "Edit"
	case 50007:
		return "Hyperlink"
	case 50008:
		return "Image"
	case 50009:
		return "ListItem"
	case 50010:
		return "List"
	case 50011:
		return "Menu"
	case 50012:
		return "MenuBar"
	case 50013:
		return "MenuItem"
	case 50014:
		return "ProgressBar"
	case 50015:
		return "RadioButton"
	case 50016:
		return "ScrollBar"
	case 50017:
		return "Slider"
	case 50018:
		return "Spinner"
	case 50019:
		return "StatusBar"
	case 50020:
		return "Tab"
	case 50021:
		return "TabItem"
	case 50022:
		return "Text"
	case 50023:
		return "ToolBar"
	case 50024:
		return "ToolTip"
	case 50025:
		return "Tree"
	case 50026:
		return "TreeItem"
	case 50031:
		return "Window"
	default:
		return "Unknown"
	}
}

func getVkKeyName(vk uint32) string {
	if vk >= 0x30 && vk <= 0x39 { // 0-9
		return string(rune(vk))
	}
	if vk >= 0x41 && vk <= 0x5A { // A-Z
		return string(rune(vk))
	}
	switch vk {
	case 0x08:
		return "BACKSPACE"
	case 0x09:
		return "TAB"
	case 0x0D:
		return "ENTER"
	case 0x10:
		return "SHIFT"
	case 0x11:
		return "CTRL"
	case 0x12:
		return "ALT"
	case 0x20:
		return "SPACE"
	case 0x25:
		return "LEFT"
	case 0x26:
		return "UP"
	case 0x27:
		return "RIGHT"
	case 0x28:
		return "DOWN"
	default:
		return ""
	}
}
