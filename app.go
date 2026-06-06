package main

import (
	"act_gram/capture"
	"act_gram/knowledge"
	"act_gram/llm"
	"act_gram/manual"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx            context.Context
	cfg            *Config
	localServer    *LocalServer
	orchestrator   *Orchestrator
	mu             sync.Mutex
	rag            *knowledge.RAGManager
	slicer         *manual.AutoSlicer
	session        *manual.ManualSession
	refactorEngine *llm.RefactorEngine
	recorder       *capture.Recorder
}

// NewApp creates a new App application struct
func NewApp() *App {
	log.Println("[App] Initializing actgram application structure...")
	a := &App{}
	a.localServer = NewLocalServer(a)
	a.orchestrator = NewOrchestrator(a)
	a.slicer = manual.NewAutoSlicer()
	a.session = manual.NewManualSession()
	a.refactorEngine = llm.NewRefactorEngine()
	return a
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	log.Println("[App.startup] App starting up...")
	a.ctx = ctx
	cfg, err := LoadConfig()
	if err != nil {
		log.Printf("[App.startup] Failed to load config: %v", err)
	}
	a.cfg = cfg

	// RAGマネージャーの初期化（exe配信時を考慮したパス解決）
	knowledgeDir := filepath.Join(a.getExecBaseDir(), "knowledge")
	log.Printf("[App.startup] Initializing RAG manager with knowledge folder: %s", knowledgeDir)
	a.rag = knowledge.NewRAGManager(knowledgeDir)
	if err := a.rag.LoadKnowledgeFiles(); err != nil {
		log.Printf("[App.startup] Failed to load knowledge files: %v", err)
	}

	// ローカルAPIサーバーをGoroutineで非同期起動
	go func() {
		log.Printf("[App.startup] Launching Local API Server on port %d...", a.cfg.Port)
		err := a.localServer.Start(a.cfg.Port)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("[App.startup] Local API Server error: %v", err)
		}
	}()
}

// shutdown is called when the app is closing.
func (a *App) shutdown(ctx context.Context) {
	log.Println("[App.shutdown] App shutting down...")
	if a.localServer != nil {
		log.Println("[App.shutdown] Shutting down Local API Server...")
		a.localServer.Shutdown()
	}
}

// GetConfig は現在の設定をフロントエンドに返します
func (a *App) GetConfig() (*Config, error) {
	log.Println("[App.GetConfig] Fetching configuration request...")
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cfg == nil {
		cfg, err := LoadConfig()
		if err != nil {
			log.Printf("[App.GetConfig] Failed to load config: %v", err)
			return nil, err
		}
		a.cfg = cfg
	}
	log.Printf("[App.GetConfig] Config loaded: %+v", a.cfg)
	return a.cfg, nil
}

// SaveConfig は指定されたレイヤーの設定を保存します
func (a *App) SaveConfig(layer string, provider string, model string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Printf("[App.SaveConfig] Request received. Layer=%s, Provider=%s, Model=%s", layer, provider, model)

	if a.cfg == nil {
		var err error
		a.cfg, err = LoadConfig()
		if err != nil {
			log.Printf("[App.SaveConfig] Failed to load config: %v", err)
			return err
		}
	}

	if a.cfg.Layers == nil {
		a.cfg.Layers = make(map[string]LayerConfig)
	}

	a.cfg.Layers[layer] = LayerConfig{
		Provider: provider,
		Model:    model,
	}

	err := SaveConfig(a.cfg)
	if err != nil {
		log.Printf("[App.SaveConfig] Failed to save config to file: %v", err)
	} else {
		log.Printf("[App.SaveConfig] Config successfully written for %s layer", layer)
	}
	return err
}

// SaveConfigs は複数のレイヤー設定を一括で保存します (並行書き込みを避けるため)
func (a *App) SaveConfigs(layersJSON string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Printf("[App.SaveConfigs] Bulk saving configuration changes...")

	var newLayers map[string]LayerConfig
	if err := json.Unmarshal([]byte(layersJSON), &newLayers); err != nil {
		log.Printf("[App.SaveConfigs] JSON unmarshal failed: %v", err)
		return fmt.Errorf("JSONのパースに失敗しました: %v", err)
	}

	if a.cfg == nil {
		var err error
		a.cfg, err = LoadConfig()
		if err != nil {
			log.Printf("[App.SaveConfigs] Failed to load config: %v", err)
			return err
		}
	}

	if a.cfg.Layers == nil {
		a.cfg.Layers = make(map[string]LayerConfig)
	}

	for layer, lcfg := range newLayers {
		a.cfg.Layers[layer] = lcfg
		log.Printf("[App.SaveConfigs] Updating Layer=%s: Provider=%s, Model=%s", layer, lcfg.Provider, lcfg.Model)
	}

	err := SaveConfig(a.cfg)
	if err != nil {
		log.Printf("[App.SaveConfigs] Failed to write config file: %v", err)
	} else {
		log.Printf("[App.SaveConfigs] Bulk configurations successfully saved")
	}
	return err
}

