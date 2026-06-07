package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type OpenAIProvider struct {
	apiKey  string
	baseURL string
	*httpClient
}

func NewOpenAIProvider(apiKey string, baseURL string) *OpenAIProvider {
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	return &OpenAIProvider{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: newHTTPClient(),
	}
}

func (p *OpenAIProvider) GetAvailableModels() ([]string, error) {
	// OpenAI (ChatGPT) で APIキーが空の場合は標準的なモデルリストを返す
	if p.baseURL == "https://api.openai.com/v1" && p.apiKey == "" {
		return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
	}
	
	// /models API を叩きに行く
	url := fmt.Sprintf("%s/models", p.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		if p.baseURL == "https://api.openai.com/v1" {
			return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if p.baseURL == "https://api.openai.com/v1" {
			return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
		}
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("openai compat api error (status %d): %s", resp.StatusCode, string(body))
	}

	type modelList struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	var list modelList
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		if p.baseURL == "https://api.openai.com/v1" {
			return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
		}
		return nil, err
	}

	var models []string
	for _, m := range list.Data {
		if p.baseURL == "https://api.openai.com/v1" {
			id := m.ID
			// テキスト生成用のチャットモデルのみを抽出
			if strings.HasPrefix(id, "gpt-") || strings.HasPrefix(id, "o1-") || strings.HasPrefix(id, "o3-") {
				models = append(models, id)
			}
		} else {
			models = append(models, m.ID)
		}
	}
	
	if len(models) == 0 {
		if p.baseURL == "https://api.openai.com/v1" {
			return []string{"gpt-4o-mini", "gpt-4o", "o1-mini", "o3-mini"}, nil
		}
		return []string{"custom-model"}, nil
	}
	return models, nil
}

func (p *OpenAIProvider) GenerateText(prompt string, imagePath string, model string) LLMResponse {
	if p.apiKey == "" {
		return LLMResponse{Error: fmt.Errorf("APIキーが設定されていません")}
	}

	url := fmt.Sprintf("%s/chat/completions", p.baseURL)

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
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

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
		return LLMResponse{Error: fmt.Errorf("openai compat api status %d: %s", resp.StatusCode, string(bodyBytes))}
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
