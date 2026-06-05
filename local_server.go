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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read body error", http.StatusBadRequest)
		return
	}

	var req AIRequest
	// 1. 通常のパース
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		// 2. Windowsコマンドラインのエスケープ問題（全体がエスケープ付きの文字列として送られてくる）へのフォールバック
		var jsonStr string
		if errStr := json.Unmarshal(bodyBytes, &jsonStr); errStr == nil {
			// ネストされたJSON文字列を再度パース
			if errRetry := json.Unmarshal([]byte(jsonStr), &req); errRetry != nil {
				http.Error(w, fmt.Sprintf("Failed to parse nested JSON: %v", errRetry), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Invalid JSON: %v. Body: %s", err, string(bodyBytes)), http.StatusBadRequest)
			return
		}
	}

	// 3. 設定情報の取得
	cfg, err := LoadConfig()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load config: %v", err), http.StatusInternalServerError)
		return
	}

	// 4. レイヤーの決定 (画像がある場合は Eye層、ない場合は Brain層)
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
	apiKey, err := GetAPIKey(layerConfig.Provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get API key: %v", err), http.StatusInternalServerError)
		return
	}

	// 6. LLMプロバイダーの初期化と実行
	var provider llm.LLMProvider
	switch layerConfig.Provider {
	case "google":
		provider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		provider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		provider = llm.NewOllamaProvider()
	default:
		http.Error(w, fmt.Sprintf("Unsupported provider: %s", layerConfig.Provider), http.StatusBadRequest)
		return
	}

	// 7. LLM推論の実行
	fmt.Printf("[Local API] Executing prompt via layer '%s' (%s - %s). Image: %s\n", 
		layerName, layerConfig.Provider, layerConfig.Model, req.ImagePath)
	
	resp := provider.GenerateText(req.Prompt, req.ImagePath, layerConfig.Model)
	if resp.Error != nil {
		http.Error(w, fmt.Sprintf("LLM inference error: %v", resp.Error), http.StatusInternalServerError)
		return
	}

	// 8. テキストレスポンス返却 (UWSCRで受け取るためプレーンテキストで返却)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.Text))
}