// SaveAPIKey はAPIキーをセキュア領域に保存します
func (a *App) SaveAPIKey(provider string, key string) error {
	log.Printf("[App.SaveAPIKey] Registering API key for provider: %s", provider)
	err := SaveAPIKey(provider, key)
	if err != nil {
		log.Printf("[App.SaveAPIKey] Failed to save credential: %v", err)
	} else {
		log.Printf("[App.SaveAPIKey] API key successfully registered in Credential Manager for %s", provider)
	}
	return err
}

// HasAPIKey は指定されたプロバイダーのAPIキーが登録されているかを判定します
func (a *App) HasAPIKey(provider string) bool {
	log.Printf("[App.HasAPIKey] Checking credential existence for provider: %s", provider)
	key, err := GetAPIKey(provider)
	if err != nil || key == "" {
		log.Printf("[App.HasAPIKey] Credential NOT found for provider: %s", provider)
		return false
	}
	log.Printf("[App.HasAPIKey] Credential exists for provider: %s", provider)
	return true
}

// FetchModels は指定されたプロバイダーの利用可能モデルを動的に取得して返します
func (a *App) FetchModels(provider string) ([]string, error) {
	log.Printf("[App.FetchModels] Fetching models list for provider: %s", provider)
	switch provider {
	case "google":
		key, err := GetAPIKey("google")
		if err != nil {
			log.Printf("[App.FetchModels] API key retrieval error: %v", err)
			return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
		}
		if key == "" {
			log.Println("[App.FetchModels] API key not registered, returning default Gemini model fallbacks")
			return []string{"gemini-flash-lite-latest", "gemini-2.5-flash", "gemini-2.5-flash-lite", "gemini-1.5-pro"}, nil
		}
		providerImpl := llm.NewGeminiProvider(key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			log.Printf("[App.FetchModels] API models request failed: %v. Falling back to default list.", err)
			return []string{"gemini-flash-lite-latest", "gemini-2.5-flash", "gemini-2.5-flash-lite", "gemini-1.5-pro"}, nil
		}
		log.Printf("[App.FetchModels] Found %d models for Google provider", len(models))
		return models, nil

	case "anthropic":
		key, err := GetAPIKey("anthropic")
		if err != nil {
			log.Printf("[App.FetchModels] API key retrieval error: %v", err)
			return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
		}
		providerImpl := llm.NewAnthropicProvider(key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			log.Printf("[App.FetchModels] API models request failed: %v", err)
			return nil, err
		}
		log.Printf("[App.FetchModels] Found %d models for Anthropic provider", len(models))
		return models, nil

	case "ollama":
		providerImpl := llm.NewOllamaProvider()
		models, err := providerImpl.GetAvailableModels()
		if err != nil || len(models) == 0 {
			log.Printf("[App.FetchModels] Ollama models query failed (is it running?): %v. Returning defaults.", err)
			return []string{"qwen2.5-coder:latest", "llama3.2-vision:latest", "gemma2:latest"}, nil
		}
		log.Printf("[App.FetchModels] Found %d models for Ollama provider", len(models))
		return models, nil

	default:
		log.Printf("[App.FetchModels] Unknown provider: %s", provider)
		return []string{}, nil
	}
}

// RunScript は指定されたパスのスクリプトをトランスパイルして非同期実行します
func (a *App) RunScript(scriptPath string) error {
	log.Printf("[App.RunScript] Preparing to run UWSCR script: %s", scriptPath)
	if a.orchestrator == nil {
		log.Println("[App.RunScript] Error: orchestrator is uninitialized")
		return fmt.Errorf("orchestrator is not initialized")
	}
	return a.orchestrator.RunScript(scriptPath)
}

// SaveUWSCRPath は明示的な UWSCRPath を保存します
func (a *App) SaveUWSCRPath(path string) error {
	log.Printf("[App.SaveUWSCRPath] Saving custom UWSCR executable path: %s", path)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cfg == nil {
		var err error
		a.cfg, err = LoadConfig()
		if err != nil {
			log.Printf("[App.SaveUWSCRPath] Failed to load config: %v", err)
			return err
		}
	}
	a.cfg.UWSCRPath = path
	err := SaveConfig(a.cfg)
	if err != nil {
		log.Printf("[App.SaveUWSCRPath] Failed to save config to file: %v", err)
	} else {
		log.Println("[App.SaveUWSCRPath] Custom UWSCR path successfully saved")
	}
	return err
}

// GetDefaultManualPath はマニュアル出力先の初期デフォルト絶対パス（カレントディレクトリ\manual）を返します
func (a *App) GetDefaultManualPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("[App.GetDefaultManualPath] Failed to get working directory: %v", err)
		return "", err
	}
	defaultPath := filepath.Join(wd, "manual")
	log.Printf("[App.GetDefaultManualPath] Default manual path resolved: %s", defaultPath)
	return defaultPath, nil
}

