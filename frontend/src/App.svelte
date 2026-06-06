<script>
  import { onMount, onDestroy } from 'svelte';
  import { fade, fly } from 'svelte/transition';
  import { GetConfig, SaveUWSCRPath, SelectFile, RunInteractiveGuide, GetDefaultManualPath, GetImageBase64 } from '../wailsjs/go/main/App.js';
  import logoDark from './assets/images/logo-dark.png';
  import logoLight from './assets/images/logo-light.png';
  import * as wailsRuntime from '../wailsjs/runtime/runtime.js';
  import CredentialInput from './components/CredentialInput.svelte';
  import LayerConfig from './components/LayerConfig.svelte';
  import Console from './components/Console.svelte';
  import ManualCreator from './components/ManualCreator.svelte';
  import ScriptDeveloper from './components/ScriptDeveloper.svelte';
  import RecSimulator from './components/RecSimulator.svelte';

  let config = null;
  let isLoading = true;
  let activeTab = 'home'; // Default to portal home screen
  let isLightMode = false;
  let showSettingsDrawer = false;
  let showHelpDrawer = false;
  let selectedLanguage = 'ja';
  let uwscrPath = '';
  let currentLogDir = ''; // Store path for auto-slicer

  let isRunningGuide = false;
  let guideError = '';
  let guideSuccess = '';

  let manualSteps = [
    {
      step: 1,
      title: "APIキーとUWSCRの設定",
      desc: "actgramを起動したら、まずは画面右上の「設定」ギアアイコンをクリックして設定ドロワーを開きます。ご自身のGoogle/Gemini等のAPIキーを設定し、uwscr.exeのパス（未指定時は自動探索）を確認・保存してください。",
      imgKey: "step_1.png",
      imgSrc: ""
    },
    {
      step: 2,
      title: "スクリプトの実行と自己リファクタリング",
      desc: "「スクリプト実行」コンソールでは、UWSCRスクリプト（.uws）をリアルタイムに実行・監視できます。実行終了後、「ボトルネック分析 & 最適化提案」をクリックすることで、実測ログ（ミリ秒ファクト）に基づいて不要な固定SLEEPを自動で検出し、非同期ループ監視コードへのリファクタリング提案とファイル適用を行うことができます。",
      imgKey: "step_2.png",
      imgSrc: ""
    },
    {
      step: 3,
      title: "オート・スライサーによるマニュアル作成",
      desc: "「マニュアル作成」タブでは、レコーダーが保存した生の操作ログフォルダ（log.jsonが含まれる場所）を指定するだけで、システムが操作の区切りを自律分割して論理的なステップに自動分割します。自動生成されたUWSCRコードと、音声案内TTS、座標ズームと連動した対話型ガイドが瞬時に構築されます。",
      imgKey: "step_3.png",
      imgSrc: ""
    }
  ];

  onMount(async () => {
    try {
      config = await GetConfig();
      if (config) {
        uwscrPath = config.UWSCRPath || '';
      }

      // レコーダー停止時のイベントを検知してマニュアル生成画面に自動遷移
      wailsRuntime.EventsOn("recording_stopped", (logDir) => {
        console.log("[App] recording_stopped received:", logDir);
        currentLogDir = logDir;
        activeTab = 'manual';
      });
    } catch (e) {
      console.error('Failed to load configuration:', e);
    } finally {
      isLoading = false;
    }
  });

  onDestroy(() => {
    wailsRuntime.EventsOff("recording_stopped");
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

  async function loadHelpImages() {
    try {
      const baseDir = await GetDefaultManualPath();
      for (let i = 0; i < manualSteps.length; i++) {
        const imgPath = baseDir + "\\act_gram_guide\\images\\" + manualSteps[i].imgKey;
        const b64 = await GetImageBase64(imgPath);
        manualSteps[i].imgSrc = b64;
      }
      manualSteps = [...manualSteps];
    } catch (e) {
      console.error("Failed to load help images:", e);
    }
  }

  function openHelp() {
    showHelpDrawer = true;
    loadHelpImages();
  }

  async function startInteractiveGuide() {
    isRunningGuide = true;
    guideError = '';
    guideSuccess = '';
    try {
      await RunInteractiveGuide();
      guideSuccess = 'ガイドを実行しました！UWSCRの操作案内（メッセージボックス）に従ってください。';
    } catch (e) {
      console.error("Failed to run interactive guide:", e);
      guideError = 'ガイドの実行に失敗しました: ' + e;
    } finally {
      isRunningGuide = false;
    }
  }
</script>

<main class="app-container">
  <header class="app-header">
    <div class="brand">
      {#if activeTab !== 'home'}
        <button class="back-btn" on:click={() => activeTab = 'home'} aria-label="ホームへ戻る">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="19" y1="12" x2="5" y2="12"/>
            <polyline points="12 19 5 12 12 5"/>
          </svg>
        </button>
      {/if}
      <img src={isLightMode ? logoLight : logoDark} alt="actgram logo" class="logo-img" style="height: 28px; width: auto; object-fit: contain;" />
      <div class="brand-text">
        <h1>::UWSCR</h1>
        <span class="subtitle">Control Center</span>
      </div>
    </div>

    <div class="header-right">
      <div class="nav-shortcut-group">
        <button 
          class="nav-shortcut-btn" 
          class:active={activeTab === 'run'} 
          on:click={() => activeTab = 'run'} 
          title="PLAY: スクリプト実行・分析"
          aria-label="PLAY screen"
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
            <polygon points="5 3 19 12 5 21 5 3"/>
          </svg>
        </button>

        <button 
          class="nav-shortcut-btn" 
          class:active={activeTab === 'rec'} 
          on:click={() => activeTab = 'rec'} 
          title="REC: 本格マクロレコーダー"
          aria-label="REC screen"
        >
          <svg viewBox="0 0 24 24" fill="currentColor">
            <circle cx="12" cy="12" r="8"/>
          </svg>
        </button>

        <button 
          class="nav-shortcut-btn" 
          class:active={activeTab === 'dev'} 
          on:click={() => activeTab = 'dev'} 
          title="DEVELOP: スクリプト開発・自動修正"
          aria-label="DEVELOP screen"
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="16 18 22 12 16 6"/>
            <polyline points="8 6 2 12 8 18"/>
          </svg>
        </button>

        <button 
          class="nav-shortcut-btn" 
          class:active={activeTab === 'manual'} 
          on:click={() => activeTab = 'manual'} 
          title="MANUAL: マニュアル自動生成・再生"
          aria-label="MANUAL screen"
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
        </button>
      </div>

      <div class="divider"></div>

      <button class="settings-btn" on:click={openHelp} aria-label="Open help">
        <svg class="settings-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"/>
          <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/>
          <line x1="12" y1="17" x2="12.01" y2="17"/>
        </svg>
      </button>

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
        {#if activeTab === 'home'}
          <div class="home-grid" in:fade={{ duration: 150 }}>
            <button class="home-card" on:click={() => activeTab = 'run'}>
              <div class="card-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="5 3 19 12 5 21 5 3"/>
                </svg>
              </div>
              <div class="card-title">PLAY</div>
              <div class="card-desc">実行コンソール</div>
            </button>

            <button class="home-card" on:click={() => activeTab = 'rec'}>
              <div class="card-icon rec-icon">
                <svg viewBox="0 0 24 24" fill="currentColor">
                  <circle cx="12" cy="12" r="8"/>
                </svg>
              </div>
              <div class="card-title">REC</div>
              <div class="card-desc">マクロレコーダー</div>
            </button>

            <button class="home-card" on:click={() => activeTab = 'dev'}>
              <div class="card-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="16 18 22 12 16 6"/>
                  <polyline points="8 6 2 12 8 18"/>
                </svg>
              </div>
              <div class="card-title">DEVELOP</div>
              <div class="card-desc">スクリプト開発</div>
            </button>

            <button class="home-card" on:click={() => activeTab = 'manual'}>
              <div class="card-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                  <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
                </svg>
              </div>
              <div class="card-title">MANUAL</div>
              <div class="card-desc">マニュアル作成</div>
            </button>
          </div>
        {:else if activeTab === 'manual'}
          <ManualCreator bind:selectedLogDir={currentLogDir} />
        {:else if activeTab === 'dev'}
          <ScriptDeveloper />
        {:else if activeTab === 'rec'}
          <RecSimulator onRecordFinished={(logDir) => {
            currentLogDir = logDir;
            activeTab = 'manual';
          }} />
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

  <!-- Help Drawer Overlay & Panel -->
  {#if showHelpDrawer}
    <div class="drawer-overlay" on:click={() => showHelpDrawer = false} transition:fade={{ duration: 200 }}></div>
    <div class="settings-drawer" transition:fly={{ x: 500, duration: 300, opacity: 1 }}>
      <div class="drawer-header">
        <div class="drawer-title">
          <svg class="drawer-header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" style="width: 20px; height: 20px;">
            <circle cx="12" cy="12" r="10"/>
            <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/>
            <line x1="12" y1="17" x2="12.01" y2="17"/>
          </svg>
          <h2>ヘルプ & 使い方マニュアル</h2>
        </div>
        <button class="btn-close" on:click={() => showHelpDrawer = false} aria-label="Close help">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="width: 16px; height: 16px;">
            <line x1="18" y1="6" x2="6" y2="18"/>
            <line x1="6" y1="6" x2="18" y2="18"/>
          </svg>
        </button>
      </div>

      <div class="drawer-body">
        <div class="drawer-section">
          <h3>actgram の基本概念</h3>
          <p style="font-size: 0.8rem; line-height: 1.5; color: var(--text-secondary); margin: 0 0 16px 0;">
            UWSCR::actgram は、次世代RPA実行・開発コントロールセンターです。
            操作の実行、ボトルネックの自動リファクタリング、録画からのマニュアル・スクリプト自動生成などの機能を提供します。
          </p>
        </div>

        {#each manualSteps as step}
          <div class="drawer-section" style="border-top: 1px solid var(--border-color); padding-top: 20px;">
            <div style="display: flex; align-items: center; gap: 8px; margin-bottom: 8px;">
              <span style="background: var(--accent-color); color: var(--bg-primary); font-size: 0.7rem; font-weight: 700; padding: 2px 8px; border-radius: 10px;">Step {step.step}</span>
              <h4 style="margin: 0; font-size: 0.9rem; font-weight: 600; color: var(--text-primary);">{step.title}</h4>
            </div>
            <p style="font-size: 0.78rem; line-height: 1.5; color: var(--text-secondary); margin: 0 0 12px 0;">
              {step.desc}
            </p>
            {#if step.imgSrc}
              <div style="border-radius: 8px; overflow: hidden; border: 1px solid var(--border-color); background: #000; display: flex; justify-content: center; margin-bottom: 8px;">
                <img src={step.imgSrc} alt={step.title} style="max-width: 100%; height: auto; display: block;" />
              </div>
            {/if}
          </div>
        {/each}

        <!-- Interactive Demo Section -->
        <div class="drawer-section" style="border-top: 1px solid var(--border-color); padding-top: 20px; margin-bottom: 20px;">
          <h3>インタラクティブ実演ガイド</h3>
          <p style="font-size: 0.78rem; line-height: 1.5; color: var(--text-secondary); margin: 0 0 16px 0;">
            UWSCRが実際にデスクトップ上を操作して、actgramの各機能の概要を対話形式で案内します。
          </p>
          
          <button 
            class="btn-theme-toggle" 
            on:click={startInteractiveGuide} 
            disabled={isRunningGuide}
            style="background: var(--accent-color); color: var(--bg-primary); border-color: var(--accent-color); font-weight: 600; width: 100%; padding: 12px; font-size: 0.85rem;"
          >
            {#if isRunningGuide}
              実演ガイド起動中...
            {:else}
              インタラクティブ・ガイドを実行する (実演)
            {/if}
          </button>

          {#if guideSuccess}
            <div style="margin-top: 12px; padding: 8px 12px; background: rgba(46, 204, 113, 0.15); border: 1px solid rgba(46, 204, 113, 0.3); border-radius: 6px; font-size: 0.75rem; color: #2ecc71;">
              {guideSuccess}
            </div>
          {/if}
          {#if guideError}
            <div style="margin-top: 12px; padding: 8px 12px; background: rgba(231, 76, 60, 0.15); border: 1px solid rgba(231, 76, 60, 0.3); border-radius: 6px; font-size: 0.75rem; color: #e74c3c;">
              {guideError}
            </div>
          {/if}
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
    flex-shrink: 0;
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .back-btn {
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

  .back-btn:hover {
    color: var(--text-primary);
    border-color: var(--border-hover);
    background: var(--accent-soft);
  }

  .back-btn svg {
    width: 16px;
    height: 16px;
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

  .nav-shortcut-group {
    display: flex;
    align-items: center;
    gap: 6px;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid var(--border-color);
    padding: 3px;
    border-radius: 10px;
  }

  .nav-shortcut-btn {
    background: transparent;
    border: none;
    color: var(--text-secondary);
    width: 32px;
    height: 32px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    padding: 0;
  }

  .nav-shortcut-btn svg {
    width: 16px;
    height: 16px;
  }

  .nav-shortcut-btn:hover {
    color: var(--text-primary);
    background: rgba(255, 255, 255, 0.05);
    transform: translateY(-1px);
  }

  .nav-shortcut-btn:active {
    transform: translateY(0);
  }

  .nav-shortcut-btn.active {
    color: var(--accent-color);
    background: var(--accent-soft);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .nav-shortcut-btn.active:hover {
    background: var(--accent-soft);
  }

  .nav-shortcut-btn[aria-label="REC screen"] {
    color: var(--accent-red);
  }

  .nav-shortcut-btn[aria-label="REC screen"].active {
    background: rgba(255, 71, 87, 0.15);
    color: #ef4444;
  }

  .header-right .divider {
    width: 1px;
    height: 20px;
    background-color: var(--border-color);
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

  /* HOME Portal Screen Grid */
  .home-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 24px;
    max-width: 800px;
    width: 100%;
    margin: auto;
    padding: 40px 20px;
    box-sizing: border-box;
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    max-height: 540px;
  }

  .home-card {
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 16px;
    padding: 32px;
    cursor: pointer;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 16px;
    transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    box-shadow: var(--shadow-sm);
    width: 100%;
    max-width: 320px;
    height: 200px;
    box-sizing: border-box;
  }

  .home-card:hover {
    border-color: var(--accent-color);
    background: var(--accent-soft);
    transform: translateY(-4px);
    box-shadow: var(--shadow-md);
  }

  .card-icon {
    width: 52px;
    height: 52px;
    color: var(--text-primary);
    display: flex;
    align-items: center;
    justify-content: center;
    transition: transform 0.25s ease;
  }

  .home-card:hover .card-icon {
    transform: scale(1.1);
  }

  .card-icon svg {
    width: 100%;
    height: 100%;
  }

  .rec-icon {
    color: var(--accent-red);
  }

  .home-card:hover .rec-icon {
    color: #ef4444;
    filter: drop-shadow(0 0 8px rgba(239, 68, 68, 0.4));
  }

  .card-title {
    font-size: 1.15rem;
    font-weight: 700;
    letter-spacing: 1px;
    color: var(--text-primary);
    margin: 0;
  }

  .card-desc {
    font-size: 0.78rem;
    color: var(--text-secondary);
    margin: 0;
  }

  .content-area {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    background-color: var(--bg-primary);
    transition: background-color 0.3s ease;
  }

  .tab-content {
    flex: 1;
    width: 100%;
    max-width: 1400px;
    margin: 0 auto;
    padding: 20px 24px;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    min-height: 0;
    height: 100%;
    overflow: hidden;
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
