<script>
  import { onMount, onDestroy } from 'svelte';
  import { fade, fly } from 'svelte/transition';
  import { GetConfig, SaveUWSCRPath, SelectFile, RunInteractiveGuide, GetDefaultManualPath, GetImageBase64, StopScript, StopRecording, SetMiniMode, SaveKnowledgeDirAndURL, SyncUWSCRReference, RestoreDefaultReference, SelectDirectory } from '../wailsjs/go/main/App.js';
  import logoDark from './assets/images/logo-dark.png';
  import logoLight from './assets/images/logo-light.png';
  import * as wailsRuntime from '../wailsjs/runtime/runtime.js';
  import CredentialInput from './components/CredentialInput.svelte';
  import LayerConfig from './components/LayerConfig.svelte';
  import Console from './components/Console.svelte';
  import ManualCreator from './components/ManualCreator.svelte';
  import ScriptDeveloper from './components/ScriptDeveloper.svelte';
  import RecSimulator from './components/RecSimulator.svelte';
  import TourHighlight from './components/TourHighlight.svelte';


  let config = null;
  let isLoading = true;
  let activeTab = 'home'; // Default to portal home screen
  let isLightMode = false;
  let showSettingsModal = false; // settings drawer -> settings modal
  let settingsActiveTab = 'general'; // 'general', 'uwscr', 'api', 'ai'
  let showHelpDrawer = false;
  let selectedLanguage = 'ja';
  let uwscrPath = '';
  let knowledgeDir = '';
  let uwscrDocURL = '';
  let currentLogDir = ''; // Store path for auto-slicer
  let isMiniMode = false; // Layout state for small control panel
  let miniModeType = 'play'; // 'play' or 'record'
  let layerConfigRef = null;
  let isSavingAISettings = false;
  let aiSettingsStatusMessage = '';
  let isSavingPathSettings = false;
  let pathSettingsStatusMessage = '';

  // ズームモーダル用ステート
  let zoomImage = "";
  let zoomScale = 1;
  let zoomPosX = 0;
  let zoomPosY = 0;
  let isDragging = false;
  let dragStart = { x: 0, y: 0 };

  function closeZoomModal() {
    zoomImage = "";
  }

  function handleWheel(e) {
    e.preventDefault();
    const zoomFactor = 0.1;
    if (e.deltaY < 0) {
      zoomScale = Math.min(zoomScale + zoomFactor, 5);
    } else {
      zoomScale = Math.max(zoomScale - zoomFactor, 0.5);
    }
  }

  function handleMouseDown(e) {
    isDragging = true;
    dragStart = { x: e.clientX - zoomPosX, y: e.clientY - zoomPosY };
  }

  function handleMouseMove(e) {
    if (!isDragging) return;
    zoomPosX = e.clientX - dragStart.x;
    zoomPosY = e.clientY - dragStart.y;
  }

  function handleMouseUp() {
    isDragging = false;
  }

  function resetZoom() {
    zoomScale = 1;
    zoomPosX = 0;
    zoomPosY = 0;
  }

  function zoomIn() {
    zoomScale = Math.min(zoomScale + 0.2, 5);
  }

  function zoomOut() {
    zoomScale = Math.max(zoomScale - 0.2, 0.5);
  }

  async function copyImageToClipboard() {
    if (!zoomImage) return;
    try {
      let blob;
      if (zoomImage.startsWith('data:')) {
        const parts = zoomImage.split(',');
        const mime = parts[0].match(/:(.*?);/)[1];
        const bstr = atob(parts[1]);
        let n = bstr.length;
        const u8arr = new Uint8Array(n);
        while (n--) {
          u8arr[n] = bstr.charCodeAt(n);
        }
        blob = new Blob([u8arr], { type: mime });
      } else {
        const response = await fetch(zoomImage);
        blob = await response.blob();
      }
      
      const item = new ClipboardItem({ [blob.type]: blob });
      await navigator.clipboard.write([item]);
      alert("画像をクリップボードにコピーしました！");
    } catch (err) {
      console.error("Failed to copy image:", err);
      alert("画像のコピーに失敗しました。ブラウザのセキュリティ設定等を確認してください。");
    }
  }

  // ズームイベントハンドラを登録するため外側で定義
  function handleZoomEvent(e) {
    zoomImage = e.detail;
    zoomScale = 1;
    zoomPosX = 0;
    zoomPosY = 0;
  }

  let isRunningGuide = false;
  let guideError = '';
  let guideSuccess = '';

  // --- UIツアー (インタラクティブガイド) ---
  let showTour = false;
  let tourStep = 0;

  const tourSteps = [
    {
      title: "actgramへようこそ",
      desc: "このガイドではactgramの主要機能を画面で案内します。画面上部のタブでPLAY / REC / DEVELOP / MANUALの各機能にアクセスできます。",
      highlight: null,
      position: 'center'
    },
    {
      title: "環境設定",
      desc: "まず右上の歯車アイコンから【環境設定】を開いてください。\n・API認証キー：Gemini / Claude / ChatGPT のAPIキーを登録\n・パス設定：uwscr.exeと知識フォルダのパスを指定\n・AIモデル設定：使用するAIモデルを選択\n設定後【設定を保存】ボタンを押してください。",
      highlight: 'settings-btn',
      position: 'bottom-left'
    },
    {
      title: "PLAY — スクリプト実行",
      desc: "PLAYタブでは、UWSCRスクリプト（.uws）を選択して実行できます。\n・実行中はリアルタイムログが表示されます\n・実行後に【ボトルネック分析 & 最適化提案】でAIが自動的に無駄なSLEEP等を検出し、改善コードを提案します",
      highlight: 'play-tab',
      position: 'bottom'
    },
    {
      title: "REC — 操作記録",
      desc: "RECタブでは、マウス・キーボード操作をキャプチャして記録できます。\n・【記録開始】ボタンでactgramが最小化され、録画スタート\n・デスクトップで操作した内容が自動的に記録されます\n・停止するとマニュアル生成タブへ自動遷移します",
      highlight: 'rec-tab',
      position: 'bottom'
    },
    {
      title: "DEVELOP — AIスクリプト開発",
      desc: "DEVELOPタブでは、画面キャプチャとテキスト指示をAIに渡してUWSCRスクリプトを自動生成します。\n・AIがコードを生成→即座にテスト実行→エラーがあれば自己修復\n・【修正指示】ボタンで生成済みコードを追加指示で修正も可能",
      highlight: 'develop-tab',
      position: 'bottom'
    },
    {
      title: "MANUAL — マニュアル作成",
      desc: "MANUALタブでは、RECで記録した操作ログからHTMLマニュアルを自動生成します。\n・ログフォルダを選択するとAIが操作を分析・分割\n・対話型HTML + UWSCRガイドスクリプトが生成されます\n・TTS（音声案内）付きのプレミアムマニュアルも作成可能",
      highlight: 'manual-tab',
      position: 'bottom'
    },
    {
      title: "設定完了",
      desc: "これでactgramの主要機能をご確認いただきました。\n\n実際にお使いいただくには：\n1. 環境設定でAPIキーとuwscr.exeのパスを設定\n2. RECで操作を記録\n3. MANUALでマニュアルを自動生成",
      highlight: null,
      position: 'center'
    }
  ];

  function startTour() {
    showHelpDrawer = false;
    setTimeout(() => {
      showTour = true;
      tourStep = 0;
    }, 300);
  }

  function nextTourStep() {
    if (tourStep < tourSteps.length - 1) {
      tourStep++;
    } else {
      closeTour();
    }
  }

  function prevTourStep() {
    if (tourStep > 0) tourStep--;
  }

  function closeTour() {
    showTour = false;
    tourStep = 0;
  }

  function goToTourStep(i) {
    tourStep = i;
  }

  let isSyncingManual = false;
  let syncStatusMessage = '';


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
    // ズームイベントリスナー登録
    window.addEventListener('zoom-image', handleZoomEvent);

    try {
      config = await GetConfig();
      if (config) {
        uwscrPath = config.UWSCRPath || '';
        knowledgeDir = config.KnowledgeDir || '';
        uwscrDocURL = config.UWSCRDocURL || '';
      }

      // レコーダー停止時のイベントを検知してマニュアル生成画面に自動遷移
      wailsRuntime.EventsOn("recording_stopped", (logDir) => {
        console.log("[App] recording_stopped received:", logDir);
        currentLogDir = logDir;
        activeTab = 'manual';
      });

      // ミニモード切り替えイベントをリッスン
      wailsRuntime.EventsOn("mini_mode_changed", (isMini, modeType) => {
        console.log("[App] mini_mode_changed received:", isMini, modeType);
        isMiniMode = !!isMini;
        miniModeType = modeType || 'play';
      });

      // マニュアル同期イベントの監視
      wailsRuntime.EventsOn("manual_sync_started", () => {
        isSyncingManual = true;
        syncStatusMessage = 'オンラインからUWSCR最新マニュアルをクロール同期しています...';
      });

      wailsRuntime.EventsOn("manual_sync_finished", (res) => {
        isSyncingManual = false;
        if (res && res.success) {
          syncStatusMessage = '同期が正常に完了しました！最新のマニュアルが読み込まれました。';
        } else {
          syncStatusMessage = '同期エラー: ' + (res?.error || '不明なエラー');
        }
      });
    } catch (e) {
      console.error('Failed to load configuration:', e);
    } finally {
      isLoading = false;
    }
  });

  onDestroy(() => {
    window.removeEventListener('zoom-image', handleZoomEvent);
    wailsRuntime.EventsOff("recording_stopped");
    wailsRuntime.EventsOff("mini_mode_changed");
    wailsRuntime.EventsOff("manual_sync_started");
    wailsRuntime.EventsOff("manual_sync_finished");
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

  async function browseKnowledgeDir() {
    try {
      const selected = await SelectDirectory("知識フォルダを選択");
      if (selected) {
        knowledgeDir = selected;
      }
    } catch (e) {
      console.error('Failed to select knowledge directory:', e);
    }
  }

  async function handleSaveKnowledgeDirAndURL() {
    try {
      await SaveKnowledgeDirAndURL(knowledgeDir, uwscrDocURL);
      if (config) {
        config.KnowledgeDir = knowledgeDir;
        config.UWSCRDocURL = uwscrDocURL;
      }
      syncStatusMessage = '設定を保存しました。';
      setTimeout(() => {
        if (syncStatusMessage === '設定を保存しました。') syncStatusMessage = '';
      }, 3000);
    } catch (e) {
      console.error('Failed to save knowledge configuration:', e);
      syncStatusMessage = '保存エラー: ' + e;
    }
  }

  async function handleSaveAISettings() {
    isSavingAISettings = true;
    aiSettingsStatusMessage = '';
    try {
      // LayerConfig の保存
      if (layerConfigRef) {
        await layerConfigRef.applyConfig();
      }
      
      aiSettingsStatusMessage = 'AIモデル設定を保存しました！';
      setTimeout(() => {
        if (aiSettingsStatusMessage === 'AIモデル設定を保存しました！') {
          aiSettingsStatusMessage = '';
        }
      }, 3000);
    } catch (e) {
      console.error('Failed to save AI configurations:', e);
      aiSettingsStatusMessage = '保存失敗: ' + (e.message || e);
    } finally {
      isSavingAISettings = false;
    }
  }

  async function handleSavePathSettings() {
    isSavingPathSettings = true;
    pathSettingsStatusMessage = '';
    try {
      // 1. UWSCR パス保存
      await SaveUWSCRPath(uwscrPath);
      // 2. 知識フォルダ・URL保存
      await SaveKnowledgeDirAndURL(knowledgeDir, uwscrDocURL);
      if (config) {
        config.UWSCRPath = uwscrPath;
        config.KnowledgeDir = knowledgeDir;
        config.UWSCRDocURL = uwscrDocURL;
      }
      
      pathSettingsStatusMessage = 'パス設定を保存しました！';
      setTimeout(() => {
        if (pathSettingsStatusMessage === 'パス設定を保存しました！') {
          pathSettingsStatusMessage = '';
        }
      }, 3000);
    } catch (e) {
      console.error('Failed to save path settings:', e);
      pathSettingsStatusMessage = '保存失敗: ' + (e.message || e);
    } finally {
      isSavingPathSettings = false;
    }
  }

  async function triggerSyncManual() {
    syncStatusMessage = '';
    try {
      await SyncUWSCRReference();
    } catch (e) {
      console.error('Failed to trigger manual sync:', e);
      syncStatusMessage = '同期起動失敗: ' + e;
    }
  }

  async function triggerRestoreManual() {
    if (!confirm('UWSCRリファレンスとマニュアルを初期状態（デフォルト）に上書き復元しますか？（カスタムの変更は失われます）')) {
      return;
    }
    syncStatusMessage = 'マニュアルを初期化中...';
    try {
      await RestoreDefaultReference();
      syncStatusMessage = 'マニュアルをデフォルト状態に復元（初期化）しました。';
    } catch (e) {
      console.error('Failed to restore manual:', e);
      syncStatusMessage = '復元失敗: ' + e;
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
    // UIツアーモードで起動（UWSCRスクリプト実行は行わない）
    startTour();
  }

  async function handleStopScript() {
    try {
      if (miniModeType === 'record') {
        const logDir = await StopRecording();
        console.log("[App] Recording stopped via mini mode, logDir =", logDir);
      } else {
        await StopScript();
        await SetMiniMode(false, 'play');
      }
    } catch (e) {
      console.error("Failed to stop script/recording:", e);
    }
  }

  let importedStepsForDev = null;
</script>

<main class="app-container">
  {#if !isMiniMode}
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
          id="play-tab"
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
          id="rec-tab"
          class="nav-shortcut-btn" 
          class:active={activeTab === 'rec'} 
          on:click={() => activeTab = 'rec'} 
          title="REC: 本格操作レコーダー"
          aria-label="REC screen"
        >
          <svg viewBox="0 0 24 24" fill="currentColor">
            <circle cx="12" cy="12" r="8"/>
          </svg>
        </button>

        <button 
          id="develop-tab"
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
          id="manual-tab"
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

      <button id="settings-btn" class="settings-btn" on:click={() => showSettingsModal = true} aria-label="Open settings">
        <svg class="settings-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="3"/>
          <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/>
        </svg>
      </button>
    </div>
  </header>
  {/if}

  <div class="content-area">
    {#if isLoading}
      <div class="loader-container">
        <div class="spinner"></div>
        <p>設定をロード中...</p>
      </div>
    {:else}
      {#if isMiniMode}
        <div class="mini-control-panel" class:record-mode={miniModeType === 'record'} in:fade={{ duration: 150 }}>
          <div class="mini-status">
            {#if miniModeType === 'record'}
              <span class="status-dot record-dot animate-pulse"></span>
              <span class="status-text record-text">操作記録中...</span>
            {:else}
              <span class="status-dot play-dot animate-pulse"></span>
              <span class="status-text play-text">UWSCR再生中...</span>
            {/if}
          </div>
          <div class="mini-actions">
            <button class="btn-mini-stop" class:btn-record-stop={miniModeType === 'record'} on:click={handleStopScript} title={miniModeType === 'record' ? '記録を停止して保存' : 'スクリプトを強制停止'}>
              <svg viewBox="0 0 24 24" fill="currentColor" width="16" height="16">
                <rect x="4" y="4" width="16" height="16" rx="1.5"/>
              </svg>
              {miniModeType === 'record' ? '記録停止' : '再生停止'}
            </button>
            <button class="btn-mini-restore" on:click={() => SetMiniMode(false, miniModeType)} title="通常サイズに復元">
              復元
            </button>
          </div>
        </div>
      {/if}

      <div class="tab-content" style={isMiniMode ? "display: none;" : ""}>
        <div class="tab-pane" class:hidden={activeTab !== 'home'}>
          <div class="home-grid">
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
              <div class="card-desc">操作レコーダー</div>
            </button>

            <button class="home-card" on:click={() => activeTab = 'dev'}>
              <div class="card-icon dev-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="16 18 22 12 16 6"/>
                  <polyline points="8 6 2 12 8 18"/>
                </svg>
              </div>
              <div class="card-title">DEVELOP</div>
              <div class="card-desc">スクリプト開発</div>
            </button>

            <button class="home-card" on:click={() => activeTab = 'manual'}>
              <div class="card-icon manual-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                  <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
                </svg>
              </div>
              <div class="card-title">MANUAL</div>
              <div class="card-desc">マニュアル作成</div>
            </button>
          </div>
        </div>

        <div class="tab-pane" class:hidden={activeTab !== 'manual'}>
          <ManualCreator 
            bind:selectedLogDir={currentLogDir} 
            on:exportToDev={(event) => {
              importedStepsForDev = event.detail;
              activeTab = 'dev';
            }}
          />
        </div>

        <div class="tab-pane" class:hidden={activeTab !== 'dev'}>
          <ScriptDeveloper bind:importedSteps={importedStepsForDev} bind:activeTab={activeTab} />
        </div>

        <div class="tab-pane" class:hidden={activeTab !== 'rec'}>
          <RecSimulator onRecordFinished={(logDir) => {
            currentLogDir = logDir;
            activeTab = 'manual';
          }} />
        </div>

        <div class="tab-pane" class:hidden={activeTab !== 'run'}>
          <Console />
        </div>
      </div>
    {/if}
  </div>

  <!-- Settings Modal Overlay & Panel -->
  {#if showSettingsModal}
    <div class="modal-overlay" on:click={() => showSettingsModal = false} transition:fade={{ duration: 200 }}></div>
    <div class="settings-modal" transition:fly={{ y: 50, duration: 300, opacity: 0.5 }}>
      <div class="modal-header">
        <div class="modal-title">
          <svg class="modal-header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="3"/>
            <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/>
          </svg>
          <h2>環境設定</h2>
        </div>
        <button class="btn-close" on:click={() => showSettingsModal = false} aria-label="Close settings">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"/>
            <line x1="6" y1="6" x2="18" y2="18"/>
          </svg>
        </button>
      </div>

      <div class="modal-body-container">
        <!-- Sidebar Tabs -->
        <div class="modal-sidebar">
          <button class="sidebar-tab-btn" class:active={settingsActiveTab === 'general'} on:click={() => settingsActiveTab = 'general'}>
            全般設定
          </button>
          <button class="sidebar-tab-btn" class:active={settingsActiveTab === 'path_manual'} on:click={() => settingsActiveTab = 'path_manual'}>
            パス・マニュアル設定
          </button>
          <button class="sidebar-tab-btn" class:active={settingsActiveTab === 'ai'} on:click={() => settingsActiveTab = 'ai'}>
            AIモデル設定
          </button>
          <button class="sidebar-tab-btn" class:active={settingsActiveTab === 'api'} on:click={() => settingsActiveTab = 'api'}>
            API認証キー
          </button>
        </div>

        <!-- Content Area -->
        <div class="modal-content-panel">
          {#if settingsActiveTab === 'general'}
            <div class="modal-tab-section" transition:fade={{ duration: 150 }}>
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
          {:else if settingsActiveTab === 'path_manual'}
            <div class="modal-tab-section" transition:fade={{ duration: 150 }}>
              <h3>パス・マニュアル設定</h3>
              
              <!-- 1. UWSCR Path -->
              <div class="path-setting-block">
                <div class="path-setting-label-row">
                  <span class="setting-name">UWSCR本体のパス (uwscr.exe)</span>
                  <span class="setting-desc">スクリプトを実行する UWSCR 実行ファイルの絶対パスを指定します。未指定時は自動探索されます。</span>
                </div>
                <div class="path-setting-input-row">
                  <input
                    id="drawer-uwscr-path"
                    type="text"
                    class="path-input-full"
                    bind:value={uwscrPath}
                    placeholder="未指定時は自動探索されます"
                  />
                  <button class="btn-path-browse" on:click={browseUWSCRPath}>選択...</button>
                </div>
              </div>

              <!-- 2. Knowledge Path -->
              <div class="path-setting-block">
                <div class="path-setting-label-row">
                  <span class="setting-name">知識フォルダのパス</span>
                  <span class="setting-desc">仕様書やマニュアル、自己学習データを保存する知識フォルダの絶対パスを指定します。</span>
                </div>
                <div class="path-setting-input-row">
                  <input
                    id="drawer-knowledge-dir"
                    type="text"
                    class="path-input-full"
                    bind:value={knowledgeDir}
                    placeholder="未指定時はデフォルト (knowledge) が使われます"
                  />
                  <button class="btn-path-browse" on:click={browseKnowledgeDir}>参照...</button>
                </div>
              </div>

              <!-- 3. Document URL -->
              <div class="path-setting-block" style="margin-bottom: 4px;">
                <div class="path-setting-label-row">
                  <span class="setting-name">UWSCRドキュメントURL</span>
                  <span class="setting-desc">最新マニュアルをクローリングするオンラインのベースURLを指定します。</span>
                </div>
                <div class="path-setting-input-row">
                  <input
                    id="drawer-uwscr-doc-url"
                    type="text"
                    class="path-input-full"
                    bind:value={uwscrDocURL}
                    placeholder="https://stuncloud.github.io/UWSCR/"
                    style="flex: 1;"
                  />
                </div>
              </div>


              <!-- Unified Path Settings Save Button -->
              <div style="display: flex; gap: 10px; justify-content: flex-end; margin-bottom: 24px; align-items: center; border-bottom: 1px solid var(--border-color); padding-bottom: 20px;">
                {#if pathSettingsStatusMessage}
                  <span class="status-msg" style="font-size: 0.8rem; color: var(--accent-color); font-weight: 500; margin-right: 10px;">{pathSettingsStatusMessage}</span>
                {/if}
                <button class="btn-theme-toggle" on:click={handleSavePathSettings} disabled={isSavingPathSettings} style="background: var(--accent-color); color: var(--bg-primary); border-color: var(--accent-color); font-weight: 600; padding: 10px 24px; border-radius: 6px;">
                  {#if isSavingPathSettings}
                    保存中...
                  {:else}
                    設定を保存
                  {/if}
                </button>
              </div>

              <!-- 4. Manual Sync and Restore -->
              <div class="setting-item flex-column" style="align-items: stretch; gap: 12px; padding-top: 10px;">
                <div class="setting-info">
                  <span class="setting-name">UWSCRマニュアル・辞書の更新</span>
                  <span class="setting-desc">ローカルのマニュアルファイルを操作します。</span>
                </div>
                <div style="display: flex; gap: 10px; margin-top: 5px;">
                  <button 
                    class="btn-theme-toggle" 
                    on:click={triggerSyncManual} 
                    disabled={isSyncingManual}
                    style="flex: 1; justify-content: center; background: rgba(16, 185, 129, 0.1); color: rgb(16, 185, 129); border-color: rgb(16, 185, 129);"
                  >
                    {#if isSyncingManual}
                      同期中...
                    {:else}
                      🔄 オンラインから同期 (最新化)
                    {/if}
                  </button>
                  <button 
                    class="btn-theme-toggle" 
                    on:click={triggerRestoreManual}
                    style="flex: 1; justify-content: center; background: rgba(239, 68, 68, 0.1); color: rgb(239, 68, 68); border-color: rgb(239, 68, 68);"
                  >
                    ⚠️ デフォルトに復元 (初期化)
                  </button>
                </div>
                {#if syncStatusMessage}
                  <div style="font-size: 0.75rem; color: var(--text-secondary); margin-top: 5px; padding: 8px; background: rgba(128, 128, 128, 0.05); border-radius: 4px; border-left: 3px solid var(--accent-color); word-break: break-all;">
                    {syncStatusMessage}
                  </div>
                {/if}
              </div>
            </div>

          {:else if settingsActiveTab === 'api'}
            <div class="modal-tab-section" transition:fade={{ duration: 150 }}>
              <CredentialInput />
            </div>

          {:else if settingsActiveTab === 'ai'}
            <div class="modal-tab-section" transition:fade={{ duration: 150 }}>
              <LayerConfig bind:this={layerConfigRef} />
              
              <div style="display: flex; gap: 10px; justify-content: flex-end; margin-top: 24px; align-items: center; border-top: 1px solid var(--border-color); padding-top: 20px;">
                {#if aiSettingsStatusMessage}
                  <span class="status-msg" style="font-size: 0.8rem; color: var(--accent-color); font-weight: 500; margin-right: 10px;">{aiSettingsStatusMessage}</span>
                {/if}
                <button class="btn-theme-toggle" on:click={handleSaveAISettings} disabled={isSavingAISettings} style="background: var(--accent-color); color: var(--bg-primary); border-color: var(--accent-color); font-weight: 600; padding: 10px 24px; border-radius: 6px;">
                  {#if isSavingAISettings}
                    保存中...
                  {:else}
                    設定を保存
                  {/if}
                </button>
              </div>
            </div>
          {/if}
        </div>
      </div>
    </div>
  {/if}

  <!-- Image Zoom Modal -->
  {#if zoomImage}
    <div class="zoom-modal-overlay" on:click={closeZoomModal} transition:fade={{ duration: 150 }}>
      <div class="zoom-modal-window" on:click|stopPropagation>
        <div class="zoom-modal-header">
          <h3>画像プレビュー (ホイールで拡大縮小、ドラッグで移動)</h3>
          <div class="zoom-controls">
            <button class="btn-zoom-action" on:click={copyImageToClipboard} title="クリップボードにコピー">
              コピー
            </button>
            <button class="btn-zoom-action" on:click={zoomIn} title="拡大">
              ＋
            </button>
            <button class="btn-zoom-action" on:click={zoomOut} title="縮小">
              －
            </button>
            <button class="btn-zoom-action" on:click={resetZoom} title="リセット">
              リセット
            </button>
            <button class="btn-zoom-close" on:click={closeZoomModal} aria-label="プレビューを閉じる">
              ✕
            </button>
          </div>
        </div>
        <div 
          class="zoom-viewport" 
          on:wheel={handleWheel}
          on:mousedown={handleMouseDown}
          on:mousemove={handleMouseMove}
          on:mouseup={handleMouseUp}
          on:mouseleave={handleMouseUp}
        >
          <img 
            src={zoomImage} 
            alt="Zoom preview" 
            style="transform: translate({zoomPosX}px, {zoomPosY}px) scale({zoomScale}); cursor: {isDragging ? 'grabbing' : 'grab'}; transition: transform {isDragging ? '0s' : '0.1s'} ease-out;"
            draggable="false"
          />
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
            actgram::UWSCR は、次世代RPA実行・開発コントロールセンターです。
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
          <h3>インタラクティブ UIガイド</h3>
          <p style="font-size: 0.78rem; line-height: 1.5; color: var(--text-secondary); margin: 0 0 16px 0;">
            actgramのUIを操作しながら、各機能を対話形式で案内します。actgramの画面上にガイドカードが表示され、ハイライトで各要素を説明します。
          </p>
          
          <button 
            class="btn-theme-toggle guide-start-btn"
            on:click={startInteractiveGuide} 
          >
            UIガイドをスタート
          </button>
        </div>
      </div>
    </div>
  {/if}
</main>

<!-- ====== インタラクティブ UI ツアー ====== -->
{#if showTour}
  {@const step = tourSteps[tourStep]}
  
  {#if step.highlight}
    <TourHighlight targetId={step.highlight} />
  {/if}

  <div
    class="tour-overlay"
    on:click|self={closeTour}
    transition:fade={{ duration: 200 }}
  >
    <div
      class="tour-card {step.position === 'center' ? 'tour-card-center' : step.position === 'bottom-left' ? 'tour-card-bottom-left' : 'tour-card-bottom'}"
      on:click|stopPropagation
      transition:fly={{ y: 20, duration: 250 }}
    >
      <div class="tour-card-header">
        <span class="tour-step-badge">{tourStep + 1} / {tourSteps.length}</span>
        <button class="tour-close-btn" on:click={closeTour} aria-label="ガイドを閉じる">✕</button>
      </div>

      <h3 class="tour-card-title">{step.title}</h3>

      <div class="tour-card-desc">
        {#each step.desc.split('\n') as line}
          <p>{line}</p>
        {/each}
      </div>

      <div class="tour-dots">
        {#each tourSteps as _, i}
          <button
            class="tour-dot {i === tourStep ? 'tour-dot-active' : ''}"
            on:click={() => goToTourStep(i)}
            aria-label="ステップ{i+1}へ"
          ></button>
        {/each}
      </div>

      <div class="tour-nav">
        <button
          class="tour-btn tour-btn-secondary"
          on:click={prevTourStep}
          disabled={tourStep === 0}
        >
          ← 前へ
        </button>
        
        <button
          class="tour-btn tour-btn-primary"
          on:click={nextTourStep}
        >
          {tourStep === tourSteps.length - 1 ? '✓ 完了' : '次へ →'}
        </button>
      </div>
    </div>
  </div>
{/if}


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
    grid-template-columns: repeat(4, 1fr);
    gap: 24px;
    max-width: 1000px;
    width: 100%;
    margin: auto;
    padding: 40px 20px;
    box-sizing: border-box;
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

  .guide-start-btn {
    width: 100%;
    padding: 11px 16px;
    font-size: 0.82rem;
    font-weight: 600;
    background: var(--accent-color);
    color: var(--bg-primary);
    border-color: var(--accent-color);
    text-align: center;
  }

  .guide-start-btn:hover {
    background: var(--accent-hover, var(--accent-color));
    border-color: var(--accent-hover, var(--accent-color));
    opacity: 0.9;
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

  .hidden {
    display: none !important;
  }

  .tab-pane {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  /* ミニモード用スタイル */
  .mini-control-panel {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100vw;
    padding: 4px 8px; /* Tighter padding */
    background: linear-gradient(135deg, #1e1e24 0%, #16161b 100%);
    box-sizing: border-box;
    border-radius: 0px;
    gap: 4px; /* Tighter gap */
    position: fixed;
    top: 0;
    left: 0;
    z-index: 99999;
  }

  :global(html.light-mode) .mini-control-panel {
    background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  }

  .mini-status {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
  }

  .status-dot.record-dot {
    background-color: #ef4444;
    box-shadow: 0 0 8px #ef4444;
  }

  .status-dot.play-dot {
    background-color: #10b981;
    box-shadow: 0 0 8px #10b981;
  }

  .animate-pulse {
    animation: pulse 1.5s infinite ease-in-out;
  }

  @keyframes pulse {
    0%, 100% { opacity: 0.4; transform: scale(0.9); }
    50% { opacity: 1; transform: scale(1.1); }
  }

  .status-text {
    font-size: 0.75rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .status-text.record-text {
    color: #ef4444;
  }

  .status-text.play-text {
    color: #10b981;
  }

  .mini-actions {
    display: flex;
    gap: 8px;
    width: 100%;
    justify-content: center;
  }

  .btn-mini-stop {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 4px;
    padding: 4px 10px; /* Slimmer button */
    background: #3b82f6;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 0.7rem;
    font-weight: 600;
    cursor: pointer;
    transition: background 0.15s ease;
  }

  .btn-mini-stop:hover {
    background: #2563eb;
  }

  .btn-mini-stop.btn-record-stop {
    background: #ef4444;
  }

  .btn-mini-stop.btn-record-stop:hover {
    background: #dc2626;
  }

  .btn-mini-restore {
    padding: 4px 10px; /* Slimmer button */
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-radius: 4px;
    font-size: 0.7rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .btn-mini-restore:hover {
    border-color: var(--border-hover);
    background: var(--accent-soft);
  }

  /* Settings Modal Styles */
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(4px);
    z-index: 1000;
  }

  .settings-modal {
    position: fixed;
    top: 10%;
    left: 50%;
    transform: translate(-50%, 0);
    width: 760px;
    height: 80vh;
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    box-shadow: var(--shadow-lg);
    z-index: 1001;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    box-sizing: border-box;
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
    border-bottom: 1px solid var(--border-color);
    background: rgba(255, 255, 255, 0.01);
  }

  .modal-title {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .modal-header-icon {
    width: 18px;
    height: 18px;
    color: var(--text-primary);
  }

  .modal-header h2 {
    margin: 0;
    font-size: 1.05rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .modal-body-container {
    display: flex;
    flex: 1;
    overflow: hidden;
    height: 100%;
  }

  .modal-sidebar {
    width: 200px;
    border-right: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.05);
    padding: 16px 8px;
    display: flex;
    flex-direction: column;
    gap: 6px;
    box-sizing: border-box;
  }

  .sidebar-tab-btn {
    background: transparent;
    border: none;
    color: var(--text-secondary);
    padding: 10px 14px;
    text-align: left;
    border-radius: 6px;
    font-size: 0.8rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .sidebar-tab-btn:hover {
    color: var(--text-primary);
    background: rgba(255, 255, 255, 0.03);
  }

  .sidebar-tab-btn.active {
    color: var(--accent-color);
    background: var(--accent-soft);
    font-weight: 600;
  }

  .modal-content-panel {
    flex: 1;
    padding: 24px;
    overflow-y: auto;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
  }

  .modal-tab-section {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: 100%;
  }

  .modal-tab-section h3 {
    margin: 0 0 4px 0;
    font-size: 0.95rem;
    font-weight: 600;
    color: var(--text-primary);
    border-bottom: 1px solid var(--border-color);
    padding-bottom: 8px;
  }

  /* Clean up embedded cards styling inside settings modal */
  .modal-content-panel :global(.card) {
    background: transparent !important;
    border: none !important;
    padding: 0 !important;
    box-shadow: none !important;
    margin: 0 !important;
    max-width: 100% !important;
  }

  .modal-content-panel :global(.card-header) {
    border-bottom: none !important;
    padding-bottom: 0 !important;
    margin-bottom: 16px !important;
  }

  .modal-content-panel :global(.header-icon) {
    display: none !important;
  }

  .modal-content-panel :global(.layers-grid) {
    grid-template-columns: 1fr !important;
    gap: 16px !important;
  }

  .modal-content-panel :global(.layer-card) {
    padding: 16px !important;
    border-radius: 8px !important;
    background: rgba(0, 0, 0, 0.08) !important;
  }

  :global(html.light-mode) .modal-content-panel :global(.layer-card) {
    background: rgba(0, 0, 0, 0.02) !important;
  }

  /* Image Zoom Modal Styles */
  .zoom-modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.8);
    backdrop-filter: blur(6px);
    z-index: 2000;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .zoom-modal-window {
    width: 85vw;
    height: 85vh;
    background: #111;
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
  }

  .zoom-modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 20px;
    background: #181818;
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  }

  .zoom-modal-header h3 {
    margin: 0;
    font-size: 0.9rem;
    font-weight: 600;
    color: #eee;
  }

  .zoom-controls {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .btn-zoom-action {
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    color: #eee;
    padding: 6px 12px;
    border-radius: 6px;
    font-size: 0.75rem;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .btn-zoom-action:hover {
    background: rgba(255, 255, 255, 0.15);
    border-color: rgba(255, 255, 255, 0.2);
  }

  .btn-zoom-close {
    background: rgba(239, 68, 68, 0.15);
    border: 1px solid rgba(239, 68, 68, 0.25);
    color: #f87171;
    width: 28px;
    height: 28px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 6px;
    font-size: 0.8rem;
    cursor: pointer;
    transition: all 0.15s ease;
    padding: 0;
  }

  .btn-zoom-close:hover {
    background: rgba(239, 68, 68, 0.3);
    color: #f87171;
  }

  .zoom-viewport {
    flex: 1;
    overflow: hidden;
    position: relative;
    background: #090909;
    display: flex;
    align-items: center;
    justify-content: center;
    user-select: none;
  }

  .zoom-viewport img {
    max-width: 90%;
    max-height: 90%;
    object-fit: contain;
    transform-origin: center center;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
  }

  /* ========== インタラクティブ UI ツアー ========== */
  .tour-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.45);
    z-index: 9990;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .tour-card {
    position: fixed;
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    padding: 24px 28px;
    width: 360px;
    max-width: calc(100vw - 40px);
    box-shadow: var(--shadow-md);
    z-index: 9995;
  }

  .tour-card-center {
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }

  .tour-card-bottom {
    top: 64px;
    right: 16px;
  }

  .tour-card-bottom-left {
    top: 64px;
    right: 16px;
  }

  .tour-card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 14px;
  }

  .tour-step-badge {
    font-size: 0.68rem;
    font-weight: 500;
    color: var(--text-secondary);
    background: transparent;
    padding: 2px 0;
    letter-spacing: 0.04em;
  }

  .tour-close-btn {
    background: none;
    border: none;
    color: var(--text-secondary);
    cursor: pointer;
    font-size: 0.85rem;
    padding: 4px 6px;
    border-radius: 4px;
    line-height: 1;
    transition: color 0.15s, background 0.15s;
    font-weight: 300;
  }

  .tour-close-btn:hover {
    color: var(--text-primary);
    background: var(--hover-bg, rgba(128,128,128,0.1));
  }

  .tour-card-title {
    font-size: 0.95rem;
    font-weight: 600;
    color: var(--text-primary);
    margin: 0 0 10px 0;
    line-height: 1.35;
  }

  .tour-card-desc {
    font-size: 0.78rem;
    color: var(--text-secondary);
    line-height: 1.65;
    margin-bottom: 20px;
  }

  .tour-card-desc p {
    margin: 0 0 3px 0;
  }

  .tour-card-desc p:empty {
    height: 6px;
  }

  .tour-dots {
    display: flex;
    gap: 6px;
    justify-content: center;
    margin-bottom: 18px;
  }

  .tour-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    border: 1px solid var(--border-color);
    background: transparent;
    cursor: pointer;
    padding: 0;
    transition: background 0.18s, border-color 0.18s;
  }

  .tour-dot-active {
    background: var(--accent-color);
    border-color: var(--accent-color);
  }

  .tour-dot:hover:not(.tour-dot-active) {
    background: var(--border-color);
  }

  .tour-nav {
    display: flex;
    gap: 8px;
    justify-content: space-between;
  }

  .tour-btn {
    flex: 1;
    padding: 9px 0;
    border-radius: 6px;
    font-size: 0.8rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .tour-btn-primary {
    background: var(--accent-color);
    color: var(--bg-primary);
    border: 1px solid var(--accent-color);
  }

  .tour-btn-primary:hover {
    background: var(--accent-hover, var(--accent-color));
    border-color: var(--accent-hover, var(--accent-color));
  }

  .tour-btn-primary:active {
    opacity: 0.9;
  }

  .tour-btn-secondary {
    background: transparent;
    color: var(--text-secondary);
    border: 1px solid var(--border-color);
  }

  .tour-btn-secondary:hover:not(:disabled) {
    background: var(--hover-bg, rgba(128,128,128,0.08));
    color: var(--text-primary);
  }

  .tour-btn-secondary:disabled {
    opacity: 0.3;
    cursor: not-allowed;
  }

  /* ========== パス設定ブロック ========== */
  .path-setting-block {
    display: flex;
    flex-direction: column;
    gap: 6px;
    padding: 14px 0;
    border-bottom: 1px solid var(--border-color);
  }

  .path-setting-block:first-of-type {
    padding-top: 4px;
  }

  .path-setting-label-row {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .path-setting-input-row {
    display: flex;
    gap: 8px;
    align-items: center;
    width: 100%;
  }

  .path-input-full {
    flex: 1;
    min-width: 0;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 8px 12px;
    font-size: 0.78rem;
    font-family: Consolas, 'Courier New', monospace;
    color: var(--text-primary);
    width: 100%;
    box-sizing: border-box;
    transition: border-color 0.15s;
  }

  .path-input-full:focus {
    outline: none;
    border-color: var(--accent-color);
  }

  .btn-path-browse {
    flex-shrink: 0;
    background: transparent;
    border: 1px solid var(--border-color);
    color: var(--text-secondary);
    border-radius: 6px;
    padding: 7px 14px;
    font-size: 0.78rem;
    cursor: pointer;
    transition: all 0.15s;
    white-space: nowrap;
  }

  .btn-path-browse:hover {
    border-color: var(--accent-color);
    color: var(--accent-color);
  }

  /* ライトモード対応 */
  :global(html.light-mode) .tour-overlay {
    background: rgba(0, 0, 0, 0.3);
  }

  :global(html.light-mode) .tour-card {
    box-shadow: 0 4px 24px rgba(0, 0, 0, 0.12);
  }
</style>

