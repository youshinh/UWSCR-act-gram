package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeminiProvider struct {
	apiKey string
	*httpClient
}

func NewGeminiProvider(apiKey string) *GeminiProvider {
	return &GeminiProvider{
		apiKey:     apiKey,
		httpClient: newHTTPClient(),
	}
}

type geminiModelList struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

func (p *GeminiProvider) GetAvailableModels() ([]string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models?key=%s", p.apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("gemini api error (status %d): %s", resp.StatusCode, string(body))
	}

	var modelList geminiModelList
	if err := json.NewDecoder(resp.Body).Decode(&modelList); err != nil {
		return nil, err
	}

	var models []string
	for _, m := range modelList.Models {
		// "models/gemini-2.5-flash" -> "gemini-2.5-flash" のようにプレフィックスをトリムして使いやすくする
		name := m.Name
		if len(name) > 7 && name[:7] == "models/" {
			name = name[7:]
		}
		models = append(models, name)
	}

	return models, nil
}

func (p *GeminiProvider) GenerateText(prompt string, imagePath string, model string) LLMResponse {
	if p.apiKey == "" {
		return LLMResponse{Error: fmt.Errorf("Gemini APIキーが設定されていません")}
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, p.apiKey)

	// リクエスト構造の組み立て
	type inlineData struct {
		MimeType string `json:"mimeType"`
		Data     string `json:"data"`
	}
	type part struct {
		Text       string      `json:"text,omitempty"`
		InlineData *inlineData `json:"inlineData,omitempty"`
	}
	type content struct {
		Parts []part `json:"parts"`
	}
	type geminiRequest struct {
		Contents []content `json:"contents"`
	}

	parts := []part{{Text: prompt}}

	// 画像の追加処理
	if imagePath != "" {
		b64, mime, err := readImageAsBase64(imagePath)
		if err != nil {
			return LLMResponse{Error: fmt.Errorf("画像の読み込みに失敗しました: %v", err)}
		}
		parts = append(parts, part{
			InlineData: &inlineData{
				MimeType: mime,
				Data:     b64,
			},
		})
	}

	reqBody := geminiRequest{
		Contents: []content{{Parts: parts}},
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
		return LLMResponse{Error: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LLMResponse{Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return LLMResponse{Error: fmt.Errorf("gemini api status %d: %s", resp.StatusCode, string(bodyBytes))}
	}

	// レスポンスパース
	type geminiResponse struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	var geminiResp geminiResponse
	if err := json.Unmarshal(bodyBytes, &geminiResp); err != nil {
		return LLMResponse{Error: err}
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return LLMResponse{Error: fmt.Errorf("返答テキストが見つかりません")}
	}

	return LLMResponse{Text: geminiResp.Candidates[0].Content.Parts[0].Text}
}
