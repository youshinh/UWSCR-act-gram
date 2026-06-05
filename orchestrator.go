package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Orchestrator はUWSCR子プロセスの実行を管理します。
type Orchestrator struct {
	app *App
}

func NewOrchestrator(app *App) *Orchestrator {
	return &Orchestrator{app: app}
}

// FindUWSCRPath は優先度に従って uwscr.exe のパスを探索します。
func (o *Orchestrator) FindUWSCRPath() (string, error) {
	// 優先度 1: 実行中の act-gram.exe と同じディレクトリ
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

	return "", fmt.Errorf("uwscr.exe が見つかりません。アプリの同階層に配置するか、設定からパスを指定してください。")
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

	// 4. 一時ファイル（temp_exec.uws）の生成
	// 実行中ディレクトリに配置
	exePath, _ := os.Executable()
	tempPath := filepath.Join(filepath.Dir(exePath), "temp_exec.uws")
	
	err = os.WriteFile(tempPath, []byte(transpiled), 0644)
	if err != nil {
		return fmt.Errorf("一時スクリプトの作成に失敗しました: %v", err)
	}

	// 5. 非同期実行 (Goroutine)
	go func() {
		defer os.Remove(tempPath) // 終了後に一時ファイルを削除

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

		o.emitLog("[System] UWSCRプロセスを起動中...", false)
		if err := cmd.Start(); err != nil {
			o.emitLog(fmt.Sprintf("[Error] プロセスの起動に失敗しました: %v", err), true)
			return
		}

		o.emitLog(fmt.Sprintf("[System] 起動成功 (PID: %d)", cmd.Process.Pid), false)



		// stdoutの読み込み
		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				o.emitLog(scanner.Text(), false)
			}
		}()

		// stderrの読み込み
		go func() {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				o.emitLog(scanner.Text(), true)
			}
		}()

		// 実行終了を待つ
		err = cmd.Wait()
		if err != nil {
			o.emitLog(fmt.Sprintf("[System] プロセスがエラーで終了しました: %v", err), true)
		} else {
			o.emitLog("[System] プロセスが正常に終了しました。", false)
		}
	}()

	return nil
}

func (o *Orchestrator) emitLog(message string, isError bool) {
	type LogLine struct {
		Message string `json:"message"`
		IsError bool   `json:"is_error"`
	}
	runtime.EventsEmit(o.app.ctx, "uwscr_log", LogLine{
		Message: message,
		IsError: isError,
	})
}
