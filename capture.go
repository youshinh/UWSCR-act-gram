package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")
)

// GetActiveWindowTitle は現在のアクティブウィンドウのタイトルを取得します
func GetActiveWindowTitle() string {
	hwnd, _, _ := procGetForegroundWindow.Call()
	if hwnd == 0 {
		return "デスクトップ"
	}

	buf := make([]uint16, 512)
	procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	title := syscall.UTF16ToString(buf)
	if title == "" {
		return "不明なウィンドウ"
	}
	return title
}

// CaptureScreen はWindowsのPowerShellを利用して画面全体をキャプチャし、指定パスに保存します
func CaptureScreen(outputPath string) error {
	// PowerShellで.NETのSystem.Drawingを使用してスクリーンショットを取得するスクリプト
	script := fmt.Sprintf(`
Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing

# Set DPI Aware to get correct physical bounds under Windows DPI Scaling
$shcore = '[DllImport("user32.dll")] public static extern bool SetProcessDPIAware();'
$type = Add-Type -MemberDefinition $shcore -Name "Win32DPI" -Namespace "Win32" -PassThru
$type::SetProcessDPIAware()

$screen = [System.Windows.Forms.Screen]::PrimaryScreen
$bitmap = New-Object System.Drawing.Bitmap $screen.Bounds.Width, $screen.Bounds.Height
$graphics = [System.Drawing.Graphics]::FromImage($bitmap)
$graphics.CopyFromScreen($screen.Bounds.X, $screen.Bounds.Y, 0, 0, $bitmap.Size)
$bitmap.Save('%s', [System.Drawing.Imaging.ImageFormat]::Png)
$graphics.Dispose()
$bitmap.Dispose()
`, filepath.Clean(outputPath))

	tempFile, err := os.CreateTemp("", "capture-*.ps1")
	if err != nil {
		return fmt.Errorf("一時スクリプトの作成に失敗しました: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(script); err != nil {
		tempFile.Close()
		return fmt.Errorf("スクリプトの書き込みに失敗しました: %v", err)
	}
	tempFile.Close()

	// PowerShellを実行して画面をキャプチャ
	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", tempFile.Name())
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("PowerShell実行エラー: %v, 詳細: %s", err, string(output))
	}
	return nil
}
