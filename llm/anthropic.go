package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AnthropicProvider struct {
	apiKey string
	*httpClient
}

func NewAnthropicProvider(apiKey string) *AnthropicProvider {
	return &AnthropicProvider{
		apiKey:     apiKey,
		httpClient: newHTTPClient(),
	}
}

// AnthropicはAPIによるモデル一覧取得を提供していないため、定義されたリストを返す
func (p *AnthropicProvider) GetAvailableModels() ([]string, error) {
	return []string{
		"claude-3-7-sonnet-20250219",
		"claude-3-5-sonnet-20241022",
		"claude-3-5-haiku-20241022",
	}, nil
}

func (p *AnthropicProvider) GenerateText(prompt string, imagePath string, model string) LLMResponse {
	if p.apiKey == "" {
		return LLMResponse{Error: fmt.Errorf("Claude APIキーが設定されていません")}
	}

	url := "https://api.anthropic.com/v1/messages"

	// リクエスト構造の組み立て
	type imageSource struct {
		Type      string `json:"type"`
		MediaType string `json:"media_type"`
		Data      string `json:"data"`
	}
	type content struct {
		Type   string       `json:"type"`
		Text   string       `json:"text,omitempty"`
		Source *imageSource `json:"source,omitempty"`
	}
	type message struct {
		Role    string    `json:"role"`
		Content []content `json:"content"`
	}
	type anthropicRequest struct {
		Model     string    `json:"model"`
		MaxTokens int       `json:"max_tokens"`
		Messages  []message `json:"messages"`
	}

	contents := []content{
		{Type: "text", Text: prompt},
	}

	if imagePath != "" {
		b64, mime, err := readImageAsBase64(imagePath)
		if err != nil {
			return LLMResponse{Error: fmt.Errorf("画像の読み込みに失敗しました: %v", err)}
		}
		contents = append(contents, content{
			Type: "image",
			Source: &imageSource{
				Type:      "base64",
				MediaType: mime,
				Data:      b64,
			},
		})
	}

	reqBody := anthropicRequest{
		Model:     model,
		MaxTokens: 4000,
		Messages: []message{
			{
				Role:    "user",
				Content: contents,
			},
		},
	}

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return LLMResponse{Error: err}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return LLMResponse{Error: err}
	}
	req.Header.Set("x-api-key", p.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return LLMResponse{Error: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LLMResponse{Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return LLMResponse{Error: fmt.Errorf("anthropic api status %d: %s", resp.StatusCode, string(bodyBytes))}
	}

	// レスポンスのパース
	type anthropicResponse struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	var anthropicResp anthropicResponse
	if err := json.Unmarshal(bodyBytes, &anthropicResp); err != nil {
		return LLMResponse{Error: err}
	}

	if len(anthropicResp.Content) == 0 {
		return LLMResponse{Error: fmt.Errorf("返答テキストが見つかりません")}
	}

	return LLMResponse{Text: anthropicResp.Content[0].Text}
}
