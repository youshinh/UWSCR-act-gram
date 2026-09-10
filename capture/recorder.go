package capture

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	goruntime "runtime"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Windows API constants
const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14
	WM_KEYDOWN     = 0x0100
	WM_LBUTTONDOWN = 0x0201
	WM_RBUTTONDOWN = 0x0204
	WM_QUIT        = 0x0012

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
	Type       string `json:"type"`      // "click", "input", "key_action", "keydown"
	X          int    `json:"x,omitempty"`
	Y          int    `json:"y,omitempty"`
	RelX       int    `json:"rel_x,omitempty"`
	RelY       int    `json:"rel_y,omitempty"`
	Key        string `json:"key,omitempty"`
	Window     string `json:"window"`
	Control    string `json:"control,omitempty"`    // e.g. "Button", "ComboBox"
	ControlID  string `json:"control_id,omitempty"` // AutomationId
	Value      string `json:"value,omitempty"`      // Value, state, or typed text
	ImagePath  string `json:"image_path,omitempty"` // Screenshot path (clean raw image)
	WindowRect *Rect  `json:"-"`                    // Active window rect for evidence output
}

// Recorder manages the Win32 hooks and UI Automation
type Recorder struct {
	ctx context.Context

	stateMu       sync.Mutex
	isRecording   bool
	channelClosed bool

	dataMu       sync.Mutex
	logDir       string
	events       []RecordEvent
	lastWindow   string
	prevX, prevY int
	startedAt    int64
	eventsFile   *os.File
	jsonlEncoder *json.Encoder

	// Typing aggregation buffer
	keyBuf   strings.Builder
	bufMu    sync.Mutex
	lastType time.Time

	// Hook handles
	mouseHook    uintptr
	kbdHook      uintptr
	hookThreadId uint32

	// Async capture channel and concurrency control
	captureChan chan RecordEvent
	captureSem  chan struct{}
	senderWG    sync.WaitGroup
	workerWG    sync.WaitGroup
	hookWG      sync.WaitGroup
}

var (
	user32                   = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx         = user32.NewProc("SetWindowsHookExW")
	callNextHookEx           = user32.NewProc("CallNextHookEx")
	unhookWindowsHookEx      = user32.NewProc("UnhookWindowsHookEx")
	getMessage               = user32.NewProc("GetMessageW")
	translateMessage         = user32.NewProc("TranslateMessage")
	dispatchMessage          = user32.NewProc("DispatchMessageW")
	postThreadMessage        = user32.NewProc("PostThreadMessageW")
	getForegroundWindow      = user32.NewProc("GetForegroundWindow")
	getWindowText            = user32.NewProc("GetWindowTextW")
	getWindowRect            = user32.NewProc("GetWindowRect")
	getWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
	getCurrentThreadId       = kernel32.NewProc("GetCurrentThreadId")
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")

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
		captureChan: make(chan RecordEvent, 200),
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
	r.stateMu.Lock()
	if r.isRecording {
		r.stateMu.Unlock()
		return fmt.Errorf("すでに記録中です。")
	}
	r.captureChan = make(chan RecordEvent, 200)
	r.channelClosed = false
	r.isRecording = true
	setGlobalRecorder(r)
	r.stateMu.Unlock()

	// Ensure evidence directories exist
	if err := os.MkdirAll(filepath.Join(r.logDir, "captures"), 0755); err != nil {
		r.stateMu.Lock()
		r.isRecording = false
		clearGlobalRecorder(r)
		r.stateMu.Unlock()
		return err
	}
	if err := os.MkdirAll(filepath.Join(r.logDir, "templates"), 0755); err != nil {
		r.stateMu.Lock()
		r.isRecording = false
		clearGlobalRecorder(r)
		r.stateMu.Unlock()
		return err
	}
	r.startedAt = time.Now().UnixNano() / int64(time.Millisecond)
	if err := r.writeSessionJSON(); err != nil {
		r.stateMu.Lock()
		r.isRecording = false
		clearGlobalRecorder(r)
		r.stateMu.Unlock()
		return err
	}
	eventsFile, err := os.OpenFile(filepath.Join(r.logDir, "events.jsonl"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		r.stateMu.Lock()
		r.isRecording = false
		clearGlobalRecorder(r)
		r.stateMu.Unlock()
		return err
	}
	r.eventsFile = eventsFile
	r.jsonlEncoder = json.NewEncoder(eventsFile)

	// Start background Goroutine to process captures and write JSON logs
	r.workerWG.Add(1)
	go r.worker(captureFunc)

	// Start Windows message loop on a dedicated, locked OS thread
	errChan := make(chan error, 1)
	r.hookWG.Add(1)
	go func() {
		goruntime.LockOSThread()
		defer goruntime.UnlockOSThread()
		defer r.hookWG.Done()

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
			r.kbdHook = 0
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
			translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
			dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
		}

		// Cleanup hooks
		if r.kbdHook != 0 {
			unhookWindowsHookEx.Call(r.kbdHook)
			r.kbdHook = 0
		}
		if r.mouseHook != 0 {
			unhookWindowsHookEx.Call(r.mouseHook)
			r.mouseHook = 0
		}
	}()

	err = <-errChan
	if err != nil {
		r.stopAcceptingEvents()
		r.hookWG.Wait()
		r.senderWG.Wait()
		r.closeCaptureChannel()
		r.workerWG.Wait()
		r.closeEventsFile()
		clearGlobalRecorder(r)
		return err
	}

	log.Println("[Recorder] Low-level keyboard and mouse hooks registered successfully.")
	return nil
}

