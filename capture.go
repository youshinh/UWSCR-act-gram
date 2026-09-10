package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")
	procGetDC               = user32.NewProc("GetDC")
	procReleaseDC           = user32.NewProc("ReleaseDC")
	procGetSystemMetrics    = user32.NewProc("GetSystemMetrics")
	procSetProcessDPIAware  = user32.NewProc("SetProcessDPIAware")

	gdi32                  = syscall.NewLazyDLL("gdi32.dll")
	procCreateCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
	procCreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
	procSelectObject       = gdi32.NewProc("SelectObject")
	procBitBlt             = gdi32.NewProc("BitBlt")
	procGetDIBits          = gdi32.NewProc("GetDIBits")
	procDeleteObject       = gdi32.NewProc("DeleteObject")
	procDeleteDC           = gdi32.NewProc("DeleteDC")
)

const (
	SM_XVIRTUALSCREEN  = 76
	SM_YVIRTUALSCREEN  = 77
	SM_CXVIRTUALSCREEN = 78
	SM_CYVIRTUALSCREEN = 79
	SM_CXSCREEN        = 0
	SM_CYSCREEN        = 1
	SRCCOPY            = 0x00CC0020
	DIB_RGB_COLORS     = 0
	BI_RGB             = 0
)

type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]uint32
}

func init() {
	// プロセス起動時にDPI Awareを有効化し、高DPIディスプレイでの座標ズレを防止
	if procSetProcessDPIAware.Find() == nil {
		procSetProcessDPIAware.Call()
	}
}

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

// CaptureScreen は Win32 GDI API を直接呼び出してデスクトップ画面を高速キャプチャ（10ms前後）し、PNG形式で保存します。
func CaptureScreen(outputPath string) error {
	// 出力先ディレクトリの確保
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("ディレクトリ作成エラー: %v", err)
	}

	// 仮想スクリーン（全モニタ合算）のサイズ取得
	x, _, _ := procGetSystemMetrics.Call(SM_XVIRTUALSCREEN)
	y, _, _ := procGetSystemMetrics.Call(SM_YVIRTUALSCREEN)
	w, _, _ := procGetSystemMetrics.Call(SM_CXVIRTUALSCREEN)
	h, _, _ := procGetSystemMetrics.Call(SM_CYVIRTUALSCREEN)

	width := int(w)
	height := int(h)
	startX := int(x)
	startY := int(y)

	// マルチモニタ非対応または取得失敗時のフォールバック
	if width <= 0 || height <= 0 {
		pw, _, _ := procGetSystemMetrics.Call(SM_CXSCREEN)
		ph, _, _ := procGetSystemMetrics.Call(SM_CYSCREEN)
		width = int(pw)
		height = int(ph)
		startX = 0
		startY = 0
	}

	if width <= 0 || height <= 0 {
		return fmt.Errorf("画面サイズの取得に失敗しました (width=%d, height=%d)", width, height)
	}

	// デスクトップDC取得
	hDesktopDC, _, _ := procGetDC.Call(0)
	if hDesktopDC == 0 {
		return fmt.Errorf("GetDC に失敗しました")
	}
	defer procReleaseDC.Call(0, hDesktopDC)

	// 互換メモリDCとビットマップ作成
	hMemDC, _, _ := procCreateCompatibleDC.Call(hDesktopDC)
	if hMemDC == 0 {
		return fmt.Errorf("CreateCompatibleDC に失敗しました")
	}
	defer procDeleteDC.Call(hMemDC)

	hBitmap, _, _ := procCreateCompatibleBitmap.Call(hDesktopDC, uintptr(width), uintptr(height))
	if hBitmap == 0 {
		return fmt.Errorf("CreateCompatibleBitmap に失敗しました")
	}
	defer procDeleteObject.Call(hBitmap)

	oldObj, _, _ := procSelectObject.Call(hMemDC, hBitmap)
	defer procSelectObject.Call(hMemDC, oldObj)

	// 画面からメモリDCへ転送 (BitBlt)
	ret, _, _ := procBitBlt.Call(
		hMemDC, 0, 0, uintptr(width), uintptr(height),
		hDesktopDC, uintptr(startX), uintptr(startY), SRCCOPY,
	)
	if ret == 0 {
		return fmt.Errorf("BitBlt に失敗しました")
	}

	// DIBits の取得用ヘッダ設定 (負の高さにすることで上から下への順序(top-down)にする)
	var bi BITMAPINFO
	bi.BmiHeader.BiSize = uint32(unsafe.Sizeof(bi.BmiHeader))
	bi.BmiHeader.BiWidth = int32(width)
	bi.BmiHeader.BiHeight = -int32(height) // top-down
	bi.BmiHeader.BiPlanes = 1
	bi.BmiHeader.BiBitCount = 32
	bi.BmiHeader.BiCompression = BI_RGB

	// ピクセルバッファ (BGRA 4バイト/ピクセル)
	bufSize := width * height * 4
	rawPixels := make([]byte, bufSize)

	ret, _, _ = procGetDIBits.Call(
		hMemDC, hBitmap, 0, uintptr(height),
		uintptr(unsafe.Pointer(&rawPixels[0])),
		uintptr(unsafe.Pointer(&bi)),
		DIB_RGB_COLORS,
	)
	if ret == 0 {
		return fmt.Errorf("GetDIBits に失敗しました")
	}

	// Goの画像オブジェクトに変換 (BGRA -> RGBA)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < len(rawPixels); i += 4 {
		b := rawPixels[i]
		g := rawPixels[i+1]
		r := rawPixels[i+2]
		a := uint8(255) // アルファは完全不透明に設定
		img.SetRGBA((i/4)%width, (i/4)/width, color.RGBA{R: r, G: g, B: b, A: a})
	}

	// PNGファイル保存
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("画像ファイル作成エラー: %v", err)
	}
	defer outFile.Close()

	if err := png.Encode(outFile, img); err != nil {
		return fmt.Errorf("PNGエンコードエラー: %v", err)
	}

	return nil
}
