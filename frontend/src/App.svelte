<script>
  import { onMount } from 'svelte';
  import { GetConfig } from '../wailsjs/go/main/App.js';
  import CredentialInput from './components/CredentialInput.svelte';
  import LayerConfig from './components/LayerConfig.svelte';
  import Console from './components/Console.svelte';

  let config = null;
  let isLoading = true;
  let activeTab = 'credentials'; // 'credentials' | 'config' | 'run'

  onMount(async () => {
    try {
      config = await GetConfig();
    } catch (e) {
      console.error('Failed to load configuration:', e);
    } finally {
      isLoading = false;
    }
  });
</script>

<main class="app-container">
  <header class="app-header">
    <div class="brand">
      <span class="logo-icon">🤖</span>
      <div class="brand-text">
        <h1>UWSCR::act-gram</h1>
        <span class="subtitle">AI-Native RPA Control Center</span>
      </div>
    </div>

    <nav class="nav-tabs">
      <button 
        class="tab-btn {activeTab === 'credentials' ? 'active' : ''}" 
        on:click={() => activeTab = 'credentials'}
      >
        🔑 認証設定
      </button>
      <button 
        class="tab-btn {activeTab === 'config' ? 'active' : ''}" 
        on:click={() => activeTab = 'config'}
        disabled={isLoading}
      >
        🧠 AIレイヤー設定
      </button>
      <button 
        class="tab-btn {activeTab === 'run' ? 'active' : ''}" 
        on:click={() => activeTab = 'run'}
        disabled={isLoading}
      >
        ⚡ スクリプト実行
      </button>
    </nav>
  </header>

  <div class="content-area">
    {#if isLoading}
      <div class="loader-container">
        <div class="spinner"></div>
        <p>設定をロード中...</p>
      </div>
    {:else}
      {#if activeTab === 'credentials'}
        <CredentialInput />
      {:else if activeTab === 'config'}
        <LayerConfig />
      {:else}
        <Console />
      {/if}
    {/if}
  </div>
</main>

<style>
  :global(body) {
    background-color: #0f111a;
    color: #eceff1;
    font-family: 'Inter', system-ui, -apple-system, sans-serif;
    margin: 0;
    padding: 0;
    overflow: hidden;
  }

  .app-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    box-sizing: border-box;
  }

  .app-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
    background: rgba(21, 25, 36, 0.8);
    backdrop-filter: blur(8px);
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .logo-icon {
    font-size: 1.8rem;
  }

  .brand-text h1 {
    margin: 0;
    font-size: 1.15rem;
    font-weight: 700;
    letter-spacing: 0.5px;
    background: linear-gradient(45deg, #535bf2, #8e44ad);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .subtitle {
    font-size: 0.75rem;
    color: #888;
  }

  .nav-tabs {
    display: flex;
    gap: 8px;
  }

  .tab-btn {
    background: transparent;
    border: none;
    color: #888;
    padding: 8px 16px;
    font-size: 0.85rem;
    font-weight: 600;
    border-radius: 6px;
    cursor: pointer;
    transition: background-color 0.2s, color 0.2s;
  }

  .tab-btn:hover:not(:disabled) {
    background: rgba(255, 255, 255, 0.05);
    color: #fff;
  }

  .tab-btn.active {
    background: rgba(83, 91, 242, 0.15);
    color: #7b83ff;
    border: 1px solid rgba(83, 91, 242, 0.25);
  }

  .tab-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .content-area {
    flex: 1;
    padding: 24px;
    overflow-y: auto;
  }

  .loader-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    gap: 16px;
    color: #888;
  }

  .spinner {
    width: 32px;
    height: 32px;
    border: 3px solid rgba(255, 255, 255, 0.1);
    border-radius: 50%;
    border-top-color: #535bf2;
    animation: spin 1s ease-in-out infinite;
  }

  .placeholder-card {
    background: rgba(255, 255, 255, 0.02);
    border: 1px dashed rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    padding: 40px;
    text-align: center;
    color: #888;
  }

  .placeholder-card h2 {
    color: #ccc;
    margin-top: 0;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }
</style>
