<script lang="ts">
  import { onMount } from 'svelte';
  import AdvancedFilterBar from '../components/AdvancedFilterBar.svelte';
  import BatchActionToolbar from '../components/BatchActionToolbar.svelte';
  import CommandPalette from '../components/CommandPalette.svelte';
  import DetailDrawer from '../components/DetailDrawer.svelte';
  import { memberBenefitsStore, uiStore } from '../stores';
  import type { MemberBenefit } from '../types';
  
  let selectedItems: MemberBenefit[] = [];
  let showDetailDrawer = false;
  let showCommandPalette = false;
  let showCreateModal = false;
  let editingItem: MemberBenefit | null = null;
  
  const filters = [
    { key: 'status', label: '状态', type: 'select' as const, options: [
      { value: 'active', label: '活跃' },
      { value: 'expired', label: '已过期' },
      { value: 'pending_review', label: '待复核' },
      { value: 'archived', label: '已归档' }
    ]},
    { key: 'member_level', label: '会员等级', type: 'select' as const, options: [
      { value: 'platinum', label: '白金卡' },
      { value: 'gold', label: '金卡' },
      { value: 'silver', label: '银卡' },
      { value: 'standard', label: '标准卡' }
    ]},
    { key: 'search', label: '搜索', type: 'text' as const },
    { key: 'airport_code', label: '机场', type: 'text' as const }
  ];
  
  const batchActions = [
    { id: 'batch_archive', label: '批量归档', icon: '📦', variant: 'secondary' as const },
    { id: 'batch_activate', label: '批量激活', icon: '✅', variant: 'primary' as const },
    { id: 'batch_delete', label: '批量删除', icon: '🗑️', variant: 'danger' as const }
  ];

  onMount(() => {
    memberBenefitsStore.fetchMemberBenefits();
    
    document.addEventListener('keydown', (e) => {
      if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
        e.preventDefault();
        showCommandPalette = true;
      }
    });
  });

  function handleFilterChange(filters: any) {
    memberBenefitsStore.setFilters(filters);
    memberBenefitsStore.fetchMemberBenefits(1, 10);
  }

  function handleSelect(item: MemberBenefit) {
    if (selectedItems.find(i => i.id === item.id)) {
      selectedItems = selectedItems.filter(i => i.id !== item.id);
    } else {
      selectedItems = [...selectedItems, item];
    }
  }

  function handleBatchAction(actionId: string) {
    switch (actionId) {
      case 'batch_archive':
        uiStore.addNotification('批量归档操作已触发', 'info');
        break;
      case 'batch_activate':
        uiStore.addNotification('批量激活操作已触发', 'success');
        break;
      case 'batch_delete':
        if (confirm('确认删除选中的项目？')) {
          uiStore.addNotification('批量删除操作已触发', 'warning');
        }
        break;
    }
  }

  function openDetail(item: MemberBenefit) {
    memberBenefitsStore.selectMemberBenefit(item);
    showDetailDrawer = true;
  }

  function openCreate() {
    editingItem = null;
    showCreateModal = true;
  }

  function openEdit(item: MemberBenefit) {
    editingItem = item;
    showCreateModal = true;
  }

  async function handleSave(data: Partial<MemberBenefit>) {
    try {
      if (editingItem?.id) {
        await memberBenefitsStore.updateMemberBenefit(editingItem.id, data);
        uiStore.addNotification('更新成功', 'success');
      } else {
        await memberBenefitsStore.createMemberBenefit(data);
        uiStore.addNotification('创建成功', 'success');
      }
      showCreateModal = false;
    } catch (error) {
      uiStore.addNotification('操作失败: ' + (error as Error).message, 'error');
    }
  }

  async function handleDelete(id: number) {
    if (confirm('确认删除此记录？')) {
      try {
        await memberBenefitsStore.deleteMemberBenefit(id);
        uiStore.addNotification('删除成功', 'success');
      } catch (error) {
        uiStore.addNotification('删除失败: ' + (error as Error).message, 'error');
      }
    }
  }

  async function handleStatusTransition(id: number, toStatus: string, action: string) {
    let reason = '';
    if (toStatus === 'rejected') {
      reason = prompt('请输入驳回原因：');
      if (!reason) return;
    }
    
    try {
      await memberBenefitsStore.transitionStatus(id, toStatus, action, reason);
      uiStore.addNotification(`状态已更新为 ${toStatus}`, 'success');
    } catch (error) {
      uiStore.addNotification('状态更新失败: ' + (error as Error).message, 'error');
    }
  }

  $: isLoading = $memberBenefitsStore.loading;
  $: benefits = $memberBenefitsStore.data;
  $: pagination = $memberBenefitsStore.pagination;

  const commands = [
    { id: 'create', title: '新建会员权益', category: '操作', shortcut: 'Ctrl+N', icon: '➕', action: () => openCreate() },
    { id: 'refresh', title: '刷新数据', category: '操作', shortcut: 'F5', icon: '🔄', action: () => memberBenefitsStore.fetchMemberBenefits() },
    { id: 'filter_active', title: '筛选活跃权益', category: '筛选', action: () => handleFilterChange({ status: 'active' }) },
    { id: 'filter_expired', title: '筛选过期权益', category: '筛选', action: () => handleFilterChange({ status: 'expired' }) },
    { id: 'export', title: '导出数据', category: '导出', icon: '📤', action: () => uiStore.addNotification('导出功能开发中', 'info') },
  ];
