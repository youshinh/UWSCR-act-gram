package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaProvider struct {
	*httpClient
}

func NewOllamaProvider() *OllamaProvider {
	return &OllamaProvider{
		httpClient: newHTTPClient(),
	}
}

type ollamaTagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

func (p *OllamaProvider) GetAvailableModels() ([]string, error) {
	url := "http://localhost:11434/api/tags"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		// Ollamaが起動していない場合はエラーにせず空リストを返す（UIでのエラー防止）
		return []string{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama api status %d: %s", resp.StatusCode, string(body))
	}

	var tagsResp ollamaTagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tagsResp); err != nil {
		return nil, err
	}

	var models []string
	for _, m := range tagsResp.Models {
		models = append(models, m.Name)
	}

	return models, nil
}

func (p *OllamaProvider) GenerateText(prompt string, imagePath string, model string) LLMResponse {
	url := "http://localhost:11434/api/generate"

	// リクエスト構造の組み立て
	type ollamaRequest struct {
		Model  string   `json:"model"`
		Prompt string   `json:"prompt"`
		Stream bool     `json:"stream"`
		Images []string `json:"images,omitempty"`
	}

	var images []string
	if imagePath != "" {
		b64, _, err := readImageAsBase64(imagePath)
		if err != nil {
			return LLMResponse{Error: fmt.Errorf("画像の読み込みに失敗しました: %v", err)}
		}
		images = append(images, b64)
	}

	reqBody := ollamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
		Images: images,
	}

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return LLMResponse{Error: err}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return LLMResponse{Error: err}
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return LLMResponse{Error: fmt.Errorf("Ollamaサーバーへの接続に失敗しました（起動しているか確認してください）: %v", err)}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LLMResponse{Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return LLMResponse{Error: fmt.Errorf("ollama api status %d: %s", resp.StatusCode, string(bodyBytes))}
	}

	// レスポンスのパース
	type ollamaResponse struct {
		Response string `json:"response"`
	}

	var ollamaResp ollamaResponse
	if err := json.Unmarshal(bodyBytes, &ollamaResp); err != nil {
		return LLMResponse{Error: err}
	}

	return LLMResponse{Text: ollamaResp.Response}
}