// GenerateInteractiveManual はフロントエンドから渡されたステップ情報をもとにマニュアルを生成します
func (a *App) GenerateInteractiveManual(outputPath string, stepsJSON string, useHighQualityTTS bool) error {
	log.Printf("[App.GenerateInteractiveManual] Generating manual. Output=%s, TTS=%v", outputPath, useHighQualityTTS)
	var steps []ManualStep
	if err := json.Unmarshal([]byte(stepsJSON), &steps); err != nil {
		log.Printf("[App.GenerateInteractiveManual] JSON parsing failed: %v", err)
		return fmt.Errorf("JSONのパースに失敗しました: %v", err)
	}
	log.Printf("[App.GenerateInteractiveManual] Processing %d steps...", len(steps))
	err := GenerateManual(outputPath, steps, useHighQualityTTS)
	if err != nil {
		log.Printf("[App.GenerateInteractiveManual] Error generating manual: %v", err)
	} else {
		log.Printf("[App.GenerateInteractiveManual] HTML manual package created successfully at %s", outputPath)
	}
	return err
}

// SessionContext は操作情報の取得結果を表します
type SessionContext struct {
	ActiveTitle      string `json:"active_title"`
	ScreenshotPath   string `json:"screenshot_path"`
	ScreenshotBase64 string `json:"screenshot_base64"`
}

