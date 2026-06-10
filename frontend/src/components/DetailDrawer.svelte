<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  export let isOpen: boolean = false;
  export let title: string = '详情';
  export let data: any = null;
  export let width: string = '480px';
  
  const dispatch = createEventDispatcher();
  
  function close() {
    dispatch('close');
  }
  
  function handleOverlayClick(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      close();
    }
  }
</script>

{#if isOpen}
  <div class="drawer-overlay" on:click={handleOverlayClick} role="dialog" aria-modal="true">
    <div class="drawer" style="width: {width}">
      <div class="drawer-header">
        <h2 class="drawer-title">{title}</h2>
        <button class="close-btn" on:click={close} aria-label="关闭">
          ✕
        </button>
      </div>
      
      <div class="drawer-content">
        {#if data}
          <slot {data} />
        {:else}
          <div class="empty-state">
            <span class="empty-icon">📋</span>
            <p>暂无数据</p>
          </div>
        {/if}
      </div>
      
      <div class="drawer-footer">
        <slot name="footer" />
      </div>
    </div>
  </div>
{/if}

<style>
  .drawer-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    z-index: 999;
    display: flex;
    justify-content: flex-end;
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

  .drawer {
    background: white;
    height: 100%;
    box-shadow: -4px 0 20px rgba(0, 0, 0, 0.15);
    display: flex;
    flex-direction: column;
    animation: slideInRight 0.3s ease-out;
  }

  @keyframes slideInRight {
    from {
      transform: translateX(100%);
    }
    to {
      transform: translateX(0);
    }
  }

  .drawer-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 24px;
    border-bottom: 1px solid #e5e7eb;
    background: #fafafa;
  }

  .drawer-title {
    font-size: 18px;
    font-weight: 600;
    color: #111827;
    margin: 0;
  }

  .close-btn {
    background: none;
    border: none;
    font-size: 20px;
    cursor: pointer;
    color: #6b7280;
    padding: 4px 8px;
    border-radius: 4px;
    transition: all 0.2s;
  }

  .close-btn:hover {
    background: #e5e7eb;
    color: #111827;
  }

  .drawer-content {
    flex: 1;
    overflow-y: auto;
    padding: 24px;
  }

  .empty-state {
    text-align: center;
    padding: 40px;
    color: #9ca3af;
  }

  .empty-icon {
    font-size: 48px;
    margin-bottom: 12px;
  }

  .drawer-footer {
    padding: 16px 24px;
    border-top: 1px solid #e5e7eb;
    background: #fafafa;
    display: flex;
    gap: 8px;
    justify-content: flex-end;
  }
</style>
