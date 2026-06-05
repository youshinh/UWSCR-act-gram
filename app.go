package main

import (
	"act_gram/llm"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	cfg *Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
	}
	a.cfg = cfg
}

// GetConfig は現在の設定をフロントエンドに返します
func (a *App) GetConfig() (*Config, error) {
	if a.cfg == nil {
		cfg, err := LoadConfig()
		if err != nil {
			return nil, err
		}
		a.cfg = cfg
	}
	return a.cfg, nil
}

// SaveConfig は指定されたレイヤーの設定を保存します
func (a *App) SaveConfig(layer string, provider string, model string) error {
	if a.cfg == nil {
		var err error
		a.cfg, err = LoadConfig()
		if err != nil {
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

	return SaveConfig(a.cfg)
}

// SaveAPIKey はAPIキーをセキュア領域に保存します
func (a *App) SaveAPIKey(provider string, key string) error {
	return SaveAPIKey(provider, key)
}

// HasAPIKey は指定されたプロバイダーのAPIキーが登録されているかを判定します
func (a *App) HasAPIKey(provider string) bool {
	key, err := GetAPIKey(provider)
	if err != nil || key == "" {
		return false
	}
	return true
}

// FetchModels は指定されたプロバイダーの利用可能モデルを動的に取得して返します
func (a *App) FetchModels(provider string) ([]string, error) {
	switch provider {
	case "google":
		key, err := GetAPIKey("google")
		if err != nil {
			return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
		}
		if key == "" {
			return []string{"gemini-2.5-flash", "gemini-2.5-flash-lite", "gemini-1.5-pro"}, nil // 未登録時は主要モデルのフォールバック
		}
		providerImpl := llm.NewGeminiProvider(key)
		models, err := providerImpl.GetAvailableModels()
		if err != nil {
			// APIエラー時はデフォルトリストにフォールバック
			return []string{"gemini-2.5-flash", "gemini-2.5-flash-lite", "gemini-1.5-pro"}, nil
		}
		return models, nil

	case "anthropic":
		key, err := GetAPIKey("anthropic")
		if err != nil {
			return nil, fmt.Errorf("APIキーの取得に失敗しました: %v", err)
		}
		providerImpl := llm.NewAnthropicProvider(key)
		return providerImpl.GetAvailableModels()

	case "ollama":
		providerImpl := llm.NewOllamaProvider()
		models, err := providerImpl.GetAvailableModels()
		if err != nil || len(models) == 0 {
			// Ollamaが未起動または取得失敗時は主要な標準モデルをフォールバック
			return []string{"qwen2.5-coder:latest", "llama3.2-vision:latest", "gemma2:latest"}, nil
		}
		return models, nil

	default:
		return []string{}, nil
	}
}
