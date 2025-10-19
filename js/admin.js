/**
 * 管理后台逻辑
 */

import storage from './storage.js';
import { showToast, formatDateTime } from './utils.js';

// 食物分类常量
const FOOD_CATEGORIES = {
  '快餐小吃': '🍔',
  '甜品饮品': '🍰',
  '大餐': '🍖',
  '水果': '🍎'
};

// 应用状态
const appState = {
  currentCategory: 'all',
  isAuthenticated: false
};

// 管理员密码
const ADMIN_PASSWORD = import.meta.env.VITE_ADMIN_PASSWORD;



/**
 * 检查登录状态
 */
function checkAuth() {
  const isLoggedIn = sessionStorage.getItem('admin_logged_in') === 'true';
  if (isLoggedIn) {
    appState.isAuthenticated = true;
    showAdminPage();
    initApp();
  } else {
    showLoginPage();
  }
}

/**
 * 显示登录页面
 */
function showLoginPage() {
  document.getElementById('login-page').classList.remove('hidden');
  document.getElementById('admin-page').classList.add('hidden');
}

/**
 * 显示管理页面
 */
function showAdminPage() {
  document.getElementById('login-page').classList.add('hidden');
  document.getElementById('admin-page').classList.remove('hidden');
}

/**
 * 处理登录
 */
function handleLogin(event) {
  event.preventDefault();
  
  const password = document.getElementById('password-input').value;
  const errorDiv = document.getElementById('login-error');
  
  if (!ADMIN_PASSWORD) {
    errorDiv.innerHTML = '<span>⚠️ 管理员密码未配置，请检查 .env 文件</span>';
    errorDiv.classList.remove('hidden');
    return;
  }
  
  if (password === ADMIN_PASSWORD) {
    // 登录成功
    sessionStorage.setItem('admin_logged_in', 'true');
    appState.isAuthenticated = true;
    errorDiv.classList.add('hidden');
    showAdminPage();
    initApp();
    showToast('登录成功', 'success');
  } else {
    // 登录失败
    errorDiv.innerHTML = '<span>❌ 密码错误，请重试</span>';
    errorDiv.classList.remove('hidden');
    document.getElementById('password-input').value = '';
    document.getElementById('password-input').focus();
  }
}

/**
 * 处理登出
 */
function handleLogout() {
  if (confirm('确定要退出管理后台吗？')) {
    sessionStorage.removeItem('admin_logged_in');
    appState.isAuthenticated = false;
    showLoginPage();
    showToast('已退出', 'info');
  }
}

/**
 * 初始化应用
 */
async function initApp() {
  if (!appState.isAuthenticated) return;
  
  try {
    // 显示加载状态
    updateSyncStatus(true, '加载中...');
    
    // 初始化数据
    await storage.initialize();
    
    // 渲染食物列表
    renderFoodList();
    
    // 更新统计信息
    updateStats();
    
    // 更新同步状态
    updateSyncStatus(false, '已同步');
    
    showToast('数据加载成功', 'success');
  } catch (error) {
    showToast('初始化失败，使用离线模式', 'warning');
    updateSyncStatus(false, '离线模式');
  }
}

/**
 * 渲染食物列表
 */
function renderFoodList() {
  const container = document.getElementById('food-list');
  const foods = appState.currentCategory === 'all' 
    ? storage.getAllFoods()
    : storage.getFoodsByCategory(appState.currentCategory);
  
  if (foods.length === 0) {
    container.innerHTML = `
      <div class="col-span-full text-center py-12 text-gray-500">
        <div class="text-6xl mb-4">🍽️</div>
        <p class="text-lg font-semibold">暂无食物数据</p>
        <p class="text-sm mt-2">添加一些食物开始使用吧！</p>
      </div>
    `;
    return;
  }
  
  container.innerHTML = foods.map(food => `
    <div class="glass-card rounded-2xl shadow-lg food-card p-6">
      <div class="flex items-start justify-between mb-3">
        <div class="flex items-center gap-3">
          <span class="text-4xl">${FOOD_CATEGORIES[food.category] || '🍽️'}</span>
          <div>
            <h3 class="font-bold text-lg">${food.name}</h3>
            ${food.location ? `<p class="text-sm text-gray-600">📍 ${food.location}</p>` : ''}
          </div>
        </div>
      </div>
      <div class="flex items-center justify-between mt-4 pt-4 border-t border-gray-200">
        <div class="text-xs text-gray-500">
          ${formatDateTime(food.createdAt)}
        </div>
        <button class="btn btn-sm btn-error btn-outline" onclick="deleteFood('${food.id}')">
          🗑️ 删除
        </button>
      </div>
    </div>
  `).join('');
}

/**
 * 添加食物
 */
