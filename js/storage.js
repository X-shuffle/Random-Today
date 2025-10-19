/**
 * 数据存储管理
 * 负责 GitHub API 和 localStorage 的数据同步
 */

import githubAPI from './github-api.js';

const CACHE_KEY = 'foods';
const LAST_SYNC_KEY = 'last-sync';

/**
 * 数据存储类
 */
export class DataStorage {
  constructor() {
    this.foods = [];
    this.isOnline = navigator.onLine;
    this.isSyncing = false;
    
    // 监听网络状态
    window.addEventListener('online', () => {
      this.isOnline = true;
      this.syncFromGitHub();
    });
    
    window.addEventListener('offline', () => {
      this.isOnline = false;
    });
  }

  /**
   * 从 localStorage 加载缓存
   */
  loadFromCache() {
    try {
      const cached = localStorage.getItem(CACHE_KEY);
      if (cached) {
        this.foods = JSON.parse(cached);
        return this.foods;
      }
    } catch (error) {
      console.error('加载缓存失败:', error);
    }
    return [];
  }

  /**
   * 保存到 localStorage 缓存
   */
  saveToCache() {
    try {
      localStorage.setItem(CACHE_KEY, JSON.stringify(this.foods));
      localStorage.setItem(LAST_SYNC_KEY, new Date().toISOString());
    } catch (error) {
      console.error('保存缓存失败:', error);
    }
  }

  /**
   * 从 GitHub 加载数据
   */
  async loadFromGitHub() {
    if (!this.isOnline) {
      throw new Error('网络不可用，使用缓存数据');
    }

    try {
      this.isSyncing = true;
      const { content } = await githubAPI.getFile();
      this.foods = content.foods || [];
      this.saveToCache();
      return this.foods;
    } catch (error) {
      console.error('从 GitHub 加载失败:', error);
      throw error;
    } finally {
      this.isSyncing = false;
    }
  }

  /**
   * 保存到 GitHub
   */
  async saveToGitHub() {
    if (!this.isOnline) {
      // 离线时只保存到缓存
      this.saveToCache();
      return false;
    }

    try {
      this.isSyncing = true;
      const content = {
        version: '1.0.0',
        updatedAt: new Date().toISOString(),
        foods: this.foods
      };

      await githubAPI.updateFile(
        content,
        `Update foods data at ${new Date().toLocaleString('zh-CN')}`
      );

      this.saveToCache();
      return true;
    } catch (error) {
      console.error('保存到 GitHub 失败:', error);
      // 失败时仍然保存到缓存
      this.saveToCache();
      return false;
    } finally {
      this.isSyncing = false;
    }
  }

  /**
   * 同步数据（从 GitHub 拉取最新）
   */
  async syncFromGitHub() {
    try {
      await this.loadFromGitHub();
      return true;
    } catch (error) {
      console.error('同步失败:', error);
      return false;
    }
  }

  /**
   * 初始化数据
   */
  async initialize() {
    // 先加载缓存，快速显示
    this.loadFromCache();
    
    // 然后从 GitHub 同步最新数据
    try {
      await this.loadFromGitHub();
    } catch (error) {
      console.log('使用缓存数据（离线模式）');
    }
    
    return this.foods;
  }

  /**
   * 添加食物
   */
  async addFood(food) {
    const newFood = {
      id: Date.now().toString(),
      name: food.name,
      location: food.location || '',
      category: food.category,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    };

    this.foods.push(newFood);
    await this.saveToGitHub();
    return newFood;
  }

  /**
   * 删除食物
   */
  async removeFood(id) {
    this.foods = this.foods.filter(food => food.id !== id);
    await this.saveToGitHub();
  }

  /**
   * 更新食物
   */
  async updateFood(id, updates) {
    const index = this.foods.findIndex(food => food.id === id);
    if (index !== -1) {
      this.foods[index] = {
        ...this.foods[index],
        ...updates,
        updatedAt: new Date().toISOString()
      };
      await this.saveToGitHub();
      return this.foods[index];
    }
    return null;
  }

  /**
   * 获取所有食物
   */
  getAllFoods() {
    return this.foods;
  }

  /**
   * 按分类获取食物
   */
  getFoodsByCategory(category) {
    if (category === 'all') {
      return this.foods;
    }
    return this.foods.filter(food => food.category === category);
  }

  /**
   * 清空所有数据
   */
  async clearAll() {
    this.foods = [];
    await this.saveToGitHub();
    localStorage.clear();
  }

  /**
   * 导出数据为 JSON 文件
   */
  exportToFile() {
    const content = {
      version: '1.0.0',
      exportedAt: new Date().toISOString(),
      foods: this.foods
    };

    const blob = new Blob([JSON.stringify(content, null, 2)], { 
      type: 'application/json' 
    });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `foods-backup-${new Date().toISOString().split('T')[0]}.json`;
    a.click();
    URL.revokeObjectURL(url);
  }

  /**
   * 获取最后同步时间
   */
  getLastSyncTime() {
    const lastSync = localStorage.getItem(LAST_SYNC_KEY);
    if (lastSync) {
      return new Date(lastSync);
    }
    return null;
  }
}

export default new DataStorage();
