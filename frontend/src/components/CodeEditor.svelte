<script>
  import { onMount } from 'svelte';

  export let value = '';
  export let placeholder = '';
  export let highlightedLine = null; // 1-indexed

  let textareaEl;
  let lineNumbersEl;

  $: lines = value ? value.split('\n') : [''];

  function handleScroll() {
    if (textareaEl && lineNumbersEl) {
      lineNumbersEl.scrollTop = textareaEl.scrollTop;
    }
  }

  onMount(() => {
    handleScroll();
  });
</script>

<div class="code-editor-container">
  <div class="line-numbers" bind:this={lineNumbersEl}>
    {#each lines as _, index}
      <div 
        class="line-number" 
        class:error-line={highlightedLine === index + 1}
      >
        {index + 1}
      </div>
    {/each}
  </div>
  <textarea
    bind:this={textareaEl}
    bind:value
    {placeholder}
    on:scroll={handleScroll}
    spellcheck="false"
  />
</div>

<style>
  .code-editor-container {
    position: relative;
    display: flex;
    font-family: 'Consolas', 'Courier New', Courier, monospace;
    font-size: 14px;
    line-height: 1.5;
    background: #121216;
    border: 1px solid #23232c;
    border-radius: 8px;
    height: 100%;
    width: 100%;
    overflow: hidden;
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.5);
  }

  .line-numbers {
    padding: 10px 0;
    width: 48px;
    background: #0b0b0e;
    color: #4c4c5a;
    text-align: right;
    user-select: none;
    border-right: 1px solid #1a1a24;
    overflow: hidden;
    display: flex;
    flex-direction: column;
  }

  .line-number {
    padding-right: 12px;
    min-height: 21px; /* 14px * 1.5 */
    box-sizing: border-box;
    transition: background-color 0.2s, color 0.2s;
  }

  .line-number.error-line {
    background: rgba(239, 68, 68, 0.2);
    color: #ef4444;
    font-weight: bold;
    border-left: 2px solid #ef4444;
    padding-left: 10px;
  }

  textarea {
    flex: 1;
    margin: 0;
    padding: 10px;
    border: none;
    resize: none;
    background: transparent;
    color: #e2e8f0;
    font-family: inherit;
    font-size: inherit;
    line-height: inherit;
    outline: none;
    tab-size: 4;
    white-space: pre;
    overflow-wrap: normal;
    overflow-x: auto;
  }

  textarea::placeholder {
    color: #4a4a5a;
  }
</style>
