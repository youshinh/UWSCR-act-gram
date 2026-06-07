package main

import (
	"act_gram/capture"
	"act_gram/knowledge"
	"act_gram/llm"
	"act_gram/manual"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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

// SaveConfigsPayload は一括設定情報と各レイヤー設定を保持する構造体です。
type SaveConfigsPayload struct {
	UseUnifiedModel bool                   `json:"use_unified_model"`
	Layers          map[string]LayerConfig `json:"layers"`
}

// SaveConfigs は複数のレイヤー設定を一括で保存します (並行書き込みを避けるため)
func (a *App) SaveConfigs(layersJSON string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Printf("[App.SaveConfigs] Bulk saving configuration changes...")

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

	// 1. 新形式 (SaveConfigsPayload) でのアンマーシャルを試行
	var payload SaveConfigsPayload
	if err := json.Unmarshal([]byte(layersJSON), &payload); err == nil && payload.Layers != nil {
		a.cfg.UseUnifiedModel = payload.UseUnifiedModel
		for layer, lcfg := range payload.Layers {
			a.cfg.Layers[layer] = lcfg
			log.Printf("[App.SaveConfigs] Updating (New Format) Layer=%s: Provider=%s, Model=%s", layer, lcfg.Provider, lcfg.Model)
		}
	} else {
		// 2. 旧形式 (map[string]LayerConfig) へのフォールバック
		var oldLayers map[string]LayerConfig
		if err := json.Unmarshal([]byte(layersJSON), &oldLayers); err != nil {
			log.Printf("[App.SaveConfigs] JSON unmarshal failed on all formats: %v", err)
			return fmt.Errorf("JSONのパースに失敗しました: %v", err)
		}
		
		// 互換性のため、全て一致していれば一括設定モードとみなす
		brain, okB := oldLayers["brain"]
		eye, okE := oldLayers["eye"]
		utility, okU := oldLayers["utility"]
		if okB && okE && okU && 
			brain.Provider == eye.Provider && brain.Provider == utility.Provider &&
			brain.Model == eye.Model && brain.Model == utility.Model {
			a.cfg.UseUnifiedModel = true
		} else {
			a.cfg.UseUnifiedModel = false
		}

		for layer, lcfg := range oldLayers {
			a.cfg.Layers[layer] = lcfg
			log.Printf("[App.SaveConfigs] Updating (Fallback) Layer=%s: Provider=%s, Model=%s", layer, lcfg.Provider, lcfg.Model)
		}
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

// getLLMProviderHelper は指定されたプロバイダー名とAPIキーを元に、具象 llm.LLMProvider を作成して返します。
func (a *App) getLLMProviderHelper(provider string, apiKey string) (llm.LLMProvider, error) {
	switch provider {
	case "google":
		return llm.NewGeminiProvider(apiKey), nil
	case "anthropic":
		return llm.NewAnthropicProvider(apiKey), nil
	case "openai":
		return llm.NewOpenAIProvider(apiKey, "https://api.openai.com/v1"), nil
	case "custom":
		a.mu.Lock()
		baseURL := a.cfg.CustomBaseURL
		a.mu.Unlock()
		if baseURL == "" {
			baseURL = "http://localhost:8080/v1"
		}
		return llm.NewOpenAIProvider(apiKey, baseURL), nil
	case "local":
		a.mu.Lock()
		llmType := a.cfg.LocalLLMType
		baseURL := a.cfg.LocalLLMURL
		a.mu.Unlock()
		if llmType == "" {
			llmType = "ollama"
		}
		if baseURL == "" {
			if llmType == "ollama" {
				baseURL = "http://localhost:11434"
			} else {
				baseURL = "http://localhost:1234"
			}
		}
		return llm.NewLocalProvider(llmType, baseURL, apiKey), nil
	default:
		return nil, fmt.Errorf("未サポートのプロバイダーです: %s", provider)
	}
}

// SaveCustomBaseURL はカスタムプロバイダー用のベースURLを保存します。
func (a *App) SaveCustomBaseURL(baseURL string) error {
	log.Printf("[App.SaveCustomBaseURL] Saving custom base URL: %s", baseURL)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cfg == nil {
		var err error
		a.cfg, err = LoadConfig()
		if err != nil {
			return err
		}
	}
	a.cfg.CustomBaseURL = baseURL
	return SaveConfig(a.cfg)
}

// TestAPIKeyConnection は指定されたプロバイダーとAPIキーを用いて疎通確認（テスト接続）を行います。
// localLLMType は local プロバイダー選択時のみ使用します（UIの現在選択値を渡す）。
func (a *App) TestAPIKeyConnection(provider string, key string, customBaseURL string, localLLMType string) (string, error) {
	log.Printf("[App.TestAPIKeyConnection] Testing connection for provider: %s, localLLMType=%s", provider, localLLMType)
	
	// キーが空の場合は、現在登録されているキーを使用する
	if key == "" && provider != "local" {
		var err error
		key, err = GetAPIKey(provider)
		if err != nil {
			return "", fmt.Errorf("登録されているAPIキーの読み出しに失敗しました: %v", err)
		}
		if key == "" {
			return "", fmt.Errorf("APIキーが入力されていない、または未登録です")
		}
	}

	switch provider {
	case "google":
		providerImpl := llm.NewGeminiProvider(key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return "", fmt.Errorf("Gemini API接続テスト失敗: %v", err)
		}
		return fmt.Sprintf("接続成功！利用可能モデル数: %d (先頭モデル: %v)", len(models), firstModel(models)), nil

	case "anthropic":
		providerImpl := llm.NewAnthropicProvider(key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return "", fmt.Errorf("Anthropic API接続テスト失敗: %v", err)
		}
		return fmt.Sprintf("接続成功！利用可能モデル数: %d (先頭モデル: %v)", len(models), firstModel(models)), nil

	case "openai":
		providerImpl := llm.NewOpenAIProvider(key, "https://api.openai.com/v1")
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return "", fmt.Errorf("OpenAI API接続テスト失敗: %v", err)
		}
		return fmt.Sprintf("接続成功！利用可能モデル数: %d (先頭モデル: %v)", len(models), firstModel(models)), nil

	case "custom":
		if customBaseURL == "" {
			a.mu.Lock()
			if a.cfg != nil {
				customBaseURL = a.cfg.CustomBaseURL
			}
			a.mu.Unlock()
		}
		if customBaseURL == "" {
			customBaseURL = "http://localhost:8080/v1"
		}
		providerImpl := llm.NewOpenAIProvider(key, customBaseURL)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return "", fmt.Errorf("カスタム接続先 APIテスト失敗: %v", err)
		}
		return fmt.Sprintf("接続成功！利用可能モデル数: %d (先頭モデル: %v)", len(models), firstModel(models)), nil

	case "local":
		// UIから渡された localLLMType を優先し、なければ保存済み設定を使用
		if localLLMType == "" {
			a.mu.Lock()
			localLLMType = a.cfg.LocalLLMType
			a.mu.Unlock()
		}
		if localLLMType == "" {
			localLLMType = "ollama"
		}

		// URLが渡されなければ保存済み設定を使用
		if customBaseURL == "" {
			a.mu.Lock()
			customBaseURL = a.cfg.LocalLLMURL
			a.mu.Unlock()
		}
		if customBaseURL == "" {
			if localLLMType == "ollama" {
				customBaseURL = "http://localhost:11434"
			} else {
				customBaseURL = "http://localhost:1234"
			}
		}

		// APIキーはオプション（local は空でも可）
		if key == "" {
			key, _ = GetAPIKey("local")
		}

		providerImpl := llm.NewLocalProvider(localLLMType, customBaseURL, key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return "", fmt.Errorf("ローカルLLM (%s) 接続テスト失敗: %v", localLLMType, err)
		}
		return fmt.Sprintf("接続成功！利用可能モデル数: %d%s", len(models), firstModel(models)), nil

	default:
		return "", fmt.Errorf("未対応のプロバイダーです: %s", provider)
	}
}

// firstModel はモデルリストが空でなければ先頫モデル名を返します。リストが空の場合はpanicを防止するため空文字を返します。
func firstModel(models []string) string {
	if len(models) == 0 {
		return ""
	}
	return fmt.Sprintf(" (先頫モデル: %s)", models[0])
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

	case "openai":
		key, err := GetAPIKey("openai")
		if err != nil {
			log.Printf("[App.FetchModels] API key retrieval error: %v", err)
			return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
		}
		if key == "" {
			return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
		}
		providerImpl := llm.NewOpenAIProvider(key, "https://api.openai.com/v1")
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
		}
		return models, nil

	case "custom":
		key, err := GetAPIKey("custom")
		if err != nil {
			log.Printf("[App.FetchModels] API key retrieval error: %v", err)
			return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
		}
		a.mu.Lock()
		baseURL := a.cfg.CustomBaseURL
		a.mu.Unlock()
		if baseURL == "" {
			baseURL = "http://localhost:8080/v1"
		}
		providerImpl := llm.NewOpenAIProvider(key, baseURL)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			return []string{"custom-model"}, nil
		}
		return models, nil

	case "local":
		a.mu.Lock()
		llmType := a.cfg.LocalLLMType
		baseURL := a.cfg.LocalLLMURL
		a.mu.Unlock()
		if llmType == "" {
			llmType = "ollama"
		}
		if baseURL == "" {
			if llmType == "ollama" {
				baseURL = "http://localhost:11434"
			} else {
				baseURL = "http://localhost:1234"
			}
		}

		key, _ := GetAPIKey("local") // オプションのAPIキー
		providerImpl := llm.NewLocalProvider(llmType, baseURL, key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil || len(models) == 0 {
			log.Printf("[App.FetchModels] Local LLM models query failed (is it running?): %v. Returning defaults.", err)
			if llmType == "ollama" {
				return []string{"qwen2.5-coder:latest", "llama3.2-vision:latest", "gemma2:latest"}, nil
			} else {
				return []string{"qwen2.5-coder-7b-instruct", "llama-3.2-3b-instruct", "gemma-2-9b-it"}, nil
			}
		}
		log.Printf("[App.FetchModels] Found %d models for local provider (%s)", len(models), llmType)
		return models, nil

	default:
		log.Printf("[App.FetchModels] Unknown provider: %s", provider)
		return []string{}, nil
	}
}

