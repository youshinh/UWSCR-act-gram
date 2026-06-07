<script>
  import { onMount, onDestroy } from 'svelte';

  export let targetId = '';

  let boxStyle = '';
  let rafId;
  let visible = false;

  function updatePosition() {
    if (!targetId) return;
    const el = document.getElementById(targetId);
    if (el) {
      const rect = el.getBoundingClientRect();
      const offsetX = -2;  // 横方向のズレ補正：マイナス値で「左」に移動
      const offsetY = -2;  // 縦方向のズレ補正：マイナス値で「上」に移動
      const padding = 4;   // アイコンの周りの見やすさ用の余白（お好みで）
      boxStyle = `
        left: ${rect.left - padding + offsetX}px;
        top: ${rect.top - padding + offsetY}px;
        width: ${rect.width + padding * 2}px;
        height: ${rect.height + padding * 2}px;
      `;
      visible = true;
    }
    rafId = requestAnimationFrame(updatePosition);
  }

  onMount(() => {
    rafId = requestAnimationFrame(updatePosition);
  });

  onDestroy(() => {
    if (rafId) cancelAnimationFrame(rafId);
  });
</script>

{#if visible}
  <div class="tour-highlight-box" style={boxStyle}></div>
{/if}

<style>
  .tour-highlight-box {
    position: fixed;
    border: 2px solid #38bdf8;
    border-radius: 10px;
    box-shadow: 0 0 0 4px rgba(56, 189, 248, 0.2), 0 0 20px rgba(56, 189, 248, 0.4);
    pointer-events: none;
    z-index: 9999;
    animation: tour-pulse 1.5s ease-in-out infinite;
    transition: all 0.15s ease-out;
  }

  @keyframes tour-pulse {
    0%, 100% {
      box-shadow: 0 0 0 4px rgba(56, 189, 248, 0.2), 0 0 20px rgba(56, 189, 248, 0.4);
    }
    50% {
      box-shadow: 0 0 0 8px rgba(56, 189, 248, 0.1), 0 0 30px rgba(56, 189, 248, 0.6);
    }
  }
</style>
