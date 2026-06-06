package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Orchestrator はUWSCR子プロセスの実行を管理します。
type Orchestrator struct {
	app       *App
	activeCmd *exec.Cmd
	mu        sync.Mutex
}

func NewOrchestrator(app *App) *Orchestrator {
	return &Orchestrator{app: app}
}

// FindUWSCRPath は優先度に従って uwscr.exe のパスを探索します。
func (o *Orchestrator) FindUWSCRPath() (string, error) {
	// 優先度 1: 実行中の actgram.exe と同じディレクトリ
	exePath, err := os.Executable()
	if err == nil {
		localPath := filepath.Join(filepath.Dir(exePath), "uwscr.exe")
		if _, err := os.Stat(localPath); err == nil {
			return localPath, nil
		}
	}

	// 優先度 2: config.yaml で指定されたパス
	if o.app.cfg != nil && o.app.cfg.UWSCRPath != "" {
		if _, err := os.Stat(o.app.cfg.UWSCRPath); err == nil {
			return o.app.cfg.UWSCRPath, nil
		}
	}

	// 優先度 3: システム環境変数 PATH
	path, err := exec.LookPath("uwscr.exe")
	if err == nil {
		return path, nil
	}

	return "", fmt.Errorf("uwscr.exe が見つかりません。本家UWSCR配布ページ (https://github.com/stuncloud/UWSCR/releases) からダウンロードし、本エージェント (actgram.exe) と同じディレクトリに配置するか、設定からパスを指定してください。")
}

// StopCurrentScript は現在実行中の UWSCR プロセスを強制終了します。
func (o *Orchestrator) StopCurrentScript() error {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.activeCmd == nil || o.activeCmd.Process == nil {
		return fmt.Errorf("実行中のスクリプトはありません。")
	}

	o.emitLog("[System] ユーザー指示によりスクリプトの実行を強制停止します...", false)
	err := o.activeCmd.Process.Kill()
	if err != nil {
		return fmt.Errorf("プロセスの強制終了に失敗しました: %v", err)
	}

	o.activeCmd = nil
	return nil
}

// RunScript は指定された UWS スクリプトをトランスパイルして非同期実行します。
func (o *Orchestrator) RunScript(scriptPath string) error {
	// 1. パスの特定
	uwscrPath, err := o.FindUWSCRPath()
	if err != nil {
		return err
	}

	// 2. スクリプトの読み込み
	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("スクリプトの読み込みに失敗しました: %v", err)
	}

	// 3. トランスパイル
	transpiler := NewTranspiler(o.app.cfg.Port)
	transpiled, err := transpiler.Transpile(string(scriptContent))
	if err != nil {
		return fmt.Errorf("トランスパイルエラー: %v", err)
	}

	// 4. 一時ファイル（temp_exec_YYYYMMDD_HHMMSS_xxx.uws）の生成
	exePath, _ := os.Executable()
	timestamp := time.Now().Format("20060102_150405_000")
	tempPath := filepath.Join(filepath.Dir(exePath), fmt.Sprintf("temp_exec_%s.uws", timestamp))
	
	err = os.WriteFile(tempPath, []byte(transpiled), 0644)
	if err != nil {
		return fmt.Errorf("一時スクリプトの作成に失敗しました: %v", err)
	}

	// 5. 非同期実行 (Goroutine)
	go func() {
		defer os.Remove(tempPath) // 終了後に一時ファイルを削除

		cmd := exec.Command(uwscrPath, tempPath)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000,
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			o.emitLog(fmt.Sprintf("[Error] Stdout pipe failed: %v", err), true)
			return
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			o.emitLog(fmt.Sprintf("[Error] Stderr pipe failed: %v", err), true)
			return
		}

		o.mu.Lock()
		o.activeCmd = cmd
		o.mu.Unlock()

		o.emitLog("[System] UWSCRプロセスを起動中...", false)
		if err := cmd.Start(); err != nil {
			o.emitLog(fmt.Sprintf("[Error] プロセスの起動に失敗しました: %v", err), true)
			o.mu.Lock()
			if o.activeCmd == cmd {
				o.activeCmd = nil
			}
			o.mu.Unlock()
			return
		}

		o.emitLog(fmt.Sprintf("[System] 起動成功 (PID: %d)", cmd.Process.Pid), false)

		var wg sync.WaitGroup
		wg.Add(2)

		// stdoutの読み込み
		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				o.emitLog(scanner.Text(), false)
			}
		}()

		// stderr of reading
		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				o.emitLog(scanner.Text(), true)
			}
		}()

		wg.Wait()
		// 実行終了を待つ
		err = cmd.Wait()

		o.mu.Lock()
		if o.activeCmd == cmd {
			o.activeCmd = nil
		}
		o.mu.Unlock()

		if err != nil {
			o.emitLog(fmt.Sprintf("[System] プロセスがエラーで終了しました: %v", err), true)
		} else {
			o.emitLog("[System] プロセスが正常に終了しました。", false)
		}
	}()

	return nil
}

