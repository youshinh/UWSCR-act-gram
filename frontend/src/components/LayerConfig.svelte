<script>
    import { onMount } from 'svelte';
    import { GetConfig, SaveConfig, FetchModels } from '../../wailsjs/go/main/App';

    let config = null;
    let isLoading = true;

    // 各レイヤーの選択状態
    let selections = {
        brain: { provider: 'anthropic', model: '' },
        eye: { provider: 'google', model: '' },
        utility: { provider: 'google', model: '' }
    };

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
            }
            // 初期プロバイダーのモデル一覧を取得
            await Promise.all([
                loadModelsForLayer('brain'),
                loadModelsForLayer('eye'),
                loadModelsForLayer('utility')
            ]);
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
            // キャッシュがない場合のみ取得
            if (!modelCache[provider] || modelCache[provider].length === 0) {
                modelCache[provider] = await FetchModels(provider);
            }

            // 現在のモデルがリストになければ、最初のモデルを選択
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
        // プロバイダー変更時にモデルをリセットし、再取得
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
</script>

{#if isLoading}
    <div class="loader-container">
        <div class="spinner"></div>
        <p>レイヤー設定を読み込み中...</p>
    </div>
{:else}
    <div class="layers-grid">
        <!-- Brain層 Card -->
        <div class="card layer-card brain-card">
            <div class="card-header">
                <span class="icon">🧠</span>
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
                <span class="icon">👁️</span>
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
                <span class="icon">🛠️</span>
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

<style>
    .layers-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
        gap: 20px;
    }

    .layer-card {
        background: rgba(255, 255, 255, 0.03);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.08);
        border-radius: 12px;
        padding: 24px;
        box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.2);
        display: flex;
        flex-direction: column;
        transition: transform 0.2s, border-color 0.2s;
    }

    .layer-card:hover {
        border-color: rgba(83, 91, 242, 0.2);
        transform: translateY(-2px);
    }

    .card-header {
        display: flex;
        align-items: center;
        gap: 12px;
        margin-bottom: 12px;
    }

    .card-header h3 {
        margin: 0;
        font-size: 1.1rem;
        color: #fff;
    }

    .icon {
        font-size: 1.8rem;
    }

    .badge {
        font-size: 0.65rem;
        background: rgba(83, 91, 242, 0.15);
        color: #7b83ff;
        border: 1px solid rgba(83, 91, 242, 0.3);
        padding: 2px 6px;
        border-radius: 4px;
        font-weight: 600;
        display: inline-block;
        margin-top: 4px;
    }

    .description {
        font-size: 0.8rem;
        color: #aaa;
        line-height: 1.4;
        margin-bottom: 20px;
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
        font-size: 0.8rem;
        color: #ccc;
    }

    select {
        background: rgba(0, 0, 0, 0.25);
        border: 1px solid rgba(255, 255, 255, 0.12);
        border-radius: 6px;
        padding: 8px 12px;
        color: #fff;
        font-size: 0.85rem;
        width: 100%;
        box-sizing: border-box;
    }

    select:focus {
        outline: none;
        border-color: #535bf2;
    }

    .select-loading {
        background: rgba(0, 0, 0, 0.15);
        border: 1px dashed rgba(255, 255, 255, 0.1);
        border-radius: 6px;
        padding: 8px 12px;
        color: #888;
        font-size: 0.85rem;
        text-align: center;
    }

    .loader-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 60px;
        color: #888;
        gap: 16px;
    }

    .spinner {
        width: 32px;
        height: 32px;
        border: 3px solid rgba(255, 255, 255, 0.1);
        border-radius: 50%;
        border-top-color: #535bf2;
        animation: spin 1s ease-in-out infinite;
    }

    @keyframes spin {
        to { transform: rotate(360deg); }
    }
</style>
