<script lang="ts">
  import { onMount } from 'svelte';
  import apiService from '../services/api';
  import type { any } from '../types';
  
  export let entityType: string = '';
  export let autoRefresh: boolean = true;
  export let refreshInterval: number = 30000;
  export let maxItems: number = 50;
  
  let logs: Array<{
    id: number;
    action: string;
    entity_type: string;
    entity_id: number;
    user_id: string;
    details: string;
    ip_address: string;
    created_at: string;
  }> = [];
  
  let loading: boolean = false;
  let error: string | null = null;
  let refreshTimer: ReturnType<typeof setInterval>;
  
  async function fetchLogs() {
    loading = true;
    error = null;
    
    try {
      const response = await apiService.getAuditLogs(entityType, 1, maxItems);
      if (response.success && response.data) {
        logs = response.data as typeof logs;
      }
    } catch (err) {
      error = (err as Error).message;
    } finally {
      loading = false;
    }
  }
  
  function getActionIcon(action: string): string {
    const iconMap: Record<string, string> = {
      CREATE: '➕',
      UPDATE: '✏️',
      DELETE: '🗑️',
      STATUS_TRANSITION: '🔄',
      LOGIN: '🔐',
      EXPORT: '📤',
      IMPORT: '📥',
      BATCH_CREATE: '📋',
      ADMIT: '✅',
    };
    return iconMap[action] || '📝';
  }
  
  function getActionColor(action: string): string {
    const colorMap: Record<string, string> = {
      CREATE: '#10b981',
      UPDATE: '#3b82f6',
      DELETE: '#ef4444',
      STATUS_TRANSITION: '#f59e0b',
      LOGIN: '#8b5cf6',
      EXPORT: '#06b6d4',
      IMPORT: '#ec4899',
    };
    return colorMap[action] || '#6b7280';
  }
  
  function formatTime(timeStr: string): string {
    const date = new Date(timeStr);
    const now = new Date();
    const diff = now.getTime() - date.getTime();
    
    if (diff < 60000) return '刚刚';
    if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`;
    if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`;
    return date.toLocaleDateString('zh-CN');
  }
  
  onMount(() => {
    fetchLogs();
    
    if (autoRefresh) {
      refreshTimer = setInterval(fetchLogs, refreshInterval);
    }
    
    return () => {
      if (refreshTimer) {
        clearInterval(refreshTimer);
      }
    };
  });
</script>

<div class="audit-log-stream">
  <div class="stream-header">
    <h3 class="stream-title">审计日志流</h3>
    {#if entityType}
      <span class="entity-badge">{entityType}</span>
    {/if}
    <button class="refresh-btn" on:click={fetchLogs} disabled={loading}>
      {#if loading}
        <span class="spinner"></span>
      {:else}
        🔄
      {/if}
      刷新
    </button>
  </div>
  
  <div class="log-container">
    {#if error}
      <div class="error-state">
        <span class="error-icon">⚠️</span>
        <p>加载失败: {error}</p>
        <button on:click={fetchLogs}>重试</button>
      </div>
    {:else if logs.length === 0 && !loading}
      <div class="empty-state">
        <span class="empty-icon">📋</span>
        <p>暂无日志记录</p>
      </div>
    {:else}
      <div class="log-list">
        {#each logs as log, i (log.id)}
          <div class="log-item" style="--action-color: {getActionColor(log.action)}">
            <div class="log-icon" style="background: var(--action-color)">
              {getActionIcon(log.action)}
            </div>
            <div class="log-content">
              <div class="log-header">
                <span class="log-action">{log.action}</span>
                <span class="log-time">{formatTime(log.created_at)}</span>
              </div>
              <div class="log-details">
                {#if log.entity_type}
                  <span class="log-entity">{log.entity_type}#{log.entity_id}</span>
                {/if}
                <span class="log-user">操作人: {log.user_id}</span>
              </div>
              {#if log.details}
                <div class="log-message">{log.details}</div>
              {/if}
            </div>
          </div>
        {/each}
        
        {#if loading}
          <div class="loading-more">
            <span class="spinner"></span>
            加载更多...
          </div>
        {/if}
      </div>
    {/if}
  </div>
</div>

<style>
  .audit-log-stream {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    overflow: hidden;
  }

  .stream-header {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 16px;
    border-bottom: 1px solid #e5e7eb;
    background: #fafafa;
  }

  .stream-title {
    font-size: 16px;
    font-weight: 600;
    color: #111827;
    margin: 0;
    flex: 1;
  }

  .entity-badge {
    background: #dbeafe;
    color: #1e40af;
    padding: 2px 8px;
    border-radius: 12px;
    font-size: 11px;
    font-weight: 500;
  }

  .refresh-btn {
    background: white;
    border: 1px solid #d1d5db;
    padding: 4px 10px;
    border-radius: 6px;
    font-size: 13px;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 4px;
    transition: all 0.2s;
  }

  .refresh-btn:hover:not(:disabled) {
    background: #f3f4f6;
    border-color: #9ca3af;
  }

  .refresh-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .log-container {
    max-height: 600px;
    overflow-y: auto;
  }

  .error-state,
  .empty-state {
    text-align: center;
    padding: 40px;
    color: #6b7280;
  }

  .error-icon,
  .empty-icon {
    font-size: 40px;
    margin-bottom: 8px;
  }

  .log-list {
    padding: 8px;
  }

  .log-item {
    display: flex;
    gap: 12px;
    padding: 12px;
    border-radius: 6px;
    transition: background 0.2s;
  }

  .log-item:hover {
    background: #f9fafb;
  }

  .log-icon {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    flex-shrink: 0;
  }

  .log-content {
    flex: 1;
    min-width: 0;
  }

  .log-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 4px;
  }

  .log-action {
    font-weight: 600;
    font-size: 13px;
    color: #111827;
  }

  .log-time {
    font-size: 11px;
    color: #9ca3af;
  }

  .log-details {
    display: flex;
    gap: 12px;
    font-size: 12px;
    color: #6b7280;
    margin-bottom: 4px;
  }

  .log-entity {
    background: #f3f4f6;
    padding: 1px 6px;
    border-radius: 4px;
    font-family: monospace;
  }

  .log-message {
    font-size: 13px;
    color: #374151;
    line-height: 1.4;
    word-break: break-word;
  }

  .loading-more {
    text-align: center;
    padding: 16px;
    color: #6b7280;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
  }

  .spinner {
    width: 14px;
    height: 14px;
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