async function addFood(event) {
  event.preventDefault();
  
  const form = event.target;
  const formData = new FormData(form);
  
  const food = {
    name: formData.get('name').trim(),
    location: formData.get('location').trim(),
    category: formData.get('category')
  };
  
  if (!food.name || !food.category) {
    showToast('请填写必填项', 'warning');
    return;
  }
  
  try {
    updateSyncStatus(true, '保存中...');
    await storage.addFood(food);
    
    form.reset();
    renderFoodList();
    updateStats();
    updateSyncStatus(false, '已同步');
    
    showToast('添加成功', 'success');
  } catch (error) {
    showToast('添加失败', 'error');
    updateSyncStatus(false, '同步失败');
  }
}

/**
 * 删除食物
 */
async function deleteFood(id) {
  if (!confirm('确定要删除这个食物吗？')) {
    return;
  }
  
  try {
    updateSyncStatus(true, '删除中...');
    await storage.removeFood(id);
    
    renderFoodList();
    updateStats();
    updateSyncStatus(false, '已同步');
    
    showToast('删除成功', 'success');
  } catch (error) {
    showToast('删除失败', 'error');
    updateSyncStatus(false, '同步失败');
  }
}

/**
 * 切换分类
 */
function switchCategory(category) {
  appState.currentCategory = category;
  renderFoodList();
}

/**
 * 导出数据
 */
function exportData() {
  try {
    storage.exportToFile();
    showToast('导出成功', 'success');
  } catch (error) {
    showToast('导出失败', 'error');
  }
}

/**
 * 手动同步
 */
async function syncData() {
  try {
    updateSyncStatus(true, '同步中...');
    const success = await storage.syncFromGitHub();
    
    if (success) {
      renderFoodList();
      updateStats();
      updateSyncStatus(false, '已同步');
      showToast('同步成功', 'success');
    } else {
      updateSyncStatus(false, '同步失败');
      showToast('同步失败', 'error');
    }
  } catch (error) {
    updateSyncStatus(false, '同步失败');
    showToast('同步失败', 'error');
  }
}

/**
 * 清空数据
 */
async function clearAllData() {
  if (!confirm('确定要清空所有数据吗？此操作不可恢复！')) {
    return;
  }
  
  if (!confirm('再次确认：真的要清空所有数据吗？')) {
    return;
  }
  
  try {
    updateSyncStatus(true, '清空中...');
    await storage.clearAll();
    
    renderFoodList();
    updateStats();
    updateSyncStatus(false, '已同步');
    
    showToast('数据已清空', 'success');
  } catch (error) {
    showToast('清空失败', 'error');
    updateSyncStatus(false, '操作失败');
  }
}

/**
 * 更新统计信息
 */
function updateStats() {
  const allFoods = storage.getAllFoods();
  const totalCount = allFoods.length;
  const lastSync = storage.getLastSyncTime();
  
  // 总数
  document.getElementById('total-count').textContent = totalCount;
  
  // 各分类统计
  const categoryMap = {
    '快餐小吃': 'count-fast',
    '甜品饮品': 'count-dessert',
    '大餐': 'count-meal',
    '水果': 'count-fruit'
  };
  
  Object.keys(categoryMap).forEach(category => {
    const count = allFoods.filter(food => food.category === category).length;
    document.getElementById(categoryMap[category]).textContent = count;
  });
  
  // 最后同步时间
  document.getElementById('last-sync').textContent = formatDateTime(lastSync);
}

/**
 * 更新同步状态
 */
function updateSyncStatus(isLoading, text) {
  const statusBtn = document.getElementById('sync-status');
  const spinner = statusBtn.querySelector('.loading');
  const statusText = statusBtn.querySelector('.status-text');
  
  if (isLoading) {
    spinner.classList.remove('hidden');
  } else {
    spinner.classList.add('hidden');
  }
  
  statusText.textContent = text;
}

// 页面加载时检查登录状态
document.addEventListener('DOMContentLoaded', () => {
  checkAuth();
  
  // 登录表单
  document.getElementById('login-form').addEventListener('submit', handleLogin);
  
  // 登出按钮
  document.getElementById('logout-btn')?.addEventListener('click', handleLogout);
  
  // 添加食物表单
  document.getElementById('add-food-form')?.addEventListener('submit', addFood);
  
  // 分类筛选按钮
  document.querySelectorAll('[data-category]').forEach(btn => {
    btn.addEventListener('click', () => {
      switchCategory(btn.dataset.category);
    });
  });
  
  // 数据管理按钮
  document.getElementById('export-btn')?.addEventListener('click', exportData);
  document.getElementById('sync-btn')?.addEventListener('click', syncData);
  document.getElementById('clear-btn')?.addEventListener('click', clearAllData);
});

// 暴露全局函数供 HTML 调用
window.deleteFood = deleteFood;
