<script>
  import { onMount } from 'svelte';
  import { fade, fly } from 'svelte/transition';
  import { GetConfig, SaveUWSCRPath, SelectFile } from '../wailsjs/go/main/App.js';
  import CredentialInput from './components/CredentialInput.svelte';
  import LayerConfig from './components/LayerConfig.svelte';
  import Console from './components/Console.svelte';
  import ManualCreator from './components/ManualCreator.svelte';
  import ScriptDeveloper from './components/ScriptDeveloper.svelte';

  let config = null;
  let isLoading = true;
  let activeTab = 'run'; // Default to daily execution tab
  let isLightMode = false;
  let showSettingsDrawer = false;
  let selectedLanguage = 'ja';
  let uwscrPath = '';

  onMount(async () => {
    try {
      config = await GetConfig();
      if (config) {
        uwscrPath = config.UWSCRPath || '';
      }
    } catch (e) {
      console.error('Failed to load configuration:', e);
    } finally {
      isLoading = false;
    }
  });

  async function browseUWSCRPath() {
    try {
      const selected = await SelectFile(
        "UWSCR実行ファイル (uwscr.exe) を選択",
        "Executables (*.exe)",
        "*.exe"
      );
      if (selected) {
        uwscrPath = selected;
      }
    } catch (e) {
      console.error('Failed to select UWSCR path:', e);
    }
  }

  async function handleSaveUWSCRPath() {
    try {
      await SaveUWSCRPath(uwscrPath);
      if (config) {
        config.UWSCRPath = uwscrPath;
      }
    } catch (e) {
      console.error('Failed to save UWSCR path:', e);
    }
  }

  function toggleTheme() {
    isLightMode = !isLightMode;
    if (isLightMode) {
      document.documentElement.classList.add('light-mode');
    } else {
      document.documentElement.classList.remove('light-mode');
    }
  }
</script>

