package main

import (
	"act_gram/llm"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// LocalServer はUWSCRからのリクエストを受け取るローカルAPIサーバーです。
type LocalServer struct {
	server *http.Server
	app    *App
}

type AIRequest struct {
	Prompt    string `json:"prompt"`
	ImagePath string `json:"image_path"`
}

func NewLocalServer(app *App) *LocalServer {
	return &LocalServer{app: app}
}

// Start はHTTPサーバーを起動します (非同期で起動してください)
func (s *LocalServer) Start(port int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ai_eval", s.handleAIEval)

	s.server = &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: mux,
	}

	fmt.Printf("Starting Local API Server on http://127.0.0.1:%d...\n", port)
	return s.server.ListenAndServe()
}

// Shutdown はサーバーを安全にシャットダウンします
func (s *LocalServer) Shutdown() error {
	if s.server == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *LocalServer) handleAIEval(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DEBUG 1] /ai_eval リクエストを受信しました")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read body error", http.StatusBadRequest)
		return
	}
	fmt.Println("[DEBUG 2] リクエストボディの読み込み完了")

	var req AIRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		var jsonStr string
		if errStr := json.Unmarshal(bodyBytes, &jsonStr); errStr == nil {
			if errRetry := json.Unmarshal([]byte(jsonStr), &req); errRetry != nil {
				http.Error(w, fmt.Sprintf("Failed to parse nested JSON: %v", errRetry), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Invalid JSON: %v. Body: %s", err, string(bodyBytes)), http.StatusBadRequest)
			return
		}
	}
	fmt.Printf("[DEBUG 3] JSONパース成功: Prompt=%s, ImagePath=%s\n", req.Prompt, req.ImagePath)

	// 3. 設定情報の取得
	fmt.Println("[DEBUG 4] LoadConfig() を呼び出します...")
	cfg, err := LoadConfig()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load config: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("[DEBUG 5] LoadConfig() 成功: Provider=%s, Model=%s, LocalLLMType=%s, LocalLLMURL=%s\n", 
		cfg.Layers["brain"].Provider, cfg.Layers["brain"].Model, cfg.LocalLLMType, cfg.LocalLLMURL)

	// 4. レイヤーの決定
	layerName := "brain"
	if req.ImagePath != "" {
		layerName = "eye"
	}

	layerConfig, exists := cfg.Layers[layerName]
	if !exists {
		http.Error(w, fmt.Sprintf("Config layer '%s' not found", layerName), http.StatusInternalServerError)
		return
	}

	// 5. APIキーのセキュア取得
	fmt.Printf("[DEBUG 6] GetAPIKey(%s) を呼び出します...\n", layerConfig.Provider)
	apiKey, err := GetAPIKey(layerConfig.Provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get API key: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println("[DEBUG 7] APIキー取得処理を通過")

	// 6. LLMプロバイダーの初期化
	var provider llm.LLMProvider
	// (switch文のロジックはそのまま維持してください)
	switch layerConfig.Provider {
	// ... (中略) ...
	case "local":
		llmType := cfg.LocalLLMType
		baseURL := cfg.LocalLLMURL
		if llmType == "" { llmType = "ollama" }
		if baseURL == "" {
			if llmType == "ollama" { baseURL = "http://localhost:11434" } else { baseURL = "http://localhost:1234" }
		}
		provider = llm.NewLocalProvider(llmType, baseURL, apiKey)
	default:
		// デフォルト処理
	}

	// 7. LLM推論の実行
	fmt.Printf("[DEBUG 8] LM Studioへのリクエストを送信直前です！ URL: %s, Model: %s\n", cfg.LocalLLMURL, layerConfig.Model)
	
	resp := provider.GenerateText(req.Prompt, req.ImagePath, layerConfig.Model)
	
	fmt.Println("[DEBUG 9] LM Studioからレスポンスが返ってきました！")
	if resp.Error != nil {
		fmt.Printf("[DEBUG ERROR] 推論エラーが発生: %v\n", resp.Error)
		http.Error(w, fmt.Sprintf("LLM inference error: %v", resp.Error), http.StatusInternalServerError)
		return
	}

	// 8. テキストレスポンス返却
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.Text))
	fmt.Println("[DEBUG 10] クライアント（UWSCR/UI）へ正常に応答を書き込みました")
}