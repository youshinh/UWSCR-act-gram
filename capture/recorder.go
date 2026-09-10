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
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Windows API constants
const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14
	WM_KEYDOWN     = 0x0100
	WM_LBUTTONDOWN = 0x0201
	WM_RBUTTONDOWN = 0x0204
	WM_QUIT        = 0x0012
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
	Timestamp int64  `json:"timestamp"` // Epoch ms
	Type      string `json:"type"`      // "click", "input", "key_action"
	X         int    `json:"x,omitempty"`
	Y         int    `json:"y,omitempty"`
	RelX      int    `json:"rel_x,omitempty"`
	RelY      int    `json:"rel_y,omitempty"`
	Key       string `json:"key,omitempty"`
	Window    string `json:"window"`
	Control   string `json:"control,omitempty"`    // e.g. "Button", "ComboBox"
	ControlID string `json:"control_id,omitempty"` // AutomationId
	Value     string `json:"value,omitempty"`      // Value, state, or typed text
	ImagePath string `json:"image_path,omitempty"` // Screenshot path
}

// Recorder manages the Win32 hooks and UI Automation
type Recorder struct {
	ctx          context.Context
	mu           sync.Mutex
	isRecording  bool
	logDir       string
	events       []RecordEvent
	lastWindow   string
	prevX, prevY int

	// Typing aggregation buffer
	keyBuf   strings.Builder
	bufMu    sync.Mutex
	lastType time.Time
	
	// Hook handles
	mouseHook    uintptr
	kbdHook      uintptr
	hookThreadId uint32

	// Async capture channel
	captureChan chan RecordEvent
	wg          sync.WaitGroup
	quitChan    chan struct{}
}

var (
	user32                     = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx           = user32.NewProc("SetWindowsHookExW")
	callNextHookEx             = user32.NewProc("CallNextHookEx")
	unhookWindowsHookEx        = user32.NewProc("UnhookWindowsHookEx")
	getMessage                 = user32.NewProc("GetMessageW")
	translateMessage           = user32.NewProc("TranslateMessage")
	dispatchMessage            = user32.NewProc("DispatchMessageW")
	postThreadMessage          = user32.NewProc("PostThreadMessageW")
	getForegroundWindow        = user32.NewProc("GetForegroundWindow")
	getWindowText              = user32.NewProc("GetWindowTextW")
	getWindowRect              = user32.NewProc("GetWindowRect")
	getWindowThreadProcessId   = user32.NewProc("GetWindowThreadProcessId")
	getCurrentThreadId         = kernel32.NewProc("GetCurrentThreadId")
	kernel32                   = syscall.NewLazyDLL("kernel32.dll")

	// Global recorder instance pointer for callback access
	globalRecorder *Recorder
	globalMu       sync.RWMutex
)

func NewRecorder(ctx context.Context, logDir string) *Recorder {
	return &Recorder{
		ctx:         ctx,
		logDir:      logDir,
		events:      make([]RecordEvent, 0),
		captureChan: make(chan RecordEvent, 200),
		quitChan:    make(chan struct{}),
	}
}