<main class="app-container">
  <header class="app-header">
    <div class="brand">
      <svg class="logo-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
        <rect x="2" y="3" width="20" height="14" rx="2" ry="2"/>
        <line x1="8" y1="21" x2="16" y2="21"/>
        <line x1="12" y1="17" x2="12" y2="21"/>
      </svg>
      <div class="brand-text">
        <h1>UWSCR::act-gram</h1>
        <span class="subtitle">AI-Native RPA Control Center</span>
      </div>
    </div>

    <div class="header-right">
      <nav class="nav-tabs">
        <button 
          class="tab-btn {activeTab === 'run' ? 'active' : ''}" 
          on:click={() => activeTab = 'run'}
          disabled={isLoading}
        >
          スクリプト実行
        </button>
        <button 
          class="tab-btn {activeTab === 'dev' ? 'active' : ''}" 
          on:click={() => activeTab = 'dev'}
          disabled={isLoading}
        >
          スクリプト開発
        </button>
        <button 
          class="tab-btn {activeTab === 'manual' ? 'active' : ''}" 
          on:click={() => activeTab = 'manual'}
          disabled={isLoading}
        >
          マニュアル作成
        </button>
      </nav>

      <button class="settings-btn" on:click={() => showSettingsDrawer = true} aria-label="Open settings">
        <svg class="settings-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="3"/>
          <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/>
        </svg>
      </button>
    </div>
  </header>

  <div class="content-area">
    {#if isLoading}
      <div class="loader-container">
        <div class="spinner"></div>
        <p>設定をロード中...</p>
      </div>
    {:else}
      <div class="tab-content">
        {#if activeTab === 'manual'}
          <ManualCreator />
        {:else if activeTab === 'dev'}
          <ScriptDeveloper />
        {:else}
          <Console />
        {/if}
      </div>
    {/if}
  </div>

  <!-- Settings Drawer Overlay & Panel -->
  {#if showSettingsDrawer}
    <div class="drawer-overlay" on:click={() => showSettingsDrawer = false} transition:fade={{ duration: 200 }}></div>
    <div class="settings-drawer" transition:fly={{ x: 500, duration: 300, opacity: 1 }}>
      <div class="drawer-header">
        <div class="drawer-title">
          <svg class="drawer-header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="3"/>
            <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/>
          </svg>
          <h2>環境設定</h2>
        </div>
        <button class="btn-close" on:click={() => showSettingsDrawer = false} aria-label="Close settings">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"/>
            <line x1="6" y1="6" x2="18" y2="18"/>
          </svg>
        </button>
      </div>

      <div class="drawer-body">
        <!-- Section 1: General Preferences -->
        <div class="drawer-section">
          <h3>全般設定</h3>
          
          <div class="setting-item">
            <div class="setting-info">
              <span class="setting-name">テーマ切り替え</span>
              <span class="setting-desc">ライトモードとダークモードを切り替えます。</span>
            </div>
            <button class="btn-theme-toggle" on:click={toggleTheme}>
              {#if isLightMode}
                ダークモードにする
              {:else}
                ライトモードにする
              {/if}
            </button>
          </div>

          <div class="setting-item">
            <div class="setting-info">
              <span class="setting-name">表示言語</span>
              <span class="setting-desc">UIの表示言語を切り替えます。</span>
            </div>
            <select class="select-lang" bind:value={selectedLanguage}>
              <option value="ja">日本語</option>
              <option value="en">English (Mock)</option>
            </select>
          </div>
        </div>

        <!-- Section 1.5: UWSCR Config -->
        <div class="drawer-section">
          <h3>UWSCR設定</h3>
          
          <div class="setting-item flex-column" style="align-items: stretch; gap: 10px; border-bottom: none;">
            <div class="setting-info">
              <span class="setting-name">UWSCR本体のパス (uwscr.exe)</span>
              <span class="setting-desc">スクリプトを実行する UWSCR 実行ファイルの絶対パスを指定します。未指定時は自動探索されます。</span>
            </div>
            <div class="input-with-btn" style="display: flex; gap: 8px;">
              <input
                id="drawer-uwscr-path"
                type="text"
                class="path-input"
                style="flex: 1; min-width: 0; background: var(--input-bg); border: 1px solid var(--border-color); border-radius: 6px; padding: 8px 12px; font-size: 0.8rem; font-family: Consolas, monospace; color: var(--text-primary);"
                bind:value={uwscrPath}
                placeholder="未指定時は自動探索されます"
              />
              <button class="btn-theme-toggle" on:click={browseUWSCRPath}>選択...</button>
              <button class="btn-theme-toggle" on:click={handleSaveUWSCRPath} style="background: var(--accent-color); color: var(--bg-primary); border-color: var(--accent-color);">適用</button>
            </div>
          </div>
        </div>

        <!-- Section 2: API Credentials -->
        <div class="drawer-section">
          <CredentialInput />
        </div>

        <!-- Section 3: AI Layers Config -->
        <div class="drawer-section">
          <LayerConfig />
        </div>
      </div>
    </div>
  {/if}
</main>

<style>
  :global(body) {
    background-color: var(--bg-primary);
    color: var(--text-primary);
    margin: 0;
    padding: 0;
    overflow: hidden;
  }

  .app-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    box-sizing: border-box;
    background-color: var(--bg-primary);
  }

  .app-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 14px 28px;
    background: var(--bg-secondary);
    border-bottom: 1px solid var(--border-color);
    z-index: 10;
    transition: background-color 0.3s ease, border-color 0.3s ease;
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .logo-svg {
    width: 24px;
    height: 24px;
    color: var(--text-primary);
  }

  .brand-text h1 {
    margin: 0;
    font-size: 1.05rem;
    font-weight: 600;
    letter-spacing: -0.2px;
    color: var(--text-primary);
  }

  .subtitle {
    font-size: 0.7rem;
    color: var(--text-secondary);
    display: block;
    margin-top: 1px;
  }

  .header-right {
    display: flex;
    align-items: center;
    gap: 16px;
  }

  .nav-tabs {
    display: flex;
    gap: 4px;
    background: rgba(0, 0, 0, 0.05);
    padding: 3px;
    border-radius: 8px;
    border: 1px solid var(--border-color);
  }

  :global(html.light-mode) .nav-tabs {
    background: rgba(0, 0, 0, 0.02);
  }

  .tab-btn {
    background: transparent;
    border: 1px solid transparent;
    color: var(--text-secondary);
    padding: 6px 14px;
    font-size: 0.8rem;
    font-weight: 500;
    border-radius: 6px;
    cursor: pointer;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .tab-btn:hover:not(:disabled) {
    color: var(--text-primary);
    background: var(--accent-soft);
  }

  .tab-btn.active {
    background: var(--accent-color);
    color: var(--bg-primary);
    font-weight: 600;
    box-shadow: var(--shadow-sm);
  }

  .tab-btn:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }

  .settings-btn {
    background: transparent;
    border: 1px solid var(--border-color);
    color: var(--text-secondary);
    width: 32px;
    height: 32px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.2s ease;
    padding: 0;
  }

  .settings-btn:hover {
    color: var(--text-primary);
    border-color: var(--border-hover);
    background: var(--accent-soft);
  }

  .settings-icon {
    width: 16px;
    height: 16px;
  }

  .content-area {
    flex: 1;
    overflow-y: auto;
    background-color: var(--bg-primary);
    transition: background-color 0.3s ease;
  }

  .tab-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 32px 24px;
    box-sizing: border-box;
  }

  .loader-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    gap: 16px;
    color: var(--text-secondary);
  }

  .spinner {
    width: 28px;
    height: 28px;
    border: 2px solid var(--border-color);
    border-radius: 50%;
    border-top-color: var(--accent-color);
    animation: spin 0.8s cubic-bezier(0.4, 0, 0.2, 1) infinite;
  }

  /* Settings Drawer Styles */
  .drawer-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.4);
    z-index: 900;
  }

  .settings-drawer {
    position: fixed;
    top: 0;
    right: 0;
    width: 100%;
    max-width: 520px;
    height: 100%;
    background: var(--bg-secondary);
    border-left: 1px solid var(--border-color);
    box-shadow: -10px 0 30px rgba(0, 0, 0, 0.15);
    z-index: 1000;
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
  }

  .drawer-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 24px;
    border-bottom: 1px solid var(--border-color);
  }

  .drawer-title {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .drawer-header-icon {
    width: 20px;
    height: 20px;
    color: var(--text-primary);
  }

  .drawer-header h2 {
    margin: 0;
    font-size: 1.1rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .btn-close {
    background: transparent;
    border: none;
    color: var(--text-secondary);
    width: 28px;
    height: 28px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 6px;
    transition: all 0.2s ease;
    padding: 0;
  }

  .btn-close:hover {
    color: var(--text-primary);
    background: var(--accent-soft);
  }

  .btn-close svg {
    width: 16px;
    height: 16px;
  }

  .drawer-body {
    flex: 1;
    overflow-y: auto;
    padding: 24px;
    display: flex;
    flex-direction: column;
    gap: 32px;
  }

  .drawer-section {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .drawer-section h3 {
    margin: 0 0 8px 0;
    font-size: 0.8rem;
    text-transform: uppercase;
    letter-spacing: 1px;
    color: var(--text-secondary);
    font-weight: 700;
    border-bottom: 1px solid var(--border-color);
    padding-bottom: 6px;
  }

  .setting-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 0;
    border-bottom: 1px dashed var(--border-color);
  }

  .setting-item:last-child {
    border-bottom: none;
  }

  .setting-info {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .setting-name {
    font-size: 0.85rem;
    font-weight: 500;
    color: var(--text-primary);
  }

  .setting-desc {
    font-size: 0.75rem;
    color: var(--text-secondary);
  }

  .btn-theme-toggle {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-radius: 6px;
    padding: 6px 12px;
    font-size: 0.75rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .btn-theme-toggle:hover {
    border-color: var(--border-hover);
    background: var(--accent-soft);
  }

  .select-lang {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-radius: 6px;
    padding: 6px 12px;
    font-size: 0.75rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .select-lang:focus {
    outline: none;
    border-color: var(--accent-color);
  }

  /* Clean up embedded cards styling inside drawer */
  .settings-drawer :global(.card) {
    background: transparent !important;
    border: none !important;
    padding: 0 !important;
    box-shadow: none !important;
    margin: 0 !important;
    max-width: 100% !important;
  }

  .settings-drawer :global(.card-header) {
    border-bottom: none !important;
    padding-bottom: 0 !important;
    margin-bottom: 16px !important;
  }

  .settings-drawer :global(.header-icon) {
    display: none !important; /* Icons already exist in drawer titles */
  }

  .settings-drawer :global(.layers-grid) {
    grid-template-columns: 1fr !important;
    gap: 16px !important;
  }

  .settings-drawer :global(.layer-card) {
    padding: 16px !important;
    border-radius: 8px !important;
    background: rgba(0, 0, 0, 0.08) !important;
  }

  :global(html.light-mode) .settings-drawer :global(.layer-card) {
    background: rgba(0, 0, 0, 0.02) !important;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }
</style>