// Stop unhooks events and closes channels
func (r *Recorder) Stop() (string, error) {
	if !r.stopAcceptingEvents() {
		return "", fmt.Errorf("記録が開始されていません。")
	}

	// Flush any pending typing buffer
	r.flushKeyBuffer()

	// Post WM_QUIT to hook thread message queue to break message loop
	if r.hookThreadId != 0 {
		postThreadMessage.Call(uintptr(r.hookThreadId), WM_QUIT, 0, 0)
	}
	r.hookWG.Wait()

	// Wait for in-flight senders before closing the channel
	r.senderWG.Wait()
	r.closeCaptureChannel()

	// Wait for the worker to drain the closed capture channel
	r.workerWG.Wait()

	r.closeEventsFile()
	clearGlobalRecorder(r)

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

func getThreadId() uint32 {
	r, _, _ := getCurrentThreadId.Call()
	return uint32(r)
}

func setHook(hookId int, callback uintptr) uintptr {
	h, _, _ := setWindowsHookEx.Call(uintptr(hookId), callback, 0, 0)
	return h
}

// Low-Level Callbacks
func keyboardCallback(code int32, wparam uintptr, kbd *KBDLLHOOKSTRUCT) uintptr {
	if code >= 0 && wparam == WM_KEYDOWN && kbd != nil {
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

		if rec := currentRecorder(); rec != nil {
			// 自アプリウィンドウでのキー入力は除外
			hwnd, _, _ := getForegroundWindow.Call()
			if !isOwnWindow(hwnd) {
				rec.handleKeyDown(kbd.VkCode)
			}
		}
	}
	ret, _, _ := callNextHookEx.Call(0, uintptr(code), wparam, uintptr(unsafe.Pointer(kbd)))
	return ret
}

func mouseCallback(code int32, wparam uintptr, mouse *MSLLHOOKSTRUCT) uintptr {
	if code >= 0 && wparam == WM_LBUTTONDOWN && mouse != nil {
		if rec := currentRecorder(); rec != nil {
			// 自アプリ（actgram本体やミニウィンドウ）のクリックは除外
			hwnd, _, _ := getForegroundWindow.Call()
			if !isOwnWindow(hwnd) {
				// 直前の文字入力を確定
				rec.flushKeyBuffer()

				rec.pushEvent(RecordEvent{
					Type: "click",
					X:    int(mouse.Pt.X),
					Y:    int(mouse.Pt.Y),
				})
			}
		}
	}
	ret, _, _ := callNextHookEx.Call(0, uintptr(code), wparam, uintptr(unsafe.Pointer(mouse)))
	return ret
}

func isOwnWindow(hwnd uintptr) bool {
	if hwnd == 0 {
		return false
	}
	var pid uint32
	getWindowThreadProcessId.Call(hwnd, uintptr(unsafe.Pointer(&pid)))
	return pid == uint32(os.Getpid())
}

func (r *Recorder) handleKeyDown(vk uint32) {
	keyName := getVkKeyName(vk)
	if keyName == "" {
		return
	}

	// 特殊キー（Enter, Tab, Esc等）はバッファを確定して単独アクション発行
	if vk == 0x0D || vk == 0x09 || vk == 0x1B {
		r.flushKeyBuffer()
		r.pushEvent(RecordEvent{
			Type: "key_action",
			Key:  keyName,
		})
		return
	}

	// Backspace
	if vk == 0x08 {
		r.bufMu.Lock()
		defer r.bufMu.Unlock()
		s := r.keyBuf.String()
		if len(s) > 0 {
			r.keyBuf.Reset()
			r.keyBuf.WriteString(s[:len(s)-1])
		}
		return
	}

	// 通常文字
	r.bufMu.Lock()
	r.keyBuf.WriteString(keyName)
	r.lastType = time.Now()
	r.bufMu.Unlock()
}

func (r *Recorder) flushKeyBuffer() {
	r.bufMu.Lock()
	defer r.bufMu.Unlock()
	if r.keyBuf.Len() > 0 {
		text := r.keyBuf.String()
		r.keyBuf.Reset()
		r.pushEvent(RecordEvent{
			Type:  "input",
			Value: text,
		})
	}
}

func (r *Recorder) pushEvent(ev RecordEvent) {
	r.stateMu.Lock()
	if !r.isRecording || r.channelClosed {
		r.stateMu.Unlock()
		return
	}
	r.senderWG.Add(1)
	ch := r.captureChan
	r.stateMu.Unlock()

	defer r.senderWG.Done()

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

	ch <- ev
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
				evidence.CaptureStatus.Errors = append(evidence.CaptureStatus.Errors, fmt.Sprintf("screenshot: %v", err))
				log.Printf("[Recorder Worker] Screen capture failed: %v", err)
			} else {
				images.BeforePath = beforeRel
				// 非破壊キャプチャ: ev.ImagePath は元のクリーン画像を設定
				ev.ImagePath = beforeRel

				// 後方互換・エビデンス用にマーカー画像も生成
				markedRel := filepath.Join("captures", fmt.Sprintf("event_%d_marked.png", ev.Timestamp))
				markedAbs := filepath.Join(r.logDir, markedRel)
				var px, py int
				r.dataMu.Lock()
				px, py = r.prevX, r.prevY
				r.prevX, r.prevY = ev.X, ev.Y
				r.dataMu.Unlock()

				if err := DrawMeasurementMarker(beforeAbs, markedAbs, ev.X, ev.Y, px, py); err != nil {
					evidence.CaptureStatus.Errors = append(evidence.CaptureStatus.Errors, fmt.Sprintf("marker: %v", err))
					log.Printf("[Recorder Worker] Marker drawing failed: %v", err)
				} else {
					images.MarkedPath = markedRel
				}
			}

			if images.BeforePath != "" || images.MarkedPath != "" {
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
		} else if r.eventsFile != nil {
			if err := r.eventsFile.Sync(); err != nil {
				log.Printf("[Recorder Worker] events.jsonl sync failed: %v", err)
			}
		}
	}
}