// Start hooks keyboard and mouse events
func (r *Recorder) Start(captureFunc func(string) error) error {
	r.mu.Lock()
	if r.isRecording {
		r.mu.Unlock()
		return fmt.Errorf("すでに記録中です。")
	}
	r.isRecording = true
	globalMu.Lock()
	globalRecorder = r
	globalMu.Unlock()
	r.mu.Unlock()

	// Ensure directories exist
	os.MkdirAll(filepath.Join(r.logDir, "captures"), 0755)

	// Start background Goroutine to process captures and write JSON logs
	r.wg.Add(1)
	go r.worker(captureFunc)

	// Start Windows message loop on a dedicated, locked OS thread
	errChan := make(chan error, 1)
	go func() {
		goruntime.LockOSThread()
		defer goruntime.UnlockOSThread()

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

	err := <-errChan
	if err != nil {
		r.mu.Lock()
		r.isRecording = false
		r.mu.Unlock()
		globalMu.Lock()
		globalRecorder = nil
		globalMu.Unlock()
		return err
	}

	log.Println("[Recorder] Low-level keyboard and mouse hooks registered successfully.")
	return nil
}

// Stop unhooks events and closes channels
func (r *Recorder) Stop() (string, error) {
	r.mu.Lock()
	if !r.isRecording {
		r.mu.Unlock()
		return "", fmt.Errorf("記録が開始されていません。")
	}
	r.isRecording = false
	r.mu.Unlock()

	// Flush any pending typing buffer
	r.flushKeyBuffer()

	// Post WM_QUIT to hook thread message queue
	if r.hookThreadId != 0 {
		postThreadMessage.Call(uintptr(r.hookThreadId), WM_QUIT, 0, 0)
	}

	// Close worker
	close(r.quitChan)
	r.wg.Wait()

	globalMu.Lock()
	globalRecorder = nil
	globalMu.Unlock()

	// Save final log.json
	logPath := filepath.Join(r.logDir, "log.json")
	data, err := json.MarshalIndent(r.events, "", "  ")
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
		wailsRuntime.EventsEmit(r.ctx, "recording_stopped", r.logDir)
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
			go func() {
				globalMu.RLock()
				rec := globalRecorder
				globalMu.RUnlock()
				if rec != nil {
					rec.Stop()
				}
			}()
			return 1 // Block event propagation
		}

		globalMu.RLock()
		rec := globalRecorder
		globalMu.RUnlock()

		if rec != nil {
			// 自アプリウィンドウでのキー入力は無視
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
	if code >= 0 && mouse != nil {
		if wparam == WM_LBUTTONDOWN {
			globalMu.RLock()
			rec := globalRecorder
			globalMu.RUnlock()

			if rec != nil {
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
	r.mu.Lock()
	if !r.isRecording {
		r.mu.Unlock()
		return
	}
	r.mu.Unlock()

	ev.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	ev.Window = getActiveWindowTitle()

	if ev.Type == "click" {
		hwnd, _, _ := getForegroundWindow.Call()
		if hwnd != 0 {
			var rect RECT
			ret, _, _ := getWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&rect)))
			if ret != 0 {
				ev.RelX = ev.X - int(rect.Left)
				ev.RelY = ev.Y - int(rect.Top)
			}
		}
	}

	select {
	case r.captureChan <- ev:
	default:
		log.Println("[Recorder] Warning: capture queue full, dropped event")
	}
}

// Background Worker processes captures and UIA scans
func (r *Recorder) worker(captureFunc func(string) error) {
	defer r.wg.Done()

	for {
		select {
		case ev := <-r.captureChan:
			// 1. Scan UI Automation element details if it's a click
			if ev.Type == "click" {
				ctrl, cid, val := inspectElementAtPoint(ev.X, ev.Y)
				ev.Control = ctrl
				ev.ControlID = cid
				ev.Value = val

				// 2. Automated clean capture on click (非破壊: 画像に直接マーカーは書き込まない)
				imgFileName := fmt.Sprintf("event_%d.png", ev.Timestamp)
				relPath := filepath.Join("captures", imgFileName)
				absPath := filepath.Join(r.logDir, relPath)

				err := captureFunc(absPath)
				if err == nil {
					ev.ImagePath = relPath
					r.mu.Lock()
					r.prevX, r.prevY = ev.X, ev.Y
					r.mu.Unlock()
				} else {
					log.Printf("[Recorder Worker] Screen capture failed: %v", err)
				}
			}

			// Append to event list
			r.mu.Lock()
			r.events = append(r.events, ev)
			r.mu.Unlock()

		case <-r.quitChan:
			// Process any remaining channel items
			for len(r.captureChan) > 0 {
				ev := <-r.captureChan
				if ev.Type == "click" {
					ctrl, cid, val := inspectElementAtPoint(ev.X, ev.Y)
					ev.Control = ctrl
					ev.ControlID = cid
					ev.Value = val
				}
				r.mu.Lock()
				r.events = append(r.events, ev)
				r.mu.Unlock()
			}
			return
		}
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
