package llm

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// GenerateSpeech は Gemini 3.1 TTS API を使用してテキストから高精度な音声データ(WAV)を生成します。
func GenerateSpeech(apiKey string, text string, voiceName string) ([]byte, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("Gemini APIキーが設定されていません")
	}

	if voiceName == "" {
		voiceName = "Zephyr" // デフォルトの音声名
	}

	model := "gemini-3.1-flash-tts-preview"
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

	// APIリクエストの組み立て
	type textPart struct {
		Text string `json:"text"`
	}
	type content struct {
		Role  string     `json:"role"`
		Parts []textPart `json:"parts"`
	}
	type prebuiltVoice struct {
		VoiceName string `json:"voice_name"`
	}
	type voiceConfig struct {
		PrebuiltVoiceConfig prebuiltVoice `json:"prebuilt_voice_config"`
	}
	type speechConfig struct {
		VoiceConfig voiceConfig `json:"voice_config"`
	}
	type genConfig struct {
		ResponseModalities []string     `json:"responseModalities"`
		Temperature        float64      `json:"temperature"`
		SpeechConfig       speechConfig `json:"speech_config"`
	}
	type requestBody struct {
		Contents         []content `json:"contents"`
		GenerationConfig genConfig `json:"generationConfig"`
	}

	reqBody := requestBody{
		Contents: []content{
			{
				Role: "user",
				Parts: []textPart{
					{Text: fmt.Sprintf("## Transcript:\n%s", text)},
				},
			},
		},
		GenerationConfig: genConfig{
			ResponseModalities: []string{"audio"},
			Temperature:        1.0,
			SpeechConfig: speechConfig{
				VoiceConfig: voiceConfig{
					PrebuiltVoiceConfig: prebuiltVoice{
						VoiceName: voiceName,
					},
				},
			},
		},
	}

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gemini TTS API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	// レスポンスのパース
	type inlineData struct {
		MimeType string `json:"mimeType"`
		Data     string `json:"data"`
	}
	type part struct {
		Text       string      `json:"text,omitempty"`
		InlineData *inlineData `json:"inlineData,omitempty"`
	}
	type candidate struct {
		Content struct {
			Parts []part `json:"parts"`
		} `json:"content"`
	}
	type responseBody struct {
		Candidates []candidate `json:"candidates"`
	}

	var geminiResp responseBody
	if err := json.Unmarshal(bodyBytes, &geminiResp); err != nil {
		return nil, err
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("音声データがレスポンスに含まれていません")
	}

	audioPart := geminiResp.Candidates[0].Content.Parts[0]
	if audioPart.InlineData == nil {
		return nil, fmt.Errorf("インラインデータがありません。テキスト返答: %s", audioPart.Text)
	}

	rawData, err := base64.StdEncoding.DecodeString(audioPart.InlineData.Data)
	if err != nil {
		return nil, fmt.Errorf("音声データのデコードに失敗しました: %v", err)
	}

	mimeType := audioPart.InlineData.MimeType

	// MIMEタイプが audio/wav などの完成されたコンテナであればそのまま返す
	if strings.Contains(strings.ToLower(mimeType), "audio/wav") || strings.Contains(strings.ToLower(mimeType), "audio/wave") {
		return rawData, nil
	}

	// PCM (L16) 生データの場合はWAVヘッダーを構築して結合
	sampleRate, numChannels, bitsPerSample := parseMimeType(mimeType)
	wavHeader := createWavHeader(len(rawData), sampleRate, numChannels, bitsPerSample)

	// ヘッダーと生PCMデータをマージ
	wavData := append(wavHeader, rawData...)
	return wavData, nil
}

// MIMEタイプからPCMオーディオ設定をパースします
func parseMimeType(mimeType string) (sampleRate int, numChannels int, bitsPerSample int) {
	// デフォルト値 (24kHz Mono 16bit)
	sampleRate = 24000
	numChannels = 1
	bitsPerSample = 16

	parts := strings.Split(mimeType, ";")
	if len(parts) == 0 {
		return
	}

	formatPart := strings.TrimSpace(parts[0])
	subParts := strings.Split(formatPart, "/")
	if len(subParts) > 1 {
		format := strings.ToLower(subParts[1])
		// x-l16 や l16 のような形式からビット深度を抽出
		if strings.Contains(format, "l") {
			idx := strings.Index(format, "l")
			bitsStr := format[idx+1:]
			var bits int
			if _, err := fmt.Sscanf(bitsStr, "%d", &bits); err == nil {
				bitsPerSample = bits
			}
		}
	}

	// パラメータのパース (例: rate=24000)
	for _, p := range parts[1:] {
		p = strings.TrimSpace(p)
		kv := strings.Split(p, "=")
		if len(kv) == 2 {
			key := strings.ToLower(strings.TrimSpace(kv[0]))
			value := strings.TrimSpace(kv[1])
			if key == "rate" {
				var r int
				if _, err := fmt.Sscanf(value, "%d", &r); err == nil {
					sampleRate = r
				}
			}
		}
	}

	return
}

// WAV(RIFF)ヘッダーを作成します
func createWavHeader(dataLength int, sampleRate int, numChannels int, bitsPerSample int) []byte {
	header := make([]byte, 44)

	// ChunkID: "RIFF"
	copy(header[0:4], "RIFF")
	// ChunkSize: 36 + Subchunk2Size
	binary.LittleEndian.PutUint32(header[4:8], uint32(36+dataLength))
	// Format: "WAVE"
	copy(header[8:12], "WAVE")
	// Subchunk1ID: "fmt "
	copy(header[12:16], "fmt ")
	// Subchunk1Size: 16 (for PCM)
	binary.LittleEndian.PutUint32(header[16:20], 16)
	// AudioFormat: 1 (for PCM)
	binary.LittleEndian.PutUint16(header[20:22], 1)
	// NumChannels
	binary.LittleEndian.PutUint16(header[22:24], uint16(numChannels))
	// SampleRate
	binary.LittleEndian.PutUint32(header[24:28], uint32(sampleRate))
	// ByteRate: SampleRate * NumChannels * BitsPerSample / 8
	byteRate := sampleRate * numChannels * bitsPerSample / 8
	binary.LittleEndian.PutUint32(header[28:32], uint32(byteRate))
	// BlockAlign: NumChannels * BitsPerSample / 8
	blockAlign := numChannels * bitsPerSample / 8
	binary.LittleEndian.PutUint16(header[32:34], uint16(blockAlign))
	// BitsPerSample
	binary.LittleEndian.PutUint16(header[34:36], uint16(bitsPerSample))
	// Subchunk2ID: "data"
	copy(header[36:40], "data")
	// Subchunk2Size: dataLength
	binary.LittleEndian.PutUint32(header[40:44], uint32(dataLength))

	return header
}
