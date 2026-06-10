<script lang="ts">
  export let events: Array<{
    id: string;
    timestamp: string;
    title: string;
    description?: string;
    type: 'status_change' | 'action' | 'system' | 'error' | 'success' | 'info';
    icon?: string;
    user?: string;
  }> = [];
  
  function getEventIcon(type: string): string {
    const iconMap: Record<string, string> = {
      status_change: '🔄',
      action: '⚡',
      system: '🖥️',
      error: '❌',
      success: '✅',
      info: 'ℹ️',
    };
    return iconMap[type] || '📌';
  }
  
  function getEventColor(type: string): string {
    const colorMap: Record<string, string> = {
      status_change: '#3b82f6',
      action: '#8b5cf6',
      system: '#6b7280',
      error: '#ef4444',
      success: '#10b981',
      info: '#06b6d4',
    };
    return colorMap[type] || '#9ca3af';
  }
  
  function formatTimestamp(timestamp: string): string {
    const date = new Date(timestamp);
    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    });
  }
  
  function getTimeAgo(timestamp: string): string {
    const date = new Date(timestamp);
    const now = new Date();
    const diff = now.getTime() - date.getTime();
    
    if (diff < 60000) return '刚刚';
    if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`;
    if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`;
    if (diff < 604800000) return `${Math.floor(diff / 86400000)}天前`;
    return formatTimestamp(timestamp);
  }
</script>

<div class="timeline-panel">
  {#if events.length === 0}
    <div class="empty-state">
      <span class="empty-icon">📅</span>
      <p>暂无时间线记录</p>
    </div>
  {:else}
    <div class="timeline">
      {#each events as event, i (event.id)}
        <div class="timeline-item" style="--event-color: {getEventColor(event.type)}">
          <div class="timeline-marker">
            <div class="marker-icon" style="background: var(--event-color)">
              {event.icon || getEventIcon(event.type)}
            </div>
            {#if i < events.length - 1}
              <div class="marker-line" style="background: var(--event-color)"></div>
            {/if}
          </div>
          
          <div class="timeline-content">
            <div class="event-header">
              <h4 class="event-title">{event.title}</h4>
              <span class="event-time" title={formatTimestamp(event.timestamp)}>
                {getTimeAgo(event.timestamp)}
              </span>
            </div>
            
            {#if event.description}
              <p class="event-description">{event.description}</p>
            {/if}
            
            {#if event.user}
              <div class="event-meta">
                <span class="user-badge">
                  👤 {event.user}
                </span>
              </div>
            {/if}
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  .timeline-panel {
    background: white;
    border-radius: 8px;
    padding: 20px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
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

  .timeline {
    position: relative;
  }

  .timeline-item {
    display: flex;
    gap: 16px;
    position: relative;
    padding-bottom: 24px;
    
    &:last-child {
      padding-bottom: 0;
    }
  }

  .timeline-marker {
    display: flex;
    flex-direction: column;
    align-items: center;
    flex-shrink: 0;
  }

  .marker-icon {
    width: 36px;
    height: 36px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 16px;
    z-index: 1;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  }

  .marker-line {
    width: 2px;
    flex: 1;
    margin-top: 4px;
    opacity: 0.3;
  }

  .timeline-content {
    flex: 1;
    min-width: 0;
    background: #f9fafb;
    border-radius: 8px;
    padding: 12px 16px;
    transition: all 0.2s;
  }

  .timeline-content:hover {
    background: #f3f4f6;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  }

  .event-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 12px;
    margin-bottom: 6px;
  }

  .event-title {
    font-size: 14px;
    font-weight: 600;
    color: #111827;
    margin: 0;
    line-height: 1.4;
  }

  .event-time {
    font-size: 11px;
    color: #9ca3af;
    white-space: nowrap;
    cursor: help;
  }

  .event-description {
    font-size: 13px;
    color: #6b7280;
    line-height: 1.5;
    margin: 0 0 8px;
  }

  .event-meta {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
  }

  .user-badge {
    display: inline-flex;
    align-items: center;
    gap: 4px;
    font-size: 11px;
    color: #374151;
    background: white;
    padding: 2px 8px;
    border-radius: 10px;
    border: 1px solid #e5e7eb;
  }
</style>
