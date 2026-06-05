<script>
    import { onMount, onDestroy } from 'svelte';
    import { RunScript, GetConfig, SaveUWSCRPath } from '../../wailsjs/go/main/App';
    import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';

    let scriptPath = '';
    let uwscrPath = '';
    let logs = [];
    let isRunning = false;
    let logContainer;
    let config = null;

    let statusMessage = '';
    let statusType = ''; // 'success' | 'error'

    onMount(async () => {
        try {
            config = await GetConfig();
            if (config) {
                uwscrPath = config.uwscr_path || '';
            }
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

    async function handleSavePath() {
        try {
            await SaveUWSCRPath(uwscrPath);
            showStatus('UWSCRのパス設定を保存しました。', 'success');
        } catch (e) {
            showStatus(`保存失敗: ${e.message || e}`, 'error');
        }
    }

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
</script>

<div class="console-grid">
    <!-- 設定・実行コントロール -->
    <div class="card control-card">
        <h2>⚡ UWSCRスクリプト実行</h2>
        
        <div class="form-group">
            <label for="uwscr-path-input">UWSCR本体のパス (オプション)</label>
            <div class="input-with-btn">
                <input 
                    id="uwscr-path-input" 
                    type="text" 
                    placeholder="未指定時は同階層またはPATHから自動探索されます" 
                    bind:value={uwscrPath}
                />
                <button on:click={handleSavePath}>適用</button>
            </div>
        </div>

        <div class="form-group">
            <label for="script-path-input">実行するスクリプト (.uws) の絶対パス</label>
            <input 
                id="script-path-input" 
                type="text" 
                placeholder="C:\scripts\myscript.uws" 
                bind:value={scriptPath}
                disabled={isRunning}
            />
        </div>

        <div class="action-bar">
            <button class="btn-clear" on:click={clearConsole} disabled={logs.length === 0}>
                コンソールクリア
            </button>
            <button class="btn-run" on:click={handleRun} disabled={isRunning}>
                {isRunning ? '実行中...' : 'スクリプト実行'}
            </button>
        </div>

        {#if statusMessage}
            <div class="status-message {statusType}">
                {statusMessage}
            </div>
        {/if}
    </div>

    <!-- ログ表示コンソール -->
    <div class="card logs-card">
        <div class="logs-header">
            <h3>💻 実行ログコンソール</h3>
            {#if isRunning}
                <span class="pulse-indicator">RUNNING</span>
            {/if}
        </div>
        
        <div class="logs-box" bind:this={logContainer}>
            {#if logs.length === 0}
                <div class="empty-logs">ここにUWSCRの実行結果やエラーがリアルタイムに表示されます。</div>
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
        gap: 20px;
    }

    @media (min-width: 800px) {
        .console-grid {
            grid-template-columns: 350px 1fr;
        }
    }

    .card {
        background: rgba(255, 255, 255, 0.03);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.08);
        border-radius: 12px;
        padding: 24px;
        box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.2);
    }

    h2, h3 {
        margin-top: 0;
        color: #fff;
    }

    .form-group {
        display: flex;
        flex-direction: column;
        gap: 6px;
        margin-bottom: 16px;
    }

    label {
        font-size: 0.8rem;
        color: #ccc;
    }

    input {
        background: rgba(0, 0, 0, 0.25);
        border: 1px solid rgba(255, 255, 255, 0.12);
        border-radius: 6px;
        padding: 10px 12px;
        color: #fff;
        font-size: 0.85rem;
        box-sizing: border-box;
        width: 100%;
    }

    input:focus {
        outline: none;
        border-color: #535bf2;
    }

    .input-with-btn {
        display: flex;
        gap: 8px;
    }

    .input-with-btn button {
        background: rgba(255, 255, 255, 0.08);
        border: 1px solid rgba(255, 255, 255, 0.12);
        color: #fff;
        border-radius: 6px;
        padding: 0 16px;
        font-size: 0.8rem;
        cursor: pointer;
        transition: background-color 0.2s;
    }

    .input-with-btn button:hover {
        background: rgba(255, 255, 255, 0.15);
    }

    .action-bar {
        display: flex;
        justify-content: space-between;
        margin-top: 24px;
    }

    button {
        border-radius: 6px;
        padding: 10px 20px;
        font-size: 0.85rem;
        font-weight: 600;
        cursor: pointer;
        transition: background-color 0.2s, transform 0.1s;
    }

    .btn-clear {
        background: transparent;
        border: 1px solid rgba(255, 255, 255, 0.15);
        color: #ccc;
    }

    .btn-clear:hover:not(:disabled) {
        background: rgba(255, 255, 255, 0.05);
        color: #fff;
    }

    .btn-run {
        background: #535bf2;
        border: none;
        color: #fff;
    }

    .btn-run:hover:not(:disabled) {
        background: #4047d9;
    }

    button:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .status-message {
        margin-top: 16px;
        font-size: 0.8rem;
        padding: 8px 12px;
        border-radius: 6px;
    }

    .status-message.success {
        background: rgba(46, 213, 115, 0.15);
        color: #2ed573;
        border-left: 4px solid #2ed573;
    }

    .status-message.error {
        background: rgba(255, 71, 87, 0.15);
        color: #ff4757;
        border-left: 4px solid #ff4757;
    }

    /* Logs Console Styles */
    .logs-card {
        display: flex;
        flex-direction: column;
        height: 450px;
    }

    .logs-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 12px;
    }

    .pulse-indicator {
        font-size: 0.65rem;
        background: rgba(255, 71, 87, 0.15);
        color: #ff4757;
        border: 1px solid rgba(255, 71, 87, 0.3);
        padding: 2px 8px;
        border-radius: 12px;
        font-weight: bold;
        animation: pulse 1.5s infinite;
    }

    .logs-box {
        background: #090a10;
        border: 1px solid rgba(255, 255, 255, 0.05);
        border-radius: 8px;
        flex-grow: 1;
        padding: 16px;
        font-family: 'Fira Code', 'Courier New', monospace;
        font-size: 0.85rem;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    .empty-logs {
        color: #555;
        text-align: center;
        margin: auto;
        font-size: 0.8rem;
    }

    .log-line {
        display: flex;
        gap: 8px;
        line-height: 1.4;
        color: #a9b7c6;
    }

    .log-line.error {
        color: #ff4757;
    }

    .timestamp {
        color: #555;
        user-select: none;
    }

    @keyframes pulse {
        0% { opacity: 0.6; }
        50% { opacity: 1; }
        100% { opacity: 0.6; }
    }
</style>
