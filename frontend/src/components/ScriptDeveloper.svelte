<script>
  import { onMount, onDestroy } from 'svelte';
  import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';
  import CodeEditor from './CodeEditor.svelte';

  const App = window.go.main.App;

  export let importedSteps = null;
  export let activeTab; // Bind activeTab to switch tabs dynamically

  // モード切り替え: 'batch' (ワンショット開発) or 'step' (マルチステップ開発)
  let mode = 'batch';

  $: if (importedSteps && importedSteps.length > 0) {
    steps = importedSteps.map((s, idx) => {
      // 1. まずベースとなるステップオブジェクトを構築
      const stepObj = {
        id: s.step_number || (idx + 1),
        title: s.title || `ステップ ${idx + 1}`,
        prompt: s.instruction || '',
        code: s.uws_code || '',
        logs: '',
        status: 'idle',
        sessionContext: null, // 初期値は一旦null
        highlightedLine: null
      };

      // 2. image_path が存在する場合、バックエンドからBase64データを取得してセッションコンテキストを復元
      // Goの s.image_path または s.ImagePath の双方を考慮したフォールバック
      const imgPath = s.image_path || s.ImagePath;
      if (imgPath) {
        App.GetImageBase64(imgPath).then(res => {
          if (res) {
            // "data:image/png;base64,..." のヘッダーから純粋なBase64文字列のみを抽出
            const base64Raw = res.includes(',') ? res.split(',')[1] : res;
            
            stepObj.sessionContext = {
              active_title: stepObj.title,
              screenshot_base64: base64Raw
            };
            steps = [...steps]; // Svelteの画面再描画をトリガー
          }
        }).catch(err => {
          console.error("[ScriptDeveloper] インポート画像の復元に失敗しました:", err);
        });
      }

      return stepObj;
    });

    activeStepIndex = 0;
    mode = 'step'; // ステップ開発モードに切り替え
    importedSteps = null; // 処理したらクリア
    showStatus('マニュアル作成から手順をインポートしました！各ステップのキャプチャ画像とコードを直接検証できます。');
  }

  // ワンショット開発用状態
  let prompt = '';
  let generatedCode = '';
  let savePath = '';
  let isLoading = false;
  let isCapturing = false;
  let statusMessage = '';
  let isError = false;
  
  // ワンショット検証用の状態
  let isBatchTestRunning = false;
  let batchLogs = '';
  let batchStatus = 'idle'; // 'idle', 'generating', 'running', 'success', 'error'
  let batchHighlightedLine = null;
  let sessionContext = null;

  // 右クリックコンテキストメニュー用
  let showContextMenu = false;
  let contextMenuPos = { x: 0, y: 0 };
  let contextMenuTarget = 'batch'; // 'batch' or 'step'

  // マルチステップ開発用状態
  let steps = [
    {
      id: 1,
      title: 'ステップ 1: アプリの起動',
      prompt: 'メモ帳を起動する',
      code: `// メモ帳を起動\nEXEC("notepad.exe")\nSLEEP(1.0)\nPRINT "[" + GETTIME() + "] メモ帳を起動しました"`,
      logs: '',
      status: 'idle', // 'idle', 'generating', 'running', 'success', 'error'
      sessionContext: null,
      highlightedLine: null
    }
  ];
  let activeStepIndex = 0;
  let isTestRunning = false;

  function closeContextMenu() {
    showContextMenu = false;
  }

  onMount(async () => {
        try {
            savePath = await App.GetDefaultScriptPath();
        } catch (err) {
            showStatus('デフォルト保存パスの取得に失敗しました', true);
        }
        window.addEventListener('click', closeContextMenu);

        // ★追加: テスト実行中のUWSCRリアルタイムログを画面側でキャッチする
        wailsRuntime.EventsOn('uwscr_log', (logLine) => {
            // 1. ワンショット開発モードのテスト実行中なら batchLogs にリアルタイム追記
            if (isBatchTestRunning) {
                batchLogs = (batchLogs ? batchLogs + "\n" : "") + logLine.message;
            }
            
            // 2. マルチステップ開発モードのテスト実行中なら、選択中のステップのlogsにリアルタイム追記
            if (isTestRunning && steps[activeStepIndex]) {
                steps[activeStepIndex].logs = (steps[activeStepIndex].logs ? steps[activeStepIndex].logs + "\n" : "") + logLine.message;
                steps = [...steps]; // Svelteのリアクティビティをトリガー
            }
        });

        return () => {
            window.removeEventListener('click', closeContextMenu);
        };
    });

    onDestroy(() => {
        window.removeEventListener('click', closeContextMenu);
        // ★追加: コンポーネント破棄時にイベント購読を解除してメモリリークを防ぐ
        wailsRuntime.EventsOff('uwscr_log');
    });

  function parseErrorLine(logs) {
    if (!logs) return null;
    const match = logs.match(/uws\[(\d+),\s*\d+\]/);
    if (match) {
      const lineNum = parseInt(match[1], 10);
      console.log("[ScriptDeveloper] Parsed error line:", lineNum);
      return lineNum;
    }
    return null;
  }

  function handleContextMenu(event, target) {
    event.preventDefault();
    contextMenuTarget = target;
    contextMenuPos = { x: event.clientX, y: event.clientY };
    showContextMenu = true;
  }

  async function handlePasteImage(event, isStepMode = false) {
    const items = (event.clipboardData || event.originalEvent?.clipboardData)?.items;
    if (!items) return;
    for (const item of items) {
      if (item.type.indexOf('image') !== -1) {
        const file = item.getAsFile();
        const reader = new FileReader();
        reader.onload = async (e) => {
          const base64 = e.target.result.split(',')[1];
          const newContext = {
            active_title: 'Clipboard Image',
            screenshot_base64: base64
          };
          if (isStepMode) {
            steps[activeStepIndex].sessionContext = newContext;
            steps = [...steps];
            showStatus('クリップボードから画像を貼り付けました (ステップ用)');
          } else {
            sessionContext = newContext;
            showStatus('クリップボードから画像を貼り付けました (ワンショット用)');
          }
        };
        reader.readAsDataURL(file);
        event.preventDefault();
        break;
      }
    }
  }

  async function pasteFromClipboard(isStepMode = false) {
    try {
      const clipboardItems = await navigator.clipboard.read();
      for (const item of clipboardItems) {
        const imageTypes = item.types.filter(type => type.startsWith('image/'));
        if (imageTypes.length > 0) {
          const blob = await item.getType(imageTypes[0]);
          const reader = new FileReader();
          reader.onload = (e) => {
            const base64 = e.target.result.split(',')[1];
            const newContext = {
              active_title: 'Clipboard Image',
              screenshot_base64: base64
            };
            if (isStepMode) {
              steps[activeStepIndex].sessionContext = newContext;
              steps = [...steps];
              showStatus('クリップボードから画像を取得しました (ステップ用)');
            } else {
              sessionContext = newContext;
              showStatus('クリップボードから画像を取得しました (ワンショット用)');
            }
          };
          reader.readAsDataURL(blob);
          return;
        }
      }
      showStatus('クリップボードに画像が見つかりませんでした', true);
    } catch (e) {
      console.error('Failed to read clipboard:', e);
      showStatus('クリップボードからの読み取り権限がないか、画像がありません', true);
    }
  }

  function showStatus(msg, err = false) {
    statusMessage = msg;
    isError = err;
    setTimeout(() => {
      if (statusMessage === msg) {
        statusMessage = '';
      }
    }, 6000);
  }

  function openLink(url) {
    wailsRuntime.BrowserOpenURL(url);
  }

  function parseMessage(msg) {
    if (!msg) return [];
    const urlRegex = /(https?:\/\/[^\s\)]+)/g;
    const parts = msg.split(urlRegex);
    return parts.map(part => {
      if (part.match(urlRegex)) {
        return { text: part, isLink: true };
      }
      return { text: part, isLink: false };
    });
  }

  // --- ワンショット開発用アクション ---
  async function handleCapture() {
    isCapturing = true;
    statusMessage = '現在の操作情報を取得中...';
    isError = false;

    setTimeout(async () => {
      try {
        const res = await App.CaptureSession();
        if (res) {
          sessionContext = res;
          showStatus('画面とアクティブウィンドウ情報をキャプチャしました');
        } else {
          showStatus('操作情報の取得に失敗しました', true);
        }
      } catch (err) {
        showStatus(`キャプチャエラー: ${err.message || err}`, true);
      } finally {
        isCapturing = false;
      }
    }, 500);
  }

  async function handleGenerate() {
    if (!prompt.trim()) {
      showStatus('自動化の指示（プロンプト）を入力してください', true);
      return;
    }

    isLoading = true;
    statusMessage = 'スクリプトを自動生成中...';
    isError = false;

    try {
      const contextJSON = sessionContext ? JSON.stringify(sessionContext) : '';
      const code = await App.GenerateScript(prompt, contextJSON);
      generatedCode = code;
      showStatus('UWSCR スクリプトを生成しました');
    } catch (err) {
      showStatus(`生成エラー: ${err.message || err}`, true);
    } finally {
      isLoading = false;
    }
  }

  async function handleLoad() {
    if (!savePath.trim()) {
      showStatus('読み込むスクリプトのパスを入力してください', true);
      return;
    }
    isLoading = true;
    try {
      const content = await App.ReadScriptFile(savePath);
      generatedCode = content;
      showStatus('スクリプトをファイルから読み込みました');
    } catch (err) {
      showStatus(`読み込みエラー: ${err.message || err}`, true);
    } finally {
      isLoading = false;
    }
  }

  async function handleSave() {
    if (!generatedCode.trim()) {
      showStatus('保存するコードがありません', true);
      return;
    }
    if (!savePath.trim()) {
      showStatus('保存先パスを入力してください', true);
      return;
    }

    try {
      await App.SaveScriptFile(savePath, generatedCode);
      showStatus('スクリプトをローカルファイルに保存しました');
    } catch (err) {
      showStatus(`保存エラー: ${err.message || err}`, true);
    }
  }

  async function handleRun() {
    if (!savePath.trim()) {
      showStatus('実行するスクリプトのパスが指定されていません', true);
      return;
    }

    try {
      if (generatedCode.trim()) {
        await App.SaveScriptFile(savePath, generatedCode);
      }
      
      showStatus('スクリプトを実行中... ログ表示へ自動遷移します');
      await App.RunScript(savePath);
      activeTab = 'run'; // 自動でコンソールタブに遷移
    } catch (err) {
      showStatus(`実行エラー: ${err.message || err}`, true);
    }
  }

  function clearCapturedContext() {
    sessionContext = null;
    showStatus('キャプチャ情報をクリアしました');
  }

  async function browseSavePath() {
    try {
      const selected = await App.SelectSaveFile(
        "保存先スクリプトを選択",
        "autoscript.uws",
        "UWSCR Script (*.uws)",
        "*.uws"
      );
      if (selected) {
        savePath = selected;
      }
    } catch (err) {
      showStatus(`保存先選択エラー: ${err.message || err}`, true);
    }
  }

  // --- マルチステップ開発用アクション ---
  function switchMode(newMode) {
    mode = newMode;
    if (mode === 'step' && steps.length === 0) {
      addStep();
    }
  }

  function addStep() {
    const nextId = steps.length > 0 ? Math.max(...steps.map(s => s.id)) + 1 : 1;
    steps = [...steps, {
      id: nextId,
      title: `ステップ ${nextId}: 新規操作`,
      prompt: '',
      code: '',
      logs: '',
      status: 'idle',
      sessionContext: null
    }];
    activeStepIndex = steps.length - 1;
  }

  function removeStep(index) {
    if (steps.length === 1) return;
    steps = steps.filter((_, i) => i !== index);
    if (activeStepIndex >= steps.length) {
      activeStepIndex = steps.length - 1;
    }
  }

  function selectStep(index) {
    activeStepIndex = index;
  }

  async function handleCaptureStep() {
    isCapturing = true;
    showStatus('ステップの操作画面をキャプチャ中...');
    
    setTimeout(async () => {
      try {
        const res = await App.CaptureSession();
        if (res) {
          steps[activeStepIndex].sessionContext = res;
          steps = [...steps];
          showStatus('画面とアクティブウィンドウ情報をキャプチャしました');
        } else {
          showStatus('操作情報の取得に失敗しました', true);
        }
      } catch (err) {
        showStatus(`キャプチャエラー: ${err.message || err}`, true);
      } finally {
        isCapturing = false;
      }
    }, 500);
  }

  function clearCapturedContextStep() {
    steps[activeStepIndex].sessionContext = null;
    steps = [...steps];
    showStatus('キャプチャ情報をクリアしました');
  }

  async function handleGenerateStep() {
    const step = steps[activeStepIndex];
    if (!step.prompt.trim()) {
      showStatus('このステップの指示（プロンプト）を入力してください', true);
      return;
    }

    step.status = 'generating';
    steps = [...steps];
    isLoading = true;
    showStatus('コードを自動生成中...');

    try {
      const contextJSON = step.sessionContext ? JSON.stringify(step.sessionContext) : '';
      const code = await App.GenerateScript(step.prompt, contextJSON);
      step.code = code;
      step.status = 'idle';
      showStatus('このステップのスクリプトを生成しました。直接編集して条件分岐を追記することも可能です。');
    } catch (err) {
      step.status = 'error';
      step.logs = `生成エラー: ${err.message || err}`;
      showStatus(`生成エラー: ${err.message || err}`, true);
    } finally {
      steps = [...steps];
      isLoading = false;
    }
  }

  async function handleTestRunStep() {
    const step = steps[activeStepIndex];
    if (!step.code.trim()) {
      showStatus('検証するコードがありません', true);
      return;
    }

    step.logs = ''; // ★実行前にこのステップのログ表示をクリア！
    step.status = 'running';
    steps = [...steps];
    isTestRunning = true;
    showStatus('UWSCRでテスト実行中...');

    try {
      const result = await App.TestRunScript(step.code);
      // ★修正: 完了時はステータスの同期のみ行う（ログはリアルタイム側で格納済み）
      step.status = result.success ? 'success' : 'error';
      step.highlightedLine = result.success ? null : parseErrorLine(result.logs);
      if (result.success) {
        showStatus('テスト実行が正常に終了しました（検証成功）');
      } else {
        showStatus('テスト実行がエラーで終了しました。自動修正を実行するか、コードを修正してください。', true);
      }
    } catch (err) {
      step.status = 'error';
      step.logs = (step.logs || '') + `\n[System Error] ${err.message || err}`;
      showStatus(`テスト実行エラー: ${err.message || err}`, true);
    } finally {
      steps = [...steps];
      isTestRunning = false;
    }
  }

  async function handleStopStep() {
    try {
      await App.StopScript();
      showStatus('UWSCRプロセスを強制終了しました。');
    } catch (err) {
      showStatus(`停止エラー: ${err.message || err}`, true);
    }
  }

  async function handleRefactorStep(instruction) {
    const step = steps[activeStepIndex];
    if (!step.code.trim()) {
      showStatus('修正対象のコードがありません', true);
      return;
    }
    if (!instruction.trim()) {
      showStatus('修正指示を入力してください', true);
      return;
    }

    step.status = 'generating';
    steps = [...steps];
    isLoading = true;
    showStatus('指示に基づいてコードを修正中...');

    try {
      const corrected = await App.CorrectScript(instruction, step.code, "");
      step.code = corrected;
      step.status = 'idle';
      step.logs = '';
      showStatus('指示に基づいてコードを修正しました。再度テスト実行を行ってください。');
    } catch (err) {
      step.status = 'error';
      showStatus(`修正エラー: ${err.message || err}`, true);
    } finally {
      steps = [...steps];
      isLoading = false;
    }
  }

  async function handleCorrectStep() {
    const step = steps[activeStepIndex];
    if (!step.code.trim() || !step.logs.trim()) {
      showStatus('修正対象のコードまたはログがありません', true);
      return;
    }

    step.status = 'generating';
    steps = [...steps];
    isLoading = true;
    showStatus('エラーログを解析してコードを自動修正中...');

    try {
      const corrected = await App.CorrectScript(step.prompt, step.code, step.logs);
      step.code = corrected;
      step.status = 'idle';
      showStatus('修正コードを自動生成しました。再度テスト実行を行ってください。');
    } catch (err) {
      step.status = 'error';
      showStatus(`自動修正エラー: ${err.message || err}`, true);
    } finally {
      steps = [...steps];
      isLoading = false;
    }
  }

  function combineAllSteps() {
    let combined = `// ==========================================\n`;
    combined += `// actgram::UWSCR 自動生成・検証済スクリプト\n`;
    combined += `// 生成日時: ${new Date().toLocaleString()}\n`;
    combined += `// ==========================================\n\n`;

    steps.forEach((step, idx) => {
      combined += `// ------------------------------------------\n`;
      combined += `// Step ${idx + 1}: ${step.title}\n`;
      combined += `// 指示: ${step.prompt}\n`;
      combined += `// ------------------------------------------\n`;
      combined += step.code + `\n\n`;
    });

    return combined.trim();
  }

  async function handleSaveCombined() {
    const combined = combineAllSteps();
    if (!combined) {
      showStatus('保存するコードがありません', true);
      return;
    }
    if (!savePath.trim()) {
      showStatus('保存先パスを入力してください', true);
      return;
    }

    try {
      await App.SaveScriptFile(savePath, combined);
      showStatus('すべてのステップを結合して保存しました！');
    } catch (err) {
      showStatus(`保存エラー: ${err.message || err}`, true);
    }
  }

  async function handleRunCombined() {
    const combined = combineAllSteps();
    if (!combined) {
      showStatus('実行するコードがありません', true);
      return;
    }
    if (!savePath.trim()) {
      showStatus('保存先パスが指定されていません', true);
      return;
    }

    try {
      await App.SaveScriptFile(savePath, combined);
      showStatus('結合スクリプトを実行中... ログ表示へ自動遷移します');
      await App.RunScript(savePath);
      activeTab = 'run'; // 自動でコンソールタブに遷移
    } catch (err) {
      showStatus(`実行エラー: ${err.message || err}`, true);
    }
  }

  async function handleTestRunBatch() {
    if (!generatedCode.trim()) {
      showStatus('検証するコードがありません', true);
      return;
    }

    batchLogs = ''; // ★最初の一括上書きではなく、実行前にまずログを綺麗にクリアする！
    batchStatus = 'running';
    isBatchTestRunning = true;
    showStatus('UWSCRでテスト実行中...');

    try {
      const result = await App.TestRunScript(generatedCode);
      // ★修正: 完了時はログ全体を再設定するのではなく、ステータス判定とハイライト処理だけを行う
      batchStatus = result.success ? 'success' : 'error';
      batchHighlightedLine = result.success ? null : parseErrorLine(result.logs);
      if (result.success) {
        showStatus('テスト実行が正常に終了しました（検証成功）');
      } else {
        showStatus('テスト実行がエラーで終了しました。自動修正を実行するか、コードを修正してください。', true);
      }
    } catch (err) {
      batchStatus = 'error';
      batchLogs = (batchLogs || '') + `\n[System Error] ${err.message || err}`;
      showStatus(`テスト実行エラー: ${err.message || err}`, true);
    } finally {
      isBatchTestRunning = false;
    }
  }

  async function handleStopBatch() {
    try {
      await App.StopScript();
      showStatus('UWSCRプロセスを強制終了しました。');
    } catch (err) {
      showStatus(`停止エラー: ${err.message || err}`, true);
    }
  }

  let showRefactorModal = false;
  let refactorPrompt = '';
  let refactorTarget = 'batch'; // 'batch' or 'step'

  function openRefactorModal(target) {
    refactorTarget = target;
    refactorPrompt = '';
    showRefactorModal = true;
  }

  async function executeRefactor() {
    if (!refactorPrompt.trim()) {
      showStatus('修正指示を入力してください', true);
      return;
    }
    showRefactorModal = false;
    if (refactorTarget === 'batch') {
      await handleRefactorBatch(refactorPrompt);
    } else {
      await handleRefactorStep(refactorPrompt);
    }
  }

  async function handleRefactorBatch(instruction) {
    if (!generatedCode.trim()) {
      showStatus('修正対象のコードがありません', true);
      return;
    }
    if (!instruction.trim()) {
      showStatus('修正指示を入力してください', true);
      return;
    }

    batchStatus = 'generating';
    isLoading = true;
    showStatus('指示に基づいてコードを修正中...');

    try {
      const corrected = await App.CorrectScript(instruction, generatedCode, "");
      generatedCode = corrected;
      batchStatus = 'idle';
      batchLogs = '';
      showStatus('指示に基づいてコードを修正しました。再度テスト実行を行ってください。');
    } catch (err) {
      batchStatus = 'error';
      showStatus(`修正エラー: ${err.message || err}`, true);
    } finally {
      isLoading = false;
    }
  }

  async function handleCorrectBatch() {
    if (!generatedCode.trim() || !batchLogs.trim()) {
      showStatus('修正対象のコードまたはログがありません', true);
      return;
    }

    batchStatus = 'generating';
    isLoading = true;
    showStatus('エラーログを解析してコードを自動修正中...');

    try {
      const corrected = await App.CorrectScript(prompt || "自動生成スクリプト", generatedCode, batchLogs);
      generatedCode = corrected;
      batchStatus = 'idle';
      showStatus('修正コードを自動生成しました。再度テスト実行を行ってください。');
    } catch (err) {
      batchStatus = 'error';
      showStatus(`自動修正エラー: ${err.message || err}`, true);
    } finally {
      isLoading = false;
    }
  }

  async function handleSaveStep(index) {
    const step = steps[index];
    if (!step || !step.code.trim()) {
      showStatus('保存するコードがありません', true);
      return;
    }
    
    try {
      const defaultName = `step_${step.id}_${step.title.replace(/[:\/\*\?"<>\|]/g, "_")}.uws`;
      const selected = await App.SelectSaveFile(
        "このステップの保存先を選択",
        defaultName,
        "UWSCR Script (*.uws)",
        "*.uws"
      );
      if (selected) {
        await App.SaveScriptFile(selected, step.code);
        showStatus(`${step.title} のスクリプトを保存しました`);
      }
    } catch (err) {
      showStatus(`ステップ保存エラー: ${err.message || err}`, true);
    }
  }
