<script>
    import { onMount } from 'svelte';
    import { GetConfig, SaveConfig, SaveConfigs, FetchModels } from '../../wailsjs/go/main/App';

    let config = null;
    let isLoading = true;

    // 各レイヤーの選択状態 (デフォルトをGoogleのgemini-flash-lite-latestにする)
    let selections = {
        brain: { provider: 'google', model: 'gemini-flash-lite-latest' },
        eye: { provider: 'google', model: 'gemini-flash-lite-latest' },
        utility: { provider: 'google', model: 'gemini-flash-lite-latest' }
    };

    // 一括設定モードの制御用
    let isUnifiedMode = true;
    let unifiedSelection = { provider: 'google', model: 'gemini-flash-lite-latest' };
    let isFetchingUnified = false;

    // 各プロバイダーのモデルリストキャッシュ
    let modelCache = {
        google: [],
        anthropic: [],
        ollama: []
    };

    // 各レイヤーのモデル取得中フラグ
    let isFetching = {
        brain: false,
        eye: false,
        utility: false
    };

    onMount(async () => {
        try {
            config = await GetConfig();
            if (config && config.layers) {
                // 既存設定の反映
                for (const layer of ['brain', 'eye', 'utility']) {
                    if (config.layers[layer]) {
                        selections[layer].provider = config.layers[layer].provider;
                        selections[layer].model = config.layers[layer].model;
                    }
                }

                // デフォルトを一括設定モードとして起動する
                isUnifiedMode = true;
                unifiedSelection.provider = selections.brain.provider;
                unifiedSelection.model = selections.brain.model;
            }

            if (isUnifiedMode) {
                await loadModelsForUnified();
            } else {
                await Promise.all([
                    loadModelsForLayer('brain'),
                    loadModelsForLayer('eye'),
                    loadModelsForLayer('utility')
                ]);
            }
        } catch (e) {
            console.error('Failed to init LayerConfig', e);
        } finally {
            isLoading = false;
        }
    });

    async function loadModelsForLayer(layer) {
        const provider = selections[layer].provider;
        if (!provider) return;

        isFetching[layer] = true;
        try {
            if (!modelCache[provider] || modelCache[provider].length === 0) {
                modelCache[provider] = await FetchModels(provider);
            }

            const models = modelCache[provider];
            if (models.length > 0 && !models.includes(selections[layer].model)) {
                selections[layer].model = models[0];
                await handleSave(layer);
            }
        } catch (e) {
            console.error(`Failed to load models for ${layer}`, e);
        } finally {
            isFetching[layer] = false;
        }
    }

    async function handleProviderChange(layer) {
        selections[layer].model = '';
        await loadModelsForLayer(layer);
    }

    async function handleSave(layer) {
        try {
            await SaveConfig(layer, selections[layer].provider, selections[layer].model);
        } catch (e) {
            console.error(`Failed to save config for ${layer}`, e);
        }
    }

    // 一括設定用
    async function loadModelsForUnified() {
        const provider = unifiedSelection.provider;
        if (!provider) return;

        isFetchingUnified = true;
        try {
            if (!modelCache[provider] || modelCache[provider].length === 0) {
                modelCache[provider] = await FetchModels(provider);
            }

            const models = modelCache[provider];
            if (models.length > 0 && !models.includes(unifiedSelection.model)) {
                unifiedSelection.model = models[0];
            }
            await handleUnifiedSave();
        } catch (e) {
            console.error(`Failed to load models for unified`, e);
        } finally {
            isFetchingUnified = false;
        }
    }

    async function handleUnifiedProviderChange() {
        unifiedSelection.model = '';
        await loadModelsForUnified();
    }

    async function handleUnifiedSave() {
        selections.brain = { ...unifiedSelection };
        selections.eye = { ...unifiedSelection };
        selections.utility = { ...unifiedSelection };

        try {
            const layers = {
                brain: selections.brain,
                eye: selections.eye,
                utility: selections.utility
            };
            await SaveConfigs(JSON.stringify(layers));
        } catch (e) {
            console.error(`Failed to save unified config`, e);
        }
    }

    async function toggleMode(unified) {
        isUnifiedMode = unified;
        if (isUnifiedMode) {
            unifiedSelection.provider = selections.brain.provider;
            unifiedSelection.model = selections.brain.model;
            await loadModelsForUnified();
        } else {
            await Promise.all([
                loadModelsForLayer('brain'),
                loadModelsForLayer('eye'),
                loadModelsForLayer('utility')
            ]);
        }
    }