</script>

<div class="member-benefits-page">
  <div class="page-header">
    <div class="header-left">
      <h1 class="page-title">会员权益管理</h1>
      <p class="page-subtitle">管理机场贵宾厅会员权益信息</p>
    </div>
    <div class="header-actions">
      <button class="btn btn-primary" on:click={openCreate}>
        ➕ 新建权益
      </button>
    </div>
  </div>

  <AdvancedFilterBar 
    bind:filters 
    on:filterChange={(e) => handleFilterChange(e.detail)}
    on:applyFilters={(e) => handleFilterChange(e.detail)}
  />

  <BatchActionToolbar 
    bind:selectedItems 
    availableActions={batchActions}
    on:action={(e) => handleBatchAction(e.detail.actionId)}
    onSelectAll={() => selectedItems = [...benefits]}
    onClearSelection={() => selectedItems = []}
  />

  <div class="table-container">
    {#if isLoading}
      <div class="loading-state">
        <div class="spinner"></div>
        <p>加载中...</p>
      </div>
    {:else if benefits.length === 0}
      <div class="empty-state">
        <span class="empty-icon">💳</span>
        <h3>暂无会员权益数据</h3>
        <p>点击"新建权益"按钮开始添加</p>
        <button class="btn btn-primary" on:click={openCreate}>立即创建</button>
      </div>
    {:else}
      <table class="data-table">
        <thead>
          <tr>
            <th class="checkbox-col">
              <input 
                type="checkbox" 
                checked={selectedItems.length === benefits.length && benefits.length > 0}
                on:change={() => selectedItems.length === benefits.length ? selectedItems = [] : selectedItems = [...benefits]}
              />
            </th>
            <th>编号</th>
            <th>名称</th>
            <th>状态</th>
            <th>会员等级</th>
            <th>剩余额度</th>
            <th>有效期至</th>
            <th>负责人</th>
            <th>批次号</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          {#each benefits as benefit (benefit.id)}
            {@const isSelected = selectedItems.find(i => i.id === benefit.id)}
            <tr class={isSelected ? 'selected' : ''}>
              <td class="checkbox-col">
                <input 
                  type="checkbox"
                  checked={!!isSelected}
                  on:change={() => handleSelect(benefit)}
                />
              </td>
              <td><code>{benefit.code}</code></td>
              <td class="name-cell">{benefit.name}</td>
              <td>
                <span class="status-badge status-{benefit.status}">
                  {getStatusText(benefit.status)}
                </span>
              </td>
              <td>{getLevelText(benefit.member_level)}</td>
              <td class="quota-cell">{benefit.remaining_quota} 次</td>
              <td>{formatDate(benefit.valid_until)}</td>
              <td>{benefit.owner}</td>
              <td><code>{benefit.batch_id}</code></td>
              <td class="actions-cell">
                <button class="btn-icon" title="查看详情" on:click={() => openDetail(benefit)}>👁️</button>
                <button class="btn-icon" title="编辑" on:click={() => openEdit(benefit)}>✏️</button>
                <button class="btn-icon danger" title="删除" on:click={() => handleDelete(benefit.id)}>🗑️</button>
                
                {#if benefit.status === 'active'}
                  <button class="btn-icon warning" title="归档" on:click={() => handleStatusTransition(benefit.id, 'archived', 'archive')}>📦</button>
                {:else if benefit.status === 'pending_review'}
                  <button class="btn-icon success" title="通过复核" on:click={() => handleStatusTransition(benefit.id, 'active', 'approve')}>✅</button>
                  <button class="btn-icon danger" title="驳回" on:click={() => handleStatusTransition(benefit.id, 'rejected', 'reject')}>❌</button>
                {/if}
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
      
      <div class="pagination">
        <button 
          class="page-btn" 
          disabled={pagination.page <= 1}
          on:click={() => memberBenefitsStore.fetchMemberBenefits(pagination.page - 1, pagination.pageSize)}
        >
          上一页
        </button>
        <span class="page-info">
          第 {pagination.page} / {pagination.totalPages} 页 (共 {pagination.totalCount} 条)
        </span>
        <button 
          class="page-btn"
          disabled={pagination.page >= pagination.totalPages}
          on:click={() => memberBenefitsStore.fetchMemberBenefits(pagination.page + 1, pagination.pageSize)}
        >
          下一页
        </button>
      </div>
    {/if}
  </div>

  <DetailDrawer 
    bind:isOpen={showDetailDrawer}
    title="会员权益详情"
    data={$memberBenefitsStore.selectedMemberBenefit}
    on:close={() => showDetailDrawer = false}
  >
    {#if $memberBenefitsStore.selectedMemberBenefit}
      <div class="detail-content">
        <div class="detail-section">
          <h4>基本信息</h4>
          <div class="detail-grid">
            <div class="detail-item">
              <label>编号</label>
              <value>{$memberBenefitsStore.selectedMemberBenefit.code}</value>
            </div>
            <div class="detail-item">
              <label>名称</label>
              <value>{$memberBenefitsStore.selectedMemberBenefit.name}</value>
            </div>
            <div class="detail-item">
              <label>状态</label>
              <value>
                <span class="status-badge status-{$memberBenefitsStore.selectedMemberBenefit.status}">
                  {getStatusText($memberBenefitsStore.selectedMemberBenefit.status)}
                </span>
              </value>
            </div>
            <div class="detail-item">
              <label>会员等级</label>
              <value>{getLevelText($memberBenefitsStore.selectedMemberBenefit.member_level)}</value>
            </div>
          </div>
        </div>
        
        <div class="detail-section">
          <h4>额度与时间</h4>
          <div class="detail-grid">
            <div class="detail-item">
              <label>剩余额度</label>
              <value>{$memberBenefitsStore.selectedMemberBenefit.remaining_quota} 次</value>
            </div>
            <div class="detail-item">
              <label>生效日期</label>
              <value>{formatDate($memberBenefitsStore.selectedMemberBenefit.valid_from)}</value>
            </div>
            <div class="detail-item">
              <label>到期日期</label>
              <value>{formatDate($memberBenefitsStore.selectedMemberBenefit.valid_until)}</value>
            </div>
          </div>
        </div>
        
        {#if $memberBenefitsStore.selectedMemberBenefit.notes}
          <div class="detail-section">
            <h4>备注</h4>
            <p class="notes-text">{$memberBenefitsStore.selectedMemberBenefit.notes}</p>
          </div>
        {/if}
      </div>
    {/if}
  </DetailDrawer>

  {#if showCommandPalette}
    <CommandPalette 
      bind:isOpen={showCommandPalette}
      commands={commands}
      on:close={() => showCommandPalette = false}
    />
  {/if}

  {#if showCreateModal}
    <div class="modal-overlay" on:click|self={() => showCreateModal = false}>
      <div class="modal-content">
        <div class="modal-header">
          <h2>{editingItem ? '编辑会员权益' : '新建会员权益'}</h2>
          <button class="close-btn" on:click={() => showCreateModal = false}>✕</button>
        </div>
        
        <form on:submit|preventDefault={(e) => {
          const formData = new FormData(e.target as HTMLFormElement);
          handleSave({
            name: formData.get('name') as string,
            member_level: formData.get('member_level') as string,
            remaining_quota: Number(formData.get('remaining_quota')),
            valid_from: formData.get('valid_from') as string,
            valid_until: formData.get('valid_until') as string,
            owner: formData.get('owner') as string,
            airport_code: formData.get('airport_code') as string,
            lounge_id: formData.get('lounge_id') as string,
            notes: formData.get('notes') as string,
          });
        }}>
          <div class="form-group">
            <label for="name">权益名称 *</label>
            <input type="text" id="name" name="name" required value={editingItem?.name || ''} placeholder="例如：2026年度白金贵宾厅权益"/>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label for="member_level">会员等级 *</label>
              <select id="member_level" name="member_level" required value={editingItem?.member_level || ''}>
                <option value="">请选择</option>
                <option value="platinum">白金卡</option>
                <option value="gold">金卡</option>
                <option value="silver">银卡</option>
                <option value="standard">标准卡</option>
              </select>
            </div>
            
            <div class="form-group">
              <label for="remaining_quota">剩余额度 *</label>
              <input type="number" id="remaining_quota" name="remaining_quota" required min="0" value={editingItem?.remaining_quota || 0}/>
            </div>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label for="valid_from">生效日期 *</label>
              <input type="date" id="valid_from" name="valid_from" required value={editingItem?.valid_from?.split('T')[0] || ''}/>
            </div>
            
            <div class="form-group">
              <label for="valid_until">到期日期 *</label>
              <input type="date" id="valid_until" name="valid_until" required value={editingItem?.valid_until?.split('T')[0] || ''}/>
            </div>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label for="owner">负责人 *</label>
              <input type="text" id="owner" name="owner" required value={editingItem?.owner || ''} placeholder="负责人姓名"/>
            </div>
            
            <div class="form-group">
              <label for="airport_code">机场代码</label>
              <input type="text" id="airport_code" name="airport_code" value={editingItem?.airport_code || ''} placeholder="例如：PEK"/>
            </div>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label for="lounge_id">贵宾厅ID</label>
              <input type="text" id="lounge_id" name="lounge_id" value={editingItem?.lounge_id || ''} placeholder="贵宾厅编号"/>
            </div>
            
            <div class="form-group">
              <label for="batch_id">批次号</label>
              <input type="text" id="batch_id" name="batch_id" readonly value={`BATCH-${Date.now()}`}/>
            </div>
          </div>
          
          <div class="form-group">
            <label for="notes">备注</label>
            <textarea id="notes" name="notes" rows="3" placeholder="可选备注信息">{editingItem?.notes || ''}</textarea>
          </div>
          
          <div class="modal-actions">
            <button type="button" class="btn btn-secondary" on:click={() => showCreateModal = false}>取消</button>
            <button type="submit" class="btn btn-primary">{editingItem ? '更新' : '创建'}</button>
          </div>
        </form>
      </div>
    </div>
  {/if}
</div>

<script lang="ts">
  function getStatusText(status: string): string {
    const map: Record<string, string> = {
      active: '活跃',
      expired: '已过期',
      pending_review: '待复核',
      archived: '已归档',
      rejected: '已驳回',
      draft: '草稿'
    };
    return map[status] || status;
  }

  function getLevelText(level: string): string {
    const map: Record<string, string> = {
      platinum: '白金卡',
      gold: '金卡',
      silver: '银卡',
      standard: '标准卡'
    };
    return map[level] || level;
  }

  function formatDate(dateStr: string): string {
    if (!dateStr) return '-';
    return new Date(dateStr).toLocaleDateString('zh-CN');
  }
</script>

<style>
  .member-benefits-page {
    padding: 24px;
  }

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 24px;
  }

  .page-title {
    font-size: 28px;
    font-weight: 700;
    color: #111827;
    margin: 0 0 8px;
  }

  .page-subtitle {
    font-size: 16px;
    color: #6b7280;
    margin: 0;
  }

  .header-actions {
    display: flex;
    gap: 12px;
  }

  .btn {
    padding: 10px 20px;
    border-radius: 8px;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    border: none;
    display: inline-flex;
    align-items: center;
    gap: 6px;
  }

  .btn-primary {
    background: #3b82f6;
    color: white;
  }

  .btn-primary:hover {
    background: #2563eb;
    transform: translateY(-1px);
  }

  .btn-secondary {
    background: #f3f4f6;
    color: #374151;
  }

  .btn-secondary:hover {
    background: #e5e7eb;
  }

  .table-container {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
    overflow-x: auto;
  }

  .loading-state,
  .empty-state {
    text-align: center;
    padding: 60px 20px;
    color: #6b7280;
  }

  .empty-icon {
    font-size: 64px;
    margin-bottom: 16px;
  }

  .spinner {
    width: 40px;
    height: 40px;
    border: 4px solid #e5e7eb;
    border-top-color: #3b82f6;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin: 0 auto 16px;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  .data-table {
    width: 100%;
    border-collapse: collapse;
    font-size: 14px;
  }

  .data-table th {
    background: #f9fafb;
    padding: 12px 16px;
    text-align: left;
    font-weight: 600;
    color: #374151;
    border-bottom: 2px solid #e5e7eb;
    white-space: nowrap;
  }

  .data-table td {
    padding: 12px 16px;
    border-bottom: 1px solid #f3f4f6;
    vertical-align: middle;
  }

  .data-table tr:hover {
    background: #fafafa;
  }

  .data-table tr.selected {
    background: #eff6ff;
  }

  .checkbox-col {
    width: 40px;
  }

  .code {
    font-family: monospace;
    font-size: 13px;
    color: #6b7280;
    background: #f3f4f6;
    padding: 2px 6px;
    border-radius: 4px;
  }

  .name-cell {
    font-weight: 500;
    color: #111827;
  }

  .status-badge {
    display: inline-block;
    padding: 4px 10px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: 500;
  }

  .status-active {
    background: #d1fae5;
    color: #065f46;
  }

  .status-expired {
    background: #fee2e2;
    color: #991b1b;
  }

  .status-pending_review {
    background: #fef3c7;
    color: #92400e;
  }

  .status-archived {
    background: #e5e7eb;
    color: #374151;
  }

  .status-rejected {
    background: #fee2e2;
    color: #991b1b;
  }

  .quota-cell {
    font-weight: 600;
    color: #3b82f6;
  }

  .actions-cell {
    white-space: nowrap;
  }

  .btn-icon {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 16px;
    padding: 4px;
    opacity: 0.6;
    transition: all 0.2s;
  }

  .btn-icon:hover {
    opacity: 1;
    transform: scale(1.1);
  }

  .btn-icon.danger:hover {
    color: #ef4444;
  }

  .btn-icon.warning:hover {
    color: #f59e0b;
  }

  .btn-icon.success:hover {
    color: #10b981;
  }

  .pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
    padding: 16px;
    border-top: 1px solid #e5e7eb;
  }

  .page-btn {
    padding: 6px 12px;
    border: 1px solid #d1d5db;
    background: white;
    border-radius: 6px;
    cursor: pointer;
    font-size: 13px;
    transition: all 0.2s;
  }

  .page-btn:hover:not(:disabled) {
    background: #f3f4f6;
    border-color: #9ca3af;
  }

  .page-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .page-info {
    font-size: 13px;
    color: #6b7280;
  }

  .detail-content {
    padding: 8px 0;
  }

  .detail-section {
    margin-bottom: 24px;
  }

  .detail-section h4 {
    font-size: 15px;
    font-weight: 600;
    color: #111827;
    margin: 0 0 12px;
    padding-bottom: 8px;
    border-bottom: 1px solid #e5e7eb;
  }

  .detail-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
  }

  .detail-item {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .detail-item label {
    font-size: 12px;
    color: #6b7280;
    font-weight: 500;
  }

  .detail-item value {
    font-size: 14px;
    color: #111827;
    font-weight: 500;
  }

  .notes-text {
    font-size: 14px;
    color: #374151;
    line-height: 1.6;
    margin: 0;
    background: #f9fafb;
    padding: 12px;
    border-radius: 6px;
  }

  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    animation: fadeIn 0.2s ease-out;
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }

  .modal-content {
    background: white;
    border-radius: 12px;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
    width: 90%;
    max-width: 700px;
    max-height: 90vh;
    overflow-y: auto;
    animation: slideUp 0.3s ease-out;
  }

  @keyframes slideUp {
    from {
      transform: translateY(30px);
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 24px;
    border-bottom: 1px solid #e5e7eb;
  }

  .modal-header h2 {
    font-size: 20px;
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
  }

  .close-btn:hover {
    background: #f3f4f6;
    color: #111827;
  }

  form {
    padding: 24px;
  }

  .form-group {
    margin-bottom: 16px;
  }

  .form-group label {
    display: block;
    font-size: 13px;
    font-weight: 600;
    color: #374151;
    margin-bottom: 6px;
  }

  .form-group input,
  .form-group select,
  .form-group textarea {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid #d1d5db;
    border-radius: 6px;
    font-size: 14px;
    transition: all 0.2s;
    box-sizing: border-box;
  }

  .form-group input:focus,
  .form-group select:focus,
  .form-group textarea:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }

  .form-row {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }

  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
    padding-top: 16px;
    border-top: 1px solid #e5e7eb;
    margin-top: 24px;
  }

  @media (max-width: 768px) {
    .form-row {
      grid-template-columns: 1fr;
    }
    
    .detail-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
