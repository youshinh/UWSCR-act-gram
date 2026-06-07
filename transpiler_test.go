package main

import (
	"strings"
	"testing"
)

func TestTranspile(t *testing.T) {
	transpiler := NewTranspiler(31415)

	tests := []struct {
		name     string
		input    string
		expected []string // 期待される部分文字列
	}{
		{
			name:  "Text only AI_EVAL",
			input: `Dim res = AI_EVAL("この伝票の合計金額は？")`,
			expected: []string{
				`DOSCMD(`,
				`curl.exe -s -X POST http://127.0.0.1:31415/ai_eval`,
				`-d "{\"prompt\":\"この伝票の合計金額は？\"}"`,
			},
		},
		{
			name:  "AI_EVAL with image capture function",
			input: `Dim res = AI_EVAL("金額を読み取って", GetScreenCapture())`,
			expected: []string{
				`DOSCMD(`,
				`curl.exe -s -X POST http://127.0.0.1:31415/ai_eval`,
				`-d "{\"prompt\":\"金額を読み取って\",\"image_path\":\"' + REPLACE(GetScreenCapture(), '\\', '\\\\') + '\"}"`,
			},
		},
		{
			name:  "AI_EVAL with path variable",
			input: `Dim res = AI_EVAL("解析して", img_path)`,
			expected: []string{
				`+ REPLACE(img_path, '\\', '\\\\')`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := transpiler.Transpile(tt.input)
			if err != nil {
				t.Fatalf("Transpile failed: %v", err)
			}

			for _, exp := range tt.expected {
				if !strings.Contains(output, exp) {
					t.Errorf("Expected output to contain:\n%s\n\nGot:\n%s", exp, output)
				}
			}
		})
	}
}
