<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  export let data: Array<Record<string, any>> = [];
  export let columns: Array<{ key: string; label: string; type?: string; required?: boolean }> = [];
  export let errors: Array<{ row: number; field: string; message: string }> = [];
  export let isValidating: boolean = false;
  
  const dispatch = createEventDispatcher();
  
  $: hasErrors = errors.length > 0;
  $: validRows = data.length - new Set(errors.map(e => e.row)).size;
  
  function getRowError(row: number): Array<{ field: string; message: string }> {
    return errors.filter(e => e.row === row);
  }
  
  function getFieldValue(row: number, column: string): any {
    return data[row]?.[column];
  }
  
  function handleImport() {
    if (hasErrors) {
      dispatch('importWithErrors', { data, errors });
    } else {
      dispatch('import', { data });
    }
  }
  
  function handleCancel() {
    dispatch('cancel');
  }

  function removeRow(index: number) {
    data = data.filter((_, i) => i !== index);
    errors = errors.filter(e => e.row !== index).map(e => ({
      ...e,
      row: e.row > index ? e.row - 1 : e.row
    }));
  }
</script>

<div class="import-preview-table">
  <div class="table-header">
    <h3 class="title">导入预览</h3>
    <div class="stats">
      <span class="stat-item total">
        总计: <strong>{data.length}</strong> 条
      </span>
      <span class="stat-item valid">
        有效: <strong>{validRows}</strong> 条
      </span>
      {#if hasErrors}
        <span class="stat-item error">
          错误: <strong>{errors.length}</strong> 处
        </span>
      {/if}
    </div>
  </div>
  
  <div class="table-container">
    <table class="preview-table">
      <thead>
        <tr>
          <th class="row-number">#</th>
          {#each columns as column}
            <th class:{required: column.required}>
              {column.label}
              {#if column.required}
                <span class="required-mark">*</span>
              {/if}
            </th>
          {/each}
          <th class="actions">操作</th>
        </tr>
      </thead>
      <tbody>
        {#each data as _, rowIndex (rowIndex)}
          {@const rowErrors = getRowError(rowIndex)}
          <tr class={rowErrors.length > 0 ? 'has-error' : ''}>
            <td class="row-number">{rowIndex + 1}</td>
            {#each columns as column}
              {@const value = getFieldValue(rowIndex, column.key)}
              {@const fieldError = rowErrors.find(e => e.field === column.key)}
              <td class={fieldError ? 'error-cell' : ''}>
                <div class="cell-content">
                  <span class="cell-value">
                    {#if value !== undefined && value !== null && value !== ''}
                      {value}
                    {:else}
                      <span class="empty-value">-</span>
                    {/if}
                  </span>
                  {#if fieldError}
                    <div class="error-tooltip">
                      <span class="error-icon">⚠️</span>
                      <span class="error-message">{fieldError.message}</span>
                    </div>
                  {/if}
                </div>
              </td>
            {/each}
            <td class="actions">
              <button class="remove-btn" on:click={() => removeRow(rowIndex)}>
                🗑️
              </button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  
  <div class="table-footer">
    <div class="error-summary">
      {#if hasErrors}
        <div class="error-list">
          <h4>错误详情:</h4>
          <ul>
            {#each errors as error (error.row + error.field)}
              <li>
                第 <strong>{error.row + 1}</strong> 行 -
                <strong>{columns.find(c => c.key === error.field)?.label || error.field}</strong>:
                {error.message}
              </li>
            {/each}
          </ul>
        </div>
      {/if}
    </div>
    
    <div class="actions">
      <button class="btn btn-secondary" on:click={handleCancel}>
        取消
      </button>
      <button 
        class="btn btn-primary" 
        on:click={handleImport} 
        disabled={isValidating || data.length === 0}
      >
        {#if isValidating}
          <span class="spinner"></span>
          验证中...
        {:else if hasErrors}
          导入有效数据 ({validRows} 条)
        {:else}
          导入全部 ({data.length} 条)
        {/if}
      </button>
    </div>
  </div>
</div>

<style>
  .import-preview-table {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    overflow: hidden;
  }

  .table-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px;
    border-bottom: 1px solid #e5e7eb;
    background: #fafafa;
  }

  .title {
    font-size: 16px;
    font-weight: 600;
    color: #111827;
    margin: 0;
  }

  .stats {
    display: flex;
    gap: 16px;
  }

  .stat-item {
    font-size: 13px;
    color: #6b7280;
  }

  .stat-item strong {
    color: #111827;
  }

  .stat-item.error strong {
    color: #ef4444;
  }

  .stat-item.valid strong {
    color: #10b981;
  }

  .table-container {
    overflow-x: auto;
    max-height: 500px;
    overflow-y: auto;
  }

  .preview-table {
    width: 100%;
    border-collapse: collapse;
    font-size: 13px;
  }

  .preview-table th {
    background: #f9fafb;
    padding: 12px;
    text-align: left;
    font-weight: 600;
    color: #374151;
    border-bottom: 2px solid #e5e7eb;
    position: sticky;
    top: 0;
    z-index: 1;
  }

  .preview-table th.required::after {
    content: '';
  }

  .required-mark {
    color: #ef4444;
    margin-left: 2px;
  }

  .preview-table td {
    padding: 10px 12px;
    border-bottom: 1px solid #f3f4f6;
    vertical-align: top;
  }

  .preview-table tr:hover {
    background: #f9fafb;
  }

  .preview-table tr.has-error {
    background: #fef2f2;
  }

  .row-number {
    width: 50px;
    font-weight: 600;
    color: #6b7280;
  }

  .cell-content {
    position: relative;
  }

  .cell-value {
    display: block;
    max-width: 200px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .empty-value {
    color: #d1d5db;
  }

  .error-cell {
    background: #fee2e2;
  }

  .error-tooltip {
    display: flex;
    align-items: center;
    gap: 4px;
    margin-top: 4px;
    font-size: 11px;
    color: #dc2626;
  }

  .error-icon {
    font-size: 12px;
  }

  .actions {
    width: 60px;
    text-align: center;
  }

  .remove-btn {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 14px;
    padding: 4px;
    opacity: 0.5;
    transition: opacity 0.2s;
  }

  .remove-btn:hover {
    opacity: 1;
  }

  .table-footer {
    padding: 16px;
    border-top: 1px solid #e5e7eb;
    background: #fafafa;
  }

  .error-summary {
    margin-bottom: 16px;
  }

  .error-list h4 {
    font-size: 13px;
    font-weight: 600;
    color: #dc2626;
    margin: 0 0 8px;
  }

  .error-list ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  .error-list li {
    font-size: 12px;
    color: #991b1b;
    padding: 4px 0;
    line-height: 1.4;
  }

  .actions {
    display: flex;
    gap: 8px;
    justify-content: flex-end;
  }

  .btn {
    padding: 8px 16px;
    border-radius: 6px;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    border: none;
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .btn-primary {
    background: #3b82f6;
    color: white;
  }

  .btn-primary:hover:not(:disabled) {
    background: #2563eb;
  }

  .btn-secondary {
    background: white;
    color: #374151;
    border: 1px solid #d1d5db;
  }

  .btn-secondary:hover {
    background: #f3f4f6;
  }

  .btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .spinner {
    width: 14px;
    height: 14px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: white;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>
