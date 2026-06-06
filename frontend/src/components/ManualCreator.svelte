<script>
    import { onMount } from 'svelte';
    import { GenerateInteractiveManual } from '../../wailsjs/go/main/App';

    let steps = [
        { step_number: 1, title: 'システムへのログイン', description: 'ブラウザを起動し、ログイン画面でユーザー名とパスワードを入力します。', image_path: '', audio_script: 'まずはログイン画面を開き、割り当てられたIDとパスワードを入力してください。' }
    ];

    let outputPath = 'C:\\'; // デフォルト出力先
    let useHighQualityTTS = true;
    let isGenerating = false;
    
    let statusMessage = '';
    let statusType = ''; // 'success' | 'error'

    function addStep() {
        const nextNum = steps.length + 1;
        steps = [...steps, {
            step_number: nextNum,
            title: `新しいステップ ${nextNum}`,
            description: '',
            image_path: '',
            audio_script: ''
        }];
    }

    function removeStep(index) {
        if (steps.length === 1) return;
        steps = steps.filter((_, i) => i !== index);
        // ステップ番号の再振り分け
        steps = steps.map((s, i) => ({ ...s, step_number: i + 1 }));
    }

    async function handleGenerate() {
        if (!outputPath.trim()) {
            showStatus('出力先フォルダパスを入力してください。', 'error');
            return;
        }

        // 入力値バリデーション
        for (const s of steps) {
            if (!s.title.trim()) {
                showStatus(`ステップ ${s.step_number} のタイトルを入力してください。`, 'error');
                return;
            }
        }

        isGenerating = true;
        showStatus('マニュアルパッケージを生成中... (TTS音声の生成には少し時間がかかります)', 'success');

        try {
            const stepsJSON = JSON.stringify(steps);
            await GenerateInteractiveManual(outputPath, stepsJSON, useHighQualityTTS);
            showStatus('インタラクティブマニュアルのパッケージ生成に成功しました！', 'success');
        } catch (e) {
            showStatus(`生成エラー: ${e.message || e}`, 'error');
        } finally {
            isGenerating = false;
        }
    }

    function showStatus(msg, type) {
        statusMessage = msg;
        statusType = type;
        if (type === 'success') {
            setTimeout(() => {
                if (statusMessage === msg) {
                    statusMessage = '';
                    statusType = '';
                }
            }, 6000);
        }
    }
</script>

