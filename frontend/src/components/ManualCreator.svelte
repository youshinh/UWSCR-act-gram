<script>
    import { onMount } from 'svelte';
    import { GenerateInteractiveManual, GetDefaultManualPath, SelectDirectory, SelectFile } from '../../wailsjs/go/main/App';

    let steps = [
        { step_number: 1, title: 'システムへのログイン', description: 'ブラウザを起動し、ログイン画面でユーザー名とパスワードを入力します。', image_path: '', audio_script: 'まずはログイン画面を開き、割り当てられたIDとパスワードを入力してください。' }
    ];

    let outputPath = ''; 
    let useHighQualityTTS = false; // デフォルトはoff
    let isGenerating = false;
    
    let statusMessage = '';
    let statusType = ''; // 'success' | 'error'

    onMount(async () => {
        try {
            outputPath = await GetDefaultManualPath();
        } catch (e) {
            console.error('Failed to get default manual path:', e);
            outputPath = 'C:\\';
        }
    });

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
            showStatus('通常マニュアル(HTML)および対話型UWSガイドの生成に成功しました！', 'success');
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

    async function browseOutputDir() {
        try {
            const selected = await SelectDirectory("出力先フォルダを選択してください");
            if (selected) {
                outputPath = selected;
            }
        } catch (e) {
            showStatus(`フォルダ選択エラー: ${e.message || e}`, 'error');
        }
    }

    async function browseStepImage(index) {
        try {
            const selected = await SelectFile(
                "画像ファイルを選択", 
                "Image Files (*.png;*.jpg;*.jpeg;*.webp)", 
                "*.png;*.jpg;*.jpeg;*.webp"
            );
            if (selected) {
                steps[index].image_path = selected;
            }
        } catch (e) {
            showStatus(`画像ファイル選択エラー: ${e.message || e}`, 'error');
        }
    }
</script>

