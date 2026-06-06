package main

import (
	"act_gram/llm"
	"fmt"
	"io"
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

	fmt.Printf("[Manual Gen] Generated package successfully at: %s\n", baseDir)
	return nil
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
                    <button class="play-btn" onclick="playAudio('audio-step-%d')">🔊 音声ガイドを再生</button>
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
    <title>自律操作マニュアル (act-gram)</title>
    <style>
        body {
            background-color: #0f111a;
            color: #eceff1;
            font-family: 'Helvetica Neue', Arial, 'Hiragino Kaku Gothic ProN', Meiryo, sans-serif;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            padding: 40px 20px;
        }
        header {
            text-align: center;
            margin-bottom: 48px;
            border-bottom: 1px solid rgba(255, 255, 255, 0.08);
            padding-bottom: 24px;
        }
        h1 {
            margin: 0;
            font-size: 2.2rem;
            background: linear-gradient(45deg, #535bf2, #8e44ad);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }
        .meta-info {
            font-size: 0.85rem;
            color: #888;
            margin-top: 8px;
        }
        .step-card {
            background: rgba(255, 255, 255, 0.03);
            border: 1px solid rgba(255, 255, 255, 0.08);
            border-radius: 12px;
            padding: 28px;
            margin-bottom: 32px;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
            transition: border-color 0.3s;
        }
        .step-card:hover {
            border-color: rgba(83, 91, 242, 0.3);
        }
        .step-header {
            display: flex;
            align-items: center;
            gap: 16px;
            margin-bottom: 12px;
        }
        .step-header h2 {
            margin: 0;
            font-size: 1.3rem;
            color: #fff;
        }
        .step-badge {
            background: #535bf2;
            color: #fff;
            font-size: 0.8rem;
            font-weight: bold;
            padding: 4px 10px;
            border-radius: 20px;
        }
        .step-desc {
            font-size: 0.95rem;
            line-height: 1.6;
            color: #ccc;
            margin-bottom: 20px;
        }
        .image-wrapper {
            margin-bottom: 20px;
            border-radius: 8px;
            overflow: hidden;
            border: 1px solid rgba(255, 255, 255, 0.1);
        }
        .image-wrapper img {
            width: 100%%;
            display: block;
        }
        .audio-control {
            display: flex;
            justify-content: flex-start;
        }
        .play-btn {
            background: rgba(83, 91, 242, 0.15);
            color: #7b83ff;
            border: 1px solid rgba(83, 91, 242, 0.3);
            border-radius: 6px;
            padding: 10px 16px;
            font-size: 0.85rem;
            font-weight: bold;
            cursor: pointer;
            transition: background-color 0.2s;
        }
        .play-btn:hover {
            background: rgba(83, 91, 242, 0.25);
            color: #fff;
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>自律操作インタラクティブマニュアル</h1>
            <div class="meta-info">Generated by UWSCR::act-gram / Generated At: %s</div>
        </header>

        <main>
            %s
        </main>
    </div>

    <script>
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
