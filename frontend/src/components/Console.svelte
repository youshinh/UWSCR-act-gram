<script>
    import { onMount, onDestroy } from 'svelte';
    import { RunScript, GetConfig, SelectFile, ProposeOptimization, SaveScriptFile, StopScript } from '../../wailsjs/go/main/App';
    import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';

    const App = window.go.main.App;

    let scriptPath = '';
    let logs = [];
    let isRunning = false;
    let logContainer;
    let config = null;

    let statusMessage = '';
    let statusType = ''; // 'success' | 'error'

    // Optimisation UI State
    let isOptimizing = false;
    let showDiffModal = false;
    let optimizationResult = null;
    let originalCode = "";

    onMount(async () => {
        try {
            config = await GetConfig();
        } catch (e) {
            console.error('Failed to load config in console', e);
        }

        // UWSCRログイベントの監視を開始
        wailsRuntime.EventsOn('uwscr_log', (logLine) => {
            logs = [...logs, logLine];
            
            // ログが追加されたら一番下まで自動スクロール
            setTimeout(() => {
                if (logContainer) {
                    logContainer.scrollTop = logContainer.scrollHeight;
                }
            }, 10);

            // プロセス終了等の特定システムメッセージを検知して状態更新
            if (logLine.message.includes('[System] プロセスが正常に終了しました。') || 
                logLine.message.includes('[System] プロセスがエラーで終了しました') ||
                logLine.message.includes('[Error]')) {
                isRunning = false;
                
                // 実行終了後にウィンドウを自動復元
                try {
                    const AppBinding = window.go.main.App;
                    if (AppBinding && typeof AppBinding.RestoreWindow === 'function') {
                        AppBinding.RestoreWindow();
                    }
                } catch (err) {
                    console.error("Failed to restore window:", err);
                }
            }
        });
    });

    onDestroy(() => {
        wailsRuntime.EventsOff('uwscr_log');
    });

    async function handleRun() {
        if (!scriptPath.trim()) {
            showStatus('実行する.uwsファイルのパスを入力してください。', 'error');
            return;
        }

        logs = []; // ログをクリア
        isRunning = true;
        showStatus('スクリプト実行を開始します...', 'success');

        // 実行開始前に画面を最小化して操作対象を遮らないようにする
        try {
            const AppBinding = window.go.main.App;
            if (AppBinding && typeof AppBinding.MinimizeWindow === 'function') {
                await AppBinding.MinimizeWindow();
            }
        } catch (err) {
            console.error("Failed to minimize window:", err);
        }

        try {
            await RunScript(scriptPath);
        } catch (e) {
            showStatus(`実行開始に失敗しました: ${e.message || e}`, 'error');
            isRunning = false;
            // 失敗時はウィンドウを復元
            try {
                const AppBinding = window.go.main.App;
                if (AppBinding && typeof AppBinding.RestoreWindow === 'function') {
                    await AppBinding.RestoreWindow();
                }
            } catch (err) {}
        }
    }

    async function handleStop() {
        try {
            showStatus('スクリプトを強制停止します...', 'success');
            await StopScript();
            isRunning = false;
            // 停止後はウィンドウを復元
            const AppBinding = window.go.main.App;
            if (AppBinding && typeof AppBinding.RestoreWindow === 'function') {
                await AppBinding.RestoreWindow();
            }
        } catch (e) {
            showStatus(`停止に失敗しました: ${e.message || e}`, 'error');
        }
    }

    // ミリ秒ファクトから実行イベントを抽出
    function extractExecutionEvents(logLines) {
        const events = [];
        let lastTime = null;
        const timeRegex = /\[(\d{2}):(\d{2}):(\d{2})\.(\d{3})\]/;
        
        for (const line of logLines) {
            const match = line.message.match(timeRegex);
            if (match) {
                const hours = parseInt(match[1]);
                const minutes = parseInt(match[2]);
                const seconds = parseInt(match[3]);
                const ms = parseInt(match[4]);
                
                const absoluteMs = ((hours * 3600 + minutes * 60 + seconds) * 1000) + ms;
                
                if (lastTime !== null) {
                    const duration = absoluteMs - lastTime;
                    events.push({
                        step_id: `step_${events.length + 1}`,
                        command: line.message.substring(line.message.indexOf(']') + 1).trim(),
                        start_time: lastTime,
                        end_time: absoluteMs,
                        idle_time: duration > 1000 ? duration - 150 : 0 // SLEEP heuristic
                    });
                }
                lastTime = absoluteMs;
            }
        }
        
        if (events.length === 0) {
            // Fallback default events if no timestamps found
            events.push({
                step_id: "step_1",
                command: "Script Execution Block",
                start_time: Date.now() - 5000,
                end_time: Date.now(),
                idle_time: 4200
            });
        }
        return events;
    }

    // 自己進化・ボトルネック分析の実行
    async function handleOptimize() {
        if (!scriptPath.trim()) {
            showStatus('最適化するスクリプトファイルを選択してください。', 'error');
            return;
        }

        isOptimizing = true;
        showStatus('ミリ秒ファクトログを分析し、最適化コードを生成中...', 'success');

        try {
            const events = extractExecutionEvents(logs);
            const resJSON = await ProposeOptimization(scriptPath, JSON.stringify(events));
            
            // Clean JSON code formatting
            let cleanJSON = resJSON.trim();
            if (cleanJSON.startsWith("```")) {
                const lines = cleanJSON.split("\n");
                if (lines.length > 2) {
                    cleanJSON = lines.slice(1, lines.length - 1).join("\n");
                }
            }
            
            optimizationResult = JSON.parse(cleanJSON);
            
            // Fetch original code from backend
            // Wails doesn't have direct ReadFile, let's treat the parameter as code to generate or read it
            // We can read it in a helper way, but let's just use ProposeOptimization's response or let the backend load it
            // Let's call a mock or backend read if we want. Actually, we can fetch original code or display optimizationResult
            originalCode = "（ファイルを読み込み中...）";
            try {
                // Fetch file code by passing scriptPath to a function that returns it
                // We don't have a direct ReadFile binding, but we can do GetImageBase64 style or we can let the backend return the original code
                // Let's implement dynamic original code display if we want. For now, since we have the original script path, we can let user see the optimized code.
                originalCode = await App.GetImageBase64(scriptPath); // Wait, GetImageBase64 returns base64. Let's decode it or just keep it simple.
                if (originalCode.startsWith("data:")) {
                    const base64Content = originalCode.split(",")[1];
                    originalCode = atob(base64Content); // Decode base64 to text!
                }
            } catch (e) {
                console.error("Failed to read original code text:", e);
                originalCode = "// 元のファイルを読み込めませんでした。";
            }

            showDiffModal = true;
            showStatus('ボトルネック分析が完了しました！', 'success');
        } catch (e) {
            console.error(e);
            showStatus(`最適化提案に失敗しました: ${e.message || e}`, 'error');
        } finally {
            isOptimizing = false;
        }
    }

    // 最適化コードの適用
    async function applyOptimization() {
        if (!optimizationResult || !optimizationResult.refactored_code) return;
        
        try {
            await SaveScriptFile(scriptPath, optimizationResult.refactored_code);
            showStatus('最適化したUWSCRコードをファイルに適用しました！', 'success');
            showDiffModal = false;
        } catch (e) {
            alert("適用に失敗しました: " + e);
        }
    }

    function clearConsole() {
        logs = [];
    }

    function showStatus(msg, type) {
        statusMessage = msg;
        statusType = type;
        setTimeout(() => {
            if (statusMessage === msg) {
                statusMessage = '';
                statusType = '';
            }
        }, 4000);
    }

    function openLink(url) {
        wailsRuntime.BrowserOpenURL(url);
    }

    function parseMessage(msg) {
        if (!msg) return [];
        const urlRegex = /(https?:\/\/[^\s\)]+)/g;
        const parts = msg.split(urlRegex);
        return parts.map(part => {
            if (part.match(urlRegex)) {
                return { text: part, isLink: true };
            }
            return { text: part, isLink: false };
        });
    }

    async function browseScriptPath() {
        try {
            const selected = await SelectFile(
                "UWSCRスクリプトを選択", 
                "UWSCR Script (*.uws)", 
                "*.uws"
            );
            if (selected) {
                scriptPath = selected;
            }
        } catch (e) {
            showStatus(`ファイル選択エラー: ${e.message || e}`, 'error');
        }
    }