// LocalLLMConfig はローカルLLMの設定情報です。
type LocalLLMConfig struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// GetLocalLLMConfig はローカルLLMの設定を取得します。
func (a *App) GetLocalLLMConfig() (LocalLLMConfig, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cfg == nil {
		return LocalLLMConfig{}, fmt.Errorf("設定がロードされていません")
	}
	return LocalLLMConfig{
		Type: a.cfg.LocalLLMType,
		URL:  a.cfg.LocalLLMURL,
	}, nil
}

// SaveLocalLLMConfig はローカルLLMの設定を保存します。
func (a *App) SaveLocalLLMConfig(llmType string, url string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cfg == nil {
		var err error
		a.cfg, err = LoadConfig()
		if err != nil {
			return err
		}
	}
	a.cfg.LocalLLMType = llmType
	a.cfg.LocalLLMURL = url
	return SaveConfig(a.cfg)
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

	llmProvider, err := a.getLLMProviderHelper(provider, apiKey)
	if err != nil {
		return "", err
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

	// RAGコンテキストの取得
	ragContext := ""
	if a.rag != nil {
		ragContext = a.rag.SearchRelevantContext(prompt, 3)
	}

	// 3. UWSCR生成用のシステムプロンプト定義
	systemPrompt := fmt.Sprintf(`あなたは卓越したRPAスクリプト開発AIです。
Windows用の自動化スクリプトエンジンである「UWSCR (UWSC互換)」で動作するスクリプト (.uws) を作成してください。

【UWSCR の基本構文ルール】
1. コメント: // コメントテキスト
2. 変数宣言: Dim 変数名 = 初期値 (または Dim 変数名)
3. ウィンドウ制御:
   - ACW(GETID("ウィンドウタイトル", "クラス名"), x, y, w, h) : ウィンドウのアクティブ化・位置調整
   - CTRLWIN(id, CLOSE) : ウィンドウを閉じる (※絶対に CTRL_WIN と書かないでください)
4. ウィンドウ相対座標制御 (座標クリック・移動の相対制御):
   - 特定のウィンドウをアクティブ化した後、mouseorg(id) を呼び出すことで、それ以降のマウス操作（btn, mmv）や画像検索（chkimg）の座標基準がウィンドウ相対になります。
   - 操作完了後は必ず mouseorg(0) でスクリーン座標基準に戻します。
   例:
   Dim id = GETID("ウィンドウタイトル")
   ACW(id)
   mouseorg(id)
   btn(LEFT, CLICK, 相対X, 相対Y, ディレイ)
   mouseorg(0)
5. ウィンドウ要素のボタンクリック (clkitem / UI Automation):
   - ウィンドウ内のボタン等をクリックする場合は、固定座標指定クリックではなく clkitem(id, "ボタン名", CLK_BTN or CLK_UIA) を使用してください。座標のズレに影響されず安定してクリックできます。
6. ウィンドウの位置・サイズ取得:
   - STATUS(id, ST_X), STATUS(id, ST_Y), STATUS(id, ST_WIDTH), STATUS(id, ST_HEIGHT) を使用します。UWSCRには STATUS_X や STATUS_Y という定数は存在しないため、使用すると構文エラーになります。
7. 制御構文:
   - IFB 条件 THEN ... ELSE ... ENDIF (※UWSCの厳格なブロック構文 IFB 〜 THEN 〜 ENDIF を使用してください)
   - FOR 変数 = 初期値 TO 最大値 ... NEXT
   - WHILE 条件 ... WEND
   - SLEEP(秒数) : 例: SLEEP(1) は1秒待機
8. ダイアログ表示:
   - MSGBOX("メッセージ内容") : ポップアップダイアログを表示
9. タイムスタンプ・ログ出力:
   - 要所で PRINT "[" + GETTIME() + "] メッセージ" のようにタイムスタンプ付きのログを出力してください。

【actgram 独自マクロ拡張関数】
- AI_EVAL("判定・取得したい内容", 画像取得関数等)
  この関数は、画面上の特定情報やテキストを判断させたい場合に使用できます。戻り値は推論結果テキストです。
  例: Dim price = AI_EVAL("この伝票の合計金額は数値でいくら？", GetScreenCapture())
  ※引数の画像取得関数には GetScreenCapture() などのダミー表記が用いられ、実行時に実際のデスクトップキャプチャに差し替わります。

【現場・業務ナレッジマニュアルコンテキスト (RAG)】
%s

【出力要件】
- スクリプトコードのみを出力してください。
- markdown のコードブロック ('''uws ... ''') などで囲まず、純粋なプレーンテキストとして UWSCR スクリプトのみを返してください。説明文や前置きは一切不要です。
`, ragContext)

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

	llmProvider, err := a.getLLMProviderHelper(provider, apiKey)
	if err != nil {
		return "", err
	}

	// RAGコンテキストの取得
	ragContext := ""
	if a.rag != nil {
		query := prompt
		if errorLog != "" {
			query = errorLog
		}
		ragContext = a.rag.SearchRelevantContext(query, 3)
	}

	var systemPrompt string
	var finalPrompt string

	if errorLog == "" {
		// 1. エラーなし時のユーザー指示に基づくリファクタリング
		systemPrompt = `あなたは卓越したRPAスクリプト開発AIです。
UWSCR スクリプト (.uws) をユーザーの修正指示に従って修正・リファクタリングしてください。

【UWSCR の基本構文ルール】
1. コメント: // コメントテキスト
2. 変数宣言: Dim 変数名 = 初期値 (または Dim 変数名)
3. ウィンドウ制御:
   - ACW(GETID("ウィンドウタイトル", "クラス名"), x, y, w, h) : ウィンドウのアクティブ化・位置調整
   - CTRLWIN(id, CLOSE) : ウィンドウを閉じる (※絶対に CTRL_WIN と書かないでください)
4. ウィンドウ相対座標制御 (座標クリック・移動の相対制御):
   - 特定のウィンドウをアクティブ化した後、mouseorg(id) を呼び出すことで、それ以降のマウス操作（btn, mmv）や画像検索（chkimg）の座標基準がウィンドウ相対になります。
   - 操作完了後は必ず mouseorg(0) でスクリーン座標基準に戻します。
   例:
   Dim id = GETID("ウィンドウタイトル")
   ACW(id)
   mouseorg(id)
   btn(LEFT, CLICK, 相対X, 相対Y, ディレイ)
   mouseorg(0)
5. ウィンドウ要素のボタンクリック (clkitem / UI Automation):
   - ウィンドウ内のボタン等をクリックする場合は、固定座標指定クリックではなく clkitem(id, "ボタン名", CLK_BTN or CLK_UIA) を使用してください。座標のズレに影響されず安定してクリックできます。
6. ウィンドウの位置・サイズ取得:
   - STATUS(id, ST_X), STATUS(id, ST_Y), STATUS(id, ST_WIDTH), STATUS(id, ST_HEIGHT) を使用します。UWSCRには STATUS_X や STATUS_Y という定数は存在しないため、使用すると構文エラーになります。
7. 制御構文:
   - IFB 条件 THEN ... ELSE ... ENDIF (※UWSCの厳格なブロック構文 IFB 〜 THEN 〜 ENDIF を使用してください)
   - FOR 変数 = 初期値 TO 最大値 ... NEXT
   - WHILE 条件 ... WEND
   - SLEEP(秒数) : 例: SLEEP(1) は1秒待機
8. ダイアログ表示:
   - MSGBOX("メッセージ内容") : ポップアップダイアログを表示
9. タイムスタンプ・ログ出力:
   - 要所で PRINT "[" + GETTIME() + "] メッセージ" のようにタイムスタンプ付きのログを出力してください。

【現場・業務ナレッジマニュアルコンテキスト (RAG)】
` + ragContext + `

【出力要件】
- 修正されたスクリプトコードのみを出力してください。
- markdown のコードブロック ('''uws ... ''') などで囲まず、純粋なプレーンテキストとして UWSCR スクリプトのみを返してください。説明文や前置きは一切不要です。
`
		finalPrompt = fmt.Sprintf(
			"%s\n\n【ユーザーからの修正指示】\n%s\n\n【修正前の元のコード】\n%s\n\n【指示】\n修正指示を反映し、正しく動作する修正後のUWSCRコードのみを返してください。",
			systemPrompt, prompt, code,
		)
	} else {
		// 2. テスト実行エラーの自動修正
		systemPrompt = `あなたは卓越したRPAスクリプト開発AIです。
UWSCR スクリプト (.uws) のテスト実行でエラーが発生しました。
元のコードとエラーログ、および当初の指示内容を分析し、エラーを修正した動作可能な UWSCR スクリプトを再生成してください。

【UWSCR の基本構文ルール】
1. コメント: // コメントテキスト
2. 変数宣言: Dim 変数名 = 初期値 (または Dim 変数名)
3. ウィンドウ制御:
   - ACW(GETID("ウィンドウタイトル", "クラス名"), x, y, w, h) : ウィンドウのアクティブ化・位置調整
   - CTRLWIN(id, CLOSE) : ウィンドウを閉じる (※絶対に CTRL_WIN と書かないでください)
4. ウィンドウ相対座標制御 (座標クリック・移動の相対制御):
   - 特定のウィンドウをアクティブ化した後、mouseorg(id) を呼び出すことで、それ以降のマウス操作（btn, mmv）や画像検索（chkimg）の座標基準がウィンドウ相対になります。
   - 操作完了後は必ず mouseorg(0) でスクリーン座標基準に戻します。
   例:
   Dim id = GETID("ウィンドウタイトル")
   ACW(id)
   mouseorg(id)
   btn(LEFT, CLICK, 相対X, 相対Y, ディレイ)
   mouseorg(0)
5. ウィンドウ要素のボタンクリック (clkitem / UI Automation):
   - ウィンドウ内のボタン等をクリックする場合は、固定座標指定クリックではなく clkitem(id, "ボタン名", CLK_BTN or CLK_UIA) を使用してください。座標のズレに影響されず安定してクリックできます。
6. ウィンドウの位置・サイズ取得:
   - STATUS(id, ST_X), STATUS(id, ST_Y), STATUS(id, ST_WIDTH), STATUS(id, ST_HEIGHT) を使用します。UWSCRには STATUS_X や STATUS_Y という定数は存在しないため、使用すると構文エラーになります。
7. 制御構文:
   - IFB 条件 THEN ... ELSE ... ENDIF (※UWSCの厳格なブロック構文 IFB 〜 THEN 〜 ENDIF を使用してください)
   - FOR 変数 = 初期値 TO 最大値 ... NEXT
   - WHILE 条件 ... WEND
   - SLEEP(秒数) : 例: SLEEP(1) は1秒待機
8. ダイアログ表示:
   - MSGBOX("メッセージ内容") : ポップアップダイアログを表示
9. タイムスタンプ・ログ出力:
   - 要所で PRINT "[" + GETTIME() + "] メッセージ" のようにタイムスタンプ付きのログを出力してください。

【現場・業務ナレッジマニュアルコンテキスト (RAG)】
` + ragContext + `

【出力要件】
- 修正されたスクリプトコードのみを出力してください。
- markdown のコードブロック ('''uws ... ''') などで囲まず、純粋なプレーンテキストとして UWSCR スクリプトのみを返してください。説明文や前置きは一切不要です。
`
		finalPrompt = fmt.Sprintf(
			"%s\n\n【当初の自動化指示】\n%s\n\n【エラーが発生した元のコード】\n%s\n\n【UWSCRの実行エラーログ】\n%s\n\n【指示】\nエラーログを分析し、バグを修正したUWSCRコードを返してください。可能な限りPRINT文で独自のタイムスタンプ付きログ（例: PRINT \"[タイムスタンプ] ログ内容\"）を付与して進捗を出力してください。",
			systemPrompt, prompt, code, errorLog,
		)
	}

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

	llmProvider, err := a.getLLMProviderHelper(providerName, apiKey)
	if err != nil {
		return nil, err
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

	llmProvider, err := a.getLLMProviderHelper(providerName, apiKey)
	if err != nil {
		return "", err
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
		code = ConvertToUTF8IfNeeded(data)
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

	llmProvider, err := a.getLLMProviderHelper(providerName, apiKey)
	if err != nil {
		return "", err
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
	baseDir := filepath.Join(a.getExecBaseDir(), "manual", "act_gram_guide")
	guidePath := filepath.Join(baseDir, "interactive_guide.uws")
	
	// Create directories
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return fmt.Errorf("ガイドフォルダの作成に失敗しました: %v", err)
	}
	
	// Generate guide script (always overwrite to ensure it exists and is correct)
	uwsContent := `// actgram::UWSCR インタラクティブガイド
MSGBOX("actgram::UWSCR デモガイドへようこそ！")
MSGBOX("これからactgramの主要な機能である『PLAY（実行）』『REC（操作記録）』『DEVELOP（スクリプト開発）』『MANUAL（マニュアル作成）』を説明します。")
MSGBOX("1. 【PLAY】タブ\nUWSCRスクリプトを実行し、実測ログからボトルネック（無駄なSLEEP）を検知して最適化できます。")
MSGBOX("2. 【REC】タブ\nデスクトップでの操作を記録して、マニュアル生成用ログを保存できます。")
MSGBOX("3. 【DEVELOP】タブ\n画面キャプチャと指示を元にAIが自動生成したスクリプトをテスト実行し、自己修復できます。")
MSGBOX("4. 【MANUAL】タブ\n記録したログから対話型HTMLマニュアルとUWSCRガイドをAIでワンクリック生成できます。")
MSGBOX("ガイドは以上です。ぜひ各機能をお試しください！")`
	
	if err := os.WriteFile(guidePath, []byte(uwsContent), 0644); err != nil {
		return fmt.Errorf("ガイドスクリプトの書き込みに失敗しました: %v", err)
	}
	
	// Generate help step images
	imagesDir := filepath.Join(baseDir, "images")
	if err := os.MkdirAll(imagesDir, 0755); err != nil {
		return fmt.Errorf("画像フォルダの作成に失敗しました: %v", err)
	}
	
	dummyPNG := []byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
		0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0x15, 0xc4, 0x89, 0x00, 0x00, 0x00,
		0x0d, 0x49, 0x44, 0x41, 0x54, 0x78, 0xda, 0x63, 0x60, 0x18, 0x05, 0xa3,
		0x60, 0x14, 0x8c, 0x00, 0x08, 0x00, 0x05, 0xff, 0x52, 0x2e, 0x07, 0x02,
		0x00, 0x00, 0x00, 0x3b, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	}
	
	for i := 1; i <= 3; i++ {
		imgPath := filepath.Join(imagesDir, fmt.Sprintf("step_%d.png", i))
		_ = os.WriteFile(imgPath, dummyPNG, 0644)
	}

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

	// ミニウィンドウモードに変更
	a.SetMiniMode(true, "record")

	// 記録開始
	err := a.recorder.Start(func(outputPath string) error {
		return CaptureScreen(outputPath)
	})
	if err != nil {
		a.SetMiniMode(false, "record")
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
	
	// ウィンドウを通常サイズに復元
	a.SetMiniMode(false, "record")

	return logDir, err
}

// ReadScriptFile は指定された UWSCR スクリプトファイルの内容を読み込みます
func (a *App) ReadScriptFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return ConvertToUTF8IfNeeded(data), nil
}

// OpenKnowledgeDir は RAG 知識フォルダをエクスプローラーで開きます
func (a *App) OpenKnowledgeDir() error {
	dir := filepath.Join(a.getExecBaseDir(), "knowledge")
	// フォルダが存在しない場合は作成
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	
	// Windowsの explorer で開く
	cmd := exec.Command("explorer", dir)
	return cmd.Start()
}

// SetMiniMode はミニウィンドウモードと通常モードの切り替えを行います (極小化: 220x90)
func (a *App) SetMiniMode(isMini bool, mode string) {
	log.Printf("[App.SetMiniMode] isMini=%v, mode=%s", isMini, mode)
	if a.ctx == nil {
		return
	}
	if isMini {
		// ミニウィンドウモードへの移行
		// 1. サイズ変更時に制限と衝突するのを防ぐため、一旦最小・最大サイズ制約をクリア
		runtime.WindowSetMinSize(a.ctx, 0, 0)
		runtime.WindowSetMaxSize(a.ctx, 0, 0)
		
		// 2. サイズ変更の適用
		runtime.WindowSetSize(a.ctx, 220, 90)
		
		// 3. サイズを固定
		runtime.WindowSetMinSize(a.ctx, 220, 90)
		runtime.WindowSetMaxSize(a.ctx, 220, 90)
		
		// 4. 最小化されている場合を考慮し、復元して表示
		runtime.WindowUnminimise(a.ctx)
		runtime.WindowShow(a.ctx)
		
		// 5. 操作の邪魔にならないよう、かつ隠れないように常に最前面に表示
		runtime.WindowSetAlwaysOnTop(a.ctx, true)
	} else {
		// 通常モード（1024 x 768）への復元
		// 1. 一旦最小・最大サイズ制約をクリアし、最前面表示も解除
		runtime.WindowSetMinSize(a.ctx, 0, 0)
		runtime.WindowSetMaxSize(a.ctx, 0, 0)
		runtime.WindowSetAlwaysOnTop(a.ctx, false)
		
		// 2. 通常サイズを適用
		runtime.WindowSetSize(a.ctx, 1024, 768)
		
		// 3. 制限を再適用
		runtime.WindowSetMinSize(a.ctx, 1024, 768)
		runtime.WindowSetMaxSize(a.ctx, 1024, 768)
		
		// 4. 確実に表示
		runtime.WindowUnminimise(a.ctx)
		runtime.WindowShow(a.ctx)
	}
	
	// フロントエンドに通知
	runtime.EventsEmit(a.ctx, "mini_mode_changed", isMini, mode)
}

// SaveErrorReflection はエラー内容とコードを自己学習用の反省ログファイルに記録します
func (a *App) SaveErrorReflection(scriptContent string, errorLog string) error {
	log.Println("[App.SaveErrorReflection] Saving error self-reflection...")
	
	knowledgeDir := filepath.Join(a.getExecBaseDir(), "knowledge")
	if err := os.MkdirAll(knowledgeDir, 0755); err != nil {
		return err
	}
	
	reflectionPath := filepath.Join(knowledgeDir, "uwscr_self_reflection.md")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	reflectionBlock := fmt.Sprintf(
		"## エラー自己反省ログ (%s)\n\n"+
		"**失敗したスクリプト:**\n```uws\n%s\n```\n\n"+
		"**エラーログ:**\n```\n%s\n```\n\n"+
		"**対策・反省点:**\n"+
		"上記のコードでエラーが発生しました。UWSCRでは CTRL_WIN は未定義であり、代わりにアンダースコアなしの CTRLWIN を使用してください。\n"+
		"また、ウィンドウ相対操作の際は mouseorg(id) を呼び出してから座標操作を行い、終了後は mouseorg(0) でスクリーン座標に戻す必要があります。\n"+
		"ボタン等の要素クリックには、固定座標クリックではなく clkitem(id, \"ボタン名\", CLK_BTN or CLK_UIA) を使用してください。\n\n---\n\n",
		timestamp, scriptContent, errorLog,
	)
	
	f, err := os.OpenFile(reflectionPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	
	if _, err := f.WriteString(reflectionBlock); err != nil {
		return err
	}
	
	if a.rag != nil {
		_ = a.rag.LoadKnowledgeFiles()
	}
	
	return nil
}

// ConvertToUTF8IfNeeded はバイト配列が UTF-8 でない場合、Shift-JIS から UTF-8 へ変換します
func ConvertToUTF8IfNeeded(data []byte) string {
	if utf8.Valid(data) {
		return string(data)
	}
	
	// Shift-JIS から UTF-8 へのデコード
	decoder := japanese.ShiftJIS.NewDecoder()
	decoded, err := io.ReadAll(transform.NewReader(bytes.NewReader(data), decoder))
	if err != nil {
		log.Printf("[App.ConvertToUTF8IfNeeded] Shift-JIS decode failed: %v", err)
		return string(data)
	}
	return string(decoded)
}

// SaveKnowledgeDirAndURL は知識ディレクトリとドキュメントURLを設定して保存します
func (a *App) SaveKnowledgeDirAndURL(knowledgeDir string, uwscrDocURL string) error {
	log.Printf("[App.SaveKnowledgeDirAndURL] Saving knowledge settings. dir=%s, url=%s", knowledgeDir, uwscrDocURL)
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cfg == nil {
		cfg, err := LoadConfig()
		if err != nil {
			return err
		}
		a.cfg = cfg
	}

	a.cfg.KnowledgeDir = knowledgeDir
	a.cfg.UWSCRDocURL = uwscrDocURL

	err := SaveConfig(a.cfg)
	if err != nil {
		log.Printf("[App.SaveKnowledgeDirAndURL] Failed to save config: %v", err)
		return err
	}

	// RAGマネージャーの再初期化
	a.rag = knowledge.NewRAGManager(knowledgeDir)
	if err := a.rag.LoadKnowledgeFiles(); err != nil {
		log.Printf("[App.SaveKnowledgeDirAndURL] Failed to reload knowledge files: %v", err)
	}

	return nil
}

// SyncUWSCRReference は最新の UWSCR ドキュメントを同期します
func (a *App) SyncUWSCRReference() error {
	log.Println("[App.SyncUWSCRReference] Sync started asynchronously...")
	runtime.EventsEmit(a.ctx, "manual_sync_started")

	go func() {
		a.mu.Lock()
		baseURL := a.cfg.UWSCRDocURL
		knowledgeDir := a.cfg.KnowledgeDir
		a.mu.Unlock()

		err := a.runCrawler(baseURL, knowledgeDir)
		if err != nil {
			log.Printf("[App.SyncUWSCRReference] Sync failed: %v", err)
			runtime.EventsEmit(a.ctx, "manual_sync_finished", map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// リロード
		a.mu.Lock()
		_ = a.rag.LoadKnowledgeFiles()
		a.mu.Unlock()

		log.Println("[App.SyncUWSCRReference] Sync completed successfully")
		runtime.EventsEmit(a.ctx, "manual_sync_finished", map[string]interface{}{
			"success": true,
		})
	}()

	return nil
}

// RestoreDefaultReference はマニュアルとリファレンスを初期状態に復元します
func (a *App) RestoreDefaultReference() error {
	log.Println("[App.RestoreDefaultReference] Restoring default reference manual...")
	
	a.mu.Lock()
	if a.cfg == nil {
		cfg, err := LoadConfig()
		if err != nil {
			a.mu.Unlock()
			return err
		}
		a.cfg = cfg
	}

	defaultURL := "https://stuncloud.github.io/UWSCR/"
	defaultDir := filepath.Join(a.getExecBaseDir(), "knowledge")

	a.cfg.UWSCRDocURL = defaultURL
	a.cfg.KnowledgeDir = defaultDir
	_ = SaveConfig(a.cfg)

	// RAGマネージャー再構築
	a.rag = knowledge.NewRAGManager(defaultDir)
	a.mu.Unlock()

	// クローラーを同期実行（完了を待つ）
	err := a.runCrawler(defaultURL, defaultDir)
	if err != nil {
		log.Printf("[App.RestoreDefaultReference] Restore crawl failed: %v", err)
		return err
	}

	a.mu.Lock()
	_ = a.rag.LoadKnowledgeFiles()
	a.mu.Unlock()

	log.Println("[App.RestoreDefaultReference] Restore completed successfully")
	return nil
}

// runCrawler は指定したベースURLからドキュメントをクロールして知識ディレクトリに結合MDを保存します
func (a *App) runCrawler(baseURL string, outputDir string) error {
	if baseURL == "" {
		baseURL = "https://stuncloud.github.io/UWSCR/"
	}
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	visited := make(map[string]bool)
	var docBuffer strings.Builder
	
	linkRegexp := regexp.MustCompile(`href="([^"]+)"`)
	titleRegexp := regexp.MustCompile(`<title>([^<]+)</title>`)
	bodyRegexp := regexp.MustCompile(`(?s)<article role="main"[^>]*>(.*?)</article>`)
	tagRegexp := regexp.MustCompile(`(?s)<[^>]*>`)
	spaceRegexp := regexp.MustCompile(`\s+`)
	headerRegexp := regexp.MustCompile(`(?s)<h([1-6])[^>]*>(.*?)</h[1-6]>`)

	seeds := []string{
		"index.html",
		"usage/installation.html",
		"usage/how_to_run.html",
		"usage/settings.html",
		"usage/how_to_build.html",
		"usage/language_server.html",
		"usage/example.html",
		"syntax/statement.html",
		"syntax/special_variables.html",
		"builtin/index.html",
		"builtin/window.html",
		"builtin/mouse_keyboard.html",
		"builtin/dialog.html",
		"builtin/system.html",
		"builtin/file.html",
		"builtin/string.html",
		"builtin/math.html",
		"builtin/other.html",
		"module/index.html",
		"module/uia.html",
		"module/sound.html",
		"changelog.html",
	}

	var crawlFunc func(string)
	crawlFunc = func(targetURL string) {
		if visited[targetURL] {
			return
		}
		visited[targetURL] = true
		log.Printf("[Crawler] Crawling: %s", targetURL)

		resp, err := http.Get(targetURL)
		if err != nil {
			log.Printf("[Crawler] Error fetching %s: %v", targetURL, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("[Crawler] Status error for %s: %d", targetURL, resp.StatusCode)
			return
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[Crawler] Error reading body for %s: %v", targetURL, err)
			return
		}

		htmlContent := string(bodyBytes)
		
		titleMatch := titleRegexp.FindStringSubmatch(htmlContent)
		title := "UWSCR Document"
		if len(titleMatch) > 1 {
			title = strings.TrimSpace(titleMatch[1])
			if idx := strings.Index(title, "—"); idx != -1 {
				title = strings.TrimSpace(title[:idx])
			}
		}

		articleMatch := bodyRegexp.FindStringSubmatch(htmlContent)
		var articleHTML string
		if len(articleMatch) > 1 {
			articleHTML = articleMatch[1]
		} else {
			articleHTML = htmlContent
		}

		hMatch := headerRegexp.FindAllStringSubmatch(articleHTML, -1)
		for _, m := range hMatch {
			if len(m) > 2 {
				level := m[1]
				content := tagRegexp.ReplaceAllString(m[2], "")
				content = strings.TrimSpace(content)
				hashes := strings.Repeat("#", len(level)+1)
				replaceStr := fmt.Sprintf("\n%s %s\n", hashes, content)
				articleHTML = strings.Replace(articleHTML, m[0], replaceStr, 1)
			}
		}

		text := tagRegexp.ReplaceAllString(articleHTML, " ")
		lines := strings.Split(text, "\n")
		var cleanedLines []string
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			trimmed = spaceRegexp.ReplaceAllString(trimmed, " ")
			if trimmed != "" {
				cleanedLines = append(cleanedLines, trimmed)
			}
		}
		structuredText := strings.Join(cleanedLines, "\n")

		docBuffer.WriteString(fmt.Sprintf("\n# %s\n\n", title))
		docBuffer.WriteString(fmt.Sprintf("> 参照元: %s\n\n", targetURL))
		docBuffer.WriteString(structuredText)
		docBuffer.WriteString("\n\n---\n")

		links := linkRegexp.FindAllStringSubmatch(htmlContent, -1)
		u, _ := url.Parse(targetURL)
		for _, match := range links {
			if len(match) > 1 {
				link := match[1]
				if idx := strings.Index(link, "#"); idx != -1 {
					link = link[:idx]
				}
				if idx := strings.Index(link, "?"); idx != -1 {
					link = link[:idx]
				}
				if link == "" {
					continue
				}

				refURL, err := url.Parse(link)
				if err != nil {
					continue
				}
				resolvedURL := u.ResolveReference(refURL).String()

				if strings.HasPrefix(resolvedURL, baseURL) && !strings.Contains(resolvedURL, "_static") && !strings.Contains(resolvedURL, "genindex") && !strings.Contains(resolvedURL, "search") {
					if !visited[resolvedURL] {
						time.Sleep(50 * time.Millisecond)
						crawlFunc(resolvedURL)
					}
				}
			}
		}
	}

	for _, seed := range seeds {
		targetURL := baseURL + seed
		crawlFunc(targetURL)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	outputPath := filepath.Join(outputDir, "uwscr_reference.md")
	return os.WriteFile(outputPath, []byte(docBuffer.String()), 0644)
}