// CaptureSession は現在のアクティブウィンドウのタイトルと、デスクトップのスクリーンショットを取得します
func (a *App) CaptureSession() (*SessionContext, error) {
	log.Println("[App.CaptureSession] Querying foreground window and capturing screen...")

	// 1. アクティブウィンドウタイトルの取得
	activeTitle := GetActiveWindowTitle()
	log.Printf("[App.CaptureSession] Active window title: %s", activeTitle)

	// 2. 一時フォルダにスクリーンショットを保存
	tempDir := os.TempDir()
	screenshotFile := filepath.Join(tempDir, fmt.Sprintf("act_gram_capture_%d.png", time.Now().UnixNano()))

	err := CaptureScreen(screenshotFile)
	if err != nil {
		log.Printf("[App.CaptureSession] Failed to capture screen: %v", err)
		return nil, fmt.Errorf("画面のキャプチャに失敗しました: %v", err)
	}
	log.Printf("[App.CaptureSession] Screenshot captured successfully: %s", screenshotFile)

	// 3. プレビュー表示用に画像をBase64エンコード
	data, err := os.ReadFile(screenshotFile)
	if err != nil {
		log.Printf("[App.CaptureSession] Failed to read screenshot file: %v", err)
		return nil, fmt.Errorf("キャプチャ画像の読み込みに失敗しました: %v", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(data)

	return &SessionContext{
		ActiveTitle:      activeTitle,
		ScreenshotPath:   screenshotFile,
		ScreenshotBase64: base64Str,
	}, nil
}

// GenerateScript はプロンプトとセッション情報を元に UWSCR スクリプトを生成します
func (a *App) GenerateScript(prompt string, sessionContextJSON string) (string, error) {
	log.Printf("[App.GenerateScript] Generating script. Prompt: %s", prompt)
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cfg == nil {
		return "", fmt.Errorf("設定がロードされていません")
	}

	// 1. Brain層（コード生成）のLLMプロバイダーとモデルを取得
	brainConfig, ok := a.cfg.Layers["brain"]
	if !ok || brainConfig.Provider == "" || brainConfig.Model == "" {
		// デフォルトフォールバック
		brainConfig = LayerConfig{
			Provider: "google",
			Model:    "gemini-flash-lite-latest",
		}
	}

	provider := brainConfig.Provider
	model := brainConfig.Model
	log.Printf("[App.GenerateScript] Selected LLM: %s (%s)", model, provider)

	apiKey, err := GetAPIKey(provider)
	if err != nil && provider != "ollama" {
		return "", fmt.Errorf("APIキーの取得に失敗しました: %v", err)
	}

	var llmProvider llm.LLMProvider
	switch provider {
	case "google":
		llmProvider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		llmProvider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		llmProvider = llm.NewOllamaProvider()
	default:
		return "", fmt.Errorf("未サポートのプロバイダーです: %s", provider)
	}

	// 2. セッション情報の解析（画像パス of 抽出）
	screenshotPath := ""
	activeTitleInfo := ""
	if sessionContextJSON != "" {
		var session SessionContext
		if err := json.Unmarshal([]byte(sessionContextJSON), &session); err == nil {
			screenshotPath = session.ScreenshotPath
			if session.ActiveTitle != "" {
				activeTitleInfo = fmt.Sprintf("現在ユーザーが開いているアクティブウィンドウ: %s\n", session.ActiveTitle)
			}
		}
	}

	// 3. UWSCR生成用のシステムプロンプト定義
	systemPrompt := `あなたは卓越したRPAスクリプト開発AIです。
Windows用の自動化スクリプトエンジンである「UWSCR (UWSC互換)」で動作するスクリプト (.uws) を作成してください。

【UWSCR の基本構文ルール】
1. コメント: // コメントテキスト
2. 変数宣言: Dim 変数名 = 初期値 (または Dim 変数名)
3. ウィンドウ制御:
   - ACW(GETID("ウィンドウタイトル", "クラス名"), x, y, w, h) : ウィンドウのアクティブ化・位置調整
   - CTRL_WIN(id, CLOSE) : ウィンドウを閉じる
4. マウス・キーボード操作:
   - BTN(LEFT, CLICK, x, y, delay) : 指定座標をクリック
   - KBD(VK_A, CLICK, delay) : キー入力 (VK_RETURN, VK_TAB, VK_BACK などの仮想キー定数を使用)
   - SENDSTR(id, "文字列") : 指定ウィンドウに文字列を直接送信
5. 制御構文:
   - IFB 条件 THEN ... ELSE ... ENDIF (※UWSCの厳格なブロック構文 IFB 〜 THEN 〜 ENDIF を使用してください)
   - FOR 変数 = 初期値 TO 最大値 ... NEXT
   - WHILE 条件 ... WEND
   - SLEEP(秒数) : 例: SLEEP(1) は1秒待機
6. ダイアログ表示:
   - MSGBOX("メッセージ内容") : ポップアップダイアログを表示
7. タイムスタンプ・ログ出力:
   - UWSCRのテスト実行時や実際の運用時に進捗状況が把握できるよう、各処理ステップの前後で PRINT "[タイムスタンプ] メッセージ" のようにタイムスタンプ付きのログを出力してください。UWSCRの経過時間や日時取得には GETTIME() を使うか、適切な文字列連結を使用してください。
     例:
     PRINT "[" + GETTIME() + "] 処理を開始します..."

【actgram 独自マクロ拡張関数】
- AI_EVAL("判定・取得したい内容", 画像取得関数等)
  この関数は、画面上の特定情報やテキストを判断させたい場合に使用できます。戻り値は推論結果テキストです。
  例: Dim price = AI_EVAL("この伝票の合計金額は数値でいくら？", GetScreenCapture())
  ※引数の画像取得関数には GetScreenCapture() などのダミー表記が用いられ、actgram側で実行時に実際のデスクトップキャプチャに差し替わります。

【出力要件】
- スクリプトコードのみを出力してください。
- markdown のコードブロック ('''uws ... ''') などで囲まず、純粋なプレーンテキストとして UWSCR スクリプトのみを返してください。説明文や前置きは一切不要です。
`

	finalPrompt := fmt.Sprintf("%s\n%s\n【指示】\n%s", systemPrompt, activeTitleInfo, prompt)

	// 4. LLMへの問い合わせ
	resp := llmProvider.GenerateText(finalPrompt, screenshotPath, model)
	if resp.Error != nil {
		log.Printf("[App.GenerateScript] Generation failed: %v", resp.Error)
		return "", resp.Error
	}

	// markdownブロック等が含まれていた場合のクレンジング
	cleanCode := cleanGeneratedCode(resp.Text)
	log.Printf("[App.GenerateScript] Script generated successfully (%d bytes)", len(cleanCode))
	return cleanCode, nil
}

// TestRunResult はテスト実行の結果をフロントエンドに返します。
type TestRunResult struct {
	Logs    string `json:"logs"`
	Success bool   `json:"success"`
}

// TestRunScript は指定された UWSCR コードを一時ファイルに保存して同期的にテスト実行し、ログと結果を返します
func (a *App) TestRunScript(code string) (*TestRunResult, error) {
	log.Println("[App.TestRunScript] Starting script test run")
	
	tempDir := os.TempDir()
	tempFile := filepath.Join(tempDir, fmt.Sprintf("act_gram_test_%d.uws", time.Now().UnixNano()))
	
	if err := os.WriteFile(tempFile, []byte(code), 0644); err != nil {
		log.Printf("[App.TestRunScript] Failed to write temporary script file: %v", err)
		return nil, fmt.Errorf("テスト用スクリプトの作成に失敗しました: %v", err)
	}
	defer os.Remove(tempFile)

	if a.orchestrator == nil {
		return nil, fmt.Errorf("orchestrator is not initialized")
	}

	// タイムアウトは15秒とする
	logBuf, success, err := a.orchestrator.RunScriptSync(tempFile, 15)
	if err != nil {
		log.Printf("[App.TestRunScript] Execution failed: %v", err)
		return &TestRunResult{Logs: logBuf, Success: false}, err
	}

	log.Printf("[App.TestRunScript] Execution finished. Success=%v", success)
	return &TestRunResult{
		Logs:    logBuf,
		Success: success,
	}, nil
}

// StopScript は現在実行中の UWSCR スクリプトを強制停止します
func (a *App) StopScript() error {
	log.Println("[App.StopScript] Request received to stop execution")
	if a.orchestrator == nil {
		return fmt.Errorf("orchestrator is not initialized")
	}
	return a.orchestrator.StopCurrentScript()
}

// CorrectScript は実行時エラーログと元のコードをもとに、AIにコードの修正指示を送り、修復されたUWSCRコードを返します
func (a *App) CorrectScript(prompt string, code string, errorLog string) (string, error) {
	log.Printf("[App.CorrectScript] Correcting script due to error. Prompt: %s", prompt)
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cfg == nil {
		return "", fmt.Errorf("設定がロードされていません")
	}

	brainConfig, ok := a.cfg.Layers["brain"]
	if !ok || brainConfig.Provider == "" || brainConfig.Model == "" {
		brainConfig = LayerConfig{
			Provider: "google",
			Model:    "gemini-flash-lite-latest",
		}
	}

	provider := brainConfig.Provider
	model := brainConfig.Model
	log.Printf("[App.CorrectScript] Selected LLM for correction: %s (%s)", model, provider)

	apiKey, err := GetAPIKey(provider)
	if err != nil && provider != "ollama" {
		return "", fmt.Errorf("APIキーの取得に失敗しました: %v", err)
	}

	var llmProvider llm.LLMProvider
	switch provider {
	case "google":
		llmProvider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		llmProvider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		llmProvider = llm.NewOllamaProvider()
	default:
		return "", fmt.Errorf("未サポートのプロバイダーです: %s", provider)
	}

	systemPrompt := `あなたは卓越したRPAスクリプト開発AIです。
UWSCR スクリプト (.uws) のテスト実行でエラーが発生しました。
元のコードとエラーログ、および当初の指示内容を分析し、エラーを修正した動作可能な UWSCR スクリプトを再生成してください。

【UWSCR の基本構文ルール】
1. コメント: // コメントテキスト
2. 変数宣言: Dim 変数名 = 初期値
3. 制御構文: IFB 条件 THEN ... ENDIF (UWSCの厳格な構文を使用してください。IFB 〜 THEN と ENDIF は必須です)
4. ループ: FOR/WHILE
5. タイムスタンプ機能: UWSCRでデバッグや進捗確認を行うため、要所（特に開始、終了、および条件判定時）に PRINT "[タイムスタンプ] メッセージ" のようにタイムスタンプ付きのログを出力してください。UWSCRの経過時間取得には GETTIME() を使うか、適切な文字列連結を使用してください。
   例:
   PRINT "[" + GETTIME() + "] 処理を開始します..."

【出力要件】
- 修正されたスクリプトコードのみを出力してください。
- markdown のコードブロック ('''uws ... ''') などで囲まず、純粋なプレーンテキストとして UWSCR スクリプトのみを返してください。説明文や前置きは一切不要です。
`

	finalPrompt := fmt.Sprintf(
		"%s\n\n【当初の自動化指示】\n%s\n\n【エラーが発生した元のコード】\n%s\n\n【UWSCRの実行エラーログ】\n%s\n\n【指示】\nエラーログを分析し、バグを修正したUWSCRコードを返してください。可能な限りPRINT文で独自のタイムスタンプ付きログ（例: PRINT \"[タイムスタンプ] ログ内容\"）を付与して進捗を出力してください。",
		systemPrompt, prompt, code, errorLog,
	)

	resp := llmProvider.GenerateText(finalPrompt, "", model)
	if resp.Error != nil {
		log.Printf("[App.CorrectScript] Generation failed: %v", resp.Error)
		return "", resp.Error
	}

	cleanCode := cleanGeneratedCode(resp.Text)
	log.Printf("[App.CorrectScript] Corrected script generated successfully (%d bytes)", len(cleanCode))
	return cleanCode, nil
}

// cleanGeneratedCode は生成されたコードから markdown のコードブロック記述などを除去します
func cleanGeneratedCode(code string) string {
	code = strings.TrimSpace(code)
	if strings.HasPrefix(code, "```") {
		lines := strings.Split(code, "\n")
		if len(lines) > 2 {
			startIdx := 1
			if strings.HasPrefix(lines[0], "```uws") || strings.HasPrefix(lines[0], "```") {
				startIdx = 1
			}
			endIdx := len(lines) - 1
			if strings.HasPrefix(lines[len(lines)-1], "```") {
				endIdx = len(lines) - 1
			}
			if endIdx > startIdx {
				code = strings.Join(lines[startIdx:endIdx], "\n")
			}
		}
	}
	return strings.TrimSpace(code)
}

// SaveScriptFile は指定されたパスに UWSCR スクリプトコードを保存します
func (a *App) SaveScriptFile(path string, code string) error {
	log.Printf("[App.SaveScriptFile] Saving script to: %s", path)
	if path == "" {
		return fmt.Errorf("保存先パスが指定されていません")
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("[App.SaveScriptFile] Failed to create directories: %v", err)
		return fmt.Errorf("保存先フォルダの作成に失敗しました: %v", err)
	}

	err := os.WriteFile(path, []byte(code), 0644)
	if err != nil {
		log.Printf("[App.SaveScriptFile] Failed to write file: %v", err)
		return fmt.Errorf("スクリプトファイルの保存に失敗しました: %v", err)
	}

	log.Println("[App.SaveScriptFile] Script successfully saved")
	return nil
}

// GetDefaultScriptPath はスクリプト保存先のデフォルト絶対パス（カレントディレクトリ\scripts\autoscript.uws）を返します
func (a *App) GetDefaultScriptPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("[App.GetDefaultScriptPath] Failed to get working directory: %v", err)
		return "", err
	}
	defaultPath := filepath.Join(wd, "scripts", "autoscript.uws")
	log.Printf("[App.GetDefaultScriptPath] Default script path resolved: %s", defaultPath)
	return defaultPath, nil
}

