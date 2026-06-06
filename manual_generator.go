package main

import (
	"act_gram/llm"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ManualStep はマニュアルの各ステップ情報を表します。
type ManualStep struct {
	StepNumber  int    `json:"step_number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	AudioScript string `json:"audio_script"`
}

// GenerateManual はインタラクティブマニュアル（HTML + 画像 + 音声）をフォルダパッケージとして出力します。
func GenerateManual(outputPath string, steps []ManualStep, useHighQualityTTS bool) error {
	// 1. 出力先ベースディレクトリの作成 (manual_YYYYMMDD_HHMMSS)
	timestamp := time.Now().Format("20060102_150405")
	manualDirName := fmt.Sprintf("manual_%s", timestamp)
	baseDir := filepath.Join(outputPath, manualDirName)

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return fmt.Errorf("マニュアル用ディレクトリの作成に失敗しました: %v", err)
	}

	imagesDir := filepath.Join(baseDir, "images")
	if err := os.MkdirAll(imagesDir, 0755); err != nil {
		return err
	}

	audioDir := filepath.Join(baseDir, "audio")
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return err
	}

	// APIキーの取得 (TTS用)
	var apiKey string
	if useHighQualityTTS {
		var err error
		apiKey, err = GetAPIKey("google")
		if err != nil {
			return fmt.Errorf("Google APIキーの取得に失敗しました: %v", err)
		}
		if apiKey == "" {
			return fmt.Errorf("高精度TTSの実行にはGoogle (Gemini) のAPIキー設定が必要です。")
		}
	}

	// 2. 各ステップの処理 (画像コピー & 音声生成)
	for i, step := range steps {
		// 画像コピー
		if step.ImagePath != "" {
			srcFile, err := os.Open(step.ImagePath)
			if err == nil {
				defer srcFile.Close()
				ext := filepath.Ext(step.ImagePath)
				if ext == "" {
					ext = ".png"
				}
				destPath := filepath.Join(imagesDir, fmt.Sprintf("step_%d%s", step.StepNumber, ext))
				destFile, err := os.Create(destPath)
				if err == nil {
					defer destFile.Close()
					io.Copy(destFile, srcFile)
					// HTMLから参照するための相対パスを更新
					steps[i].ImagePath = fmt.Sprintf("images/step_%d%s", step.StepNumber, ext)
				}
			}
		}

		// 音声生成 (TTS)
		if step.AudioScript != "" && useHighQualityTTS {
			audioData, err := llm.GenerateSpeech(apiKey, step.AudioScript, "Zephyr")
			if err == nil {
				destAudioPath := filepath.Join(audioDir, fmt.Sprintf("step_%d.wav", step.StepNumber))
				err = os.WriteFile(destAudioPath, audioData, 0644)
				if err == nil {
					// HTMLから参照するための相対パスを更新
					steps[i].AudioScript = fmt.Sprintf("audio/step_%d.wav", step.StepNumber)
				} else {
					fmt.Printf("[Manual Gen] Audio write error: %v\n", err)
					steps[i].AudioScript = "" // 失敗時は空にして再生ボタンを非表示にする
				}
			} else {
				fmt.Printf("[Manual Gen] TTS error on step %d: %v\n", step.StepNumber, err)
				steps[i].AudioScript = ""
			}
		} else {
			steps[i].AudioScript = "" // TTSを使わない、またはスクリプトが空の場合
		}
	}

	// 3. index.html の生成
	htmlContent := buildHTMLContent(steps)
	htmlPath := filepath.Join(baseDir, "index.html")
	if err := os.WriteFile(htmlPath, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("index.html の生成に失敗しました: %v", err)
	}

	// 4. インタラクティブマニュアル（UWSスクリプト）の生成と同梱
	fmt.Println("[Manual Gen] Generating interactive guide UWS script...")
	uwsContent, err := GenerateInteractiveScript(steps)
	if err != nil {
		fmt.Printf("[Manual Gen] Warning: Failed to generate interactive guide .uws: %v\n", err)
	} else {
		uwsPath := filepath.Join(baseDir, "interactive_guide.uws")
		if err := os.WriteFile(uwsPath, []byte(uwsContent), 0644); err != nil {
			fmt.Printf("[Manual Gen] Warning: Failed to write interactive guide file: %v\n", err)
		} else {
			fmt.Printf("[Manual Gen] Embedded interactive guide successfully at: %s\n", uwsPath)
		}
	}

	fmt.Printf("[Manual Gen] Generated package successfully at: %s\n", baseDir)
	return nil
}

// GenerateInteractiveScript はマニュアルのステップを元に、対話的にユーザーを案内する UWSCR (.uws) スクリプトをAIで生成します。
func GenerateInteractiveScript(steps []ManualStep) (string, error) {
	// 1. 設定の読み込みとAIプロバイダーの特定
	cfg, err := LoadConfig()
	if err != nil {
		return "", fmt.Errorf("設定のロードに失敗しました: %v", err)
	}

	brainConfig, ok := cfg.Layers["brain"]
	if !ok || brainConfig.Provider == "" || brainConfig.Model == "" {
		brainConfig = LayerConfig{
			Provider: "google",
			Model:    "gemini-flash-lite-latest",
		}
	}

	provider := brainConfig.Provider
	model := brainConfig.Model
	log.Printf("[Manual Gen] Selected AI for guide: %s (%s)", model, provider)

	apiKey, err := GetAPIKey(provider)
	if err != nil && provider != "ollama" {
		return "", fmt.Errorf("APIキーの取得に失敗しました: %v", err)
	}

	var llmProvider llm.LLMProvider
	switch provider {
	case "google":
		llmProvider = llm.NewGeminiProvider(apiKey)
	case "anthropic":
		llmProvider = llm.NewAnthropicProvider(apiKey)
	case "ollama":
		llmProvider = llm.NewOllamaProvider()
	default:
		return "", fmt.Errorf("未サポートのプロバイダーです: %s", provider)
	}

	// 2. ステップ情報のテキスト化
	var stepsText strings.Builder
	for _, s := range steps {
		stepsText.WriteString(fmt.Sprintf("Step %d: %s\n説明: %s\n音声ガイド: %s\n\n", s.StepNumber, s.Title, s.Description, s.AudioScript))
	}

	// 3. プロンプトの構築
	systemPrompt := `あなたは卓越したRPAスクリプト開発AIです。
入力されたマニュアルのステップ情報を元に、PC上でユーザーと対話しつつ、自動操作を行う UWSCR (UWSC互換) スクリプト (.uws) を作成してください。

【UWSCR インタラクティブ・ガイドスクリプトの書き方】
1. 各ステップの開始時に MSGBOX("ステップ番号: タイトル\n\n説明文", BTN_OK) を用いてユーザーに操作手順を表示し、OKが押されるのを待ちます。
2. 自動で操作できる部分がある場合は、MSGBOXの後（または前）に UWSCR の自動操作コマンド (ACW, CLK, KBD, SLEEP) を実行します。
3. もし、ユーザーによる手動操作を待つ必要がある場合は、MSGBOX で「〇〇の操作を行ってからOKを押してください」と案内します。
4. act-gramの AI_EVALマクロを組み込んで、ユーザーの操作が成功したかを簡易確認する判定を入れることも可能です。
   例:
   IFB AI_EVAL("ログイン成功画面になっていますか？", GetScreenCapture()) <> "はい" THEN
       MSGBOX("ログインが確認できません。手順をやり直してください。")
       EXIT
   ENDIF

【出力要件】
- スクリプトコードのみを出力してください。
- markdown のコードブロック ('''uws ... ''') などで囲まず、純粋なプレーンテキストとして UWSCR スクリプトのみを返してください。説明文や前置きは一切不要です。
`

	finalPrompt := fmt.Sprintf("%s\n【入力マニュアルステップ】\n%s", systemPrompt, stepsText.String())

	// 4. LLMへの問い合わせ
	resp := llmProvider.GenerateText(finalPrompt, "", model)
	if resp.Error != nil {
		return "", fmt.Errorf("AIでのインタラクティブスクリプト生成に失敗しました: %v", resp.Error)
	}

	cleanCode := cleanGeneratedCode(resp.Text)
	return cleanCode, nil
}

// 可視化用の美しい HTML テンプレートを組み立てます。
func buildHTMLContent(steps []ManualStep) string {
	var stepsHTML strings.Builder

	for _, step := range steps {
		audioHTML := ""
		if step.AudioScript != "" {
			audioHTML = fmt.Sprintf(`
                <div class="audio-control">
                    <audio id="audio-step-%d" src="%s"></audio>
                    <button class="play-btn" onclick="playAudio('audio-step-%d')">
                        <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right:6px;"><polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon><path d="M19.07 4.93a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"></path></svg>
                        音声ガイドを再生
                    </button>
                </div>`, step.StepNumber, step.AudioScript, step.StepNumber)
		}

		imageHTML := ""
		if step.ImagePath != "" {
			imageHTML = fmt.Sprintf(`
                <div class="image-wrapper">
                    <img src="%s" alt="ステップ %d の画面キャプチャ" />
                </div>`, step.ImagePath, step.StepNumber)
		}

		stepsHTML.WriteString(fmt.Sprintf(`
        <div class="step-card">
            <div class="step-header">
                <span class="step-badge">Step %d</span>
                <h2>%s</h2>
            </div>
            <p class="step-desc">%s</p>
            %s
            %s
        </div>`, step.StepNumber, step.Title, step.Description, imageHTML, audioHTML))
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>操作プロセス指示書 (act-gram)</title>
    <style>
        :root {
            --bg-color: #ffffff;
            --text-color: #0f172a;
            --text-secondary: #475569;
            --card-bg: #f8fafc;
            --border-color: #e2e8f0;
            --accent-color: #0f172a;
            --accent-text: #ffffff;
            --btn-bg: rgba(15, 23, 42, 0.04);
            --btn-border: #cbd5e1;
            --btn-text: #0f172a;
            --shadow-color: rgba(0, 0, 0, 0.04);
        }

        html.dark-mode {
            --bg-color: #0f172a;
            --text-color: #f1f5f9;
            --text-secondary: #94a3b8;
            --card-bg: #1e293b;
            --border-color: #334155;
            --accent-color: #f1f5f9;
            --accent-text: #0f172a;
            --btn-bg: rgba(241, 245, 249, 0.08);
            --btn-border: #475569;
            --btn-text: #f1f5f9;
            --shadow-color: rgba(0, 0, 0, 0.3);
        }

        body {
            background-color: var(--bg-color);
            color: var(--text-color);
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
            margin: 0;
            padding: 0;
            transition: background-color 0.2s ease, color 0.2s ease;
        }

        .container {
            max-width: 800px;
            margin: 0 auto;
            padding: 48px 24px;
        }

        header {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin-bottom: 40px;
            border-bottom: 1px solid var(--border-color);
            padding-bottom: 24px;
        }

        .header-left h1 {
            margin: 0;
            font-size: 1.6rem;
            font-weight: 700;
            letter-spacing: -0.5px;
        }

        .meta-info {
            font-size: 0.8rem;
            color: var(--text-secondary);
            margin-top: 6px;
        }

        .theme-toggle-btn {
            background: var(--btn-bg);
            border: 1px solid var(--btn-border);
            color: var(--btn-text);
            border-radius: 6px;
            padding: 8px 12px;
            font-size: 0.75rem;
            font-weight: 600;
            cursor: pointer;
            display: flex;
            align-items: center;
            gap: 6px;
            transition: all 0.2s ease;
        }

        .theme-toggle-btn:hover {
            border-color: var(--text-color);
            background: rgba(128, 128, 128, 0.08);
        }

        .step-card {
            background: var(--card-bg);
            border: 1px solid var(--border-color);
            border-radius: 8px;
            padding: 24px;
            margin-bottom: 24px;
            box-shadow: 0 4px 12px var(--shadow-color);
            transition: transform 0.2s ease, border-color 0.2s ease;
        }

        .step-card:hover {
            border-color: var(--text-color);
        }

        .step-header {
            display: flex;
            align-items: center;
            gap: 12px;
            margin-bottom: 12px;
        }

        .step-header h2 {
            margin: 0;
            font-size: 1.15rem;
            font-weight: 600;
        }

        .step-badge {
            background: var(--accent-color);
            color: var(--accent-text);
            font-size: 0.75rem;
            font-weight: 700;
            padding: 3px 8px;
            border-radius: 4px;
        }

        .step-desc {
            font-size: 0.9rem;
            line-height: 1.55;
            color: var(--text-secondary);
            margin: 0 0 20px 0;
        }

        .image-wrapper {
            margin-bottom: 20px;
            border-radius: 6px;
            overflow: hidden;
            border: 1px solid var(--border-color);
            background: #000;
            max-height: 400px;
            display: flex;
            justify-content: center;
        }

        .image-wrapper img {
            max-width: 100%%;
            height: auto;
            object-fit: contain;
            display: block;
        }

        .audio-control {
            display: flex;
            justify-content: flex-start;
        }

        .play-btn {
            background: var(--btn-bg);
            border: 1px solid var(--btn-border);
            color: var(--btn-text);
            border-radius: 6px;
            padding: 8px 14px;
            font-size: 0.8rem;
            font-weight: 600;
            cursor: pointer;
            display: inline-flex;
            align-items: center;
            transition: all 0.2s ease;
        }

        .play-btn:hover {
            border-color: var(--text-color);
            background: rgba(128, 128, 128, 0.08);
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div class="header-left">
                <h1>操作プロセス指示書</h1>
                <div class="meta-info">作成元: UWSCR::act-gram / 作成日時: %s</div>
            </div>
            <button class="theme-toggle-btn" onclick="toggleTheme()" aria-label="Toggle dark mode">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"></circle><line x1="12" y1="1" x2="12" y2="3"></line><line x1="12" y1="21" x2="12" y2="23"></line><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line><line x1="1" y1="12" x2="3" y2="12"></line><line x1="21" y1="12" x2="23" y2="12"></line><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line></svg>
                テーマ切り替え
            </button>
        </header>

        <main>
            %s
        </main>
    </div>

    <script>
        // 初期テーマ適用 (localStorage またはシステムのデフォルト)
        if (localStorage.getItem('theme') === 'dark') {
            document.documentElement.classList.add('dark-mode');
        } else if (localStorage.getItem('theme') === 'light') {
            document.documentElement.classList.remove('dark-mode');
        } else if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
            // システム初期設定に従う
            document.documentElement.classList.add('dark-mode');
        }

        function toggleTheme() {
            const isDark = document.documentElement.classList.toggle('dark-mode');
            localStorage.setItem('theme', isDark ? 'dark' : 'light');
        }

        function playAudio(audioId) {
            const audios = document.querySelectorAll('audio');
            audios.forEach(a => {
                if (a.id !== audioId) {
                    a.pause();
                    a.currentTime = 0;
                }
            });
            const audio = document.getElementById(audioId);
            if (audio) {
                audio.play();
            }
        }
    </script>
</body>
</html>`, time.Now().Format("2006-01-02 15:04:05"), stepsHTML.String())
}