<div class="creator-container">
    <div class="header-settings card">
        <h2>📂 パッケージ出力設定</h2>
        <div class="settings-grid">
            <div class="form-group">
                <label for="output-path-input">出力先ディレクトリ</label>
                <input 
                    id="output-path-input" 
                    type="text" 
                    bind:value={outputPath}
                    placeholder="例: C:\Users\Desktop"
                    disabled={isGenerating}
                />
            </div>
            
            <div class="checkbox-group">
                <label class="switch-container">
                    <input type="checkbox" bind:checked={useHighQualityTTS} disabled={isGenerating} />
                    <span class="switch-label">🎙️ Geminiによる高精度音声ガイド (TTS) を生成する</span>
                </label>
            </div>
        </div>
    </div>

    <div class="steps-section">
        <div class="section-header">
            <h2>📝 操作ステップ編集</h2>
            <button class="btn-add" on:click={addStep} disabled={isGenerating}>＋ ステップ追加</button>
        </div>

        <div class="steps-list">
            {#each steps as step, i (step.step_number)}
                <div class="card step-editor-card">
                    <div class="step-card-header">
                        <span class="step-badge">Step {step.step_number}</span>
                        <button class="btn-delete" on:click={() => removeStep(i)} disabled={isGenerating || steps.length === 1}>
                            削除
                        </button>
                    </div>

                    <div class="editor-grid">
                        <div class="form-group">
                            <label for="step-title-{i}">タイトル</label>
                            <input 
                                id="step-title-{i}"
                                type="text"
                                bind:value={step.title}
                                placeholder="例: ファイルの選択"
                                disabled={isGenerating}
                            />
                        </div>

                        <div class="form-group">
                            <label for="step-img-{i}">画面キャプチャ画像パス (ローカル絶対パス)</label>
                            <input 
                                id="step-img-{i}"
                                type="text"
                                bind:value={step.image_path}
                                placeholder="例: C:\temp\screen.png"
                                disabled={isGenerating}
                            />
                        </div>

                        <div class="form-group col-span-2">
                            <label for="step-desc-{i}">ステップの詳しい説明</label>
                            <textarea 
                                id="step-desc-{i}"
                                bind:value={step.description}
                                placeholder="ユーザーへの指示や操作の詳細を記述してください..."
                                rows="3"
                                disabled={isGenerating}
                            ></textarea>
                        </div>

                        <div class="form-group col-span-2">
                            <label for="step-audio-{i}">🔊 音声ガイド読み上げ台本</label>
                            <textarea 
                                id="step-audio-{i}"
                                bind:value={step.audio_script}
                                placeholder="音声合成で読み上げる文章を記述してください（例: ボタンをクリックして保存ダイアログを開きます）"
                                rows="2"
                                disabled={isGenerating}
                            ></textarea>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>

    <div class="footer-actions">
        <button class="btn-generate" on:click={handleGenerate} disabled={isGenerating}>
            {isGenerating ? '生成中...' : '🛠️ インタラクティブマニュアルを生成'}
        </button>
    </div>

    {#if statusMessage}
        <div class="status-overlay {statusType}">
            <p>{statusMessage}</p>
        </div>
    {/if}
</div>

<style>
    .creator-container {
        max-width: 900px;
        margin: 0 auto;
        padding-bottom: 80px;
    }

    .card {
        background: rgba(255, 255, 255, 0.03);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.08);
        border-radius: 12px;
        padding: 24px;
        box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.2);
        margin-bottom: 24px;
    }

    h2 {
        margin-top: 0;
        font-size: 1.15rem;
        color: #fff;
    }

    .settings-grid {
        display: grid;
        grid-template-columns: 1fr;
        gap: 20px;
        align-items: center;
    }

    @media (min-width: 600px) {
        .settings-grid {
            grid-template-columns: 2fr 1.5fr;
        }
    }

    .form-group {
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    .col-span-2 {
        grid-column: span 1;
    }

    @media (min-width: 600px) {
        .col-span-2 {
            grid-column: span 2;
        }
    }

    label {
        font-size: 0.8rem;
        color: #ccc;
    }

    input[type="text"], textarea {
        background: rgba(0, 0, 0, 0.25);
        border: 1px solid rgba(255, 255, 255, 0.12);
        border-radius: 6px;
        padding: 10px 12px;
        color: #fff;
        font-size: 0.85rem;
        box-sizing: border-box;
        width: 100%;
        font-family: inherit;
    }

    input:focus, textarea:focus {
        outline: none;
        border-color: #535bf2;
    }

    .checkbox-group {
        margin-top: 16px;
    }

    .switch-container {
        display: flex;
        align-items: center;
        gap: 10px;
        cursor: pointer;
        user-select: none;
        font-size: 0.85rem;
        color: #ddd;
    }

    .section-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16px;
    }

    .section-header h2 {
        margin: 0;
    }

    .btn-add {
        background: rgba(83, 91, 242, 0.15);
        border: 1px solid rgba(83, 91, 242, 0.3);
        color: #7b83ff;
        border-radius: 6px;
        padding: 8px 16px;
        font-size: 0.8rem;
        font-weight: 600;
        cursor: pointer;
        transition: background-color 0.2s;
    }

    .btn-add:hover {
        background: rgba(83, 91, 242, 0.25);
        color: #fff;
    }

    .step-editor-card {
        border-left: 4px solid #535bf2;
    }

    .step-card-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16px;
    }

    .step-badge {
        font-size: 0.75rem;
        font-weight: bold;
        background: #535bf2;
        color: #fff;
        padding: 4px 10px;
        border-radius: 20px;
    }

    .btn-delete {
        background: transparent;
        border: 1px solid rgba(255, 71, 87, 0.3);
        color: #ff4757;
        font-size: 0.75rem;
        padding: 4px 10px;
        border-radius: 4px;
        cursor: pointer;
        transition: background-color 0.2s;
    }

    .btn-delete:hover:not(:disabled) {
        background: rgba(255, 71, 87, 0.15);
    }

    .editor-grid {
        display: grid;
        grid-template-columns: 1fr;
        gap: 16px;
    }

    @media (min-width: 600px) {
        .editor-grid {
            grid-template-columns: 1fr 1fr;
        }
    }

    .footer-actions {
        display: flex;
        justify-content: flex-end;
        margin-top: 24px;
    }

    .btn-generate {
        background: #535bf2;
        border: none;
        color: #fff;
        border-radius: 6px;
        padding: 12px 24px;
        font-size: 0.95rem;
        font-weight: bold;
        cursor: pointer;
        box-shadow: 0 4px 15px rgba(83, 91, 242, 0.3);
        transition: background-color 0.2s, transform 0.1s;
    }

    .btn-generate:hover:not(:disabled) {
        background: #4047d9;
    }

    .btn-generate:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }

    .status-overlay {
        position: fixed;
        bottom: 24px;
        left: 50%;
        transform: translateX(-50%);
        padding: 12px 24px;
        border-radius: 8px;
        color: #fff;
        font-size: 0.9rem;
        box-shadow: 0 10px 30px rgba(0,0,0,0.5);
        z-index: 1000;
        background: rgba(21, 25, 36, 0.9);
        border: 1px solid rgba(255, 255, 255, 0.1);
        text-align: center;
    }

    .status-overlay.success {
        border-left: 4px solid #2ed573;
        color: #2ed573;
    }

    .status-overlay.error {
        border-left: 4px solid #ff4757;
        color: #ff4757;
    }
</style>
