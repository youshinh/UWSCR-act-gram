<script>
    import { onMount } from 'svelte';
    import { SaveAPIKey, HasAPIKey } from '../../wailsjs/go/main/App';

    let selectedProvider = 'google';
    let apiKey = '';
    let isSaving = false;
    let statusMessage = '';
    let statusType = ''; // 'success' | 'error'

    // 各プロバイダーのAPIキー登録状態
    let keyStatus = {
        google: false,
        anthropic: false,
        openai: false // Ollamaはキー不要のため除外
    };

    onMount(async () => {
        await checkKeyStatus();
    });

    async function checkKeyStatus() {
        try {
            keyStatus.google = await HasAPIKey('google');
            keyStatus.anthropic = await HasAPIKey('anthropic');
            keyStatus.openai = await HasAPIKey('openai');
        } catch (e) {
            console.error('Failed to check API key status', e);
        }
    }

    async function handleSave() {
        if (!apiKey.trim()) {
            showStatus('APIキーを入力してください。', 'error');
            return;
        }

        isSaving = true;
        showStatus('保存中...', '');

        try {
            await SaveAPIKey(selectedProvider, apiKey);
            apiKey = ''; // 入力欄をクリア
            await checkKeyStatus();
            showStatus(`${selectedProvider.toUpperCase()} のAPIキーをセキュアに保存しました。`, 'success');
        } catch (e) {
            showStatus(`保存に失敗しました: ${e.message || e}`, 'error');
        } finally {
            isSaving = false;
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
            }, 4000);
        }
    }
</script>

<div class="card credential-card">
    <div class="card-header">
        <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
        </svg>
        <div class="header-text">
            <h2>API認証キー設定</h2>
            <p class="description">
                APIキーはWindows資格情報マネージャー（セキュアストレージ）へ暗号化して安全に保存されます。
            </p>
        </div>
    </div>

    <div class="status-summary">
        <div class="status-badge {keyStatus.google ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            Gemini: {keyStatus.google ? '登録済み' : '未登録'}
        </div>
        <div class="status-badge {keyStatus.anthropic ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            Claude: {keyStatus.anthropic ? '登録済み' : '未登録'}
        </div>
    </div>

    <div class="form-body">
        <div class="form-group">
            <label for="provider-select">プロバイダー</label>
            <select id="provider-select" bind:value={selectedProvider} on:change={() => { statusMessage = ''; }}>
                <option value="google">Google (Gemini)</option>
                <option value="anthropic">Anthropic (Claude)</option>
                <option value="openai">OpenAI / 互換サーバー</option>
            </select>
        </div>

        <div class="form-group">
            <label for="apikey-input">APIキー</label>
            <input 
                id="apikey-input" 
                type="password" 
                placeholder="APIキーを入力..." 
                bind:value={apiKey}
                disabled={isSaving}
            />
        </div>

        <div class="action-bar">
            <button class="btn-primary" on:click={handleSave} disabled={isSaving}>
                {isSaving ? '保存中...' : 'セキュアに保存'}
            </button>
        </div>
    </div>

    {#if statusMessage}
        <div class="status-message {statusType}">
            {statusMessage}
        </div>
    {/if}
</div>

<style>
    .credential-card {
        background: var(--bg-secondary);
        backdrop-filter: var(--glass-blur);
        border: 1px solid var(--border-color);
        border-radius: 12px;
        padding: 32px;
        box-shadow: var(--shadow-md);
        max-width: 600px;
        margin: 0 auto;
        transition: background-color 0.3s ease, border-color 0.3s ease;
    }

    .card-header {
        display: flex;
        gap: 16px;
        align-items: flex-start;
        margin-bottom: 24px;
        border-bottom: 1px solid var(--border-color);
        padding-bottom: 20px;
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
        font-size: 1.15rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .description {
        font-size: 0.8rem;
        color: var(--text-secondary);
        margin: 4px 0 0 0;
        line-height: 1.4;
    }

    .status-summary {
        display: flex;
        gap: 12px;
        margin-bottom: 24px;
    }

    .status-badge {
        font-size: 0.75rem;
        padding: 6px 12px;
        border-radius: 20px;
        font-weight: 500;
        display: flex;
        align-items: center;
        gap: 6px;
    }

    .status-badge .indicator-dot {
        width: 6px;
        height: 6px;
        border-radius: 50%;
        display: inline-block;
    }

    .status-badge.registered {
        background-color: var(--accent-soft);
        color: var(--accent-color);
        border: 1px solid var(--accent-border);
    }

    .status-badge.registered .indicator-dot {
        background-color: var(--accent-color);
    }

    .status-badge.unregistered {
        background-color: rgba(255, 71, 87, 0.08);
        color: #ff4757;
        border: 1px solid rgba(255, 71, 87, 0.2);
    }

    .status-badge.unregistered .indicator-dot {
        background-color: #ff4757;
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

    select, input {
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

    select:focus, input:focus {
        outline: none;
        border-color: var(--accent-color);
        box-shadow: 0 0 0 2px var(--accent-soft);
    }

    .action-bar {
        display: flex;
        justify-content: flex-end;
        margin-top: 8px;
    }

    .btn-primary {
        background-color: var(--accent-color);
        color: var(--bg-primary);
        border: none;
        border-radius: 8px;
        padding: 10px 20px;
        font-size: 0.85rem;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
        box-shadow: var(--shadow-sm);
    }

    .btn-primary:hover:not(:disabled) {
        background-color: var(--accent-hover);
    }

    .btn-primary:active:not(:disabled) {
        transform: scale(0.97);
    }

    .btn-primary:disabled {
        opacity: 0.4;
        cursor: not-allowed;
    }

    .status-message {
        margin-top: 20px;
        font-size: 0.8rem;
        padding: 10px 14px;
        border-radius: 8px;
        background: rgba(255, 255, 255, 0.05);
        color: var(--text-primary);
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
</style>