</script>

<div class="developer-container">
  <div class="header-section">
    <h1>スクリプト開発</h1>
    <p class="subtitle">自然言語と現在の画面状況を組み合わせてUWSCRスクリプトを作成・編集・実行します</p>
    
    <div class="mode-tabs">
      <button class="mode-tab-btn {mode === 'batch' ? 'active' : ''}" on:click={() => switchMode('batch')}>ワンショット開発</button>
      <button class="mode-tab-btn {mode === 'step' ? 'active' : ''}" on:click={() => switchMode('step')}>マルチステップ開発</button>
    </div>
  </div>

  {#if mode === 'batch'}
    <!-- ワンショット開発モード -->
    <div class="workspace-grid">
      <!-- 左ペイン: コンテキスト取得 & プロンプト指示 -->
      <div class="input-panel">
        <!-- 画面キャプチャ部分 -->
        <div class="panel-section">
          <div class="section-header">
            <h3>1. 操作コンテキスト（オプション）</h3>
            {#if sessionContext}
              <button class="btn-ghost btn-small danger-accent" on:click={clearCapturedContext}>クリア</button>
            {/if}
          </div>
          
          <p class="section-desc">
            現在操作中の画面キャプチャとウィンドウ名を取得し、生成されるスクリプトの精度を劇的に向上させます。
          </p>

          <div 
            class="capture-box" 
            on:paste={(e) => handlePasteImage(e, false)}
            on:contextmenu={(e) => handleContextMenu(e, 'batch')}
            tabindex="0"
            style="outline: none;"
          >
            {#if isCapturing}
              <div class="capturing-indicator">
                <span class="spinner"></span>
                <p>画面とタイトルを取得中...</p>
              </div>
            {:else if sessionContext}
              <div class="context-info">
                <div class="info-row">
                  <span class="label">アクティブウィンドウ:</span>
                  <span class="value">{sessionContext.active_title}</span>
                </div>
                <div class="image-preview">
                  <img src="data:image/png;base64,{sessionContext.screenshot_base64}" alt="Captured Desktop" on:click={() => window.dispatchEvent(new CustomEvent('zoom-image', { detail: 'data:image/png;base64,' + sessionContext.screenshot_base64 }))} style="cursor: zoom-in;" />
                </div>
              </div>
            {:else}
              <div class="capture-placeholder" style="display: flex; flex-direction: column; gap: 8px;">
                <button class="btn-outlined w-100" on:click={handleCapture}>
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"></path>
                    <circle cx="12" cy="13" r="4"></circle>
                  </svg>
                  現在の画面状況をキャプチャ
                </button>
                <button class="btn-outlined w-100" style="border-style: dashed; font-size: 0.8rem;" on:click={() => pasteFromClipboard(false)}>
                  📋 クリップボードから貼付
                </button>
              </div>
            {/if}
          </div>
        </div>

        <!-- プロンプト入力部分 -->
        <div class="panel-section">
          <h3>2. 自動化の指示</h3>
          <p class="section-desc">行いたい自動化の操作内容を日本語で指示してください。</p>
          <textarea
            class="prompt-textarea"
            placeholder="例: メモ帳を起動して、現在のアクティブウィンドウのタイトルを書き込み、ファイルをデスクトップに保存してください。"
            bind:value={prompt}
            disabled={isLoading}
          ></textarea>

          <button 
            class="btn-solid btn-large w-100 generate-btn" 
            on:click={handleGenerate} 
            disabled={isLoading || isCapturing}
          >
            {#if isLoading}
              <span class="spinner white-spinner"></span>
              スクリプトを生成中...
            {:else}
              スクリプトを自動生成する
            {/if}
          </button>
        </div>
      </div>

      <!-- 右ペイン: スクリプトエディタ & 実行 -->
      <div class="editor-panel">
        <div class="panel-section h-100 flex-column">
          <div class="section-header">
            <h3>生成された UWSCR スクリプト (.uws)</h3>
          </div>

          <div class="editor-container">
            <CodeEditor
              bind:value={generatedCode}
              placeholder="// 生成されたコードがここに表示されます。直接編集も可能です。"
              highlightedLine={batchHighlightedLine}
              on:input={() => batchHighlightedLine = null}
            />
          </div>

          <!-- ワンショット検証用のテスト実行とデバッグ -->
          <div class="step-test-section" style="margin-top: 16px; border-top: 1px solid var(--border-color); padding-top: 16px; margin-bottom: 16px;">
            <div class="step-section-title-row" style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 8px;">
              <h4 style="margin: 0; font-size: 0.8rem; font-weight: 600; color: var(--text-primary);">テスト実行と自動デバッグ</h4>
              <div class="test-btn-group" style="display: flex; gap: 8px;">
                <button class="btn-outlined btn-small" on:click={handleTestRunBatch} disabled={isBatchTestRunning || !generatedCode} style="font-size: 0.7rem; padding: 4px 8px;">
                  ▶ テスト実行
                </button>
                <button class="btn-outlined btn-small danger-accent" on:click={handleStopBatch} disabled={!isBatchTestRunning} style="font-size: 0.7rem; padding: 4px 8px;">
                  ■ 強制停止
                </button>
                <button class="btn-outlined btn-small" on:click={() => openRefactorModal('batch')} disabled={isLoading || !generatedCode} style="font-size: 0.7rem; padding: 4px 8px;">
                  🪄 指示して修正
                </button>
                <button class="btn-solid btn-small correct-btn" on:click={handleCorrectBatch} disabled={isLoading || batchStatus !== 'error' || !batchLogs} style="font-size: 0.7rem; padding: 4px 8px;">
                  🪄 エラーを自動修正
                </button>
              </div>
            </div>
            
            <div class="step-logs-box" style="height: 100px; min-height: 100px;">
              {#if batchLogs}
                <pre class="logs-content {batchStatus === 'error' ? 'log-error' : ''}">{batchLogs}</pre>
              {:else}
                <div class="empty-logs">検証のテスト実行結果（ログおよびタイムスタンプ）がここに表示されます。</div>
              {/if}
            </div>
          </div>

          <!-- 保存と実行の設定 -->
          <div class="action-footer">
            <div class="path-input-group">
              <label for="script-save-path">保存先パス:</label>
              <div class="input-with-btn">
                <input
                  id="script-save-path"
                  type="text"
                  class="path-input"
                  bind:value={savePath}
                  placeholder="C:\path\to\script.uws"
                />
                <button class="btn-outlined btn-browse" on:click={browseSavePath} disabled={isLoading}>
                  選択...
                </button>
              </div>
            </div>

            <div class="button-group">
              <button class="btn-outlined" on:click={handleLoad} disabled={isLoading}>
                読み込む
              </button>
              <button class="btn-outlined" on:click={handleSave} disabled={isLoading || !generatedCode}>
                保存する
              </button>
              <button class="btn-solid" on:click={handleRun} disabled={isLoading || !generatedCode}>
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="5 3 19 12 5 21 5 3"></polygon>
                </svg>
                保存して実行
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  {:else}
    <!-- マルチステップ開発モード -->
    <div class="step-workspace-grid">
      <!-- 左ペイン: ステップ一覧 -->
      <div class="step-sidebar">
        <div class="sidebar-header">
          <h3>操作シナリオステップ</h3>
          <button class="btn-outlined btn-small" on:click={addStep}>＋ 追加</button>
        </div>

        <div class="step-list-container">
          {#each steps as step, idx}
            <div class="step-item-card {activeStepIndex === idx ? 'active' : ''}" on:click={() => selectStep(idx)}>
              <div class="step-item-header">
                <span class="step-number-badge">Step {idx + 1}</span>
                {#if step.status === 'success'}
                  <span class="status-badge success">✓ 検証済</span>
                {:else if step.status === 'error'}
                  <span class="status-badge error">⚠ エラー</span>
                {:else if step.status === 'generating'}
                  <span class="status-badge loading">生成中</span>
                {:else if step.status === 'running'}
                  <span class="status-badge loading">実行中</span>
                {/if}
              </div>
              <input type="text" class="step-item-title-input" bind:value={step.title} placeholder="ステップのタイトル" />
              <button class="btn-delete-step" on:click|stopPropagation={() => removeStep(idx)} disabled={steps.length === 1} title="削除">
                ×
              </button>
            </div>
          {/each}
        </div>
      </div>

      <!-- 右ペイン: 選択中ステップの編集・検証 -->
      <div class="step-main-panel">
        {#if steps[activeStepIndex]}
          {@const currentStep = steps[activeStepIndex]}
          <div class="step-detail-card">
            <!-- 指示とキャプチャ -->
            <div class="step-detail-section">
              <div class="step-section-title">
                <h4>1. このステップの自動化指示</h4>
              </div>
              <p class="section-desc">このステップで行う操作の指示と、対象画面のキャプチャ（オプション）を指定します。</p>
              
              <div class="step-input-row">
                <textarea
                  class="prompt-textarea step-prompt-textarea"
                  placeholder="例: ID入力フィールドをクリックして「user」と入力する。"
                  bind:value={steps[activeStepIndex].prompt}
                  disabled={isLoading}
                ></textarea>

                <div 
                  class="step-capture-box" 
                  on:paste={(e) => handlePasteImage(e, true)}
                  on:contextmenu={(e) => handleContextMenu(e, 'step')}
                  tabindex="0"
                  style="outline: none;"
                >
                  {#if currentStep.sessionContext}
                    <div class="step-context-info">
                      <div class="step-image-preview">
                        <img src="data:image/png;base64,{currentStep.sessionContext.screenshot_base64}" alt="Captured Desktop" on:click={() => window.dispatchEvent(new CustomEvent('zoom-image', { detail: 'data:image/png;base64,' + currentStep.sessionContext.screenshot_base64 }))} style="cursor: zoom-in;" />
                      </div>
                      <div class="step-context-actions">
                        <span class="window-title-tag" title={currentStep.sessionContext.active_title}>{currentStep.sessionContext.active_title}</span>
                        <button class="btn-ghost btn-small danger-accent" on:click={clearCapturedContextStep}>クリア</button>
                      </div>
                    </div>
                  {:else}
                    <div class="step-capture-placeholder" style="display: flex; flex-direction: column; gap: 6px;">
                      <button class="btn-outlined btn-small" on:click={handleCaptureStep} disabled={isCapturing}>
                        📸 キャプチャ
                      </button>
                      <button class="btn-outlined btn-small" style="border-style: dashed;" on:click={() => pasteFromClipboard(true)}>
                        📋 クリップボードから貼付
                      </button>
                    </div>
                  {/if}
                </div>
              </div>

              <button
                class="btn-solid generate-btn w-100"
                on:click={handleGenerateStep}
                disabled={isLoading || isCapturing}
              >
                {#if currentStep.status === 'generating'}
                  <span class="spinner white-spinner"></span> コード生成中...
                {:else}
                  このステップのコードを自動生成する
                {/if}
              </button>
            </div>

            <!-- コードエディタ -->
            <div class="step-code-section">
              <div class="step-section-title-row" style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 8px;">
                <h4 style="margin: 0; font-size: 0.8rem; font-weight: 600; color: var(--text-primary);">2. UWSCRスクリプト（直接編集して条件分岐を追記可能）</h4>
                <button class="btn-outlined btn-small" on:click={() => handleSaveStep(activeStepIndex)} disabled={isLoading || !steps[activeStepIndex].code} style="font-size: 0.7rem; padding: 4px 8px;">
                  💾 このステップを保存
                </button>
              </div>
              <p class="section-desc">生成されたスクリプトです。ユーザーが直接エディタ内で `IFB` などの条件分岐を追記して教えることができます。</p>
              
              <div class="editor-container step-editor-container">
                <CodeEditor
                  bind:value={steps[activeStepIndex].code}
                  placeholder="// コードがここに生成されます。直接条件分岐（IFB 〜 ENDIF）などを書き足してください。"
                  highlightedLine={steps[activeStepIndex].highlightedLine}
                  on:input={() => {
                    steps[activeStepIndex].highlightedLine = null;
                    steps = [...steps];
                  }}
                />
              </div>
            </div>

            <!-- 検証とデバッグ -->
            <div class="step-test-section">
              <div class="step-section-title-row">
                <h4>3. テスト実行と自動デバッグ</h4>
                <div class="test-btn-group">
                  <button class="btn-outlined btn-small" on:click={handleTestRunStep} disabled={isTestRunning || !currentStep.code}>
                    ▶ テスト実行
                  </button>
                  <button class="btn-outlined btn-small danger-accent" on:click={handleStopStep} disabled={!isTestRunning}>
                    ■ 強制停止
                  </button>
                  <button class="btn-outlined btn-small" on:click={() => openRefactorModal('step')} disabled={isLoading || !currentStep.code}>
                    🪄 指示して修正
                  </button>
                  <button class="btn-solid btn-small correct-btn" on:click={handleCorrectStep} disabled={isLoading || currentStep.status !== 'error' || !currentStep.logs}>
                    🪄 エラーを自動修正
                  </button>
                </div>
              </div>
              
              <div class="step-logs-box">
                {#if currentStep.logs}
                  <pre class="logs-content {currentStep.status === 'error' ? 'log-error' : ''}">{currentStep.logs}</pre>
                {:else}
                  <div class="empty-logs">検証のテスト実行結果（ログおよびタイムスタンプ）がここに表示されます。</div>
                {/if}
              </div>
            </div>
          </div>
        {/if}

        <!-- 結合保存フッター -->
        <div class="step-action-footer">
          <div class="path-input-group">
            <label for="combined-save-path">結合スクリプト保存先:</label>
            <div class="input-with-btn">
              <input
                id="combined-save-path"
                type="text"
                class="path-input"
                bind:value={savePath}
                placeholder="C:\path\to\combined_script.uws"
              />
              <button class="btn-outlined btn-browse" on:click={browseSavePath}>
                選択...
              </button>
            </div>
          </div>

          <div class="button-group">
            <button class="btn-outlined" on:click={handleSaveCombined} disabled={isLoading || isTestRunning}>
              全ステップを結合して保存
            </button>
            <button class="btn-solid" on:click={handleRunCombined} disabled={isLoading || isTestRunning}>
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="5 3 19 12 5 21 5 3"></polygon>
              </svg>
              結合して実行
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}

  <!-- ステータスバー（画面下部） -->
  {#if statusMessage}
    <div class="status-toast {isError ? 'status-error' : 'status-success'}">
      <p>
        {#each parseMessage(statusMessage) as part}
          {#if part.isLink}
            <a href="javascript:void(0)" on:click|preventDefault={() => openLink(part.text)} class="msg-link">{part.text}</a>
          {:else}
            {part.text}
          {/if}
        {/each}
      </p>
    </div>
  {/if}

  {#if showContextMenu}
    <div 
      class="custom-context-menu" 
      style="position: fixed; top: {contextMenuPos.y}px; left: {contextMenuPos.x}px; z-index: 10000; background: #1e1e24; border: 1px solid #2d2d3c; border-radius: 6px; box-shadow: 0 4px 12px rgba(0,0,0,0.5); padding: 4px 0;"
    >
      <button 
        style="background: transparent; border: none; color: #e2e8f0; padding: 8px 16px; font-size: 0.85rem; width: 100%; text-align: left; cursor: pointer; display: flex; align-items: center; gap: 8px;"
        on:click={() => {
          pasteFromClipboard(contextMenuTarget === 'step');
          closeContextMenu();
        }}
      >
        📋 クリップボードから画像を貼り付け
      </button>
    </div>
  {/if}
  <!-- 修正指示入力ダイアログ (モーダル小窓) -->
  {#if showRefactorModal}
    <div class="refactor-modal-overlay" on:click={() => showRefactorModal = false}>
      <div class="refactor-modal-window" on:click|stopPropagation>
        <div class="refactor-modal-header">
          <h3>🪄 スクリプトの修正指示</h3>
          <button class="refactor-modal-close" on:click={() => showRefactorModal = false}>✕</button>
        </div>
        <div class="refactor-modal-body">
          <p class="refactor-modal-desc">
            既存のスクリプトに対して、AIへの修正・調整指示を入力してください。
          </p>
          <textarea
            class="refactor-textarea"
            placeholder="例: ウィンドウの座標がズレているので mouseorg(id) を適切に呼び出すように修正して。 / ボタンのクリック処理を追加して。"
            bind:value={refactorPrompt}
          ></textarea>
        </div>
        <div class="refactor-modal-footer">
          <button class="btn-outlined" on:click={() => showRefactorModal = false}>キャンセル</button>
          <button class="btn-solid refactor-exec-btn" on:click={executeRefactor} disabled={!refactorPrompt.trim()}>修正を実行</button>
        </div>
      </div>
    </div>
  {/if}
</div>

<style>
  .developer-container {
    --text-color: var(--text-primary);
    --bg-color: var(--bg-primary);
    --card-bg: var(--bg-card);
    display: flex;
    flex-direction: column;
    height: 100%;
    box-sizing: border-box;
    overflow: hidden;
  }

  .header-section {
    margin-bottom: 24px;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .header-section h1 {
    font-size: 1.5rem;
    font-weight: 600;
    margin: 0;
    color: var(--text-color);
  }

  .subtitle {
    font-size: 0.85rem;
    color: var(--text-secondary);
    margin: 0;
  }

  /* モード切り替えタブ */
  .mode-tabs {
    display: flex;
    gap: 8px;
    border-bottom: 1px solid var(--border-color);
    padding-bottom: 8px;
    margin-top: 8px;
  }

  .mode-tab-btn {
    background: transparent;
    border: none;
    color: var(--text-secondary);
    padding: 8px 16px;
    font-size: 0.85rem;
    font-weight: 500;
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.2s ease;
  }

  .mode-tab-btn:hover {
    background: rgba(128, 128, 128, 0.05);
    color: var(--text-color);
  }

  .mode-tab-btn.active {
    background: var(--accent-color);
    color: var(--bg-secondary);
    font-weight: 600;
  }

  .workspace-grid {
    display: grid;
    grid-template-columns: 1fr 1.2fr;
    gap: 24px;
    flex: 1;
    min-height: 0;
  }

  .input-panel {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: 100%;
    overflow-y: auto;
    padding-right: 4px;
    min-height: 0;
  }

  .editor-panel {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: 100%;
    overflow-y: auto;
    padding-right: 4px;
    min-height: 0;
  }

  .panel-section {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 20px;
    box-sizing: border-box;
  }

  .flex-column {
    display: flex;
    flex-direction: column;
  }

  .h-100 {
    height: 100%;
  }

  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
  }

  h3 {
    font-size: 0.95rem;
    font-weight: 600;
    margin: 0;
    color: var(--text-color);
  }

  .section-desc {
    font-size: 0.8rem;
    color: var(--text-secondary);
    margin: 0 0 12px 0;
  }

  .capture-box {
    border: 1px dashed var(--border-color);
    border-radius: 6px;
    padding: 16px;
    min-height: 100px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: var(--input-bg);
  }

  .capture-placeholder {
    width: 100%;
  }

  .capturing-indicator {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
    color: var(--text-secondary);
    font-size: 0.8rem;
  }

  .context-info {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .info-row {
    font-size: 0.8rem;
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .info-row .label {
    color: var(--text-secondary);
    font-weight: 500;
  }

  .info-row .value {
    color: var(--text-color);
    font-family: Consolas, Monaco, monospace;
    font-weight: 600;
    background: var(--card-bg);
    padding: 6px 10px;
    border-radius: 4px;
    border: 1px solid var(--border-color);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .image-preview {
    border: 1px solid var(--border-color);
    border-radius: 4px;
    overflow: hidden;
    max-height: 140px;
    display: flex;
    justify-content: center;
    background: #000;
  }

  .image-preview img {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
  }

  .prompt-textarea {
    width: 100%;
    height: 110px;
    border: 1px solid var(--border-color);
    border-radius: 6px;
    background: var(--input-bg);
    color: var(--text-color);
    padding: 12px;
    font-family: inherit;
    font-size: 0.85rem;
    resize: none;
    box-sizing: border-box;
    margin-bottom: 16px;
    transition: border-color 0.2s, outline 0.2s;
  }

  .prompt-textarea:focus, .code-textarea:focus, .path-input:focus {
    outline: none;
    border-color: var(--text-color);
  }

  .generate-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
  }

  .editor-container {
    flex: 1;
    margin-bottom: 20px;
    min-height: 280px;
  }

  .code-textarea {
    width: 100%;
    height: 100%;
    border: 1px solid var(--border-color);
    border-radius: 6px;
    background: var(--input-bg);
    color: var(--text-color);
    padding: 16px;
    font-family: Consolas, Monaco, 'Courier New', monospace;
    font-size: 0.85rem;
    line-height: 1.5;
    resize: none;
    box-sizing: border-box;
    white-space: pre;
    overflow-x: auto;
  }

  .action-footer {
    display: flex;
    flex-direction: column;
    gap: 16px;
    border-top: 1px solid var(--border-color);
    padding-top: 16px;
  }

  .path-input-group {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .path-input-group label {
    font-size: 0.8rem;
    font-weight: 500;
    color: var(--text-secondary);
    white-space: nowrap;
  }

  .path-input {
    flex: 1;
    border: 1px solid var(--border-color);
    border-radius: 6px;
    background: var(--input-bg);
    color: var(--text-color);
    padding: 8px 12px;
    font-size: 0.8rem;
    font-family: Consolas, Monaco, monospace;
  }

  .input-with-btn {
    display: flex;
    gap: 8px;
    flex: 1;
    width: 100%;
  }

  .btn-browse {
    padding: 8px 12px !important;
    white-space: nowrap;
    height: 100%;
  }

  .button-group {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
  }

  .btn-solid, .btn-outlined, .btn-ghost {
    font-family: inherit;
    font-size: 0.8rem;
    font-weight: 500;
    border-radius: 6px;
    cursor: pointer;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    padding: 8px 16px;
    box-sizing: border-box;
  }

  .btn-large {
    padding: 12px 24px;
    font-size: 0.85rem;
  }

  .btn-small {
    padding: 6px 12px;
    font-size: 0.75rem;
  }

  .btn-solid {
    background: var(--text-color);
    color: var(--bg-color);
    border: 1px solid var(--text-color);
  }

  .btn-solid:hover:not(:disabled) {
    opacity: 0.9;
  }

  .btn-outlined {
    background: transparent;
    color: var(--text-color);
    border: 1px solid var(--border-color);
  }

  .btn-outlined:hover:not(:disabled) {
    background: rgba(128, 128, 128, 0.05);
    border-color: var(--text-color);
  }

  .btn-ghost {
    background: transparent;
    color: var(--text-secondary);
    border: 1px solid transparent;
  }

  .btn-ghost:hover:not(:disabled) {
    background: rgba(128, 128, 128, 0.05);
    color: var(--text-color);
  }

  .danger-accent {
    color: #ef4444 !important;
  }

  .danger-accent:hover {
    background: rgba(239, 68, 68, 0.08) !important;
  }

  .w-100 {
    width: 100%;
  }

  button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .spinner {
    width: 16px;
    height: 16px;
    border: 2px solid rgba(128, 128, 128, 0.2);
    border-radius: 50%;
    border-top-color: var(--text-color);
    animation: spin 0.8s linear infinite;
    display: inline-block;
  }

  .white-spinner {
    border-top-color: var(--bg-color);
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }

  /* ステータス通知トースト */
  .status-toast {
    position: fixed;
    bottom: 24px;
    right: 24px;
    padding: 12px 20px;
    border-radius: 6px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    font-size: 0.8rem;
    font-weight: 500;
    max-width: 380px;
    z-index: 1000;
    animation: slideUp 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .status-toast p {
    margin: 0;
  }

  .status-success {
    background: var(--text-color);
    color: var(--bg-color);
    border: 1px solid var(--text-color);
  }

  .status-error {
    background: #ef4444;
    color: #ffffff;
    border: 1px solid #ef4444;
  }

  .msg-link {
    color: inherit;
    text-decoration: underline;
    cursor: pointer;
    font-weight: 600;
    transition: opacity 0.2s ease;
  }

  .msg-link:hover {
    opacity: 0.8;
  }

  @keyframes slideUp {
    from { transform: translateY(100%); opacity: 0; }
    to { transform: translateY(0); opacity: 1; }
  }

  /* --- 伴走型ステップ開発モードスタイル --- */
  .step-workspace-grid {
    display: grid;
    grid-template-columns: 260px 1fr;
    gap: 24px;
    flex: 1;
    height: 100%;
    min-height: 0;
    align-items: stretch;
    overflow: hidden;
  }

  .step-sidebar {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 16px;
    height: 100%;
    min-height: 0;
    overflow-y: auto;
  }

  .sidebar-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid var(--border-color);
    padding-bottom: 8px;
  }

  .sidebar-header h3 {
    margin: 0;
    font-size: 0.9rem;
    font-weight: 600;
    color: var(--text-color);
  }

  .step-list-container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    overflow-y: auto;
    flex: 1;
  }

  .step-item-card {
    background: var(--bg-color);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 12px;
    position: relative;
    cursor: pointer;
    transition: all 0.2s ease;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .step-item-card:hover {
    border-color: var(--text-color);
    transform: translateY(-1px);
  }

  .step-item-card.active {
    border-color: var(--text-color);
    background: var(--accent-soft);
    border-left: 6px solid var(--accent-color);
  }

  .step-item-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .step-number-badge {
    font-size: 0.75rem;
    font-weight: 600;
    background: rgba(128, 128, 128, 0.1);
    padding: 2px 6px;
    border-radius: 4px;
    color: var(--text-color);
  }

  .status-badge {
    font-size: 0.7rem;
    font-weight: 600;
    padding: 2px 6px;
    border-radius: 4px;
  }

  .status-badge.success {
    background: rgba(46, 213, 115, 0.12);
    color: #2ed573;
  }

  .status-badge.error {
    background: rgba(239, 68, 68, 0.12);
    color: #ef4444;
  }

  .status-badge.loading {
    background: rgba(255, 159, 67, 0.12);
    color: #ff9f43;
    animation: pulse 1.5s infinite;
  }

  .step-item-title-input {
    background: transparent;
    border: none;
    border-bottom: 1px solid transparent;
    color: var(--text-color);
    font-size: 0.8rem;
    font-weight: 500;
    padding: 2px 0;
    width: calc(100% - 20px);
  }

  .step-item-title-input:focus {
    outline: none;
    border-bottom-color: var(--text-color);
  }

  .btn-delete-step {
    position: absolute;
    top: 8px;
    right: 8px;
    background: transparent;
    border: none;
    color: var(--text-secondary);
    font-size: 1.1rem;
    cursor: pointer;
    line-height: 1;
    padding: 2px 6px;
    border-radius: 4px;
    transition: background 0.2s;
  }

  .btn-delete-step:hover:not(:disabled) {
    background: rgba(239, 68, 68, 0.1);
    color: #ef4444;
  }

  .step-main-panel {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: 100%;
    overflow-y: auto;
    padding-right: 4px;
    min-height: 0;
  }

  .step-detail-card {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 24px;
    display: flex;
    flex-direction: column;
    gap: 24px;
    flex-shrink: 0;
  }

  .step-detail-section, .step-code-section, .step-test-section {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .step-section-title {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .step-section-title-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid var(--border-color);
    padding-bottom: 8px;
    margin-bottom: 8px;
  }

  h4 {
    margin: 0;
    font-size: 0.9rem;
    font-weight: 600;
    color: var(--text-color);
  }

  .step-input-row {
    display: grid;
    grid-template-columns: 1fr 280px;
    gap: 16px;
    align-items: stretch;
  }

  .step-prompt-textarea {
    height: 120px;
    margin-bottom: 0;
  }

  .step-capture-box {
    border: 1px dashed var(--border-color);
    border-radius: 6px;
    background: var(--input-bg);
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 8px;
    box-sizing: border-box;
    height: 120px;
  }

  .step-capture-placeholder {
    display: flex;
    width: 100%;
    justify-content: center;
  }

  .step-context-info {
    display: flex;
    flex-direction: column;
    gap: 6px;
    width: 100%;
    height: 100%;
  }

  .step-image-preview {
    border: 1px solid var(--border-color);
    border-radius: 4px;
    overflow: hidden;
    flex: 1;
    display: flex;
    justify-content: center;
    background: #000;
    max-height: 80px;
  }

  .step-image-preview img {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
  }

  .step-context-actions {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 8px;
  }

  .window-title-tag {
    font-size: 0.7rem;
    color: var(--text-secondary);
    background: rgba(128, 128, 128, 0.1);
    padding: 2px 6px;
    border-radius: 4px;
    max-width: 140px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .step-editor-container {
    height: 180px;
    min-height: 150px;
  }

  .test-btn-group {
    display: flex;
    gap: 8px;
  }

  .correct-btn {
    background: var(--accent-red);
    border-color: var(--accent-red);
    color: var(--bg-primary);
  }

  .correct-btn:hover:not(:disabled) {
    background: var(--accent-red-hover);
    border-color: var(--accent-red-hover);
  }

  .step-logs-box {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 12px 16px;
    min-height: 80px;
    max-height: 180px;
    overflow-y: auto;
    font-family: Consolas, Monaco, monospace;
    font-size: 0.8rem;
    line-height: 1.4;
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
  }

  .empty-logs {
    color: var(--text-secondary);
    font-size: 0.78rem;
    opacity: 0.7;
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
  }

  .logs-content {
    margin: 0;
    white-space: pre-wrap;
    color: var(--text-color);
  }

  .logs-content.log-error {
    color: #ef4444;
  }

  .step-action-footer {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  @keyframes pulse {
    0% { opacity: 0.6; }
    50% { opacity: 1; }
    100% { opacity: 0.6; }
  }

  /* 修正指示ダイアログ用スタイル */
  .refactor-modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(4px);
    z-index: 10000;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .refactor-modal-window {
    width: 480px;
    background: #1e1e24;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.4);
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }

  :global(html.light-mode) .refactor-modal-window {
    background: #ffffff;
  }

  .refactor-modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 16px;
    border-bottom: 1px solid var(--border-color);
    background: rgba(255, 255, 255, 0.02);
  }

  .refactor-modal-header h3 {
    margin: 0;
    font-size: 0.95rem;
    font-weight: 600;
    color: var(--text-color);
  }

  .refactor-modal-close {
    background: transparent;
    border: none;
    color: var(--text-secondary);
    cursor: pointer;
    font-size: 0.9rem;
  }

  .refactor-modal-close:hover {
    color: var(--text-color);
  }

  .refactor-modal-body {
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .refactor-modal-desc {
    font-size: 0.78rem;
    color: var(--text-secondary);
    margin: 0;
  }

  .refactor-textarea {
    width: 100%;
    height: 100px;
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 6px;
    padding: 10px;
    font-size: 0.8rem;
    font-family: inherit;
    color: var(--text-color);
    resize: none;
    box-sizing: border-box;
  }

  .refactor-textarea:focus {
    outline: none;
    border-color: var(--accent-color);
  }

  .refactor-modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
    padding: 12px 16px;
    border-top: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.05);
  }

  .refactor-exec-btn {
    background: var(--accent-color);
    border-color: var(--accent-color);
    color: var(--bg-primary);
  }
</style>