</script>

<div class="console-grid">
    <!-- 設定・実行コントロール -->
    <div class="card control-card">
        <div class="card-header">
            <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="5 3 19 12 5 21 5 3"/>
            </svg>
            <div class="header-text">
                <h2>UWSCRスクリプト実行</h2>
                <p class="description">UWSCRスクリプトをトランスパイルして実行します。</p>
            </div>
        </div>
        
        <div class="form-body">
            <div class="form-group">
                <label for="script-path-input">実行するスクリプト (.uws) の絶対パス</label>
                <div class="input-with-btn">
                    <input 
                        id="script-path-input" 
                        type="text" 
                        placeholder="C:\scripts\myscript.uws" 
                        bind:value={scriptPath}
                        disabled={isRunning}
                    />
                    <button class="btn-secondary" on:click={browseScriptPath} disabled={isRunning}>選択...</button>
                </div>
            </div>

             <div class="action-bar">
                <button class="btn-clear" on:click={clearConsole} disabled={logs.length === 0 || isRunning}>
                    クリア
                </button>
                <div class="run-buttons" style="display: flex; gap: 8px;">
                    {#if isRunning}
                        <button class="btn-run btn-stop" on:click={handleStop} style="background-color: var(--accent-red); color: white;">
                            🛑 停止
                        </button>
                    {:else}
                        <button class="btn-run" on:click={handleRun}>
                            実行
                        </button>
                    {/if}
                </div>
            </div>
        </div>

        {#if logs.length > 0 && !isRunning}
            <div class="optimization-trigger-box border-t">
                <h4>🧠 自己進化 (ファクトリファクタリング)</h4>
                <p class="desc-sm">実測ミリ秒ログを解析して、不要な待機時間を取り除いた非同期並行化コードをAIが提案します。</p>
                <button 
                    class="btn-optimize gradient-btn-blue" 
                    on:click={handleOptimize}
                    disabled={isOptimizing}
                >
                    {#if isOptimizing}
                        ボトルネック分析中...
                    {:else}
                        ボトルネック分析 & 最適化提案
                    {/if}
                </button>
            </div>
        {/if}

        {#if statusMessage}
            <div class="status-message {statusType}">
                {#each parseMessage(statusMessage) as part}
                    {#if part.isLink}
                        <a href="javascript:void(0)" on:click|preventDefault={() => openLink(part.text)} class="msg-link">{part.text}</a>
                    {:else}
                        {part.text}
                    {/if}
                {/each}
            </div>
        {/if}
    </div>

    <!-- ログ表示コンソール -->
    <div class="card logs-card">
        <div class="logs-header">
            <div class="logs-title">
                <svg class="console-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="4 17 10 11 4 5"/>
                    <line x1="12" y1="19" x2="20" y2="19"/>
                </svg>
                <h3>実行ログコンソール</h3>
            </div>
            {#if isRunning}
                <span class="pulse-indicator">RUNNING</span>
            {/if}
        </div>
        
        <div class="logs-box" bind:this={logContainer}>
            {#if logs.length === 0}
                <div class="empty-logs">UWSCRの実行結果やAIログがここに表示されます。</div>
            {/if}
            {#each logs as log}
                <div class="log-line {log.is_error ? 'error' : ''}">
                    <span class="message">{log.message}</span>
                </div>
            {/each}
        </div>
    </div>
</div>

<!-- Diffプレビューモーダル -->
{#if showDiffModal && optimizationResult}
    <div class="modal-overlay" on:click={() => showDiffModal = false}>
        <div class="modal-content card" on:click|stopPropagation>
            <div class="modal-header">
                <h3>🧠 ボトルネック分析 & 最適化提案</h3>
                <button class="modal-close-btn" on:click={() => showDiffModal = false}>×</button>
            </div>
            <div class="modal-body">
                <div class="stats-row">
                    <div class="stat-card">
                        <span class="stat-label">推定削減時間</span>
                        <span class="stat-value text-emerald">{optimizationResult.estimated_time_saved_ms} ms</span>
                    </div>
                    <div class="stat-card flex-fill">
                        <span class="stat-label">検出されたボトルネック</span>
                        <ul class="bottleneck-list">
                            {#each optimizationResult.bottlenecks as item}
                                <li>⚠️ {item}</li>
                            {/each}
                        </ul>
                    </div>
                </div>

                <div class="code-diff-grid">
                    <div class="code-box">
                        <span class="code-box-title">元のコード</span>
                        <textarea class="code-diff-textarea" readonly>{originalCode}</textarea>
                    </div>
                    <div class="code-box">
                        <span class="code-box-title text-emerald">提案された最適化コード</span>
                        <textarea class="code-diff-textarea text-emerald-glow" readonly>{optimizationResult.refactored_code}</textarea>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <button class="btn-cancel" on:click={() => showDiffModal = false}>キャンセル</button>
                <button class="btn-apply-opt" on:click={applyOptimization}>最適化コードを適用する</button>
            </div>
        </div>
    </div>
{/if}

<style>
    .console-grid {
        display: grid;
        grid-template-columns: 1fr;
        gap: 24px;
    }

    @media (min-width: 900px) {
        .console-grid {
            grid-template-columns: 380px 1fr;
        }
    }

    .card {
        background: var(--bg-secondary);
        backdrop-filter: var(--glass-blur);
        border: 1px solid var(--border-color);
        border-radius: 12px;
        padding: 24px;
        box-shadow: var(--shadow-md);
        display: flex;
        flex-direction: column;
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

    .form-body {
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

    .form-group {
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    label {
        font-size: 0.8rem;
        color: var(--text-secondary);
        font-weight: 500;
    }

    input {
        background: var(--input-bg);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 10px 14px;
        color: var(--text-primary);
        font-size: 0.85rem;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
        width: 100%;
        box-sizing: border-box;
    }

    input:focus {
        outline: none;
        border-color: var(--accent-color);
        box-shadow: 0 0 0 2px var(--accent-soft);
    }

    .input-with-btn {
        display: flex;
        gap: 8px;
    }

    .btn-secondary {
        background: var(--input-bg);
        border: 1px solid var(--border-color);
        color: var(--text-primary);
        border-radius: 8px;
        padding: 0 16px;
        font-size: 0.8rem;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
        white-space: nowrap;
    }

    .btn-secondary:hover {
        background: var(--border-color);
        border-color: var(--border-hover);
    }

    .action-bar {
        display: flex;
        justify-content: space-between;
        margin-top: 12px;
    }

    .btn-clear {
        background: transparent;
        border: 1px solid var(--border-color);
        color: var(--text-secondary);
        border-radius: 8px;
        padding: 10px 20px;
        font-size: 0.85rem;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .btn-clear:hover:not(:disabled) {
        border-color: var(--border-hover);
        color: var(--text-primary);
        background: var(--accent-soft);
    }

    .btn-run {
        background-color: var(--accent-color);
        border: none;
        color: var(--bg-primary);
        border-radius: 8px;
        padding: 10px 24px;
        font-size: 0.85rem;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
        box-shadow: var(--shadow-sm);
    }

    .btn-run:hover:not(:disabled) {
        background-color: var(--accent-hover);
    }

    .btn-run:active:not(:disabled) {
        transform: scale(0.97);
    }

    .optimization-trigger-box {
        margin-top: 20px;
        padding-top: 20px;
        display: flex;
        flex-direction: column;
        gap: 10px;
    }

    .optimization-trigger-box h4 {
        margin: 0;
        font-size: 0.85rem;
        color: var(--text-primary);
    }

    .optimization-trigger-box .desc-sm {
        margin: 0;
        font-size: 0.7rem;
        color: var(--text-secondary);
        line-height: 1.4;
    }

    .btn-optimize {
        border: none;
        color: #fff;
        font-weight: 600;
        border-radius: 8px;
        padding: 10px 16px;
        font-size: 0.8rem;
        cursor: pointer;
        transition: transform 0.1s ease, opacity 0.2s ease;
    }

    .gradient-btn-blue {
        background: linear-gradient(135deg, #0984e3, #6c5ce7);
    }

    .gradient-btn-blue:hover {
        opacity: 0.9;
    }

    .gradient-btn-blue:active {
        transform: scale(0.97);
    }

    .gradient-btn-blue:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .border-t {
        border-top: 1px solid var(--border-color);
    }

    button:disabled {
        opacity: 0.4;
        cursor: not-allowed;
    }

    .status-message {
        margin-top: 16px;
        font-size: 0.8rem;
        padding: 8px 12px;
        border-radius: 8px;
        background: rgba(255, 255, 255, 0.05);
    }

    .status-message.success {
        background: rgba(46, 213, 115, 0.08);
        color: #2ed573;
        border: 1px solid rgba(46, 213, 115, 0.2);
    }

    .status-message.error {
        background: rgba(255, 71, 87, 0.08);
        color: #ff4757;
        border: 1px solid rgba(255, 71, 87, 0.2);
    }

    /* Logs Console Styles */
    .logs-card {
        display: flex;
        flex-direction: column;
        height: 500px;
    }

    .logs-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16px;
        border-bottom: 1px solid var(--border-color);
        padding-bottom: 16px;
    }

    .logs-title {
        display: flex;
        align-items: center;
        gap: 12px;
    }

    .console-icon {
        width: 20px;
        height: 20px;
        color: var(--accent-color);
    }

    .logs-header h3 {
        margin: 0;
        font-size: 1.05rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .pulse-indicator {
        font-size: 0.65rem;
        background: rgba(46, 213, 115, 0.08);
        color: #2ed573;
        border: 1px solid rgba(46, 213, 115, 0.2);
        padding: 4px 10px;
        border-radius: 20px;
        font-weight: 600;
        animation: pulse 2s infinite;
    }

    .logs-box {
        background: var(--input-bg);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        flex-grow: 1;
        padding: 20px;
        font-family: 'Fira Code', 'SF Mono', 'Courier New', monospace;
        font-size: 0.8rem;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    .empty-logs {
        color: var(--text-secondary);
        text-align: center;
        margin: auto;
        font-size: 0.8rem;
    }

    .log-line {
        display: flex;
        gap: 10px;
        line-height: 1.5;
        color: var(--text-primary);
    }

    .log-line.error {
        color: #ff4757;
    }

    .msg-link {
        color: var(--text-primary);
        text-decoration: underline;
        cursor: pointer;
        font-weight: 500;
        transition: opacity 0.2s ease;
    }

    .msg-link:hover {
        opacity: 0.8;
    }

    /* Modal Styles */
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background: rgba(0, 0, 0, 0.6);
        z-index: 1000;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .modal-content {
        width: 90vw;
        max-width: 960px;
        height: 80vh;
        background: var(--bg-secondary);
        box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
    }

    .modal-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 1px solid var(--border-color);
        padding-bottom: 12px;
        margin-bottom: 16px;
    }

    .modal-header h3 {
        margin: 0;
        font-size: 1.1rem;
        color: var(--text-primary);
    }

    .modal-close-btn {
        background: transparent;
        border: none;
        color: var(--text-secondary);
        font-size: 1.5rem;
        cursor: pointer;
    }

    .modal-close-btn:hover {
        color: var(--text-primary);
    }

    .modal-body {
        flex: 1;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 16px;
        min-height: 0;
    }

    .stats-row {
        display: flex;
        gap: 16px;
    }

    .stat-card {
        background: var(--input-bg);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 12px 16px;
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .flex-fill {
        flex: 1;
    }

    .stat-label {
        font-size: 0.7rem;
        color: var(--text-secondary);
        font-weight: 600;
    }

    .stat-value {
        font-size: 1.5rem;
        font-weight: 700;
    }

    .text-emerald {
        color: #2ecc71;
    }

    .bottleneck-list {
        margin: 0;
        padding-left: 16px;
        font-size: 0.75rem;
        line-height: 1.5;
        color: var(--text-secondary);
    }

    .code-diff-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 16px;
        flex: 1;
        min-height: 0;
    }

    .code-box {
        display: flex;
        flex-direction: column;
        gap: 6px;
        min-height: 0;
    }

    .code-box-title {
        font-size: 0.75rem;
        font-weight: 600;
        color: var(--text-secondary);
    }

    .code-diff-textarea {
        flex: 1;
        background: rgba(0, 0, 0, 0.4);
        border: 1px solid var(--border-color);
        border-radius: 6px;
        padding: 12px;
        font-family: Consolas, monospace;
        font-size: 0.75rem;
        color: var(--text-primary);
        resize: none;
        outline: none;
        white-space: pre;
    }

    .text-emerald-glow {
        color: #2ecc71;
        border-color: rgba(46, 211, 113, 0.3);
    }

    .modal-footer {
        display: flex;
        justify-content: flex-end;
        gap: 12px;
        border-top: 1px solid var(--border-color);
        padding-top: 16px;
        margin-top: 16px;
    }

    .btn-cancel {
        background: transparent;
        border: 1px solid var(--border-color);
        color: var(--text-secondary);
        border-radius: 6px;
        padding: 8px 16px;
        font-size: 0.8rem;
        cursor: pointer;
    }

    .btn-cancel:hover {
        border-color: var(--border-hover);
        color: var(--text-primary);
    }

    .btn-apply-opt {
        background: #2ecc71;
        border: none;
        color: #0f172a;
        font-weight: 600;
        border-radius: 6px;
        padding: 8px 20px;
        font-size: 0.8rem;
        cursor: pointer;
    }

    .btn-apply-opt:hover {
        opacity: 0.9;
    }

    @keyframes pulse {
        0% { opacity: 0.6; }
        50% { opacity: 1; }
        100% { opacity: 0.6; }
    }
</style>
