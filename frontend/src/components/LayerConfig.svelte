<script>
    import { onMount } from 'svelte';
    import { GetConfig, SaveConfigs, FetchModels } from '../../wailsjs/go/main/App';
    import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime';

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

    // 各プロバイダー of モデルリストキャッシュ
    let modelCache = {
        google: [],
        anthropic: [],
        openai: [],
        custom: [],
        local: []
    };

    // 各レイヤー of モデル取得中フラグ
    let isFetching = {
        brain: false,
        eye: false,
        utility: false
    };

    onMount(async () => {
        let cleanup = () => {};
        try {
            config = await GetConfig();
            if (config) {
                // UseUnifiedModel フラグの復元 (WailsはGoのPascalCaseフィールド名でJSONを返す)
                if (config.UseUnifiedModel !== undefined) {
                    isUnifiedMode = !!config.UseUnifiedModel;
                } else if (config.Layers) {
                    // 古い設定ファイルの互換性：全て同じなら一括設定と判定
                    const layers = config.Layers;
                    if (layers.brain && layers.eye && layers.utility &&
                        layers.brain.Provider === layers.eye.Provider && layers.brain.Provider === layers.utility.Provider &&
                        layers.brain.Model === layers.eye.Model && layers.brain.Model === layers.utility.Model) {
                        isUnifiedMode = true;
                    } else {
                        isUnifiedMode = false;
                    }
                }

                // 既存設定の反映
                if (config.Layers) {
                    for (const layer of ['brain', 'eye', 'utility']) {
                        if (config.Layers[layer]) {
                            selections[layer].provider = config.Layers[layer].Provider;
                            selections[layer].model = config.Layers[layer].Model;
                        }
                    }
                }

                // 一括設定用 of 初期値 of 決定
                if (isUnifiedMode) {
                    unifiedSelection.provider = selections.brain.provider;
                    unifiedSelection.model = selections.brain.model;
                } else {
                    unifiedSelection.provider = selections.brain.provider;
                    unifiedSelection.model = selections.brain.model;
                }
            }

            // モデル of ロード
            if (isUnifiedMode) {
                await loadModelsForUnified();
            } else {
                await Promise.all([
                    loadModelsForLayer('brain'),
                    loadModelsForLayer('eye'),
                    loadModelsForLayer('utility')
                ]);
            }

            // APIキー更新イベント of リスニング登録
            EventsOn('llm-key-updated', async (provider) => {
                console.log(`LLM key updated for ${provider}, clearing cache and reloading models...`);
                modelCache[provider] = [];
                
                if (isUnifiedMode) {
                    if (unifiedSelection.provider === provider) {
                        await loadModelsForUnified();
                    }
                } else {
                    for (const layer of ['brain', 'eye', 'utility']) {
                        if (selections[layer].provider === provider) {
                            await loadModelsForLayer(layer);
                        }
                    }
                }
            });

            cleanup = () => {
                EventsOff('llm-key-updated');
            };
        } catch (e) {
            console.error('Failed to init LayerConfig', e);
        } finally {
            isLoading = false;
        }

        return () => {
            cleanup();
        };
    });

    let statusMessage = '';
    let isSaving = false;

    async function loadModelsForLayer(layer) {
        const provider = selections[layer].provider;
        if (!provider) return;

        isFetching[layer] = true;
        try {
            if (!modelCache[provider] || modelCache[provider].length === 0) {
                modelCache[provider] = await FetchModels(provider);
            }
            // 不要な自動リセット（models[0]への上書き）を削除し、保存済みの設定値を保護します
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

    // 一括設定用
    async function loadModelsForUnified() {
        const provider = unifiedSelection.provider;
        if (!provider) return;

        isFetchingUnified = true;
        try {
            if (!modelCache[provider] || modelCache[provider].length === 0) {
                modelCache[provider] = await FetchModels(provider);
            }
            // 不要な自動リセット（models[0]への上書き）を削除し、保存済みの設定値を保護します
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

    export async function applyConfig() {
        isSaving = true;
        statusMessage = '';
        try {
            let layers = {};
            if (isUnifiedMode) {
                selections.brain = { ...unifiedSelection };
                selections.eye = { ...unifiedSelection };
                selections.utility = { ...unifiedSelection };
            }
            layers = {
                brain: selections.brain,
                eye: selections.eye,
                utility: selections.utility
            };
            
            // 新形式のパラメータ（use_unified_modelフラグを同梱）でGoに送信
            const payload = {
                use_unified_model: isUnifiedMode,
                layers: layers
            };
            await SaveConfigs(JSON.stringify(payload));
        } catch (e) {
            console.error(`Failed to save configurations`, e);
            throw e;
        } finally {
            isSaving = false;
        }
    }

    async function handleModeChange() {
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

    // 一括設定が有効な時、リアクティブに個別設定へ値を同期 (初期ロード完了後のみ実行)
    $: if (isUnifiedMode && !isLoading) {
        selections.brain.provider = unifiedSelection.provider;
        selections.brain.model = unifiedSelection.model;
        selections.eye.provider = unifiedSelection.provider;
        selections.eye.model = unifiedSelection.model;
        selections.utility.provider = unifiedSelection.provider;
        selections.utility.model = unifiedSelection.model;
    }
</script>

{#if isLoading}
    <div class="loader-container">
        <div class="spinner"></div>
        <p>レイヤー設定を読み込み中...</p>
    </div>
{:else}
    <!-- 一括設定のチェックボックス -->
    <div class="mode-selection-bar">
        <label class="checkbox-container">
            <input type="checkbox" bind:checked={isUnifiedMode} on:change={handleModeChange} />
            <span class="checkmark"></span>
            すべてのレイヤーに同じモデルを適用する (一括設定)
        </label>
    </div>

    <div class="config-sections">
        <!-- 1. 全レイヤー共通設定 (一括設定がONのときのみアクティブ) -->
        <div class="card unified-card" class:disabled={!isUnifiedMode}>
            <div class="card-header">
                <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
                    <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
                </svg>
                <h3>全レイヤー共通設定</h3>
            </div>
            
            <div class="form-row">
                <div class="control-group">
                    <label for="unified-provider">共通プロバイダー</label>
                    <select id="unified-provider" bind:value={unifiedSelection.provider} on:change={handleUnifiedProviderChange} disabled={!isUnifiedMode}>
                        <option value="google">Google (Gemini)</option>
                        <option value="anthropic">Anthropic (Claude)</option>
                        <option value="openai">OpenAI (ChatGPT)</option>
                        <option value="custom">カスタム互換サーバー</option>
                        <option value="local">ローカルLLM (Ollama/LM Studio)</option>
                    </select>
                </div>

                <div class="control-group">
                    <label for="unified-model">共通AIモデル</label>
                    {#if isFetchingUnified || !modelCache[unifiedSelection.provider] || modelCache[unifiedSelection.provider].length === 0}
                        <div class="select-loading">取得中...</div>
                    {:else}
                        <select id="unified-model" bind:value={unifiedSelection.model} disabled={!isUnifiedMode}>
                            {#each modelCache[unifiedSelection.provider] as m}
                                <option value={m}>{m}</option>
                            {/each}
                            {#if unifiedSelection.model && !modelCache[unifiedSelection.provider].includes(unifiedSelection.model)}
                                <option value={unifiedSelection.model}>{unifiedSelection.model}</option>
                            {/if}
                        </select>
                    {/if}
                </div>
            </div>
        </div>

        <div class="section-title" class:dimmed={isUnifiedMode}>
            <h4>個別設定</h4>
            {#if isUnifiedMode}
                <span class="status-hint">一括設定が有効なため、以下は読み取り専用になっています。</span>
            {/if}
        </div>

        <!-- 2. 個別設定の各レイヤー (一括設定がOFFのときのみアクティブ) -->
        <div class="layers-grid" class:disabled={isUnifiedMode}>
            <!-- Brain層 Card -->
            <div class="card layer-card brain-card" class:disabled-item={isUnifiedMode}>
                <div class="card-header">
                    <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M22 12h-4l-3 9L9 3l-3 9H2"/>
                    </svg>
                    <div>
                        <h3>Brain層 (コード生成・計画)</h3>
                        <span class="badge">賢さと安定性優先</span>
                    </div>
                </div>
                
                <div class="form-row">
                    <div class="control-group">
                        <label for="brain-provider">プロバイダー</label>
                        <select id="brain-provider" bind:value={selections.brain.provider} on:change={() => handleProviderChange('brain')} disabled={isUnifiedMode}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="openai">OpenAI (ChatGPT)</option>
                            <option value="custom">カスタム互換サーバー</option>
                            <option value="local">ローカルLLM (Ollama/LM Studio)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="brain-model">AIモデル</label>
                        {#if isFetching.brain || !modelCache[selections.brain.provider] || modelCache[selections.brain.provider].length === 0}
                            <div class="select-loading">取得中...</div>
                        {:else}
                            <select id="brain-model" bind:value={selections.brain.model} disabled={isUnifiedMode}>
                                {#each modelCache[selections.brain.provider] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                                {#if selections.brain.model && !modelCache[selections.brain.provider].includes(selections.brain.model)}
                                    <option value={selections.brain.model}>{selections.brain.model}</option>
                                {/if}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>

            <!-- Eye層 Card -->
            <div class="card layer-card eye-card" class:disabled-item={isUnifiedMode}>
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
                
                <div class="form-row">
                    <div class="control-group">
                        <label for="eye-provider">プロバイダー</label>
                        <select id="eye-provider" bind:value={selections.eye.provider} on:change={() => handleProviderChange('eye')} disabled={isUnifiedMode}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="openai">OpenAI (ChatGPT)</option>
                            <option value="custom">カスタム互換サーバー</option>
                            <option value="local">ローカルLLM (Ollama/LM Studio)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="eye-model">AIモデル</label>
                        {#if isFetching.eye || !modelCache[selections.eye.provider] || modelCache[selections.eye.provider].length === 0}
                            <div class="select-loading">取得中...</div>
                        {:else}
                            <select id="eye-model" bind:value={selections.eye.model} disabled={isUnifiedMode}>
                                {#each modelCache[selections.eye.provider] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                                {#if selections.eye.model && !modelCache[selections.eye.provider].includes(selections.eye.model)}
                                    <option value={selections.eye.model}>{selections.eye.model}</option>
                                {/if}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>

            <!-- Hand/Utility層 Card -->
            <div class="card layer-card utility-card" class:disabled-item={isUnifiedMode}>
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
                
                <div class="form-row">
                    <div class="control-group">
                        <label for="utility-provider">プロバイダー</label>
                        <select id="utility-provider" bind:value={selections.utility.provider} on:change={() => handleProviderChange('utility')} disabled={isUnifiedMode}>
                            <option value="google">Google (Gemini)</option>
                            <option value="anthropic">Anthropic (Claude)</option>
                            <option value="openai">OpenAI (ChatGPT)</option>
                            <option value="custom">カスタム互換サーバー</option>
                            <option value="local">ローカルLLM (Ollama/LM Studio)</option>
                        </select>
                    </div>

                    <div class="control-group">
                        <label for="utility-model">AIモデル</label>
                        {#if isFetching.utility || !modelCache[selections.utility.provider] || modelCache[selections.utility.provider].length === 0}
                            <div class="select-loading">取得中...</div>
                        {:else}
                            <select id="utility-model" bind:value={selections.utility.model} disabled={isUnifiedMode}>
                                {#each modelCache[selections.utility.provider] as m}
                                    <option value={m}>{m}</option>
                                {/each}
                                {#if selections.utility.model && !modelCache[selections.utility.provider].includes(selections.utility.model)}
                                    <option value={selections.utility.model}>{selections.utility.model}</option>
                                {/if}
                            </select>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
    </div>
{/if}

<style>
    .mode-selection-bar {
        display: flex;
        align-items: center;
        background: var(--bg-secondary);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 14px 20px;
        margin-bottom: 24px;
    }

    .checkbox-container {
        display: flex;
        align-items: center;
        position: relative;
        padding-left: 28px;
        cursor: pointer;
        font-size: 0.88rem;
        font-weight: 500;
        color: var(--text-primary);
        user-select: none;
    }

    .checkbox-container input {
        position: absolute;
        opacity: 0;
        cursor: pointer;
        height: 0;
        width: 0;
    }

    .checkmark {
        position: absolute;
        top: 50%;
        left: 0;
        transform: translateY(-50%);
        height: 16px;
        width: 16px;
        background-color: var(--input-bg);
        border: 1px solid var(--border-color);
        border-radius: 4px;
        transition: all 0.2s ease;
    }

    .checkbox-container:hover input ~ .checkmark {
        border-color: var(--border-hover);
    }

    .checkbox-container input:checked ~ .checkmark {
        background-color: var(--accent-color);
        border-color: var(--accent-color);
    }

    .checkmark:after {
        content: "";
        position: absolute;
        display: none;
    }

    .checkbox-container input:checked ~ .checkmark:after {
        display: block;
    }

    .checkbox-container .checkmark:after {
        left: 5px;
        top: 2px;
        width: 4px;
        height: 8px;
        border: solid var(--bg-primary);
        border-width: 0 2px 2px 0;
        transform: rotate(45deg);
    }

    :global(html.light-mode) .checkbox-container input:checked ~ .checkmark:after {
        border-color: var(--bg-secondary);
    }

    .config-sections {
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

    .section-title {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-top: 10px;
        margin-bottom: 0px;
        padding: 0 4px;
    }

    .section-title h4 {
        margin: 0;
        font-size: 0.95rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .section-title.dimmed h4 {
        color: var(--text-secondary);
    }

    .status-hint {
        font-size: 0.72rem;
        color: var(--text-secondary);
        background: rgba(255, 255, 255, 0.04);
        padding: 2px 10px;
        border-radius: 4px;
    }

    :global(html.light-mode) .status-hint {
        background: rgba(0, 0, 0, 0.03);
    }

    .unified-card {
        width: 100%;
    }

    .layers-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
        gap: 24px;
        transition: opacity 0.25s ease;
    }

    .layers-grid.disabled {
        opacity: 0.5;
        pointer-events: none;
    }

    .layer-card, .unified-card {
        background: var(--bg-secondary);
        backdrop-filter: var(--glass-blur);
        border: 1px solid var(--border-color);
        border-radius: 12px;
        padding: 20px;
        box-shadow: var(--shadow-md);
        display: flex;
        flex-direction: column;
        gap: 16px;
        transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .layer-card:hover:not(.disabled-item), .unified-card:hover:not(.disabled) {
        border-color: var(--border-hover);
        transform: translateY(-2px);
    }

    .card.disabled {
        opacity: 0.5;
        pointer-events: none;
    }

    .card-header {
        display: flex;
        align-items: center;
        gap: 12px;
        border-bottom: 1px solid var(--border-color);
        padding-bottom: 12px;
    }

    .header-icon {
        width: 20px;
        height: 20px;
        color: var(--accent-color);
        flex-shrink: 0;
    }

    .card-header h3 {
        margin: 0;
        font-size: 0.95rem;
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
        margin-left: 6px;
    }

    /* Form rows for 2-column select inputs */
    .form-row {
        display: flex;
        gap: 16px;
        width: 100%;
    }

    .control-group {
        display: flex;
        flex-direction: column;
        gap: 6px;
        flex: 1;
        min-width: 0;
    }

    label {
        font-size: 0.72rem;
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

    select:disabled {
        background-color: var(--border-color);
        opacity: 0.8;
        color: var(--text-secondary);
        cursor: not-allowed;
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