// RunScriptSync はスクリプトを指定時間タイムアウト付きで同期実行し、実行結果のログと終了コードを返します。
func (o *Orchestrator) RunScriptSync(scriptPath string, timeoutSec int) (string, bool, error) {
	uwscrPath, err := o.FindUWSCRPath()
	if err != nil {
		return "", false, err
	}

	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", false, fmt.Errorf("スクリプトの読み込みに失敗しました: %v", err)
	}

	transpiler := NewTranspiler(o.app.cfg.Port)
	transpiled, err := transpiler.Transpile(string(scriptContent))
	if err != nil {
		return "", false, fmt.Errorf("トランスパイルエラー: %v", err)
	}

	exePath, _ := os.Executable()
	timestamp := time.Now().Format("20060102_150405_000")
	tempPath := filepath.Join(filepath.Dir(exePath), fmt.Sprintf("temp_test_%s.uws", timestamp))

	err = os.WriteFile(tempPath, []byte(transpiled), 0644)
	if err != nil {
		return "", false, fmt.Errorf("一時テストスクリプトの作成に失敗しました: %v", err)
	}
	defer os.Remove(tempPath)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, uwscrPath, tempPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", false, err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", false, err
	}

	o.mu.Lock()
	o.activeCmd = cmd
	o.mu.Unlock()
	defer func() {
		o.mu.Lock()
		if o.activeCmd == cmd {
			o.activeCmd = nil
		}
		o.mu.Unlock()
	}()

	o.emitLog("[System] テスト実行プロセスを起動中...", false)
	if err := cmd.Start(); err != nil {
		o.emitLog(fmt.Sprintf("[Error] テスト起動に失敗: %v", err), true)
		return "", false, err
	}

	var logBuf bytes.Buffer
	var mu sync.Mutex

	// ログ書き込み用のヘルパー
	writeLog := func(text string, isError bool) {
		mu.Lock()
		logBuf.WriteString(text + "\n")
		mu.Unlock()
		o.emitLog(text, isError)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			writeLog(scanner.Text(), false)
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			writeLog(scanner.Text(), true)
		}
	}()

	wg.Wait()
	err = cmd.Wait()

	success := true
	if err != nil {
		success = false
		if ctx.Err() == context.DeadlineExceeded {
			writeLog("[System] テスト実行がタイムアウト（規定秒超過）しました。", true)
			return logBuf.String() + "\n[System] タイムアウトにより強制終了されました。", false, nil
		}
		writeLog(fmt.Sprintf("[System] テスト実行がエラーで終了しました: %v", err), true)
	} else {
		writeLog("[System] テスト実行が正常に終了しました。", false)
	}

	return logBuf.String(), success, nil
}

func (o *Orchestrator) emitLog(message string, isError bool) {
	type LogLine struct {
		Message string `json:"message"`
		IsError bool   `json:"is_error"`
	}
	// ログメッセージの先頭にタイムスタンプを付加
	timestamp := time.Now().Format("15:04:05.000")
	formattedMsg := fmt.Sprintf("[%s] %s", timestamp, message)

	runtime.EventsEmit(o.app.ctx, "uwscr_log", LogLine{
		Message: formattedMsg,
		IsError: isError,
	})
}