// SelectDirectory はOSネイティブのフォルダ選択ダイアログを開き、選択されたフォルダパスを返します
func (a *App) SelectDirectory(title string) (string, error) {
	log.Printf("[App.SelectDirectory] Opening directory dialog. Title: %s", title)
	selected, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
	})
	if err != nil {
		log.Printf("[App.SelectDirectory] Dialog error: %v", err)
		return "", err
	}
	log.Printf("[App.SelectDirectory] User selected: %s", selected)
	return selected, nil
}

// SelectFile はOSネイティブのファイル選択ダイアログを開き、選択されたファイルパスを返します
func (a *App) SelectFile(title string, displayName string, pattern string) (string, error) {
	log.Printf("[App.SelectFile] Opening file open dialog. Title: %s, Filter: %s (%s)", title, displayName, pattern)
	
	filters := []runtime.FileFilter{}
	if displayName != "" && pattern != "" {
		filters = append(filters, runtime.FileFilter{
			DisplayName: displayName,
			Pattern:     pattern,
		})
	}

	selected, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:   title,
		Filters: filters,
	})
	if err != nil {
		log.Printf("[App.SelectFile] Dialog error: %v", err)
		return "", err
	}
	log.Printf("[App.SelectFile] User selected: %s", selected)
	return selected, nil
}