func (r *Recorder) callCapture(captureFunc func(string) error, outputPath string) (err error) {
	if captureFunc == nil {
		return fmt.Errorf("capture function is nil")
	}

	select {
	case r.captureSem <- struct{}{}:
		// acquired semaphore
	case <-time.After(captureTimeout):
		return fmt.Errorf("capture skipped because previous capture is still running after %s", captureTimeout)
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

	select {
	case err := <-done:
		return err
	case <-time.After(captureTimeout):
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
	case 50018:
		return "Tab"
	case 50019:
		return "TabItem"
	case 50020:
		return "Text"
	case 50021:
		return "ToolBar"
	case 50023:
		return "Tree"
	case 50024:
		return "TreeItem"
	case 50025:
		return "Custom"
	case 50026:
		return "Group"
	case 50028:
		return "Pane"
	case 50030:
		return "Window"
	case 50032:
		return "Document"
	case 50033:
		return "SplitButton"
	default:
		return "Unknown"
	}
}

// Convert VK codes to readable key representations
func getVkKeyName(vk uint32) string {
	switch vk {
	case 0x08:
		return "BACKSPACE"
	case 0x09:
		return "TAB"
	case 0x0D:
		return "ENTER"
	case 0x14:
		return "CAPSLOCK"
	case 0x1B:
		return "ESC"
	case 0x20:
		return "SPACE"
	case 0x21:
		return "PGUP"
	case 0x22:
		return "PGDN"
	case 0x23:
		return "END"
	case 0x24:
		return "HOME"
	case 0x25:
		return "LEFT"
	case 0x26:
		return "UP"
	case 0x27:
		return "RIGHT"
	case 0x28:
		return "DOWN"
	case 0x2D:
		return "INSERT"
	case 0x2E:
		return "DEL"
	case 0x70:
		return "F1"
	case 0x71:
		return "F2"
	case 0x72:
		return "F3"
	case 0x73:
		return "F4"
	case 0x74:
		return "F5"
	case 0x75:
		return "F6"
	case 0x76:
		return "F7"
	case 0x77:
		return "F8"
	case 0x78:
		return "F9"
	case 0x79:
		return "F10"
	case 0x7A:
		return "F11"
	case 0x7B:
		return "F12"
	case 0x10, 0x11, 0x12, 0x5B, 0x5C:
		return "" // Shift, Ctrl, Alt, Win keys
	}

	// Alphanumeric keys
	if (vk >= '0' && vk <= '9') || (vk >= 'A' && vk <= 'Z') {
		return string(rune(vk))
	}

	return ""
}
