package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	exePath, err := os.Executable()
	if err == nil {
		localPath := filepath.Join(filepath.Dir(exePath), "uwscr.exe")
		if _, err := os.Stat(localPath); err == nil {
			return localPath, nil
		}
	}

	if o.app.cfg != nil && o.app.cfg.UWSCRPath != "" {
		if _, err := os.Stat(o.app.cfg.UWSCRPath); err == nil {
			return o.app.cfg.UWSCRPath, nil
		}
	}

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
		// プロセスがない場合でも正常終了合図を送り、UIのロックを解除させます
		o.emitLog("[System] プロセスが正常に終了しました。", false)
		if o.app != nil {
			o.app.SetMiniMode(false, "play")
			runtime.EventsEmit(o.app.ctx, "script_finished", true)
		}
		return nil
	}

	o.emitLog("[System] ユーザー指示によりスクリプトの実行を強制停止します...", false)
	err := o.activeCmd.Process.Kill()
	if err != nil {
		return fmt.Errorf("プロセスの強制終了に失敗しました: %v", err)
	}

	o.activeCmd = nil
	o.emitLog("[System] プロセスが正常に終了しました。", false)
	if o.app != nil {
		runtime.EventsEmit(o.app.ctx, "script_finished", true)
	}
	return nil
}

// RunScript は指定された UWS スクリプトをトランスパイルして非同期実行します（通常再生）。
func (o *Orchestrator) RunScript(scriptPath string) error {
	uwscrPath, err := o.FindUWSCRPath()
	if err != nil {
		return err
	}

	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("スクリプトの読み込みに失敗しました: %v", err)
	}

	transpiler := NewTranspiler(o.app.cfg.Port)
	transpiled, err := transpiler.Transpile(string(scriptContent))
	if err != nil {
		return fmt.Errorf("トランスパイルエラー: %v", err)
	}

	exePath, _ := os.Executable()
	timestamp := time.Now().Format("20060102_150405_000")
	tempPath := filepath.Join(filepath.Dir(exePath), fmt.Sprintf("temp_exec_%s.uws", timestamp))
	
	err = os.WriteFile(tempPath, []byte(transpiled), 0644)
	if err != nil {
		return fmt.Errorf("一時スクリプトの作成に失敗しました: %v", err)
	}

	isGuide := strings.Contains(scriptPath, "interactive_guide.uws")

	go func() {
		defer os.Remove(tempPath)

		cmd := exec.Command(uwscrPath, tempPath)

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
		if !isGuide {
			o.app.SetMiniMode(true, "play")
		}

		var wg sync.WaitGroup
		wg.Add(2)

		// 💡 app.goの ConvertToUTF8IfNeeded を使って、行ごとに安全にデコードします
		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				text := ConvertToUTF8IfNeeded(scanner.Bytes())
				o.emitLog(text, false)
			}
		}()

		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				text := ConvertToUTF8IfNeeded(scanner.Bytes())
				o.emitLog(text, true)
			}
		}()

		wg.Wait()
		err = cmd.Wait()

		o.mu.Lock()
		if o.activeCmd == cmd {
			o.activeCmd = nil
		}
		o.mu.Unlock()

		if !isGuide {
			o.app.SetMiniMode(false, "play")
		}

		if err != nil {
			o.emitLog(fmt.Sprintf("[System] プロセスがエラーで終了しました: %v", err), true)
		} else {
			o.emitLog("[System] プロセスが正常に終了しました。", false)
		}
		
		if o.app != nil {
			runtime.EventsEmit(o.app.ctx, "script_finished", true)
		}
	}()

	return nil
}

// RunScriptSync はスクリプトを指定時間タイムアウト付きで同期実行します（テスト実行）。
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
		return "", false, fmt.Errorf("Stdout pipe failed: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", false, fmt.Errorf("Stderr pipe failed: %v", err)
	}

	o.mu.Lock()
	o.activeCmd = cmd
	o.mu.Unlock()

	var logBuf strings.Builder
	var logMu sync.Mutex

	writeLog := func(msg string, isErr bool) {
		logMu.Lock()
		logBuf.WriteString(msg + "\n")
		logMu.Unlock()
		o.emitLog(msg, isErr)
	}

	writeLog("[System] テスト実行プロセスを起動中...", false)
	if err := cmd.Start(); err != nil {
		writeLog(fmt.Errorf("[Error] プロセスの起動に失敗しました: %v", err).Error(), true)
		o.mu.Lock()
		if o.activeCmd == cmd {
			o.activeCmd = nil
		}
		o.mu.Unlock()
		return logBuf.String(), false, err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			writeLog(ConvertToUTF8IfNeeded(scanner.Bytes()), false)
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			writeLog(ConvertToUTF8IfNeeded(scanner.Bytes()), true)
		}
	}()

	wg.Wait()
	err = cmd.Wait()

	o.mu.Lock()
	if o.activeCmd == cmd {
		o.activeCmd = nil
	}
	o.mu.Unlock()

	success := true
	if err != nil {
		success = false
		if ctx.Err() == context.DeadlineExceeded {
			writeLog("[System] テスト実行がタイムアウト（規定秒超過）しました。", true)
			return logBuf.String() + "\n[System] タイムアウトにより強制終了されました。", false, nil
		}
		writeLog(fmt.Sprintf("[System] テスト実行がエラーで終了しました: %v", err), true)
	} else {
		logStr := logBuf.String()
		if strings.Contains(logStr, "構文エラー") || strings.Contains(logStr, "ありません") || strings.Contains(logStr, "Error:") || strings.Contains(logStr, "エラー") {
			success = false
			writeLog("[System] 警告: ログ内に構文エラーまたは未定義定数を検出したため、テスト失敗と判定します。", true)
		} else {
			writeLog("[System] テスト実行が正常に終了しました。", false)
		}
	}

	if !success {
		_ = o.app.SaveErrorReflection(string(scriptContent), logBuf.String())
	}

	return logBuf.String(), success, nil
}

// 💡 修正の要諦：JSクラッシュ防止のため、必ずタイムスタンプを付与し、LogLineオブジェクトとして送信します。
func (o *Orchestrator) emitLog(message string, isError bool) {
	if o.app == nil || o.app.ctx == nil {
		return
	}

	type LogLine struct {
		Message string `json:"message"`
		IsError bool   `json:"is_error"`
	}

	// フロントエンドのパース処理を壊さないよう、必ず先頭に [HH:MM:SS.mmm] を付与します
	timestamp := time.Now().Format("15:04:05.000")
	formattedMsg := fmt.Sprintf("[%s] %s", timestamp, message)

	runtime.EventsEmit(o.app.ctx, "uwscr_log", LogLine{
		Message: formattedMsg,
		IsError: isError,
	})
}