<script>
    import { onMount, onDestroy } from 'svelte';
    import { RunScript, GetConfig, SelectFile } from '../../wailsjs/go/main/App';
    import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';

    let scriptPath = '';
    let logs = [];
    let isRunning = false;
    let logContainer;
    let config = null;

    let statusMessage = '';
    let statusType = ''; // 'success' | 'error'

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
            }
        });
    });

    onDestroy(() => {
        // イベントリスナーの解除
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

        try {
            await RunScript(scriptPath);
        } catch (e) {
            showStatus(`実行開始に失敗しました: ${e.message || e}`, 'error');
            isRunning = false;
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
                <button class="btn-clear" on:click={clearConsole} disabled={logs.length === 0}>
                    クリア
                </button>
                <button class="btn-run" on:click={handleRun} disabled={isRunning}>
                    {isRunning ? '実行中...' : '実行'}
                </button>
            </div>
        </div>

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
                    <span class="timestamp">[{new Date().toLocaleTimeString()}]</span>
                    <span class="message">{log.message}</span>
                </div>
            {/each}
        </div>
    </div>
</div>

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
        padding: 28px;
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

    .timestamp {
        color: var(--text-secondary);
        opacity: 0.7;
        user-select: none;
        font-size: 0.75rem;
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

    @keyframes pulse {
        0% { opacity: 0.6; }
        50% { opacity: 1; }
        100% { opacity: 0.6; }
    }
</style>
