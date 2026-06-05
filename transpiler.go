package main

import (
	"fmt"
	"strings"
)

// Transpiler は .uws スクリプトを標準UWSCR互換スクリプトへトランスパイルします。
type Transpiler struct {
	port int
}

func NewTranspiler(port int) *Transpiler {
	return &Transpiler{port: port}
}

// Transpile はスクリプトテキストを受け取り、AI_EVALマクロを置換して返します。
func (t *Transpiler) Transpile(script string) (string, error) {
	var result strings.Builder
	currentIndex := 0

	for {
		// AI_EVAL( の開始位置を探す
		startIndex := strings.Index(script[currentIndex:], "AI_EVAL(")
		if startIndex == -1 {
			// これ以上 AI_EVAL がない場合は残りを書き込んで終了
			result.WriteString(script[currentIndex:])
			break
		}

		// 絶対インデックスに変換
		startIndex += currentIndex
		result.WriteString(script[currentIndex:startIndex])

		// 引数ブロックの解析
		// AI_EVAL( の直後の文字位置から走査開始
		argsStartIndex := startIndex + len("AI_EVAL(")
		depth := 1
		endIndex := -1

		for i := argsStartIndex; i < len(script); i++ {
			char := script[i]
			if char == '(' {
				depth++
			} else if char == ')' {
				depth--
				if depth == 0 {
					endIndex = i
					break
				}
			}
		}

		if endIndex == -1 {
			// 閉じカッコが見つからない場合はエラーにせずそのまま書き込む
			result.WriteString(script[startIndex : startIndex+len("AI_EVAL(")])
			currentIndex = argsStartIndex
			continue
		}

		// 引数文字列を抽出
		argsStr := script[argsStartIndex:endIndex]

		// 引数の分割 (ネストしたカッコ外のカンマで分割)
		var args []string
		var currentArg strings.Builder
		parenDepth := 0
		inQuotes := false

		for i := 0; i < len(argsStr); i++ {
			char := argsStr[i]
			if char == '"' {
				if i == 0 || argsStr[i-1] != '\\' {
					inQuotes = !inQuotes
				}
			}

			if char == ',' && parenDepth == 0 && !inQuotes {
				args = append(args, strings.TrimSpace(currentArg.String()))
				currentArg.Reset()
			} else {
				if char == '(' && !inQuotes {
					parenDepth++
				} else if char == ')' && !inQuotes {
					parenDepth--
				}
				currentArg.WriteByte(char)
			}
		}
		args = append(args, strings.TrimSpace(currentArg.String()))

		// トランスパイルの適用
		if len(args) > 0 {
			prompt := args[0]
			if len(prompt) >= 2 && prompt[0] == '"' && prompt[len(prompt)-1] == '"' {
				prompt = prompt[1 : len(prompt)-1]
			}

			escapedPrompt := strings.ReplaceAll(prompt, `"`, `\"`)
			imageExpr := ""
			if len(args) > 1 {
				imageExpr = args[1]
			}

			if imageExpr != "" {
				// 画像あり (DOSCMD)
				// Windows cmd.exeのパース仕様に合わせ、\" でエスケープしたダブルクォーテーションでJSONを囲む
				transpiled := fmt.Sprintf(
					`DOSCMD("curl.exe -s -X POST http://127.0.0.1:%d/ai_eval -H ""Content-Type: application/json"" -d ""{\""prompt\"":\""%s\"",\""image_path\"":\""\"" + REPLACE(%s, ""\"", ""\\"" ) + ""\""}""", true)`,
					t.port,
					escapedPrompt,
					imageExpr,
				)
				result.WriteString(transpiled)
			} else {
				// 画像なし (DOSCMD)
				transpiled := fmt.Sprintf(
					`DOSCMD("curl.exe -s -X POST http://127.0.0.1:%d/ai_eval -H ""Content-Type: application/json"" -d ""{\""prompt\"":\""%s\""}""", true)`,
					t.port,
					escapedPrompt,
				)
				result.WriteString(transpiled)
			}
		}

		currentIndex = endIndex + 1
	}

	return result.String(), nil
}
