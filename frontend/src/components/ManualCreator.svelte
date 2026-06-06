<script>
  import { onMount } from 'svelte';
  import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';

  const App = window.go.main.App;

  let currentStep = 0;
  let steps = []; // Goの slicer からフェッチされるステップ
  let selectedLogDir = ""; // 操作ログディレクトリのパス
  let isGenerating = false;
  let isExecuting = false;
  let isTTSPlaying = false;
  let userQuestion = "";
  let chatHistory = [];

  let activeStepBase64 = "";

  // アクティブステップの切り替え監視
  $: activeStep = steps[currentStep] || null;

  $: {
    if (activeStep && activeStep.image_path) {
      App.GetImageBase64(activeStep.image_path)
        .then(res => {
          activeStepBase64 = res;
        })
        .catch(err => {
          console.error("Failed to load base64 image:", err);
          activeStepBase64 = "";
        });
    } else {
      activeStepBase64 = "";
    }
  }

  // オートズーム用座標スタイル
  $: zoomStyle = activeStep && activeStep.click_x > 0
    ? `transform: scale(1.6) translate(${50 - (activeStep.click_x / 19.2)}%, ${50 - (activeStep.click_y / 10.8)}%);`
    : 'transform: scale(1) translate(0, 0);';

  // AI自動スライス呼び出し
  async function handleAutoGenerate() {
    if (!selectedLogDir) {
      alert("録画ログの保存されているフォルダを入力または選択してください。");
      return;
    }

    isGenerating = true;
    chatHistory = [];
    try {
      // Wailsバインド経由で自動スライス実行
      const result = await App.GenerateScenarioFromLog(selectedLogDir);
      if (result && result.length > 0) {
        steps = result;
        currentStep = 0;
        playTTS("録画ログからシナリオを自動生成しました。マニュアルに沿ってステップ実行を開始してください。");
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

  // ステップ実行
  async function handleExecuteStep() {
    if (isExecuting) return;
    isExecuting = true;

    try {
      const result = await App.ExecuteStep(currentStep);
      if (result) {
        playTTS(result.instruction);
        if (currentStep < steps.length - 1) {
          currentStep++;
        }
      }
    } catch (err) {
      console.error("ステップ実行エラー: ", err);
      alert("実行中にエラーが発生しました: " + (err.message || err));
    } finally {
      isExecuting = false;
    }
  }

  // 現場コパイロットへの質問
  async function askAI() {
    if (!userQuestion.trim()) return;
    const q = userQuestion;
    chatHistory = [...chatHistory, { role: "user", text: q }];
    userQuestion = "";

    try {
      const answer = await App.AskManualContext(q, activeStep?.image_path || "");
      chatHistory = [...chatHistory, { role: "assistant", text: answer }];
      playTTS(answer);
    } catch (err) {
      chatHistory = [...chatHistory, { role: "assistant", text: "回答の取得中にエラーが発生しました。" }];
    }
  }

  // Web Speech API による音声読み上げ (TTS)
  function playTTS(text) {
    if (!text) return;
    try {
      window.speechSynthesis.cancel();
      // 余計な記号などをマイルドに間引き
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

<div class="manual-creator-layout">
  <!-- 1. AIオート・スライサー制御パネル -->
  <div class="control-panel card">
    <div class="panel-header">
      <div class="header-main">
        <svg class="header-icon animate-pulse" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"/>
          <polygon points="10 8 16 12 10 16 10 8"/>
        </svg>
        <div class="header-text">
          <h2>AI シナリオ自動生成（オート・スライサー）</h2>
          <p class="description">レコーダーが保存した生の操作ログフォルダを指定すると、AIが論理的ステップに自動分割し、自動再生用コードを含む対話型ガイドを生成します。</p>
        </div>
      </div>
    </div>

    <div class="settings-row">
      <div class="form-group flex-fill">
        <div class="input-with-btn">
          <input 
            type="text" 
            placeholder="録画ログフォルダパス (C:/data/log...)" 
            class="path-input"
            bind:value={selectedLogDir}
            disabled={isGenerating}
          />
          <button class="btn-browse" on:click={browseLogDir} disabled={isGenerating}>
            参照...
          </button>
        </div>
      </div>
      
      <button 
        class="btn-generate gradient-btn"
        on:click={handleAutoGenerate}
        disabled={isGenerating}
      >
        {#if isGenerating}
          <span class="spinner-inline">🌀</span> 解析記述中...
        {:else}
          録画から自動生成
        {/if}
      </button>
    </div>
  </div>

  <!-- 2. メイン並走ビューワー & AIチャット -->
  <div class="main-viewer-grid">
    <!-- 左：オートズームプレビュー -->
    <div class="preview-panel card">
      {#if activeStep}
        <div class="viewport bg-slate-950">
          {#if activeStepBase64}
            <img 
              src={activeStepBase64} 
              alt="操作箇所のプレビュー" 
              class="viewport-img transition-transform duration-700 ease-out"
              style={zoomStyle}
            />
          {:else}
            <div class="empty-viewport">
              <span class="spinner"></span>
              <p>画像を読込中...</p>
            </div>
          {/if}
          
          {#if activeStep.click_x > 0}
            <div 
              class="pulse-ring"
              class:pulsing={isTTSPlaying}
              style="left: {(activeStep.click_x / 1920) * 100}%; top: {(activeStep.click_y / 1080) * 100}%;"
            ></div>
          {/if}
        </div>

        <div class="step-info-bar">
          <div class="info-content">
            <div class="badge-row">
              <span class="step-badge">ステップ {currentStep + 1}</span>
              {#if isTTSPlaying}
                <button class="tts-stop-btn" on:click={stopTTS}>🔊 音声を停止</button>
              {/if}
            </div>
            <h3 class="step-title">{activeStep.title}</h3>
            <p class="step-desc">{activeStep.instruction}</p>
          </div>
          <div class="code-preview">
            <label class="code-label">自動生成された UWSCR</label>
            <textarea 
              class="code-textarea"
              bind:value={activeStep.uws_code}
              placeholder="// UWSCRコードがここに表示されます"
            ></textarea>
          </div>
        </div>
      {:else}
        <div class="empty-state">
          <span class="empty-icon">📼</span>
          <h3>録画データがロードされていません</h3>
          <p>録画ログディレクトリを指定して、上の「録画から自動生成」を実行してください。</p>
        </div>
      {/if}

      <!-- 下部同期コントローラー -->
      <div class="sync-controller border-t">
        <button 
          class="btn-nav" 
          disabled={currentStep === 0 || steps.length === 0}
          on:click={() => currentStep--}
        >
          前へ
        </button>
        
        <span class="step-progress-text">
          ステップ進捗: {steps.length > 0 ? currentStep + 1 : 0} / {steps.length}
        </span>

        <button 
          class="btn-execute"
          disabled={steps.length === 0 || isExecuting}
          on:click={handleExecuteStep}
        >
          {#if isExecuting}
            実行中...
          {:else}
            このステップを実行して次へ
          {/if}
        </button>
      </div>
    </div>

    <!-- 右：現場コパイロット（RAG・自由質疑応答） -->
    <div class="copilot-panel card">
      <div class="copilot-header">
        <span class="copilot-dot"></span>
        <h3>現場業務コパイロット (RAG連携)</h3>
      </div>

      <div class="chat-container">
        {#if chatHistory.length === 0}
          <div class="chat-empty">
            <p>現在のステップや業務知識について、AIに質問することができます。</p>
            <div class="suggested-chips">
              <button class="chip" on:click={() => { userQuestion = "現在の操作手順について教えてください。"; askAI(); }}>手順を質問</button>
              <button class="chip" on:click={() => { userQuestion = "この画面の入力値の根拠は何ですか？"; askAI(); }}>入力の根拠を質問</button>
            </div>
          </div>
        {/if}
        {#each chatHistory as chat}
          <div class="chat-bubble-wrapper {chat.role === 'user' ? 'user-align' : 'assistant-align'}">
            <span class="chat-sender">{chat.role === 'user' ? '現場担当者' : 'AIコパイロット'}</span>
            <div class="chat-bubble {chat.role === 'user' ? 'user-bubble' : 'assistant-bubble'}">
              {chat.text}
            </div>
          </div>
        {/each}
      </div>

      <div class="chat-input-row border-t">
        <input 
          type="text" 
          placeholder="例: 框の厚みの入力仕様は？" 
          class="chat-input"
          bind:value={userQuestion}
          on:keydown={(e) => e.key === 'Enter' && askAI()}
        />
        <button 
          class="btn-chat-send" 
          on:click={askAI}
          disabled={!userQuestion.trim()}
        >
          質問
        </button>
      </div>
    </div>
  </div>
</div>

<style>
  .manual-creator-layout {
    display: flex;
    flex-direction: column;
    height: 100%;
    gap: 16px;
    box-sizing: border-box;
    overflow: hidden;
  }

  .card {
    background: var(--bg-secondary);
    backdrop-filter: var(--glass-blur);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    box-shadow: var(--shadow-md);
  }

  .control-panel {
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .header-main {
    display: flex;
    gap: 12px;
    align-items: flex-start;
  }

  .header-icon {
    width: 24px;
    height: 24px;
    color: var(--accent-color);
    margin-top: 3px;
  }

  .header-text h2 {
    margin: 0;
    font-size: 1rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .description {
    font-size: 0.75rem;
    color: var(--text-secondary);
    margin: 4px 0 0 0;
    line-height: 1.4;
  }

  .settings-row {
    display: flex;
    gap: 12px;
    align-items: center;
  }

  .flex-fill {
    flex: 1;
  }

  .input-with-btn {
    display: flex;
    gap: 8px;
    width: 100%;
  }

  .path-input {
    flex: 1;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 8px 12px;
    color: var(--text-primary);
    font-size: 0.8rem;
    outline: none;
    transition: border-color 0.2s ease;
  }

  .path-input:focus {
    border-color: var(--accent-color);
  }

  .btn-browse {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-radius: 6px;
    padding: 8px 16px;
    font-size: 0.8rem;
    font-weight: 500;
    cursor: pointer;
    white-space: nowrap;
    transition: all 0.2s ease;
  }

  .btn-browse:hover {
    border-color: var(--border-hover);
    background: var(--accent-soft);
  }

  .gradient-btn {
    background: var(--accent-color);
    border: 1px solid var(--border-color);
    color: var(--bg-primary);
    font-weight: 600;
    border-radius: 6px;
    padding: 9px 20px;
    font-size: 0.8rem;
    cursor: pointer;
    transition: background-color 0.2s ease, transform 0.1s ease;
    white-space: nowrap;
  }

  .gradient-btn:hover:not(:disabled) {
    background: var(--accent-hover);
  }

  .gradient-btn:active:not(:disabled) {
    transform: scale(0.97);
  }

  .gradient-btn:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }

  /* Main Viewer Grid */
  .main-viewer-grid {
    display: grid;
    grid-template-columns: 1fr 320px;
    gap: 20px;
    flex: 1;
    min-height: 0;
  }

  /* Left Panel */
  .preview-panel {
    display: flex;
    flex-direction: column;
    min-height: 0;
    overflow: hidden;
  }

  .viewport {
    flex: 1;
    min-height: 0;
    position: relative;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .viewport-img {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
  }

  .empty-viewport {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 12px;
    color: var(--text-secondary);
  }

  .spinner {
    width: 24px;
    height: 24px;
    border: 2px solid var(--border-color);
    border-top-color: var(--accent-color);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
  }

  .spinner-inline {
    display: inline-block;
    animation: spin 1s linear infinite;
  }

  .pulse-ring {
    position: absolute;
    width: 60px;
    height: 60px;
    margin-left: -30px;
    margin-top: -30px;
    border: 3px solid rgba(255, 71, 87, 0.8);
    border-radius: 50%;
    pointer-events: none;
    box-shadow: 0 0 10px rgba(255, 71, 87, 0.4);
    opacity: 0;
  }

  .pulse-ring.pulsing {
    animation: wave 1.5s infinite ease-out;
  }

  @keyframes wave {
    0% { transform: scale(0.4); opacity: 1; }
    100% { transform: scale(1.4); opacity: 0; }
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }

  .step-info-bar {
    padding: 16px;
    background: rgba(0, 0, 0, 0.2);
    border-top: 1px solid var(--border-color);
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 16px;
  }

  .info-content {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .badge-row {
    display: flex;
    gap: 10px;
    align-items: center;
  }

  .step-badge {
    background: var(--accent-soft);
    border: 1px solid var(--accent-border);
    color: var(--accent-color);
    font-size: 0.7rem;
    font-weight: 700;
    padding: 3px 8px;
    border-radius: 20px;
  }

  .tts-stop-btn {
    background: rgba(255, 71, 87, 0.1);
    border: 1px solid rgba(255, 71, 87, 0.2);
    color: #ff4757;
    font-size: 0.65rem;
    padding: 2px 8px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .tts-stop-btn:hover {
    background: rgba(255, 71, 87, 0.2);
  }

  .step-title {
    margin: 0;
    font-size: 0.95rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .step-desc {
    margin: 0;
    font-size: 0.8rem;
    color: var(--text-secondary);
    line-height: 1.4;
  }

  .code-preview {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .code-label {
    font-size: 0.65rem;
    color: var(--text-secondary);
    font-weight: 600;
    text-transform: uppercase;
  }

  .code-textarea {
    flex: 1;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 8px;
    font-family: Consolas, monospace;
    font-size: 0.75rem;
    color: var(--text-primary);
    resize: none;
    outline: none;
    min-height: 70px;
  }

  .empty-state {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    color: var(--text-secondary);
    padding: 40px;
    text-align: center;
  }

  .empty-icon {
    font-size: 2.5rem;
    margin-bottom: 12px;
  }

  .empty-state h3 {
    margin: 0 0 6px 0;
    font-size: 1rem;
    color: var(--text-primary);
  }

  .empty-state p {
    margin: 0;
    font-size: 0.75rem;
    max-width: 300px;
    line-height: 1.4;
  }

  .sync-controller {
    padding: 12px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: rgba(0, 0, 0, 0.1);
  }

  .border-t {
    border-top: 1px solid var(--border-color);
  }

  .btn-nav {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-radius: 6px;
    padding: 6px 16px;
    font-size: 0.75rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .btn-nav:hover:not(:disabled) {
    border-color: var(--border-hover);
    background: var(--accent-soft);
  }

  .btn-nav:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }

  .step-progress-text {
    font-size: 0.75rem;
    color: var(--text-secondary);
    font-weight: 500;
  }

  .btn-execute {
    background: var(--accent-color);
    border: 1px solid var(--border-color);
    color: var(--bg-primary);
    font-weight: 600;
    border-radius: 20px;
    padding: 7px 20px;
    font-size: 0.75rem;
    cursor: pointer;
    transition: background-color 0.2s ease, transform 0.1s ease;
  }

  .btn-execute:hover:not(:disabled) {
    background: var(--accent-hover);
  }

  .btn-execute:active:not(:disabled) {
    transform: scale(0.97);
  }

  .btn-execute:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }

  /* Right Copilot Panel */
  .copilot-panel {
    display: flex;
    flex-direction: column;
    min-height: 0;
    overflow: hidden;
  }

  .copilot-header {
    padding: 14px 16px;
    display: flex;
    align-items: center;
    gap: 8px;
    background: rgba(0, 0, 0, 0.15);
    border-bottom: 1px solid var(--border-color);
  }

  .copilot-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background-color: var(--accent-color);
    box-shadow: 0 0 8px var(--accent-color);
    animation: flash 1.5s infinite alternate;
  }

  @keyframes flash {
    from { opacity: 0.4; }
    to { opacity: 1; }
  }

  .copilot-header h3 {
    margin: 0;
    font-size: 0.8rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .chat-container {
    flex: 1;
    overflow-y: auto;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .chat-empty {
    margin: auto;
    text-align: center;
    color: var(--text-secondary);
    font-size: 0.75rem;
    max-width: 240px;
    line-height: 1.4;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .suggested-chips {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .chip {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    padding: 6px 12px;
    border-radius: 12px;
    font-size: 0.7rem;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .chip:hover {
    background: var(--accent-soft);
    border-color: var(--accent-color);
  }

  .chat-bubble-wrapper {
    display: flex;
    flex-direction: column;
    gap: 4px;
    max-width: 85%;
  }

  .user-align {
    align-self: flex-end;
    align-items: flex-end;
  }

  .assistant-align {
    align-self: flex-start;
    align-items: flex-start;
  }

  .chat-sender {
    font-size: 0.6rem;
    color: var(--text-secondary);
  }

  .chat-bubble {
    padding: 10px 14px;
    border-radius: 12px;
    font-size: 0.75rem;
    line-height: 1.4;
  }

  .user-bubble {
    background: var(--accent-color);
    color: var(--bg-primary);
    border-bottom-right-radius: 2px;
  }

  .assistant-bubble {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    border-bottom-left-radius: 2px;
  }

  .chat-input-row {
    padding: 12px 16px;
    display: flex;
    gap: 8px;
    background: rgba(0, 0, 0, 0.1);
  }

  .chat-input {
    flex: 1;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 6px 12px;
    color: var(--text-primary);
    font-size: 0.75rem;
    outline: none;
    transition: border-color 0.2s ease;
  }

  .chat-input:focus {
    border-color: var(--accent-color);
  }

  .btn-chat-send {
    background: var(--accent-color);
    border: 1px solid var(--border-color);
    color: var(--bg-primary);
    font-weight: 600;
    border-radius: 6px;
    padding: 6px 14px;
    font-size: 0.75rem;
    cursor: pointer;
    transition: background-color 0.2s ease, opacity 0.2s ease;
  }

  .btn-chat-send:hover:not(:disabled) {
    background: var(--accent-hover);
  }

  .btn-chat-send:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }
</style>
