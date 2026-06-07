package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LocalProvider struct {
	llmType string // "ollama" or "lmstudio"
	baseURL string
	apiKey  string
	*httpClient
}

func NewLocalProvider(llmType string, baseURL string, apiKey string) *LocalProvider {
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
	return &LocalProvider{
		llmType:    llmType,
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: newHTTPClient(),
	}
}

// getHostURL は /v1 や /api/v1 を除いたホストのベースURLを返します。
func (p *LocalProvider) getHostURL() string {
	u := strings.TrimRight(p.baseURL, "/")
	u = strings.TrimSuffix(u, "/api/v1")
	u = strings.TrimSuffix(u, "/v1")
	return u
}

type ollamaModelsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

type openAIModelsResponse struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (p *LocalProvider) GetAvailableModels() ([]string, error) {
	if p.llmType == "ollama" {
		url := fmt.Sprintf("%s/api/tags", p.getHostURL())
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := p.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("Ollamaサーバーに接続できません: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("ollama api status %d: %s", resp.StatusCode, string(body))
		}

		var tagsResp ollamaModelsResponse
		if err := json.NewDecoder(resp.Body).Decode(&tagsResp); err != nil {
			return nil, err
		}

		var models []string
		for _, m := range tagsResp.Models {
			models = append(models, m.Name)
		}
		return models, nil
	}

	// LM Studio 又は OpenAI互換ローカルLLM
	hostURL := p.getHostURL()
	// 推奨の /api/v1/models と互換の /v1/models を試す
	endpoints := []string{
		hostURL + "/api/v1/models",
		hostURL + "/v1/models",
	}

	var lastErr error
	for _, url := range endpoints {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			lastErr = err
			continue
		}
		if p.apiKey != "" {
			req.Header.Set("Authorization", "Bearer "+p.apiKey)
		}

		resp, err := p.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("LM Studio / ローカルLLM サーバーに接続できません (%s): %v", url, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			lastErr = fmt.Errorf("api status %d from %s: %s", resp.StatusCode, url, string(body))
			continue
		}

		var modelsResp openAIModelsResponse
		if err := json.NewDecoder(resp.Body).Decode(&modelsResp); err != nil {
			lastErr = err
			continue
		}

		var models []string
		for _, m := range modelsResp.Data {
			models = append(models, m.ID)
		}
		return models, nil
	}

	return nil, fmt.Errorf("モデル一覧の取得に失敗しました: %v", lastErr)
}

func (p *LocalProvider) GenerateText(prompt string, imagePath string, model string) LLMResponse {
	if p.llmType == "ollama" {
		url := fmt.Sprintf("%s/api/generate", p.getHostURL())

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
			return LLMResponse{Error: fmt.Errorf("Ollamaサーバーへの接続に失敗しました: %v", err)}
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return LLMResponse{Error: err}
		}

		if resp.StatusCode != http.StatusOK {
			return LLMResponse{Error: fmt.Errorf("ollama api status %d: %s", resp.StatusCode, string(bodyBytes))}
		}

		type ollamaResponse struct {
			Response string `json:"response"`
		}

		var ollamaResp ollamaResponse
		if err := json.Unmarshal(bodyBytes, &ollamaResp); err != nil {
			return LLMResponse{Error: err}
		}

		return LLMResponse{Text: ollamaResp.Response}
	}

	// LM Studio / OpenAI互換ローカルLLM
	hostURL := p.getHostURL()
	// /v1/chat/completions または /api/v1/chat/completions を使用
	url := hostURL + "/v1/chat/completions"

	type textContent struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
	type imageURL struct {
		URL string `json:"url"`
	}
	type imageContent struct {
		Type     string   `json:"type"`
		ImageURL imageURL `json:"image_url"`
	}
	type message struct {
		Role    string        `json:"role"`
		Content []interface{} `json:"content"`
	}
	type openAIRequest struct {
		Model    string    `json:"model"`
		Messages []message `json:"messages"`
	}

	var contentItems []interface{}
	contentItems = append(contentItems, textContent{
		Type: "text",
		Text: prompt,
	})

	if imagePath != "" {
		b64, mime, err := readImageAsBase64(imagePath)
		if err != nil {
			return LLMResponse{Error: fmt.Errorf("画像の読み込みに失敗しました: %v", err)}
		}
		contentItems = append(contentItems, imageContent{
			Type: "image_url",
			ImageURL: imageURL{
				URL: fmt.Sprintf("data:%s;base64,%s", mime, b64),
			},
		})
	}

	reqBody := openAIRequest{
		Model: model,
		Messages: []message{
			{
				Role:    "user",
				Content: contentItems,
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
	req.Header.Set("Content-Type", "application/json")
	if p.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+p.apiKey)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		// /v1 がだめなら /api/v1 を試す
		url = hostURL + "/api/v1/chat/completions"
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			return LLMResponse{Error: err}
		}
		req.Header.Set("Content-Type", "application/json")
		if p.apiKey != "" {
			req.Header.Set("Authorization", "Bearer "+p.apiKey)
		}
		resp, err = p.client.Do(req)
		if err != nil {
			return LLMResponse{Error: fmt.Errorf("ローカルLLMサーバーへの接続に失敗しました: %v", err)}
		}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LLMResponse{Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return LLMResponse{Error: fmt.Errorf("local llm api status %d: %s", resp.StatusCode, string(bodyBytes))}
	}

	type openAIResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	var openAIResp openAIResponse
	if err := json.Unmarshal(bodyBytes, &openAIResp); err != nil {
		return LLMResponse{Error: err}
	}

	if len(openAIResp.Choices) == 0 {
		return LLMResponse{Error: fmt.Errorf("返答テキストが見つかりません")}
	}

	return LLMResponse{Text: openAIResp.Choices[0].Message.Content}
}