// SelectSaveFile はOSネイティブのファイル保存ダイアログを開き、選択された保存先パスを返します
func (a *App) SelectSaveFile(title string, defaultName string, displayName string, pattern string) (string, error) {
	log.Printf("[App.SelectSaveFile] Opening file save dialog. Title: %s, Default: %s", title, defaultName)
	
	filters := []runtime.FileFilter{}
	if displayName != "" && pattern != "" {
		filters = append(filters, runtime.FileFilter{
			DisplayName: displayName,
			Pattern:     pattern,
		})
	}

	selected, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           title,
		DefaultFilename: defaultName,
		Filters:         filters,
	})
	if err != nil {
		log.Printf("[App.SelectSaveFile] Dialog error: %v", err)
		return "", err
	}
	log.Printf("[App.SelectSaveFile] User selected: %s", selected)
	return selected, nil
}

// getExecBaseDir は実行ファイルの位置（開発環境ならcwd）を基準としたベースディレクトリを解決します
func (a *App) getExecBaseDir() string {
	exePath, err := os.Executable()
	if err != nil {
		wd, _ := os.Getwd()
		return wd
	}
	dir := filepath.Dir(exePath)
	// Tempフォルダやwails開発ビルド内であれば、cwd（プロジェクトルート）を基準とする
	if strings.Contains(dir, "Temp") || strings.Contains(dir, "wails") || strings.Contains(dir, "Appdata") {
		wd, _ := os.Getwd()
		return wd
	}

	// フォールバック: exePathと同じ階層に manual/act_gram_guide がない場合
	if _, err := os.Stat(filepath.Join(dir, "manual", "act_gram_guide")); err != nil {
		// cwdを確認
		wd, _ := os.Getwd()
		if _, err := os.Stat(filepath.Join(wd, "manual", "act_gram_guide")); err == nil {
			return wd
		}
		// 親の親（build/bin から見たプロジェクトルート）を確認
		parentOfParent := filepath.Dir(filepath.Dir(dir))
		if _, err := os.Stat(filepath.Join(parentOfParent, "manual", "act_gram_guide")); err == nil {
			return parentOfParent
		}
	}
	return dir
}

// GenerateScenarioFromLog はレコーダー生ログからスライス・シナリオとコードを自律生成します
func (a *App) GenerateScenarioFromLog(logDir string) ([]manual.ManualStep, error) {
	log.Printf("[App.GenerateScenarioFromLog] Request received. logDir=%s", logDir)
	if a.cfg == nil {
		return nil, fmt.Errorf("設定がロードされていません")
	}

	// RAGコンテキストの取得
	ragContext := ""
	if a.rag != nil {
		ragContext = a.rag.SearchRelevantContext("UWSCR code logic generation manual scenario", 3)
	}

	// LLMプロバイダーの取得
	brainConfig, ok := a.cfg.Layers["brain"]
	if !ok || brainConfig.Provider == "" || brainConfig.Model == "" {
		brainConfig = LayerConfig{
			Provider: "google",
			Model:    "gemini-flash-lite-latest",
		}
	}
	providerName := brainConfig.Provider
	model := brainConfig.Model
	apiKey, err := GetAPIKey(providerName)
	if err != nil && providerName != "ollama" {
		return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
	}

	var llmProvider llm.LLMProvider
	switch providerName {
	case "google":
		llmProvider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		llmProvider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		llmProvider = llm.NewOllamaProvider()
	default:
		return nil, fmt.Errorf("未サポートのプロバイダーです: %s", providerName)
	}

	steps, err := a.slicer.SliceLogToScenario(logDir, ragContext, llmProvider, model)
	if err != nil {
		log.Printf("[App.GenerateScenarioFromLog] Error slicing log: %v", err)
		return nil, err
	}

	// Bake click markers into step images
	for i := range steps {
		if steps[i].ClickX > 0 && steps[i].ClickY > 0 && steps[i].ImagePath != "" {
			markedPath, err := a.DrawMarker(steps[i].ImagePath, steps[i].ClickX, steps[i].ClickY)
			if err == nil {
				steps[i].ImagePath = markedPath
			}
		}
	}

	a.session.Steps = steps
	return steps, nil
}