<div class="creator-container">
    <div class="header-settings card">
        <div class="card-header">
            <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"/>
            </svg>
            <div class="header-text">
                <h2>パッケージ出力設定</h2>
                <p class="description">生成されるマニュアルHTML、音声アセット、および対話案内用のUWSCRスクリプト (.uws) の保存先を指定します。</p>
            </div>
        </div>
        
        <div class="settings-grid">
            <div class="form-group">
                <label for="output-path-input">出力先ディレクトリ</label>
                <div class="input-with-btn">
                    <input 
                        id="output-path-input" 
                        type="text" 
                        bind:value={outputPath}
                        placeholder="例: C:\Users\Desktop"
                        disabled={isGenerating}
                    />
                    <button class="btn-add btn-browse" on:click={browseOutputDir} disabled={isGenerating}>選択...</button>
                </div>
            </div>
            
            <div class="checkbox-group">
                <label class="switch-container">
                    <input type="checkbox" class="minimal-checkbox" bind:checked={useHighQualityTTS} disabled={isGenerating} />
                    <span class="switch-label">Geminiによる高精度音声ガイド (TTS) を生成する</span>
                </label>
            </div>
        </div>
    </div>

    <div class="steps-section">
        <div class="section-header">
            <div class="section-title">
                <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                    <path d="M18.5 2.5a2.121 2.121 0 1 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
                <h2>操作ステップ編集</h2>
            </div>
            <button class="btn-add" on:click={addStep} disabled={isGenerating}>ステップ追加</button>
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
                            <div class="input-with-btn">
                                <input 
                                    id="step-img-{i}"
                                    type="text"
                                    bind:value={step.image_path}
                                    placeholder="例: C:\temp\screen.png"
                                    disabled={isGenerating}
                                />
                                <button class="btn-add btn-browse" on:click={() => browseStepImage(i)} disabled={isGenerating}>選択...</button>
                            </div>
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
                            <label for="step-audio-{i}">音声ガイド読み上げ台本</label>
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
            {isGenerating ? '生成中...' : 'インタラクティブマニュアルを生成'}
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
        background: var(--bg-secondary);
        backdrop-filter: var(--glass-blur);
        border: 1px solid var(--border-color);
        border-radius: 12px;
        padding: 28px;
        box-shadow: var(--shadow-md);
        margin-bottom: 24px;
        transition: background-color 0.3s ease, border-color 0.3s ease;
    }

    .card-header {
        display: flex;
        gap: 16px;
        align-items: flex-start;
        margin-bottom: 20px;
        border-bottom: 1px solid var(--border-color);
        padding-bottom: 16px;
    }

    .header-icon {
        width: 24px;
        height: 24px;
        color: var(--accent-color);
        flex-shrink: 0;
        margin-top: 2px;
    }

    .header-text h2 {
        margin: 0;
        font-size: 1.1rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .description {
        font-size: 0.8rem;
        color: var(--text-secondary);
        margin: 4px 0 0 0;
        line-height: 1.4;
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
        gap: 8px;
    }

    .input-with-btn {
        display: flex;
        gap: 8px;
        align-items: center;
        width: 100%;
    }

    .btn-browse {
        white-space: nowrap;
        padding: 10px 14px !important;
        border-radius: 8px !important;
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
        color: var(--text-secondary);
        font-weight: 500;
    }

    input[type="text"], textarea {
        background: var(--input-bg);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 10px 14px;
        color: var(--text-primary);
        font-size: 0.85rem;
        box-sizing: border-box;
        width: 100%;
        font-family: inherit;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    input:focus, textarea:focus {
        outline: none;
        border-color: var(--accent-color);
        box-shadow: 0 0 0 2px var(--accent-soft);
    }

    .checkbox-group {
        margin-top: 8px;
    }

    .switch-container {
        display: flex;
        align-items: center;
        gap: 10px;
        cursor: pointer;
        user-select: none;
        font-size: 0.8rem;
        color: var(--text-secondary);
    }

    .minimal-checkbox {
        width: auto !important;
        cursor: pointer;
        margin: 0;
        accent-color: var(--accent-color);
    }

    .section-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 20px;
    }

    .section-title {
        display: flex;
        align-items: center;
        gap: 12px;
    }

    .section-title h2 {
        margin: 0;
        font-size: 1.1rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .btn-add {
        background: var(--accent-soft);
        border: 1px solid var(--accent-border);
        color: var(--accent-color);
        border-radius: 8px;
        padding: 8px 16px;
        font-size: 0.8rem;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .btn-add:hover {
        background: var(--accent-color);
        color: var(--bg-primary);
    }

    .step-editor-card {
        border-left: 3px solid var(--accent-color) !important;
    }

    .step-card-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 20px;
        border-bottom: 1px solid var(--border-color);
        padding-bottom: 12px;
    }

    .step-badge {
        font-size: 0.75rem;
        font-weight: 500;
        background: var(--accent-soft);
        color: var(--accent-color);
        border: 1px solid var(--accent-border);
        padding: 4px 10px;
        border-radius: 20px;
    }

    .btn-delete {
        background: transparent;
        border: 1px solid var(--accent-red-border);
        color: var(--accent-red);
        font-size: 0.75rem;
        padding: 4px 12px;
        border-radius: 6px;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .btn-delete:hover:not(:disabled) {
        background: var(--accent-red-soft);
        border-color: var(--accent-red);
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
        background: var(--accent-color);
        border: none;
        color: var(--bg-primary);
        border-radius: 8px;
        padding: 12px 24px;
        font-size: 0.85rem;
        font-weight: 600;
        cursor: pointer;
        box-shadow: var(--shadow-sm);
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .btn-generate:hover:not(:disabled) {
        background: var(--accent-hover);
    }

    .btn-generate:active:not(:disabled) {
        transform: scale(0.97);
    }

    .btn-generate:disabled {
        opacity: 0.4;
        cursor: not-allowed;
    }

    .status-overlay {
        position: fixed;
        bottom: 24px;
        left: 50%;
        transform: translateX(-50%);
        padding: 12px 24px;
        border-radius: 8px;
        color: var(--text-primary);
        font-size: 0.85rem;
        box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
        z-index: 1000;
        background: var(--bg-secondary);
        backdrop-filter: var(--glass-blur);
        border: 1px solid var(--border-color);
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
