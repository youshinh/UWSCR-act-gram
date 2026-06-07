package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

// getHostURL はユーザーが入力したURLから scheme+host+port だけを抽出して返します。
// 例: "http://localhost:1234/api/v1/chat" -> "http://localhost:1234"
func (p *LocalProvider) getHostURL() string {
	u := strings.TrimRight(p.baseURL, "/")
	// net/url.Parse でホスト部分だけを取り出す
	parsed, err := url.Parse(u)
	if err != nil || parsed.Host == "" {
		// パースに失敗した場合は文字列ベースのフォールバック
		u = strings.TrimSuffix(u, "/api/v1")
		u = strings.TrimSuffix(u, "/v1")
		return u
	}
	return parsed.Scheme + "://" + parsed.Host
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

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close() // defer不使用: ループ外でも即時解放
		if err != nil {
			return nil, fmt.Errorf("Ollamaレスポンス読み込み失敗: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("ollama api status %d: %s", resp.StatusCode, string(body))
		}

		var tagsResp ollamaModelsResponse
		if err := json.Unmarshal(body, &tagsResp); err != nil {
			return nil, fmt.Errorf("Ollamaレスポンスパース失敗: %v", err)
		}

		var models []string
		for _, m := range tagsResp.Models {
			models = append(models, m.Name)
		}
		return models, nil
	}

	// LM Studio 又は OpenAI互換ローカルLLM
	hostURL := p.getHostURL()
	// 推奨の /api/v1/models (LM Studio native) と互換の /v1/models (OpenAI互換) を試す
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

		// io.ReadAll で一括読み込み: json.NewDecoder streaming はKeep-Alive接続でブロックする恐れがある
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close() // defer不使用: ループ内では即時Close必須
		if err != nil {
			lastErr = fmt.Errorf("レスポンスボディ読み込み失敗 (%s): %v", url, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("api status %d from %s: %s", resp.StatusCode, url, string(body))
			continue
		}

		var modelsResp openAIModelsResponse
		if err := json.Unmarshal(body, &modelsResp); err != nil {
			lastErr = fmt.Errorf("JSONパース失敗 (%s): %v", url, err)
			continue
		}

		var models []string
		for _, m := range modelsResp.Data {
			models = append(models, m.ID)
		}
		if len(models) > 0 {
			return models, nil
		}
		// modelsが空の場合は次のエンドポイントを試す
		lastErr = fmt.Errorf("モデルリストが空でした (%s)", url)
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