// ExecuteStep はマニュアル並走における特定のステップを同期実行します
func (a *App) ExecuteStep(stepIdx int) (*manual.ManualStep, error) {
	log.Printf("[App.ExecuteStep] Executing step index: %d", stepIdx)
	if a.orchestrator == nil {
		return nil, fmt.Errorf("orchestrator is not initialized")
	}

	runSyncHelper := func(code string) (string, bool, error) {
		tempDir := os.TempDir()
		tempFile := filepath.Join(tempDir, fmt.Sprintf("act_gram_step_exec_%d.uws", time.Now().UnixNano()))
		if err := os.WriteFile(tempFile, []byte(code), 0644); err != nil {
			return "", false, err
		}
		defer os.Remove(tempFile)

		return a.orchestrator.RunScriptSync(tempFile, 15)
	}

	step, err := a.session.ExecuteStep(stepIdx, runSyncHelper)
	if err != nil {
		log.Printf("[App.ExecuteStep] Execution failed: %v", err)
		return nil, err
	}
	return step, nil
}

// AskManualContext はマニュアルに紐づく質問に対してRAGと画面キャプチャをベースに推論回答します
func (a *App) AskManualContext(question, imagePath string) (string, error) {
	log.Printf("[App.AskManualContext] Question received. ImagePath=%s", imagePath)
	if a.cfg == nil {
		return "", fmt.Errorf("設定がロードされていません")
	}

	// 1. RAG知識検索
	ragContext := ""
	if a.rag != nil {
		ragContext = a.rag.SearchRelevantContext(question, 3)
	}

	// 2. LLMプロバイダーの取得
	brainConfig, ok := a.cfg.Layers["brain"]
	if !ok || brainConfig.Provider == "" || brainConfig.Model == "" {
		brainConfig = LayerConfig{
			Provider: "google",
			Model:    "gemini-flash-lite-latest",
		}
	}
	providerName := brainConfig.Provider
	model := brainConfig.Model
	apiKey, err := GetAPIKey(providerName)
	if err != nil && providerName != "ollama" {
		return "", fmt.Errorf("APIキーの取得に失敗しました: %v", err)
	}

	var llmProvider llm.LLMProvider
	switch providerName {
	case "google":
		llmProvider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		llmProvider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		llmProvider = llm.NewOllamaProvider()
	default:
		return "", fmt.Errorf("未サポートのプロバイダーです: %s", providerName)
	}

	systemPrompt := `あなたは優秀なRPAコパイロットAIエージェントです。
指示された画像、ユーザーからの質問、および以下の現場業務知識コンテキスト（RAG）をもとに、業務に即した的確な回答を行ってください。

【現場業務知識コンテキスト（RAG）】
` + ragContext + `

【回答ガイドライン】
- 専門用語や社内ルールをふまえ、現場担当者がわかりやすい日本語で簡潔に回答してください。
- 必要に応じて、現在の画面（画像）の状態が良いか悪いかを画像認識した上で判定してください。
- 最後の回答は日本語を返してください。
`

	finalPrompt := fmt.Sprintf("%s\n\n【ユーザーの質問】\n%s", systemPrompt, question)
	resp := llmProvider.GenerateText(finalPrompt, imagePath, model)
	if resp.Error != nil {
		return "", resp.Error
	}
	return resp.Text, nil
}

// ProposeOptimization は実行ログ（ミリ秒）からコード最適化を提案します
func (a *App) ProposeOptimization(scriptPath string, eventsJSON string) (string, error) {
	log.Println("[App.ProposeOptimization] Processing optimization proposal request")
	if a.cfg == nil {
		return "", fmt.Errorf("設定がロードされていません")
	}

	var code string
	if _, err := os.Stat(scriptPath); err == nil {
		data, err := os.ReadFile(scriptPath)
		if err != nil {
			return "", fmt.Errorf("スクリプトファイルの読み込みに失敗しました: %v", err)
		}
		code = string(data)
	} else {
		// Fallback: If it's not a file path, treat as raw code
		code = scriptPath
	}

	var events []llm.ExecutionEvent
	if err := json.Unmarshal([]byte(eventsJSON), &events); err != nil {
		return "", fmt.Errorf("ファクトログJSONのパースに失敗しました: %v", err)
	}

	ragContext := ""
	if a.rag != nil {
		ragContext = a.rag.SearchRelevantContext("UWSCR code optimization performance concurrency", 3)
	}

	brainConfig, ok := a.cfg.Layers["brain"]
	if !ok || brainConfig.Provider == "" || brainConfig.Model == "" {
		brainConfig = LayerConfig{
			Provider: "google",
			Model:    "gemini-flash-lite-latest",
		}
	}
	providerName := brainConfig.Provider
	model := brainConfig.Model
	apiKey, err := GetAPIKey(providerName)
	if err != nil && providerName != "ollama" {
		return "", fmt.Errorf("APIキーの取得に失敗しました: %v", err)
	}

	var llmProvider llm.LLMProvider
	switch providerName {
	case "google":
		llmProvider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		llmProvider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		llmProvider = llm.NewOllamaProvider()
	default:
		return "", fmt.Errorf("未サポートのプロバイダーです: %s", providerName)
	}

	result, err := a.refactorEngine.ProposeOptimization(code, events, ragContext, llmProvider, model)
	if err != nil {
		return "", err
	}
	return result, nil
}

