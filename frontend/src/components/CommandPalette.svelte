<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  export let isOpen: boolean = false;
  export let commands: Array<{
    id: string;
    title: string;
    category: string;
    shortcut?: string;
    icon?: string;
    action: () => void;
  }> = [];
  export let query: string = '';
  
  const dispatch = createEventDispatcher();
  
  $: filteredCommands = commands.filter(cmd =>
    cmd.title.toLowerCase().includes(query.toLowerCase()) ||
    cmd.category.toLowerCase().includes(query.toLowerCase())
  );
  
  $: groupedCommands = filteredCommands.reduce((groups, cmd) => {
    const category = cmd.category;
    if (!groups[category]) {
      groups[category] = [];
    }
    groups[category].push(cmd);
    return groups;
  }, {} as Record<string, typeof commands>);
  
  function handleSelect(command: (typeof commands)[0]) {
    command.action();
    dispatch('close');
    query = '';
  }
  
  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      dispatch('close');
      query = '';
    } else if (e.key === 'Enter' && filteredCommands.length > 0) {
      handleSelect(filteredCommands[0]);
    }
  }
</script>

{#if isOpen}
  <div class="command-palette-overlay" on:click={() => dispatch('close')} role="dialog" aria-modal="true">
    <div class="command-palette" on:click|stopPropagation on:keydown={handleKeydown}>
      <div class="search-container">
        <span class="search-icon">🔍</span>
        <input
          type="text"
          class="search-input"
          placeholder="输入命令搜索... (Esc 关闭)"
          bind:value={query}
          autofocus
        />
        {#if query}
          <button class="clear-btn" on:click={() => query = ''}>✕</button>
        {/if}
      </div>
      
      <div class="commands-list">
        {#each Object.entries(groupedCommands) as [category, cmds]}
          <div class="command-group">
            <div class="group-label">{category}</div>
            {#each cmds as cmd}
              <button
                class="command-item"
                on:click={() => handleSelect(cmd)}
              >
                <span class="command-icon">{cmd.icon || '📋'}</span>
                <span class="command-title">{cmd.title}</span>
                {#if cmd.shortcut}
                  <span class="shortcut">{cmd.shortcut}</span>
                {/if}
              </button>
            {/each}
          </div>
        {:else}
          <div class="no-results">
            <span class="no-results-icon">🔍</span>
            <p>未找到匹配的命令</p>
            <p class="hint">尝试其他关键词</p>
          </div>
        {/each}
      </div>
      
      <div class="footer">
        <span class="footer-hint">
          <kbd>↑↓</kbd> 导航
          <kbd>Enter</kbd> 选择
          <kbd>Esc</kbd> 关闭
        </span>
      </div>
    </div>
  </div>
{/if}

<style>
  .command-palette-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    padding-top: 15vh;
    z-index: 1000;
    animation: fadeIn 0.2s ease-out;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  .command-palette {
    background: white;
    border-radius: 12px;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
    width: 90%;
    max-width: 600px;
    max-height: 70vh;
    display: flex;
    flex-direction: column;
    animation: slideUp 0.2s ease-out;
  }

  @keyframes slideUp {
    from {
      transform: translateY(20px);
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }

  .search-container {
    display: flex;
    align-items: center;
    padding: 16px;
    border-bottom: 1px solid #e5e7eb;
    gap: 12px;
  }

  .search-icon {
    font-size: 18px;
    color: #9ca3af;
  }

  .search-input {
    flex: 1;
    border: none;
    font-size: 16px;
    outline: none;
    color: #111827;
  }

  .search-input::placeholder {
    color: #9ca3af;
  }

  .clear-btn {
    background: none;
    border: none;
    cursor: pointer;
    color: #9ca3af;
    font-size: 14px;
    padding: 4px;
  }

  .clear-btn:hover {
    color: #374151;
  }

  .commands-list {
    flex: 1;
    overflow-y: auto;
    padding: 8px;
  }

  .command-group {
    margin-bottom: 16px;
  }

  .group-label {
    font-size: 11px;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    color: #6b7280;
    padding: 8px 12px 4px;
  }

  .command-item {
    display: flex;
    align-items: center;
    gap: 12px;
    width: 100%;
    padding: 10px 12px;
    border: none;
    background: transparent;
    border-radius: 6px;
    cursor: pointer;
    transition: all 0.15s;
    text-align: left;
  }

  .command-item:hover {
    background: #f3f4f6;
  }

  .command-icon {
    font-size: 18px;
  }

  .command-title {
    flex: 1;
    font-size: 14px;
    color: #111827;
  }

  .shortcut {
    font-size: 11px;
    color: #6b7280;
    background: #f3f4f6;
    padding: 2px 6px;
    border-radius: 4px;
    font-family: monospace;
  }

  .no-results {
    text-align: center;
    padding: 32px;
    color: #6b7280;
  }

  .no-results-icon {
    font-size: 48px;
    margin-bottom: 8px;
  }

  .hint {
    font-size: 13px;
    color: #9ca3af;
    margin-top: 4px;
  }

  .footer {
    padding: 12px 16px;
    border-top: 1px solid #e5e7eb;
    display: flex;
    justify-content: center;
  }

  .footer-hint {
    display: flex;
    gap: 16px;
    font-size: 12px;
    color: #9ca3af;
  }

  kbd {
    background: #f3f4f6;
    padding: 2px 6px;
    border-radius: 4px;
    font-family: monospace;
    font-size: 11px;
    border: 1px solid #d1d5db;
  }
</style>
