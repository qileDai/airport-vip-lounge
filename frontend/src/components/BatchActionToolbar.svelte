<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  export let selectedItems: any[] = [];
  export let availableActions: Array<{ id: string; label: string; icon?: string; variant?: 'primary' | 'secondary' | 'danger'; disabled?: boolean }> = [];
  export let batchProcessing: boolean = false;
  
  const dispatch = createEventDispatcher();
  
  function handleAction(actionId: string) {
    dispatch('action', { actionId, selectedItems });
  }
  
  function selectAll() {
    dispatch('selectAll');
  }
  
  function clearSelection() {
    dispatch('clearSelection');
  }
  
  $: hasSelection = selectedItems.length > 0;
</script>

<div class="batch-action-toolbar">
  <div class="selection-info">
    <div class="checkbox-container">
      <input
        type="checkbox"
        id="select-all"
        checked={hasSelection}
        on:change={(e) => (e.target as HTMLInputElement).checked ? selectAll() : clearSelection()}
      />
      <label for="select-all">全选</label>
    </div>
    
    {#if hasSelection}
      <span class="selection-count">
        已选择 <strong>{selectedItems.length}</strong> 项
      </span>
    {/if}
  </div>
  
  {#if hasSelection}
    <div class="actions">
      {#each availableActions as action}
        <button
          class="btn btn-{action.variant || 'secondary'}"
          disabled={action.disabled || batchProcessing}
          on:click={() => handleAction(action.id)}
        >
          {#if action.icon}
            <span class="icon">{action.icon}</span>
          {/if}
          {action.label}
        </button>
      {/each}
    </div>
  {/if}
  
  {#if batchProcessing}
    <div class="processing-indicator">
      <div class="spinner"></div>
      <span>批量处理中...</span>
    </div>
  {/if}
</div>

<style>
  .batch-action-toolbar {
    background: white;
    border-radius: 8px;
    padding: 12px 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    margin-bottom: 16px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .selection-info {
    display: flex;
    align-items: center;
    gap: 16px;
  }

  .checkbox-container {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .checkbox-container input[type="checkbox"] {
    width: 16px;
    height: 16px;
    cursor: pointer;
  }

  .checkbox-container label {
    font-size: 14px;
    color: #374151;
    cursor: pointer;
  }

  .selection-count {
    font-size: 14px;
    color: #6b7280;
  }

  .selection-count strong {
    color: #3b82f6;
    font-weight: 600;
  }

  .actions {
    display: flex;
    gap: 8px;
  }

  .btn {
    padding: 6px 12px;
    border-radius: 6px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    border: none;
    display: flex;
    align-items: center;
    gap: 4px;
  }

  .btn-primary {
    background: #3b82f6;
    color: white;
  }

  .btn-primary:hover:not(:disabled) {
    background: #2563eb;
  }

  .btn-secondary {
    background: #f3f4f6;
    color: #374151;
  }

  .btn-secondary:hover:not(:disabled) {
    background: #e5e7eb;
  }

  .btn-danger {
    background: #ef4444;
    color: white;
  }

  .btn-danger:hover:not(:disabled) {
    background: #dc2626;
  }

  .btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .icon {
    font-size: 14px;
  }

  .processing-indicator {
    display: flex;
    align-items: center;
    gap: 8px;
    color: #3b82f6;
    font-size: 14px;
  }

  .spinner {
    width: 16px;
    height: 16px;
    border: 2px solid #e5e7eb;
    border-top-color: #3b82f6;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>
