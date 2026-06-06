<script>
  import { onMount, onDestroy } from 'svelte';
  import { StartRecording, StopRecording } from '../../wailsjs/go/main/App.js';
  import * as wailsRuntime from '../../wailsjs/runtime/runtime.js';

  export let onRecordFinished = (logDir) => {};

  let isRecording = false;
  let statusMessage = "待機中";
  let elapsedSeconds = 0;
  let timerInterval = null;
  let errorMessage = "";

  onMount(() => {
    // F8キー押下などによる録画停止イベントのハンドラ
    wailsRuntime.EventsOn("recording_stopped", (logDir) => {
      console.log("[RecSimulator] recording_stopped event received:", logDir);
      stopTimer();
      isRecording = false;
      statusMessage = "記録終了";
      onRecordFinished(logDir);
    });
  });

  onDestroy(() => {
    stopTimer();
    wailsRuntime.EventsOff("recording_stopped");
  });

  function startTimer() {
    elapsedSeconds = 0;
    timerInterval = setInterval(() => {
      elapsedSeconds++;
    }, 1000);
  }

  function stopTimer() {
    if (timerInterval) {
      clearInterval(timerInterval);
      timerInterval = null;
    }
  }

  function formatTime(sec) {
    const m = Math.floor(sec / 60).toString().padStart(2, '0');
    const s = (sec % 60).toString().padStart(2, '0');
    return `${m}:${s}`;
  }

  async function handleStart() {
    errorMessage = "";
    statusMessage = "起動中...";
    try {
      isRecording = true;
      await StartRecording();
      startTimer();
      statusMessage = "記録中";
    } catch (err) {
      isRecording = false;
      errorMessage = err.message || err;
      statusMessage = "起動失敗";
    }
  }

  async function handleStop() {
    errorMessage = "";
    statusMessage = "停止処理中...";
    try {
      const logDir = await StopRecording();
      stopTimer();
      isRecording = false;
      statusMessage = "記録終了";
      onRecordFinished(logDir);
    } catch (err) {
      errorMessage = err.message || err;
      statusMessage = "停止失敗";
    }
  }
</script>

<div class="rec-container">
  <div class="rec-card">
    <div class="rec-header">
      <div class="rec-dot" class:active={isRecording}></div>
      <h2>操作レコーダー</h2>
    </div>

    <p class="rec-desc">
      開始ボタンを押すとアプリが自動で最小化され、デスクトップ上でのすべてのマウスクリック、キー入力、および対象のUI要素（トグル状態や選択値など）のタイムスタンプ付き記録を開始します。
    </p>

    <div class="rec-status-area">
      {#if isRecording}
        <div class="recording-ui">
          <span class="timer">{formatTime(elapsedSeconds)}</span>
          <span class="pulse-text">RECORDING</span>
          <p class="stop-guide">キーボードの <strong>F8</strong> キーを押すか、この画面で停止ボタンを押すと記録を終了します。</p>
        </div>
      {:else}
        <div class="idle-ui">
          <span class="timer">00:00</span>
          <span class="idle-text">READY</span>
        </div>
      {/if}
    </div>

    {#if errorMessage}
      <div class="error-box">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="error-icon">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        <span>{errorMessage}</span>
      </div>
    {/if}

    <div class="rec-actions">
      {#if !isRecording}
        <button class="btn-rec-start" on:click={handleStart}>
          <svg viewBox="0 0 24 24" fill="currentColor" class="btn-icon">
            <circle cx="12" cy="12" r="8"/>
          </svg>
          記録を開始 (最小化)
        </button>
      {:else}
        <button class="btn-rec-stop" on:click={handleStop}>
          <svg viewBox="0 0 24 24" fill="currentColor" class="btn-icon">
            <rect x="6" y="6" width="12" height="12" rx="1"/>
          </svg>
          記録を停止 (復元)
        </button>
      {/if}
    </div>
  </div>
</div>

<style>
  .rec-container {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    background: var(--bg-primary);
  }

  .rec-card {
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    padding: 32px;
    max-width: 480px;
    width: 100%;
    box-shadow: var(--shadow-md);
    display: flex;
    flex-direction: column;
    gap: 20px;
    text-align: center;
  }

  .rec-header {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
  }

  .rec-dot {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    background-color: var(--text-secondary);
    transition: background-color 0.3s ease;
  }

  .rec-dot.active {
    background-color: #ef4444;
    box-shadow: 0 0 10px #ef4444;
    animation: pulse 1.5s infinite;
  }

  .rec-header h2 {
    margin: 0;
    font-size: 1.2rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .rec-desc {
    font-size: 0.8rem;
    line-height: 1.5;
    color: var(--text-secondary);
    margin: 0;
  }

  .rec-status-area {
    background: var(--input-bg);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 24px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 120px;
  }

  .timer {
    font-family: Consolas, monospace;
    font-size: 2.2rem;
    font-weight: 700;
    color: var(--text-primary);
    line-height: 1;
    letter-spacing: -1px;
  }

  .pulse-text {
    font-size: 0.7rem;
    font-weight: 700;
    color: #ef4444;
    letter-spacing: 1.5px;
    margin-top: 8px;
    display: inline-block;
    animation: pulse 1.5s infinite;
  }

  .idle-text {
    font-size: 0.7rem;
    font-weight: 700;
    color: var(--text-secondary);
    letter-spacing: 1.5px;
    margin-top: 8px;
    display: inline-block;
  }

  .stop-guide {
    font-size: 0.7rem;
    color: var(--text-secondary);
    margin: 12px 0 0 0;
    line-height: 1.4;
  }

  .error-box {
    display: flex;
    align-items: center;
    gap: 8px;
    background: rgba(239, 68, 68, 0.08);
    border: 1px solid rgba(239, 68, 68, 0.2);
    border-radius: 6px;
    padding: 10px 14px;
    color: #ef4444;
    font-size: 0.75rem;
    text-align: left;
  }

  .error-icon {
    width: 16px;
    height: 16px;
    flex-shrink: 0;
  }

  .rec-actions {
    display: flex;
    justify-content: center;
  }

  .btn-rec-start {
    background: var(--accent-color);
    color: var(--bg-primary);
    border: 1px solid var(--accent-color);
    font-weight: 600;
    font-size: 0.85rem;
    border-radius: 24px;
    padding: 12px 28px;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 8px;
    transition: all 0.2s ease;
  }

  .btn-rec-start:hover {
    background: var(--accent-hover);
    transform: translateY(-1px);
  }

  .btn-rec-stop {
    background: #ef4444;
    color: #ffffff;
    border: 1px solid #ef4444;
    font-weight: 600;
    font-size: 0.85rem;
    border-radius: 24px;
    padding: 12px 28px;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 8px;
    transition: all 0.2s ease;
  }

  .btn-rec-stop:hover {
    background: #dc2626;
    transform: translateY(-1px);
  }

  .btn-icon {
    width: 16px;
    height: 16px;
  }

  @keyframes pulse {
    0% { opacity: 0.5; }
    50% { opacity: 1; }
    100% { opacity: 0.5; }
  }
</style>
