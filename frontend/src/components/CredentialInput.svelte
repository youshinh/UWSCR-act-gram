<script>
    import { onMount } from 'svelte';
    import { SaveAPIKey, HasAPIKey, GetConfig, SaveCustomBaseURL, TestAPIKeyConnection, GetLocalLLMConfig, SaveLocalLLMConfig } from '../../wailsjs/go/main/App';
    import { EventsEmit } from '../../wailsjs/runtime/runtime';

    let selectedProvider = 'google';
    let apiKey = '';
    let customBaseURL = '';
    let localLLMType = 'ollama';
    let localLLMURL = '';
    let isSaving = false;
    let isTesting = false;
    let statusMessage = '';
    let statusType = ''; // 'success' | 'error' | 'testing'

    // 各プロバイダーのAPIキー登録状態
    let keyStatus = {
        google: false,
        anthropic: false,
        openai: false,
        custom: false,
        local: false
    };

    onMount(async () => {
        await checkKeyStatus();
        await loadConfig();
        await loadLocalConfig();
    });

    async function loadConfig() {
        try {
            const config = await GetConfig();
            if (config && config.CustomBaseURL) {
                customBaseURL = config.CustomBaseURL;
            } else {
                customBaseURL = 'http://localhost:8080/v1';
            }
        } catch (e) {
            console.error('Failed to load config for CredentialInput', e);
        }
    }

    async function loadLocalConfig() {
        try {
            const localConfig = await GetLocalLLMConfig();
            if (localConfig) {
                localLLMType = localConfig.type || 'ollama';
                localLLMURL = localConfig.url || '';
            }
        } catch (e) {
            console.error('Failed to load local config', e);
        }
    }

    async function checkKeyStatus() {
        try {
            keyStatus.google = await HasAPIKey('google');
            keyStatus.anthropic = await HasAPIKey('anthropic');
            keyStatus.openai = await HasAPIKey('openai');
            keyStatus.custom = await HasAPIKey('custom');
            keyStatus.local = await HasAPIKey('local');
        } catch (e) {
            console.error('Failed to check API key status', e);
        }
    }

    async function handleSave() {
        if (selectedProvider !== 'local' && !apiKey.trim()) {
            showStatus('APIキーを入力してください。', 'error');
            return;
        }

        isSaving = true;
        showStatus('保存中...', '');

        try {
            if (selectedProvider === 'custom') {
                await SaveCustomBaseURL(customBaseURL);
            }
            if (selectedProvider === 'local') {
                await SaveLocalLLMConfig(localLLMType, localLLMURL);
                if (apiKey.trim()) {
                    await SaveAPIKey('local', apiKey);
                } else {
                    await SaveAPIKey('local', '');
                }
            } else {
                await SaveAPIKey(selectedProvider, apiKey);
            }
            apiKey = ''; // 入力欄をクリア
            await checkKeyStatus();
            EventsEmit('llm-key-updated', selectedProvider);
            
            let displayName = selectedProvider === 'openai' ? 'ChatGPT' : 
                              selectedProvider === 'custom' ? 'カスタム互換' : 
                              selectedProvider === 'local' ? 'ローカルLLM' : 
                              selectedProvider.toUpperCase();
            showStatus(`${displayName} の設定を保存しました。`, 'success');
        } catch (e) {
            showStatus(`保存に失敗しました: ${e.message || e}`, 'error');
        } finally {
            isSaving = false;
        }
    }

    async function handleTestConnection() {
        isTesting = true;
        showStatus('テスト接続中...', 'testing');

        try {
            let testURL = '';
            let testLLMType = '';
            if (selectedProvider === 'custom') {
                testURL = customBaseURL;
            } else if (selectedProvider === 'local') {
                testURL = localLLMURL;
                if (!testURL) {
                    testURL = localLLMType === 'ollama' ? 'http://localhost:11434' : 'http://localhost:1234';
                }
                testLLMType = localLLMType; // UIの現在選択値を渡す
            }
            const res = await TestAPIKeyConnection(selectedProvider, apiKey, testURL, testLLMType);
            showStatus(`${res}`, 'success');
            EventsEmit('llm-key-updated', selectedProvider);
        } catch (e) {
            showStatus(`接続失敗: ${e.message || e}`, 'error');
        } finally {
            isTesting = false;
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

<div class="card credential-card">
    <div class="card-header">
        <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
        </svg>
        <div class="header-text">
            <h2>API認証キー設定</h2>
            <p class="description">
                APIキーはWindows内で暗号化して安全に保存されます。
            </p>
        </div>
    </div>

    <div class="status-summary" style="display: flex; flex-wrap: wrap; gap: 12px; margin-bottom: 24px;">
        <div class="status-badge {keyStatus.google ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            Gemini: {keyStatus.google ? '登録済み' : '未登録'}
        </div>
        <div class="status-badge {keyStatus.anthropic ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            Claude: {keyStatus.anthropic ? '登録済み' : '未登録'}
        </div>
        <div class="status-badge {keyStatus.openai ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            chatGPT: {keyStatus.openai ? '登録済み' : '未登録'}
        </div>
        <div class="status-badge {keyStatus.custom ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            カスタム: {keyStatus.custom ? '登録済み' : '未登録'}
        </div>
        <div class="status-badge {keyStatus.local ? 'registered' : 'unregistered'}">
            <span class="indicator-dot"></span>
            ローカルLLM: {keyStatus.local ? 'キー登録あり' : 'キーなし（又は未登録）'}
        </div>
    </div>

    <div class="form-body">
        <div class="form-group">
            <label for="provider-select">プロバイダー</label>
            <select id="provider-select" bind:value={selectedProvider} on:change={() => { statusMessage = ''; }}>
                <option value="google">Google (Gemini)</option>
                <option value="anthropic">Anthropic (Claude)</option>
                <option value="openai">OpenAI (ChatGPT)</option>
                <option value="custom">カスタム互換サーバー</option>
                <option value="local">ローカルLLM (Ollama/LM Studio)</option>
            </select>
        </div>

        {#if selectedProvider === 'custom'}
            <div class="form-group">
                <label for="custom-url-input">接続先ベースURL</label>
                <input 
                    id="custom-url-input" 
                    type="text" 
                    placeholder="http://localhost:8080/v1" 
                    bind:value={customBaseURL}
                    disabled={isSaving || isTesting}
                />
            </div>
        {/if}

        {#if selectedProvider === 'local'}
            <div class="form-group">
                <label for="local-llm-type">ローカルLLMタイプ</label>
                <select id="local-llm-type" bind:value={localLLMType} disabled={isSaving || isTesting}>
                    <option value="ollama">Ollama</option>
                    <option value="lmstudio">LM Studio (OpenAI互換)</option>
                </select>
            </div>
            
            <div class="form-group">
                <label for="local-url-input">接続先ベースURL</label>
                <input 
                    id="local-url-input" 
                    type="text" 
                    placeholder={localLLMType === 'ollama' ? 'http://localhost:11434' : 'http://localhost:1234'} 
                    bind:value={localLLMURL}
                    disabled={isSaving || isTesting}
                />
                <span class="input-hint">
                    {#if localLLMType === 'ollama'}
                        Ollamaの標準: http://localhost:11434
                    {:else}
                        LM Studioの標準: http://localhost:1234 (ネイティブAPI /api/v1/models を自動検出します)
                    {/if}
                </span>
            </div>
        {/if}

        <div class="form-group">
            <label for="apikey-input">
                {selectedProvider === 'local' ? 'APIキー / トークン (オプショナル)' : 'APIキー'}
            </label>
            <input 
                id="apikey-input" 
                type="password" 
                placeholder={selectedProvider === 'local' ? '不要な場合は空欄のままにしてください' : 'APIキーを入力...'} 
                bind:value={apiKey}
                disabled={isSaving || isTesting}
            />
        </div>

        <div class="action-bar" style="display: flex; gap: 12px; justify-content: flex-end;">
            <button 
                class="btn-secondary" 
                on:click={handleTestConnection} 
                disabled={isSaving || isTesting}
                style="background: transparent; color: var(--accent-color); border: 1px solid var(--accent-color); padding: 10px 20px; border-radius: 8px; font-weight: 600; cursor: pointer; font-size: 0.85rem;"
            >
                {isTesting ? '接続確認中...' : 'テスト接続'}
            </button>
            <button class="btn-primary" on:click={handleSave} disabled={isSaving || isTesting}>
                {isSaving ? '保存中...' : '設定を保存'}
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

    .input-hint {
        font-size: 0.72rem;
        color: var(--text-secondary);
        margin-top: 4px;
        display: block;
    }
</style>
