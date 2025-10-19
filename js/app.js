/**
 * 主应用逻辑 - 首页（仅随机选择功能）
 */

import storage from './storage.js';
import { RandomSelector } from './random.js';
import { showToast } from './utils.js';

// 食物分类常量
const FOOD_CATEGORIES = {
  '快餐小吃': '🍔',
  '甜品饮品': '🍰',
  '大餐': '🍖',
  '水果': '🍎'
};

// 应用状态
const appState = {
  isRandomizing: false
};

/**
 * 初始化应用
 */
async function initApp() {
  try {
    // 显示加载状态
    updateSyncStatus(true, '加载中...');
    
    // 初始化数据
    await storage.initialize();
    
    // 更新同步状态
    updateSyncStatus(false, '已同步');
  } catch (error) {
    console.error('初始化失败:', error);
    updateSyncStatus(false, '离线模式');
  }
}

/**
 * 随机选择食物
 */
async function randomSelect() {
  const categorySelect = document.getElementById('random-category');
  const category = categorySelect.value;
  
  if (!category) {
    showToast('请先选择分类', 'warning');
    return;
  }
  
  const foods = storage.getFoodsByCategory(category);
  
  if (foods.length === 0) {
    showToast('该分类下没有食物', 'warning');
    return;
  }
  
  if (appState.isRandomizing) {
    return;
  }
  
  appState.isRandomizing = true;
  const resultDiv = document.getElementById('random-result');
  const btn = document.getElementById('random-btn');
  
  btn.disabled = true;
  btn.innerHTML = '<span class="loading loading-spinner"></span> 选择中...';
  
  // 带动画的随机选择
  await RandomSelector.selectWithAnimation(
    foods,
    (food, isFinal) => {
      if (isFinal) {
        resultDiv.innerHTML = `
          <div class="result-animation">
            <div class="text-6xl mb-4">${FOOD_CATEGORIES[food.category]}</div>
            <h3 class="text-3xl font-bold mb-2">${food.name}</h3>
            ${food.location ? `<p class="text-xl">📍 ${food.location}</p>` : ''}
          </div>
        `;
      } else {
        resultDiv.innerHTML = `
          <div class="shake-animation">
            <div class="text-4xl">${FOOD_CATEGORIES[food.category]}</div>
            <p class="text-xl mt-2">${food.name}</p>
          </div>
        `;
      }
    },
    2000
  );
  
  appState.isRandomizing = false;
  btn.disabled = false;
  btn.innerHTML = '<span class="text-2xl">🎯</span> 再来一次';
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

// 绑定事件
document.addEventListener('DOMContentLoaded', () => {
  // 初始化应用
  initApp();
  
  // 随机选择按钮
  document.getElementById('random-btn').addEventListener('click', randomSelect);
});
