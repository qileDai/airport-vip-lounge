<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  export let filters: Array<{ key: string; label: string; type: 'text' | 'select' | 'date' | 'number'; options?: Array<{ value: string; label: string }> }> = [];
  export let activeFilters: Record<string, any> = {};
  
  const dispatch = createEventDispatcher();
  
  function handleFilterChange(key: string, value: any) {
    activeFilters = { ...activeFilters, [key]: value };
    dispatch('filterChange', activeFilters);
  }
  
  function clearFilters() {
    activeFilters = {};
    dispatch('filterChange', {});
  }
  
  function applyFilters() {
    dispatch('applyFilters', activeFilters);
  }
</script>

<div class="advanced-filter-bar">
  <div class="filter-container">
    {#each filters as filter}
      <div class="filter-item">
        <label class="filter-label">{filter.label}</label>
        
        {#if filter.type === 'text'}
          <input
            type="text"
            class="filter-input"
            placeholder={`搜索${filter.label}...`}
            value={activeFilters[filter.key] || ''}
            on:change={(e) => handleFilterChange(filter.key, (e.target as HTMLInputElement).value)}
          />
        {:else if filter.type === 'select'}
          <select
            class="filter-select"
            value={activeFilters[filter.key] || ''}
            on:change={(e) => handleFilterChange(filter.key, (e.target as HTMLSelectElement).value)}
          >
            <option value="">全部</option>
            {#each filter.options || [] as option}
              <option value={option.value}>{option.label}</option>
            {/each}
          </select>
        {:else if filter.type === 'date'}
          <input
            type="date"
            class="filter-input"
            value={activeFilters[filter.key] || ''}
            on:change={(e) => handleFilterChange(filter.key, (e.target as HTMLInputElement).value)
          />
        {:else if filter.type === 'number'}
          <input
            type="number"
            class="filter-input"
            placeholder="输入数值"
            value={activeFilters[filter.key] || ''}
            on:change={(e) => handleFilterChange(filter.key, Number((e.target as HTMLInputElement).value))}
          />
        {/if}
      </div>
    {/each}
    
    <div class="filter-actions">
      <button class="btn btn-primary" on:click={applyFilters}>
        应用筛选
      </button>
      <button class="btn btn-secondary" on:click={clearFilters}>
        清除筛选
      </button>
    </div>
  </div>
</div>

<style>
  .advanced-filter-bar {
    background: white;
    border-radius: 8px;
    padding: 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    margin-bottom: 16px;
  }

  .filter-container {
    display: flex;
    flex-wrap: wrap;
    gap: 16px;
    align-items: flex-end;
  }

  .filter-item {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 180px;
  }

  .filter-label {
    font-size: 12px;
    font-weight: 600;
    color: #374151;
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }

  .filter-input,
  .filter-select {
    padding: 8px 12px;
    border: 1px solid #d1d5db;
    border-radius: 6px;
    font-size: 14px;
    transition: all 0.2s;
  }

  .filter-input:focus,
  .filter-select:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }

  .filter-actions {
    display: flex;
    gap: 8px;
    margin-left: auto;
  }

  .btn {
    padding: 8px 16px;
    border-radius: 6px;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    border: none;
  }

  .btn-primary {
    background: #3b82f6;
    color: white;
  }

  .btn-primary:hover {
    background: #2563eb;
  }

  .btn-secondary {
    background: #f3f4f6;
    color: #374151;
  }

  .btn-secondary:hover {
    background: #e5e7eb;
  }
</style>