// DrawMarker は指定の座標に赤丸マークを上書き描画した画像を生成し、そのフルパスを返します
func (a *App) DrawMarker(imagePath string, x, y int) (string, error) {
	log.Printf("[App.DrawMarker] Request. ImagePath=%s, X=%d, Y=%d", imagePath, x, y)
	if imagePath == "" {
		return "", fmt.Errorf("画像パスが空です")
	}

	ext := filepath.Ext(imagePath)
	if ext == "" {
		ext = ".png"
	}

	dir := filepath.Dir(imagePath)
	filename := filepath.Base(imagePath)
	outFilename := fmt.Sprintf("marked_%d_%s", time.Now().UnixNano(), filename)
	outputPath := filepath.Join(dir, outFilename)

	err := capture.DrawClickMarker(imagePath, outputPath, x, y)
	if err != nil {
		log.Printf("[App.DrawMarker] Drawing click marker failed: %v", err)
		return "", err
	}

	log.Printf("[App.DrawMarker] Marked image created successfully: %s", outputPath)
	return outputPath, nil
}

// GetImageBase64 は指定された画像ファイルを読み込み、Base64エンコードしたデータURL文字列を返します
func (a *App) GetImageBase64(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[App.GetImageBase64] Error reading image: %v", err)
		return "", err
	}
	ext := strings.ToLower(filepath.Ext(path))
	mimeType := "image/png"
	switch ext {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".gif":
		mimeType = "image/gif"
	case ".webp":
		mimeType = "image/webp"
	}
	return fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(data)), nil
}

// RunInteractiveGuide は act_gram_guide/interactive_guide.uws を非同期実行します
func (a *App) RunInteractiveGuide() error {
	guidePath := filepath.Join(a.getExecBaseDir(), "manual", "act_gram_guide", "interactive_guide.uws")
	log.Printf("[App.RunInteractiveGuide] Preparing to run guide script: %s", guidePath)
	return a.RunScript(guidePath)
}

// MinimizeWindow はアプリのウィンドウを最小化します
func (a *App) MinimizeWindow() {
	runtime.WindowMinimise(a.ctx)
}

// RestoreWindow はアプリのウィンドウを復元し最前面に表示します
func (a *App) RestoreWindow() {
	runtime.WindowUnminimise(a.ctx)
	runtime.WindowShow(a.ctx)
}

// CheckAIConnection は設定されているGeminiなどの接続テストを行い、疎通に成功しているかを返します
func (a *App) CheckAIConnection() bool {
	if a.cfg == nil {
		return false
	}
	brainCfg, exists := a.cfg.Layers["brain"]
	if !exists {
		return false
	}
	apiKey, err := GetAPIKey(brainCfg.Provider)
	if err != nil || apiKey == "" {
		return false
	}
	return true
}

// StartRecording はマクロ記録を開始します（ウィンドウの自動最小化をトリガー）
func (a *App) StartRecording() error {
	log.Println("[App.StartRecording] Starting record session...")
	
	timestamp := time.Now().Format("20060102_150405")
	baseDir := a.getExecBaseDir()
	logDir := filepath.Join(baseDir, "manual", fmt.Sprintf("recording_%s", timestamp))
	
	a.mu.Lock()
	a.recorder = capture.NewRecorder(a.ctx, logDir)
	a.mu.Unlock()

	// 最小化
	a.MinimizeWindow()

	// 記録開始
	err := a.recorder.Start(func(outputPath string) error {
		return CaptureScreen(outputPath)
	})
	if err != nil {
		a.RestoreWindow()
		return fmt.Errorf("レコーダーの起動に失敗: %v", err)
	}

	return nil
}

// StopRecording は記録を停止し、記録されたログディレクトリを返します（ウィンドウの自動復元をトリガー）
func (a *App) StopRecording() (string, error) {
	log.Println("[App.StopRecording] Stopping record session...")
	a.mu.Lock()
	rec := a.recorder
	a.mu.Unlock()

	if rec == nil {
		return "", fmt.Errorf("レコーディングが開始されていません")
	}

	logDir, err := rec.Stop()
	
	// 復元
	a.RestoreWindow()

	return logDir, err
}