</script>

{#if isLoading}
    <div class="loader-container">
        <div class="spinner"></div>
        <p>レイヤー設定を読み込み中...</p>
    </div>
{:else}
    <div class="setting-controls">
        <div class="mode-selector">
            <button class="mode-btn" class:active={!isUnifiedMode} on:click={() => toggleMode(false)}>
                個別設定
            </button>
            <button class="mode-btn" class:active={isUnifiedMode} on:click={() => toggleMode(true)}>
                一括設定
            </button>
        </div>
    </div>

    {#if isUnifiedMode}
        <div class="unified-container">
            <div class="card unified-card">
                <div class="card-header">
                    <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
                        <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
                    </svg>
                    <div>
                        <h3>全レイヤー共通設定</h3>
                        <span class="badge">すべてのレイヤーに同じモデルを適用</span>
                    </div>
                </div>
                <p class="description">Brain（計画）、Eye（画像解析）、Utility（データ処理）のすべてのレイヤーに同一のモデルを一括して適用します。</p>
                
                <div class="form-body">
                    <div class="control-group">
                        <label for="unified-provider">共通プロバイダー</label>
                        <select id="unified-provider" bind:value={unifiedSelection.provider} on:change={handleUnifiedProviderChange}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="ollama">Ollama (ローカルLLM)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="unified-model">共通AIモデル</label>
                        {#if isFetchingUnified}
                            <div class="select-loading">モデル取得中...</div>
                        {:else}
                            <select id="unified-model" bind:value={unifiedSelection.model} on:change={handleUnifiedSave}>
                                {#each modelCache[unifiedSelection.provider] || [] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
    {:else}
        <div class="layers-grid">
            <!-- Brain層 Card -->
            <div class="card layer-card brain-card">
                <div class="card-header">
                    <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M22 12h-4l-3 9L9 3l-3 9H2"/>
                    </svg>
                    <div>
                        <h3>Brain層 (コード生成・計画)</h3>
                        <span class="badge">賢さと安定性優先</span>
                    </div>
                </div>
                <p class="description">UWSCRコードの組み立て、実行計画の作成などを司るコアモジュール。</p>
                
                <div class="form-body">
                    <div class="control-group">
                        <label for="brain-provider">プロバイダー</label>
                        <select id="brain-provider" bind:value={selections.brain.provider} on:change={() => handleProviderChange('brain')}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="ollama">Ollama (ローカルLLM)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="brain-model">AIモデル</label>
                        {#if isFetching.brain}
                            <div class="select-loading">モデル取得中...</div>
                        {:else}
                            <select id="brain-model" bind:value={selections.brain.model} on:change={() => handleSave('brain')}>
                                {#each modelCache[selections.brain.provider] || [] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>

            <!-- Eye層 Card -->
            <div class="card layer-card eye-card">
                <div class="card-header">
                    <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                        <circle cx="12" cy="12" r="3"/>
                    </svg>
                    <div>
                        <h3>Eye層 (画面解析・要素特定)</h3>
                        <span class="badge">視覚能力優先</span>
                    </div>
                </div>
                <p class="description">デスクトップ画面をキャプチャし、UI要素の位置や画像内の情報を解析するモジュール。</p>
                
                <div class="form-body">
                    <div class="control-group">
                        <label for="eye-provider">プロバイダー</label>
                        <select id="eye-provider" bind:value={selections.eye.provider} on:change={() => handleProviderChange('eye')}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="ollama">Ollama (ローカルLLM)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="eye-model">AIモデル</label>
                        {#if isFetching.eye}
                            <div class="select-loading">モデル取得中...</div>
                        {:else}
                            <select id="eye-model" bind:value={selections.eye.model} on:change={() => handleSave('eye')}>
                                {#each modelCache[selections.eye.provider] || [] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>

            <!-- Hand/Utility層 Card -->
            <div class="card layer-card utility-card">
                <div class="card-header">
                    <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="3"/>
                        <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/>
                    </svg>
                    <div>
                        <h3>Utility / Hand層 (データ処理・反射)</h3>
                        <span class="badge">速度とコスト優先</span>
                    </div>
                </div>
                <p class="description">文字列の置換、正規表現の生成、データ形式の高速な変換を担うモジュール。</p>
                
                <div class="form-body">
                    <div class="control-group">
                        <label for="utility-provider">プロバイダー</label>
                        <select id="utility-provider" bind:value={selections.utility.provider} on:change={() => handleProviderChange('utility')}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="ollama">Ollama (ローカルLLM)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="utility-model">AIモデル</label>
                        {#if isFetching.utility}
                            <div class="select-loading">モデル取得中...</div>
                        {:else}
                            <select id="utility-model" bind:value={selections.utility.model} on:change={() => handleSave('utility')}>
                                {#each modelCache[selections.utility.provider] || [] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
    {/if}
{/if}

<style>
    .setting-controls {
        display: flex;
        justify-content: flex-end;
        margin-bottom: 24px;
    }

    .mode-selector {
        display: flex;
        background: rgba(0, 0, 0, 0.05);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 3px;
        gap: 4px;
    }

    :global(html.light-mode) .mode-selector {
        background: rgba(0, 0, 0, 0.02);
    }

    .mode-btn {
        background: transparent;
        border: none;
        color: var(--text-secondary);
        padding: 6px 14px;
        border-radius: 6px;
        font-size: 0.8rem;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .mode-btn:hover {
        color: var(--text-primary);
    }

    .mode-btn.active {
        color: var(--bg-primary);
        background: var(--accent-color);
        font-weight: 600;
        box-shadow: var(--shadow-sm);
    }

    .unified-container {
        display: flex;
        justify-content: center;
        width: 100%;
    }

    .unified-card {
        max-width: 500px;
        width: 100%;
    }

    .layers-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
        gap: 24px;
    }

    .layer-card, .unified-card {
        background: var(--bg-secondary);
        backdrop-filter: var(--glass-blur);
        border: 1px solid var(--border-color);
        border-radius: 12px;
        padding: 28px;
        box-shadow: var(--shadow-md);
        display: flex;
        flex-direction: column;
        transition: transform 0.2s cubic-bezier(0.4, 0, 0.2, 1), border-color 0.2s ease;
    }

    .layer-card:hover, .unified-card:hover {
        border-color: var(--border-hover);
        transform: translateY(-2px);
    }

    .card-header {
        display: flex;
        align-items: flex-start;
        gap: 16px;
        margin-bottom: 16px;
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

    .card-header h3 {
        margin: 0;
        font-size: 1.05rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .badge {
        font-size: 0.65rem;
        background: var(--accent-soft);
        color: var(--accent-color);
        border: 1px solid var(--accent-border);
        padding: 2px 8px;
        border-radius: 20px;
        font-weight: 500;
        display: inline-block;
        margin-top: 6px;
    }

    .description {
        font-size: 0.8rem;
        color: var(--text-secondary);
        line-height: 1.5;
        margin-bottom: 24px;
        margin-top: 0;
        flex-grow: 1;
    }

    .form-body {
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .control-group {
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    label {
        font-size: 0.75rem;
        color: var(--text-secondary);
        font-weight: 500;
    }

    select {
        background: var(--input-bg);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 8px 12px;
        color: var(--text-primary);
        font-size: 0.85rem;
        width: 100%;
        box-sizing: border-box;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    }

    select:focus {
        outline: none;
        border-color: var(--accent-color);
        box-shadow: 0 0 0 2px var(--accent-soft);
    }

    .select-loading {
        background: var(--input-bg);
        border: 1px dashed var(--border-color);
        border-radius: 8px;
        padding: 8px 12px;
        color: var(--text-secondary);
        font-size: 0.85rem;
        text-align: center;
    }

    .loader-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 60px;
        color: var(--text-secondary);
        gap: 16px;
    }

    .spinner {
        width: 28px;
        height: 28px;
        border: 2px solid var(--border-color);
        border-radius: 50%;
        border-top-color: var(--accent-color);
        animation: spin 0.8s cubic-bezier(0.4, 0, 0.2, 1) infinite;
    }

    @keyframes spin {
        to { transform: rotate(360deg); }
    }
</style>
