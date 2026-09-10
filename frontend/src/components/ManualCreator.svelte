<script>
  import { onMount, createEventDispatcher } from 'svelte';
  import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';
  import appIcon from '../assets/images/appicon.png';

  const App = window.go.main.App;
  const dispatch = createEventDispatcher();

  export let selectedLogDir = ""; // 操作ログディレクトリのパス

  let currentStep = 0;
  let steps = []; // ステップ一覧
  let isGenerating = false;
  let isExecuting = false;
  let isExecutingAll = false;
  let isSaving = false;
  let isExporting = false;
  let isTTSPlaying = false;
  let showCopilot = true; // 業務アシスタントペイン表示切り替え
  let userQuestion = "";
  let chatHistory = [];
  let statusMessage = "";

  let activeStepBase64 = "";
  let imageLoadError = false;
  let lastProcessedDir = "";

  $: if (selectedLogDir && selectedLogDir !== lastProcessedDir) {
    lastProcessedDir = selectedLogDir;
    loadOrGenerate();
  }

  // アクティブステップの切り替え監視
  $: activeStep = steps[currentStep] || null;

  $: {
    if (activeStep && activeStep.image_path) {
      imageLoadError = false;
      App.GetImageBase64(activeStep.image_path)
        .then(res => {
          if (res) {
            activeStepBase64 = res;
          } else {
            activeStepBase64 = "";
            imageLoadError = true;
          }
        })
        .catch(err => {
          console.error("Failed to load base64 image:", err);
          activeStepBase64 = "";
          imageLoadError = true;
        });
    } else {
      activeStepBase64 = "";
      imageLoadError = activeStep ? true : false;
    }
  }

  async function loadOrGenerate() {
    if (!selectedLogDir) return;
    try {
      // 1. まず保存済みの scenario.json があるか確認
      const savedSteps = await App.LoadManualScenario(selectedLogDir);
      if (savedSteps && savedSteps.length > 0) {
        steps = savedSteps;
        currentStep = 0;
        showStatus("保存済みシナリオを読み込みました");
        return;
      }
    } catch (e) {
      // 未保存の場合は自動生成を実行
    }
    handleAutoGenerate();
  }

  // AI自動スライス呼び出し
  async function handleAutoGenerate() {
    if (!selectedLogDir) {
      alert("録画ログの保存されているフォルダを入力または選択してください。");
      return;
    }

    isGenerating = true;
    chatHistory = [];
    try {
      const result = await App.GenerateScenarioFromLog(selectedLogDir);
      if (result && result.length > 0) {
        steps = result;
        currentStep = 0;
        showStatus("操作ログからシナリオを自動生成しました");
      } else {
        alert("操作ログから有効なステップを検出できませんでした。");
      }
    } catch (error) {
      console.error("シナリオ自動生成に失敗しました: ", error);
      alert("自動生成中にエラーが発生しました: " + (error.message || error));
    } finally {
      isGenerating = false;
    }
  }

  // ディレクトリ選択ダイアログ
  async function browseLogDir() {
    try {
      const selected = await App.SelectDirectory("録画ログ（log.jsonのあるフォルダ）を選択してください");
      if (selected) {
        selectedLogDir = selected;
      }
    } catch (err) {
      console.error("フォルダ選択エラー:", err);
    }
  }

  // シナリオの保存
  async function handleSaveScenario() {
    if (!selectedLogDir || steps.length === 0) {
      alert("保存するステップまたはログフォルダが指定されていません。");
      return;
    }
    isSaving = true;
    try {
      await App.SaveManualScenario(selectedLogDir, JSON.stringify(steps));
      showStatus("シナリオを保存しました！ (scenario.json)");
    } catch (err) {
      console.error("シナリオ保存エラー:", err);
      alert("保存に失敗しました: " + (err.message || err));
    } finally {
      isSaving = false;
    }
  }

  // ステップ単体のテスト実行
  async function handleExecuteStepOnly() {
    if (isExecuting || !activeStep) return;
    isExecuting = true;
    try {
      const result = await App.ExecuteStep(currentStep);
      if (result) {
        showStatus(`ステップ ${currentStep + 1} のテスト実行が完了しました。`);
      }
    } catch (err) {
      console.error("ステップ実行エラー: ", err);
      alert("実行中にエラーが発生しました: " + (err.message || err));
    } finally {
      isExecuting = false;
    }
  }

  // 全ステップの一括自動実行（シミュレーション）
  async function handleExecuteAll() {
    if (isExecutingAll || steps.length === 0) return;
    if (!confirm(`全 ${steps.length} ステップを一連の自動化フローとして連続実行しますか？`)) return;

    isExecutingAll = true;
    try {
      await App.ExecuteAllSteps(JSON.stringify(steps));
      showStatus("全ステップの自動実行を開始しました。コンソールで進捗を確認できます。");
    } catch (err) {
      console.error("一括実行エラー:", err);
      alert("一括実行に失敗しました: " + (err.message || err));
    } finally {
      isExecutingAll = false;
    }
  }

  // HTMLマニュアルパッケージのエクスポート
  async function handleExportHtml() {
    if (isExporting || steps.length === 0) return;
    isExporting = true;
    try {
      const outDir = await App.ExportManualPackage("", JSON.stringify(steps), false);
      if (confirm(`HTMLマニュアルパッケージを生成しました！\n出力先: ${outDir}\n\n今すぐブラウザで開きますか？`)) {
        await App.OpenFileInBrowser(`${outDir}\\index.html`);
      }
    } catch (err) {
      console.error("HTML出力エラー:", err);
      alert("HTMLパッケージ出力に失敗しました: " + (err.message || err));
    } finally {
      isExporting = false;
    }
  }

  // 統合UWSCRスクリプトのクリップボードコピー
  async function handleCopyCombinedScript() {
    if (steps.length === 0) return;
    try {
      const script = await App.ExportCombinedScript(JSON.stringify(steps));
      await navigator.clipboard.writeText(script);
      showStatus("全ステップを結合したUWSCRスクリプトをコピーしました！");
    } catch (err) {
      console.error("スクリプトコピーエラー:", err);
      alert("コピーに失敗しました: " + (err.message || err));
    }
  }

  // スクリプト開発（DEVELOP）タブへ引き継ぎ
  function handleExportToDev() {
    if (steps.length === 0) {
      alert("エクスポートする手順がありません。");
      return;
    }
    dispatch('exportToDev', steps);
  }

  // 画像クリックによるマーカー座標の直感的修正＆UWSCRコード連動
  function handleImageClick(e) {
    if (!activeStep) return;
    const rect = e.currentTarget.getBoundingClientRect();
    const clickXInElement = e.clientX - rect.left;
    const clickYInElement = e.clientY - rect.top;

    const imgElem = e.currentTarget;
    const naturalW = imgElem.naturalWidth || 1920;
    const naturalH = imgElem.naturalHeight || 1080;

    const actualX = Math.round((clickXInElement / rect.width) * naturalW);
    const actualY = Math.round((clickYInElement / rect.height) * naturalH);

    activeStep.click_x = actualX;
    activeStep.click_y = actualY;

    if (activeStep.rel_x && activeStep.rel_y) {
      activeStep.rel_x = actualX;
      activeStep.rel_y = actualY;
    }

    if (activeStep.uws_code) {
      const btnRegex = /(btn\s*\(\s*LEFT\s*,\s*CLICK\s*,\s*)\d+(\s*,\s*)\d+/i;
      if (btnRegex.test(activeStep.uws_code)) {
        activeStep.uws_code = activeStep.uws_code.replace(btnRegex, `$1${actualX}$2${actualY}`);
      }
    }

    steps = [...steps];
    showStatus(`マーカー座標を (${actualX}, ${actualY}) に更新しました`);
  }

  // ステップの追加
  function handleAddStep() {
    const newStepNum = steps.length + 1;
    const newStep = {
      step_id: newStepNum,
      step_number: newStepNum,
      title: `ステップ ${newStepNum}: 新規手順`,
      instruction: "操作内容を具体的に記述してください。",
      description: "操作内容を具体的に記述してください。",
      window_title: activeStep ? activeStep.window_title : "",
      target_element: "操作対象要素",
      action_type: "click",
      click_x: 100,
      click_y: 100,
      uws_code: `// ステップ ${newStepNum} 自動制御コード\n`,
      image_path: activeStep ? activeStep.image_path : ""
    };
    steps = [...steps, newStep];
    currentStep = steps.length - 1;
    showStatus("新規ステップを追加しました");
  }

  // ステップの削除
  function handleDeleteStep(idx) {
    if (steps.length <= 1) {
      alert("これ以上ステップを削除できません。");
      return;
    }
    if (!confirm(`ステップ ${idx + 1} を削除してもよろしいですか？`)) return;

    steps.splice(idx, 1);
    steps.forEach((s, i) => {
      s.step_id = i + 1;
      s.step_number = i + 1;
    });
    if (currentStep >= steps.length) {
      currentStep = steps.length - 1;
    }
    steps = [...steps];
    showStatus(`ステップ ${idx + 1} を削除しました`);
  }

  // ステップの上移動
  function handleMoveUp(idx) {
    if (idx <= 0) return;
    const temp = steps[idx];
    steps[idx] = steps[idx - 1];
    steps[idx - 1] = temp;
    steps.forEach((s, i) => {
      s.step_id = i + 1;
      s.step_number = i + 1;
    });
    currentStep = idx - 1;
    steps = [...steps];
  }

  // ステップの下移動
  function handleMoveDown(idx) {
    if (idx >= steps.length - 1) return;
    const temp = steps[idx];
    steps[idx] = steps[idx + 1];
    steps[idx + 1] = temp;
    steps.forEach((s, i) => {
      s.step_id = i + 1;
      s.step_number = i + 1;
    });
    currentStep = idx + 1;
    steps = [...steps];
  }

  function showStatus(msg) {
    statusMessage = msg;
    setTimeout(() => {
      if (statusMessage === msg) statusMessage = "";
    }, 3500);
  }

  // 現場コパイロット (AIチャット)
  async function askAI() {
    if (!userQuestion.trim()) return;
    const q = userQuestion;
    chatHistory = [...chatHistory, { role: "user", text: q }];
    userQuestion = "";

    try {
      const answer = await App.AskManualContext(q, activeStep?.image_path || "");
      chatHistory = [...chatHistory, { role: "assistant", text: answer }];
    } catch (err) {
      chatHistory = [...chatHistory, { role: "assistant", text: "回答の取得中にエラーが発生しました。" }];
    }
  }

  function clearChat() {
    chatHistory = [];
    userQuestion = "";
  }

  async function openRAGFolder() {
    try {
      await App.OpenKnowledgeDir();
    } catch (err) {
      alert("知識フォルダを開けませんでした: " + err);
    }
  }

  function playTTS(text) {
    if (!text) return;
    try {
      window.speechSynthesis.cancel();
      const cleanText = text.replace(/[\*\[\]`#\-_]/g, '');
      const utterance = new SpeechSynthesisUtterance(cleanText);
      utterance.lang = 'ja-JP';
      utterance.rate = 1.0;
      utterance.onstart = () => { isTTSPlaying = true; };
      utterance.onend = () => { isTTSPlaying = false; };
      utterance.onerror = () => { isTTSPlaying = false; };
      window.speechSynthesis.speak(utterance);
    } catch (e) {
      console.error("TTS playback failed:", e);
      isTTSPlaying = false;
    }
  }

  function stopTTS() {
    window.speechSynthesis.cancel();
    isTTSPlaying = false;
  }
</script>

<div class="manual-studio-container">
  <!-- 1. トップアクションバー -->
  <header class="top-action-bar card">
    <div class="top-left">
      <div class="header-badge">MANUAL STUDIO</div>
      <div class="log-dir-group">
        <input 
          type="text" 
          placeholder="録画ログフォルダパス (manual/recording_...)" 
          class="path-input"
          bind:value={selectedLogDir}
          disabled={isGenerating}
        />
        <button class="btn-secondary" on:click={browseLogDir} disabled={isGenerating}>
          📁 参照...
        </button>
        <button class="btn-primary" on:click={handleAutoGenerate} disabled={isGenerating || !selectedLogDir}>
          {#if isGenerating}
            <span class="spinner-mini"></span> AI解析中...
          {:else}
            ⚡ 記録から自動生成
          {/if}
        </button>
      </div>
    </div>

    <div class="top-right">
      {#if statusMessage}
        <span class="status-toast animate-fade">{statusMessage}</span>
      {/if}

      <div class="action-btn-group">
        <button class="btn-action" on:click={handleSaveScenario} disabled={steps.length === 0 || isSaving} title="編集内容を scenario.json に保存">
          💾 シナリオ保存
        </button>
        <button class="btn-action btn-execute-all" on:click={handleExecuteAll} disabled={steps.length === 0 || isExecutingAll} title="全ステップを連続自動実行">
          ▶ 全自動実行
        </button>
        <button class="btn-action" on:click={handleExportHtml} disabled={steps.length === 0 || isExporting} title="ブラウザ閲覧用HTMLマニュアルを出力">
          📦 HTML出力
        </button>
        <button class="btn-action" on:click={handleCopyCombinedScript} disabled={steps.length === 0} title="全ステップ結合スクリプトをコピー">
          📋 統合コード
        </button>
        <button class="btn-action btn-accent" on:click={handleExportToDev} disabled={steps.length === 0} title="DEVELOPタブで詳細スクリプト編集">
          シナリオを編集 →
        </button>
        <button 
          class="btn-action btn-toggle-copilot" 
          class:active={showCopilot}
          on:click={() => showCopilot = !showCopilot} 
          title={showCopilot ? "業務アシスタントを非表示にしてエディタを広く使う" : "業務アシスタントを表示"}
        >
          🤖 {showCopilot ? "アシスタント非表示" : "アシスタント表示"}
        </button>
      </div>
    </div>
  </header>

  <!-- 2. 3ペインメインワークスペース -->
  <div class="main-workspace-grid" class:collapsed-copilot={!showCopilot}>
    <!-- 左ペイン: ステップ一覧 (リスト & 並び替え & 追加・削除) -->
    <aside class="steps-sidebar card">
      <div class="sidebar-header">
        <div class="sidebar-title">
          <span>手順リスト</span>
          <span class="step-count">{steps.length} 件</span>
        </div>
        <button class="btn-add-step" on:click={handleAddStep} title="新しいステップを末尾に追加">
          ＋ 追加
        </button>
      </div>

      <div class="steps-scroll-list">
        {#if steps.length === 0}
          <div class="empty-steps">
            <p>手順がありません。<br>上の「記録から自動生成」を実行してください。</p>
          </div>
        {:else}
          {#each steps as s, idx}
            <div 
              class="step-item-card"
              class:active={currentStep === idx}
              on:click={() => currentStep = idx}
            >
              <div class="step-item-left">
                <span class="step-num-badge">{idx + 1}</span>
              </div>
              <div class="step-item-center">
                <div class="step-item-title" title={s.title || `ステップ ${idx + 1}`}>{s.title || `ステップ ${idx + 1}`}</div>
                <div class="step-item-sub">
                  {#if s.window_title}
                    <span class="sub-tag" title={s.window_title}>🪟 {s.window_title.substring(0, 16)}...</span>
                  {/if}
                  {#if s.target_element}
                    <span class="sub-tag elem-tag" title={s.target_element}>🎯 {s.target_element}</span>
                  {/if}
                </div>
              </div>
              <div class="step-item-right" on:click|stopPropagation>
                <div class="reorder-group">
                  <button class="btn-order" on:click={() => handleMoveUp(idx)} disabled={idx === 0} title="上へ移動">▲</button>
                  <button class="btn-order" on:click={() => handleMoveDown(idx)} disabled={idx === steps.length - 1} title="下へ移動">▼</button>
                </div>
                <button class="btn-delete" on:click={() => handleDeleteStep(idx)} title="このステップを削除">🗑️</button>
              </div>
            </div>
          {/each}
        {/if}
      </div>
    </aside>

    <!-- 中央ペイン: ステップエディタ ＆ インタラクティブ画像プレビュー -->
    <section class="step-editor-panel card">
      {#if activeStep}
        <div class="editor-header-bar">
          <div class="step-identity">
            <span class="step-tag">STEP {currentStep + 1} / {steps.length}</span>
            <input 
              type="text" 
              class="step-title-input" 
              bind:value={activeStep.title} 
              placeholder="ステップタイトル（例: 基幹システムへのログイン）"
            />
          </div>
          <div class="editor-header-actions">
            {#if isTTSPlaying}
              <button class="btn-tts stop" on:click={stopTTS}>🔊 音声停止</button>
            {:else}
              <button class="btn-tts" on:click={() => playTTS(activeStep.instruction || activeStep.description)}>
                ▶ 音声案内
              </button>
            {/if}
            <button class="btn-test-step" on:click={handleExecuteStepOnly} disabled={isExecuting}>
              {#if isExecuting}
                <span class="spinner-mini"></span> 実行中...
              {:else}
                ▶ このステップをテスト再生
              {/if}
            </button>
          </div>
        </div>

        <div class="editor-content-split">
          <!-- 上段: プレビュー画像 & インタラクティブマーカー -->
          <div class="preview-viewport-box">
            {#if activeStepBase64}
              <div class="image-wrapper">
                <img 
                  src={activeStepBase64} 
                  alt="ステッププレビュー" 
                  class="preview-image"
                  on:click={handleImageClick}
                  title="クリックした位置に赤丸マーカーを移動します"
                />
                {#if activeStep.click_x > 0 && activeStep.click_y > 0}
                  <div 
                    class="interactive-marker pulse-ring"
                    style="left: {(activeStep.click_x / 1920) * 100}%; top: {(activeStep.click_y / 1080) * 100}%;"
                    title={`クリック位置: (${activeStep.click_x}, ${activeStep.click_y})`}
                  ></div>
                {/if}
              </div>
              <div class="viewport-tip-bar">
                <span>💡 画像上の操作箇所を直接クリックすると、マーカー座標と自動制御コードが自動連動して修正されます</span>
                <span class="coord-label">座標: X={activeStep.click_x}, Y={activeStep.click_y}</span>
              </div>
            {:else if imageLoadError}
              <div class="empty-preview">
                <span class="empty-icon">📷</span>
                <p>キャプチャ画像がありません</p>
              </div>
            {:else}
              <div class="empty-preview">
                <span class="spinner"></span>
                <p>画像を読込中...</p>
              </div>
            {/if}
          </div>

          <!-- 下段: 手順指示・メタデータ・自動制御スクリプト編集 -->
          <div class="editor-form-box">
            <div class="form-row-grid">
              <div class="form-field">
                <label>対象ウィンドウ名</label>
                <input 
                  type="text" 
                  class="form-input" 
                  bind:value={activeStep.window_title} 
                  placeholder="例: 基幹業務システム"
                />
              </div>
              <div class="form-field">
                <label>操作対象要素</label>
                <input 
                  type="text" 
                  class="form-input" 
                  bind:value={activeStep.target_element} 
                  placeholder="例: ログインボタン, ユーザーID入力欄"
                />
              </div>
              <div class="form-field-coord">
                <label>操作座標 (X, Y)</label>
                <div class="coord-inputs">
                  <input type="number" class="coord-input" bind:value={activeStep.click_x} />
                  <input type="number" class="coord-input" bind:value={activeStep.click_y} />
                </div>
              </div>
            </div>

            <div class="form-field full-width">
              <label>作業手順・操作説明 (操作者が迷わない指示)</label>
              <textarea 
                class="form-textarea desc-textarea" 
                bind:value={activeStep.instruction}
                placeholder="誰が見てもわかる具体的な作業手順を記述してください。"
                rows="2"
              ></textarea>
            </div>

            <div class="form-field full-width">
              <div class="code-label-row">
                <label>自動制御スクリプト (UWSCR)</label>
                <span class="code-hint">※ステップ実行時および統合エクスポート時に自動実行されます</span>
              </div>
              <textarea 
                class="form-textarea code-textarea font-mono" 
                bind:value={activeStep.uws_code}
                placeholder="// UWSCRスクリプトを記述"
                rows="4"
              ></textarea>
            </div>
          </div>
        </div>
      {:else}
        <div class="empty-editor">
          <img src={appIcon} alt="App Icon" class="empty-app-icon" />
          <h3>ステップが選択されていません</h3>
          <p>左側のリストから編集したいステップを選択するか、記録ログから自動生成してください。</p>
        </div>
      {/if}
    </section>

    <!-- 右ペイン: 業務アシスタント (AIチャット & RAGナレッジ) -->
    {#if showCopilot}
      <aside class="copilot-sidebar card">
        <div class="copilot-header">
          <div class="copilot-title-group">
            <span class="copilot-dot"></span>
            <h3>業務アシスタント</h3>
          </div>
          <button on:click={openRAGFolder} class="btn-knowledge" title="マニュアルや業務仕様書を追加">
            📂 知識フォルダ
          </button>
        </div>

        <div class="rag-info-banner">
          <b>業務ナレッジ参照中:</b> 知識フォルダ内の仕様書やマニュアルに基づいてAIがアドバイスします。
        </div>

        <div class="chat-messages-container">
          {#if chatHistory.length === 0}
            <div class="chat-welcome">
              <p>この画面の操作手順や入力内容について、AIアシスタントにいつでも質問できます。</p>
              <div class="quick-questions">
                <button class="quick-btn" on:click={() => { userQuestion = "現在の操作手順の注意点を教えてください。"; askAI(); }}>
                  注意点を質問
                </button>
                <button class="quick-btn" on:click={() => { userQuestion = "この画面での入力値の仕様を教えてください。"; askAI(); }}>
                  入力仕様を質問
                </button>
              </div>
            </div>
          {/if}

          {#each chatHistory as chat}
            <div class="chat-bubble-row {chat.role === 'user' ? 'user-row' : 'assistant-row'}">
              <span class="chat-role">{chat.role === 'user' ? 'あなた' : 'AI'}</span>
              <div class="chat-bubble {chat.role === 'user' ? 'user-bubble' : 'assistant-bubble'}">
                {chat.text}
              </div>
            </div>
          {/each}
        </div>

        <div class="chat-input-bar">
          <input 
            type="text" 
            placeholder="業務手順や仕様について質問..." 
            class="chat-input"
            bind:value={userQuestion}
            on:keydown={(e) => e.key === 'Enter' && askAI()}
          />
          <button class="btn-send" on:click={askAI} disabled={!userQuestion.trim()}>
            送信
          </button>
        </div>
      </aside>
    {/if}
  </div>
</div>

<style>
  .manual-studio-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    flex: 1;
    min-height: 0;
    gap: 10px;
    box-sizing: border-box;
    overflow: hidden;
    padding: 0;
  }

  .card {
    background: var(--bg-secondary);
    backdrop-filter: var(--glass-blur);
    border: 1px solid var(--border-color);
    border-radius: 10px;
    box-shadow: var(--shadow-sm);
  }

  /* 1. トップアクションバー */
  .top-action-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 12px;
    gap: 10px;
    flex-shrink: 0;
    flex-wrap: wrap;
  }

  .top-left {
    display: flex;
    align-items: center;
    gap: 8px;
    flex: 1 1 320px;
    min-width: 240px;
  }

  .header-badge {
    background: var(--accent-color);
    color: var(--bg-primary);
    font-size: 0.65rem;
    font-weight: 700;
    padding: 4px 8px;
    border-radius: 4px;
    letter-spacing: 0.05em;
    white-space: nowrap;
  }

  .log-dir-group {
    display: flex;
    gap: 6px;
    align-items: center;
    flex: 1;
    max-width: 440px;
  }

  .path-input {
    flex: 1;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 5px 8px;
    color: var(--text-primary);
    font-size: 0.78rem;
    outline: none;
    transition: border-color 0.2s;
  }

  .path-input:focus {
    border-color: var(--accent-color);
  }

  .top-right {
    display: flex;
    align-items: center;
    gap: 8px;
    flex-wrap: wrap;
    justify-content: flex-end;
  }

  .status-toast {
    font-size: 0.72rem;
    color: #10b981;
    background: rgba(16, 185, 129, 0.1);
    border: 1px solid rgba(16, 185, 129, 0.3);
    padding: 3px 8px;
    border-radius: 4px;
    font-weight: 500;
  }

  .action-btn-group {
    display: flex;
    gap: 5px;
    flex-wrap: wrap;
    align-items: center;
  }

  .btn-primary, .btn-secondary, .btn-action {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 5px;
    padding: 5px 10px;
    font-size: 0.72rem;
    font-weight: 500;
    border-radius: 6px;
    cursor: pointer;
    transition: all 0.2s ease;
    white-space: nowrap;
  }

  .btn-primary {
    background: var(--accent-color);
    color: var(--bg-primary);
    border: 1px solid var(--accent-color);
  }

  .btn-primary:hover:not(:disabled) {
    opacity: 0.9;
  }

  .btn-secondary, .btn-action {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
  }

  .btn-secondary:hover:not(:disabled), .btn-action:hover:not(:disabled) {
    background: var(--accent-soft);
    border-color: var(--accent-color);
  }

  .btn-execute-all {
    background: rgba(16, 185, 129, 0.1);
    border-color: rgba(16, 185, 129, 0.4);
    color: #10b981;
    font-weight: 600;
  }

  .btn-execute-all:hover:not(:disabled) {
    background: #10b981;
    color: #fff;
  }

  .btn-accent {
    background: var(--accent-soft);
    border-color: var(--accent-color);
    color: var(--accent-color);
    font-weight: 600;
  }

  .btn-toggle-copilot.active {
    border-color: var(--accent-color);
    color: var(--accent-color);
    background: var(--accent-soft);
  }

  button:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }

  /* 2. メイン3ペイングリッド */
  .main-workspace-grid {
    display: grid;
    grid-template-columns: 240px minmax(0, 1fr) 280px;
    gap: 10px;
    flex: 1;
    min-height: 0;
    overflow: hidden;
    width: 100%;
  }

  .main-workspace-grid.collapsed-copilot {
    grid-template-columns: 240px minmax(0, 1fr);
  }

  /* 左ペイン: ステップ一覧 */
  .steps-sidebar {
    display: flex;
    flex-direction: column;
    min-height: 0;
    overflow: hidden;
  }

  .sidebar-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 12px;
    border-bottom: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.05);
  }

  .sidebar-title {
    font-size: 0.8rem;
    font-weight: 600;
    color: var(--text-primary);
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .step-count {
    font-size: 0.65rem;
    background: var(--input-bg);
    padding: 1px 6px;
    border-radius: 10px;
    border: 1px solid var(--border-color);
  }

  .btn-add-step {
    background: transparent;
    border: 1px dashed var(--accent-color);
    color: var(--accent-color);
    font-size: 0.7rem;
    padding: 2px 8px;
    border-radius: 4px;
    cursor: pointer;
  }

  .btn-add-step:hover {
    background: var(--accent-soft);
  }

  .steps-scroll-list {
    flex: 1;
    overflow-y: auto;
    padding: 8px;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .empty-steps {
    padding: 24px 12px;
    text-align: center;
    color: var(--text-secondary);
    font-size: 0.75rem;
    line-height: 1.5;
  }

  .step-item-card {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .step-item-card:hover {
    border-color: var(--accent-color);
  }

  .step-item-card.active {
    border-color: var(--accent-color);
    background: var(--accent-soft);
    box-shadow: 0 0 0 1px var(--accent-color);
  }

  .step-num-badge {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 22px;
    height: 22px;
    background: var(--border-color);
    color: var(--text-primary);
    border-radius: 50%;
    font-size: 0.7rem;
    font-weight: 700;
  }

  .step-item-card.active .step-num-badge {
    background: var(--accent-color);
    color: var(--bg-primary);
  }

  .step-item-center {
    flex: 1;
    min-width: 0;
  }

  .step-item-title {
    font-size: 0.75rem;
    font-weight: 600;
    color: var(--text-primary);
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    line-height: 1.35;
    word-break: break-all;
  }

  .step-item-sub {
    display: flex;
    gap: 4px;
    margin-top: 2px;
    flex-wrap: wrap;
  }

  .sub-tag {
    font-size: 0.6rem;
    color: var(--text-secondary);
    background: rgba(0, 0, 0, 0.05);
    padding: 1px 4px;
    border-radius: 3px;
  }

  .elem-tag {
    color: #2563eb;
    background: rgba(37, 99, 235, 0.1);
  }

  .step-item-right {
    display: flex;
    align-items: center;
    gap: 2px;
  }

  .reorder-group {
    display: flex;
    flex-direction: column;
    gap: 1px;
  }

  .btn-order {
    background: transparent;
    border: none;
    font-size: 0.55rem;
    padding: 1px 3px;
    color: var(--text-secondary);
    cursor: pointer;
  }

  .btn-order:hover:not(:disabled) {
    color: var(--accent-color);
  }

  .btn-delete {
    background: transparent;
    border: none;
    font-size: 0.75rem;
    cursor: pointer;
    padding: 2px 4px;
    opacity: 0.6;
    transition: opacity 0.2s;
  }

  .btn-delete:hover {
    opacity: 1;
  }

  /* 中央ペイン: ステップエディタ */
  .step-editor-panel {
    display: flex;
    flex-direction: column;
    min-height: 0;
    min-width: 0;
    overflow: hidden;
    height: 100%;
  }

  .editor-header-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 6px 12px;
    border-bottom: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.04);
    gap: 8px;
    flex-shrink: 0;
    flex-wrap: wrap;
  }

  .step-identity {
    display: flex;
    align-items: center;
    gap: 8px;
    flex: 1 1 200px;
    min-width: 0;
  }

  .step-tag {
    font-size: 0.7rem;
    font-weight: 700;
    color: var(--accent-color);
    background: var(--accent-soft);
    padding: 2px 6px;
    border-radius: 4px;
    border: 1px solid var(--border-color);
    white-space: nowrap;
    flex-shrink: 0;
  }

  .step-title-input {
    flex: 1;
    min-width: 0;
    font-size: 0.85rem;
    font-weight: 600;
    color: var(--text-primary);
    background: transparent;
    border: 1px solid transparent;
    border-radius: 4px;
    padding: 4px 6px;
    outline: none;
  }

  .step-title-input:hover, .step-title-input:focus {
    background: var(--input-bg);
    border-color: var(--border-color);
  }

  .editor-header-actions {
    display: flex;
    align-items: center;
    gap: 6px;
    flex-shrink: 0;
  }

  .btn-tts {
    font-size: 0.7rem;
    padding: 4px 8px;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-radius: 4px;
    cursor: pointer;
  }

  .btn-tts.stop {
    color: #ef4444;
    border-color: #ef4444;
  }

  .btn-test-step {
    font-size: 0.72rem;
    font-weight: 600;
    padding: 5px 10px;
    background: var(--accent-color);
    color: var(--bg-primary);
    border: none;
    border-radius: 4px;
    cursor: pointer;
    white-space: nowrap;
  }

  .editor-content-split {
    display: flex;
    flex-direction: column;
    flex: 1;
    min-height: 0;
    overflow-y: auto;
    padding: 10px;
    gap: 10px;
  }

  /* 画像プレビュー & マーカー */
  .preview-viewport-box {
    position: relative;
    background: #000;
    border-radius: 8px;
    overflow: hidden;
    border: 1px solid var(--border-color);
    min-height: 200px;
    flex: 1 1 360px;
    max-height: 52vh;
    display: flex;
    flex-direction: column;
  }

  .image-wrapper {
    position: relative;
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    cursor: crosshair;
    background: #0f172a;
  }

  .preview-image {
    width: 100%;
    height: 100%;
    object-fit: contain;
    user-select: none;
  }

  .interactive-marker {
    position: absolute;
    width: 36px;
    height: 36px;
    border: 3px solid #ef4444;
    border-radius: 50%;
    transform: translate(-50%, -50%);
    pointer-events: none;
    box-shadow: 0 0 12px rgba(239, 68, 68, 0.8);
  }

  .pulse-ring::before {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    width: 6px;
    height: 6px;
    background: #ef4444;
    border-radius: 50%;
    transform: translate(-50%, -50%);
  }

  .viewport-tip-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: rgba(0, 0, 0, 0.85);
    color: #e2e8f0;
    font-size: 0.68rem;
    padding: 4px 10px;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    flex-wrap: wrap;
    gap: 6px;
    flex-shrink: 0;
  }

  .coord-label {
    font-family: monospace;
    color: #38bdf8;
    white-space: nowrap;
  }

  .empty-preview, .empty-editor {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    color: var(--text-secondary);
    gap: 8px;
    text-align: center;
  }

  .empty-app-icon {
    width: 48px;
    height: 48px;
    opacity: 0.6;
  }

  /* 下段フォーム */
  .editor-form-box {
    display: flex;
    flex-direction: column;
    gap: 8px;
    flex-shrink: 0;
  }

  .form-row-grid {
    display: grid;
    grid-template-columns: 1fr 1fr 130px;
    gap: 8px;
  }

  .form-field {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }

  .form-field label, .code-label-row label {
    font-size: 0.7rem;
    font-weight: 600;
    color: var(--text-secondary);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .form-input {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 5px;
    padding: 5px 8px;
    font-size: 0.75rem;
    color: var(--text-primary);
    outline: none;
    width: 100%;
    box-sizing: border-box;
  }

  .form-input:focus, .form-textarea:focus {
    border-color: var(--accent-color);
  }

  .form-field-coord {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }

  .form-field-coord label {
    font-size: 0.7rem;
    font-weight: 600;
    color: var(--text-secondary);
    white-space: nowrap;
  }

  .coord-inputs {
    display: flex;
    gap: 4px;
  }

  .coord-input {
    width: 100%;
    max-width: 60px;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 5px;
    padding: 5px 4px;
    font-size: 0.75rem;
    color: var(--text-primary);
    text-align: center;
    outline: none;
    box-sizing: border-box;
  }

  .form-textarea {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 5px;
    padding: 6px 8px;
    font-size: 0.75rem;
    color: var(--text-primary);
    outline: none;
    resize: vertical;
    line-height: 1.4;
    width: 100%;
    box-sizing: border-box;
  }

  .desc-textarea {
    min-height: 44px;
  }

  .code-label-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    gap: 4px;
  }

  .code-hint {
    font-size: 0.65rem;
    color: var(--text-secondary);
  }

  .code-textarea {
    font-family: 'Consolas', monospace;
    background: #1e1e1e;
    color: #9cdcfe;
    border-color: #333;
    min-height: 70px;
  }

  /* 右ペイン: 業務アシスタント */
  .copilot-sidebar {
    display: flex;
    flex-direction: column;
    min-height: 0;
    overflow: hidden;
  }

  .copilot-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 12px;
    border-bottom: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.05);
  }

  .copilot-title-group {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .copilot-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: #10b981;
    box-shadow: 0 0 6px #10b981;
  }

  .copilot-header h3 {
    margin: 0;
    font-size: 0.8rem;
    font-weight: 600;
  }

  .btn-knowledge {
    background: transparent;
    border: 1px solid var(--border-color);
    color: var(--text-secondary);
    font-size: 0.65rem;
    padding: 2px 6px;
    border-radius: 4px;
    cursor: pointer;
  }

  .rag-info-banner {
    padding: 6px 10px;
    background: rgba(0, 0, 0, 0.02);
    border-bottom: 1px solid var(--border-color);
    font-size: 0.65rem;
    color: var(--text-secondary);
    line-height: 1.3;
  }

  .chat-messages-container {
    flex: 1;
    overflow-y: auto;
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .chat-welcome {
    margin: auto;
    text-align: center;
    font-size: 0.7rem;
    color: var(--text-secondary);
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .quick-questions {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .quick-btn {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    padding: 4px 8px;
    border-radius: 6px;
    font-size: 0.68rem;
    cursor: pointer;
  }

  .quick-btn:hover {
    background: var(--accent-soft);
    border-color: var(--accent-color);
  }

  .chat-bubble-row {
    display: flex;
    flex-direction: column;
    gap: 2px;
    max-width: 90%;
  }

  .user-row {
    align-self: flex-end;
    align-items: flex-end;
  }

  .assistant-row {
    align-self: flex-start;
    align-items: flex-start;
  }

  .chat-role {
    font-size: 0.6rem;
    color: var(--text-secondary);
  }

  .chat-bubble {
    padding: 6px 10px;
    border-radius: 8px;
    font-size: 0.72rem;
    line-height: 1.4;
  }

  .user-bubble {
    background: var(--accent-color);
    color: var(--bg-primary);
  }

  .assistant-bubble {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
  }

  .chat-input-bar {
    display: flex;
    padding: 8px;
    gap: 6px;
    border-top: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.03);
  }

  .chat-input {
    flex: 1;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 5px;
    padding: 5px 8px;
    font-size: 0.72rem;
    color: var(--text-primary);
    outline: none;
  }

  .btn-send {
    background: var(--accent-color);
    color: var(--bg-primary);
    border: none;
    border-radius: 5px;
    padding: 0 10px;
    font-size: 0.72rem;
    font-weight: 600;
    cursor: pointer;
  }

  .spinner-mini {
    display: inline-block;
    width: 12px;
    height: 12px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-radius: 50%;
    border-top-color: #fff;
    animation: spin 0.8s linear infinite;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }
</style>
