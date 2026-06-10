<script lang="ts">
  import { onMount } from 'svelte';
  import AuditLogStream from './components/AuditLogStream.svelte';
  import Dashboard from './pages/Dashboard.svelte';
  import MemberBenefits from './pages/MemberBenefits.svelte';
  import { uiStore } from './stores';

  let activeTab = 'dashboard';

  const menuItems = [
    { id: 'dashboard', label: '统计驾驶舱', icon: '📊' },
    { id: 'member-benefits', label: '会员权益', icon: '💳' },
    { id: 'appointments', label: '预约记录', icon: '📅' },
    { id: 'flight-slots', label: '航班时段', icon: '✈️' },
    { id: 'companions', label: '同行人管理', icon: '👥' },
    { id: 'waiting-list', label: '候补名单', icon: '⏳' },
    { id: 'exceptions', label: '异常事件', icon: '⚠️' },
    { id: 'audit-log', label: '审计日志', icon: '📋' }
  ];

  function handleMenuClick(tabId: string) {
    activeTab = tabId;
    uiStore.setActiveTab(tabId);
  }

  onMount(() => {
    document.addEventListener('keydown', (e) => {
      if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
        e.preventDefault();
        uiStore.toggleCommandPalette();
      }
    });
  });
</script>

<div class="app-layout">
  <aside class="sidebar" class:collapsed={$uiStore.sidebarCollapsed}>
    <div class="sidebar-header">
      <div class="logo">
        <span class="logo-icon">✈️</span>
        <span class="logo-text">贵宾厅协同控制台</span>
      </div>
      <button 
        class="toggle-btn"
        on:click={() => uiStore.toggleSidebar()}
        title="折叠侧边栏"
      >
        ☰
      </button>
    </div>

    <nav class="nav-menu">
      {#each menuItems as item}
        <button
          class="nav-item"
          class:active={activeTab === item.id}
          on:click={() => handleMenuClick(item.id)}
        >
          <span class="nav-icon">{item.icon}</span>
          <span class="nav-label">{item.label}</span>
        </button>
      {/each}
    </nav>

    <div class="sidebar-footer">
      <div class="version-info">v1.0.0</div>
    </div>
  </aside>

  <main class="main-content">
    <header class="top-bar">
      <div class="top-bar-left">
        <button 
          class="menu-toggle"
          on:click={() => uiStore.toggleSidebar()}
        >
          ☰
        </button>
        
        <div class="search-box">
          <input 
            type="text" 
            placeholder="搜索... (Ctrl+K)" 
            readonly
            on:click={() => uiStore.toggleCommandPalette()}
          />
          <kbd>Ctrl+K</kbd>
        </div>
      </div>

      <div class="top-bar-right">
        <div class="notification-area">
          {#if $uiStore.notifications.length > 0}
            <span class="notification-badge">{$uiStore.notifications.length}</span>
          {/if}
          🔔
        </div>
        
        <div class="user-info">
          <div class="user-avatar">管</div>
          <span class="user-name">管理员</span>
        </div>
      </div>
    </header>

    <div class="content-area">
      {#if activeTab === 'dashboard'}
        <svelte:component this={Dashboard} />
      {:else if activeTab === 'member-benefits'}
        <svelte:component this={MemberBenefits} />
      {:else if activeTab === 'appointments'}
        <div class="coming-soon">
          <span class="coming-icon">🚧</span>
          <h2>功能开发中</h2>
          <p>预约记录模块正在开发中，敬请期待...</p>
        </div>
      {:else if activeTab === 'flight-slots'}
        <div class="coming-soon">
          <span class="coming-icon">🚧</span>
          <h2>功能开发中</h2>
          <p>航班时段模块正在开发中，敬请期待...</p>
        </div>
      {:else if activeTab === 'companions'}
        <div class="coming-soon">
          <span class="coming-icon">🚧</span>
          <h2>功能开发中</h2>
          <p>同行人管理模块正在开发中，敬请期待...</p>
        </div>
      {:else if activeTab === 'waiting-list'}
        <div class="coming-soon">
          <span class="coming-icon">🚧</span>
          <h2>功能开发中</h2>
          <p>候补名单模块正在开发中，敬请期待...</p>
        </div>
      {:else if activeTab === 'exceptions'}
        <div class="coming-soon">
          <span class="coming-icon">🚧</span>
          <h2>功能开发中</h2>
          <p>异常事件模块正在开发中，敬请期待...</p>
        </div>
      {:else if activeTab === 'audit-log'}
        <svelte:component this={AuditLogStream} entityType="" />
      {:else}
        <div class="coming-soon">
          <span class="coming-icon">🔍</span>
          <h2>页面未找到</h2>
          <p>请从侧边栏选择一个有效的菜单项</p>
        </div>
      {/if}

      <!-- Notification Toast Container -->
      {#each $uiStore.notifications as notification (notification.id)}
        <div class="toast notification-{notification.type}">
          <span class="toast-icon">
            {#if notification.type === 'success'} ✅
            {:else if notification.type === 'error'} ❌
            {:else if notification.type === 'warning'} ⚠️
            {:else} ℹ️
            {/if}
          </span>
          <span class="toast-message">{notification.message}</span>
        </div>
      {/each}
    </div>
  </main>
</div>

<style>
  :global(*) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  :global(body) {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    background: #f3f4f6;
    color: #111827;
  }

  .app-layout {
    display: flex;
    height: 100vh;
    overflow: hidden;
  }

  .sidebar {
    width: 260px;
    background: linear-gradient(180deg, #1e3a8a 0%, #1e40af 100%);
    color: white;
    display: flex;
    flex-direction: column;
    transition: width 0.3s ease;
    flex-shrink: 0;
  }

  .sidebar.collapsed {
    width: 70px;
  }

  .sidebar-header {
    padding: 20px;
    border-bottom: rgba(255, 255, 255, 0.1) solid 1px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .logo-icon {
    font-size: 28px;
  }

  .logo-text {
    font-size: 16px;
    font-weight: 700;
    line-height: 1.2;
  }

  .sidebar.collapsed .logo-text {
    display: none;
  }

  .toggle-btn {
    background: rgba(255, 255, 255, 0.1);
    border: none;
    color: white;
    cursor: pointer;
    padding: 6px;
    border-radius: 4px;
    font-size: 18px;
  }

  .toggle-btn:hover {
    background: rgba(255, 255, 255, 0.2);
  }

  .nav-menu {
    flex: 1;
    padding: 12px 8px;
    overflow-y: auto;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 16px;
    border: none;
    background: transparent;
    color: rgba(255, 255, 255, 0.85);
    cursor: pointer;
    border-radius: 8px;
    margin-bottom: 4px;
    transition: all 0.2s;
    text-align: left;
    width: calc(100% - 8px);
  }

  .nav-item:hover {
    background: rgba(255, 255, 255, 0.15);
    color: white;
  }

  .nav-item.active {
    background: rgba(255, 255, 255, 0.25);
    color: white;
    font-weight: 600;
  }

  .nav-icon {
    font-size: 18px;
    flex-shrink: 0;
  }

  .nav-label {
    font-size: 14px;
  }

  .sidebar.collapsed .nav-label {
    display: none;
  }

  .sidebar-footer {
    padding: 16px 20px;
    border-top: rgba(255, 255, 255, 0.1) solid 1px;
    text-align: center;
  }

  .version-info {
    font-size: 11px;
    opacity: 0.7;
  }

  .main-content {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }

  .top-bar {
    height: 64px;
    background: white;
    border-bottom: 1px solid #e5e7eb;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 24px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  }

  .top-bar-left {
    display: flex;
    align-items: center;
    gap: 16px;
  }

  .menu-toggle {
    background: none;
    border: none;
    font-size: 20px;
    cursor: pointer;
    color: #374151;
    padding: 4px;
  }

  .search-box {
    position: relative;
  }

  .search-box input {
    width: 320px;
    padding: 8px 36px 8px 14px;
    border: 1px solid #d1d5db;
    border-radius: 8px;
    font-size: 14px;
    transition: all 0.2s;
  }

  .search-box input:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }

  .search-box kbd {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
    background: #f3f4f6;
    padding: 2px 6px;
    border-radius: 4px;
    font-family: monospace;
    font-size: 11px;
    color: #6b7280;
    border: 1px solid #d1d5db;
  }

  .top-bar-right {
    display: flex;
    align-items: center;
    gap: 20px;
  }

  .notification-area {
    position: relative;
    font-size: 20px;
    cursor: pointer;
  }

  .notification-badge {
    position: absolute;
    top: -4px;
    right: -4px;
    background: #ef4444;
    color: white;
    font-size: 10px;
    font-weight: 600;
    padding: 2px 5px;
    border-radius: 10px;
    min-width: 16px;
    text-align: center;
  }

  .user-info {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .user-avatar {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    background: #3b82f6;
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 13px;
    font-weight: 600;
  }

  .user-name {
    font-size: 14px;
    font-weight: 500;
    color: #374151;
  }

  .content-area {
    flex: 1;
    overflow-y: auto;
    position: relative;
  }

  .coming-soon {
    text-align: center;
    padding: 80px 20px;
    color: #6b7280;
  }

  .coming-icon {
    font-size: 72px;
    margin-bottom: 20px;
  }

  .coming-soon h2 {
    font-size: 24px;
    font-weight: 600;
    color: #374151;
    margin-bottom: 12px;
  }

  .toast {
    position: fixed;
    bottom: 24px;
    right: 24px;
    padding: 14px 20px;
    border-radius: 10px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
    display: flex;
    align-items: center;
    gap: 10px;
    z-index: 9999;
    animation: slideInUp 0.3s ease-out;
    max-width: 400px;
    font-size: 14px;
    font-weight: 500;
  }

  @keyframes slideInUp {
    from {
      transform: translateY(20px);
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }

  .toast-icon {
    font-size: 18px;
  }

  .notification-success {
    background: #d1fae5;
    color: #065f46;
    border-left: 4px solid #10b981;
  }

  .notification-error {
    background: #fee2e2;
    color: #991b1b;
    border-left: 4px solid #ef4444;
  }

  .notification-warning {
    background: #fef3c7;
    color: #92400e;
    border-left: 4px solid #f59e0b;
  }

  .notification-info {
    background: #dbeafe;
    color: #1e40af;
    border-left: 4px solid #3b82f6;
  }

  @media (max-width: 1024px) {
    .sidebar {
      width: 70px;
    }
    
    .sidebar .logo-text,
    .sidebar .nav-label {
      display: none;
    }
    
    .search-box input {
      width: 200px;
    }
  }

  @media (max-width: 768px) {
    .sidebar {
      position: fixed;
      left: 0;
      top: 0;
      bottom: 0;
      z-index: 1000;
      transform: translateX(-100%);
    }
    
    .sidebar.open {
      transform: translateX(0);
    }
    
    .search-box {
      display: none;
    }
    
    .user-name {
      display: none;
    }
  }
</style>
