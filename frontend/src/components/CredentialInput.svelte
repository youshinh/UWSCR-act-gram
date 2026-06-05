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
    <h2>🔑 API認証キー設定</h2>
    <p class="description">
        APIキーは平文ファイルには保存されず、Windows資格情報マネージャー（セキュアストレージ）へ暗号化保存されます。
    </p>

    <div class="status-summary">
        <span class="status-badge {keyStatus.google ? 'registered' : 'unregistered'}">
            Gemini: {keyStatus.google ? '登録済み' : '未登録'}
        </span>
        <span class="status-badge {keyStatus.anthropic ? 'registered' : 'unregistered'}">
            Claude: {keyStatus.anthropic ? '登録済み' : '未登録'}
        </span>
    </div>

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
        <button on:click={handleSave} disabled={isSaving}>
            {isSaving ? '保存中...' : 'セキュアに保存'}
        </button>
    </div>

    {#if statusMessage}
        <div class="status-message {statusType}">
            {statusMessage}
        </div>
    {/if}
</div>

<style>
    .credential-card {
        background: rgba(255, 255, 255, 0.05);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 12px;
        padding: 24px;
        box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.3);
        margin-bottom: 20px;
    }

    h2 {
        margin-top: 0;
        font-size: 1.25rem;
        color: #fff;
        display: flex;
        align-items: center;
        gap: 8px;
    }

    .description {
        font-size: 0.85rem;
        color: #aaa;
        margin-bottom: 20px;
    }

    .status-summary {
        display: flex;
        gap: 12px;
        margin-bottom: 20px;
    }

    .status-badge {
        font-size: 0.75rem;
        padding: 4px 8px;
        border-radius: 6px;
        font-weight: 600;
    }

    .status-badge.registered {
        background-color: rgba(46, 213, 115, 0.2);
        color: #2ed573;
        border: 1px solid rgba(46, 213, 115, 0.3);
    }

    .status-badge.unregistered {
        background-color: rgba(255, 71, 87, 0.2);
        color: #ff4757;
        border: 1px solid rgba(255, 71, 87, 0.3);
    }

    .form-group {
        display: flex;
        flex-direction: column;
        gap: 8px;
        margin-bottom: 16px;
    }

    label {
        font-size: 0.85rem;
        color: #ccc;
        font-weight: 500;
    }

    select, input {
        background: rgba(0, 0, 0, 0.2);
        border: 1px solid rgba(255, 255, 255, 0.15);
        border-radius: 6px;
        padding: 10px 12px;
        color: #fff;
        font-size: 0.9rem;
        transition: border-color 0.2s, box-shadow 0.2s;
    }

    select:focus, input:focus {
        outline: none;
        border-color: #535bf2;
        box-shadow: 0 0 0 2px rgba(83, 91, 242, 0.3);
    }

    .action-bar {
        display: flex;
        justify-content: flex-end;
    }

    button {
        background-color: #535bf2;
        color: white;
        border: none;
        border-radius: 6px;
        padding: 10px 20px;
        font-size: 0.9rem;
        font-weight: 600;
        cursor: pointer;
        transition: background-color 0.2s, transform 0.1s;
    }

    button:hover:not(:disabled) {
        background-color: #4047d9;
    }

    button:active:not(:disabled) {
        transform: scale(0.98);
    }

    button:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }

    .status-message {
        margin-top: 16px;
        font-size: 0.85rem;
        padding: 8px 12px;
        border-radius: 6px;
        background: rgba(255, 255, 255, 0.1);
        color: #fff;
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
</style>
