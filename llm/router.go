package llm

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LLMResponse は各プロバイダーからの統一された応答構造体です。
type LLMResponse struct {
	Text  string
	Error error
}

// LLMProvider は異なるLLMプロバイダーへのリクエストを抽象化するインターフェースです。
type LLMProvider interface {
	GetAvailableModels() ([]string, error)
	GenerateText(prompt string, imagePath string, model string) LLMResponse
}

// 共通クライアント
type httpClient struct {
	client *http.Client
}

func newHTTPClient() *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// helper: 画像ファイルをBase64エンコードし、MimeTypeを判別します
func readImageAsBase64(imagePath string) (string, string, error) {
	if imagePath == "" {
		return "", "", nil
	}

	data, err := os.ReadFile(imagePath)
	if err != nil {
		return "", "", err
	}

	ext := strings.ToLower(filepath.Ext(imagePath))
	mimeType := "image/png"
	switch ext {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".gif":
		mimeType = "image/gif"
	case ".webp":
		mimeType = "image/webp"
	case ".pdf":
		mimeType = "application/pdf"
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, mimeType, nil
}
