/**
 * 随机选择算法
 */

/**
 * 随机选择器类
 */
export class RandomSelector {
  /**
   * 从数组中随机选择一个元素
   * @param {Array} items - 待选择的数组
   * @returns {*} 随机选中的元素
   */
  static selectRandom(items) {
    if (!items || items.length === 0) {
      return null;
    }

    const randomIndex = Math.floor(Math.random() * items.length);
    return items[randomIndex];
  }

  /**
   * 从指定分类中随机选择食物
   * @param {Array} foods - 所有食物列表
   * @param {string} category - 分类名称
   * @returns {Object|null} 随机选中的食物
   */
  static selectFoodByCategory(foods, category) {
    if (!category) {
      return null;
    }

    const categoryFoods = foods.filter(food => food.category === category);
    return this.selectRandom(categoryFoods);
  }

  /**
   * 带动画效果的随机选择
   * @param {Array} items - 待选择的数组
   * @param {Function} onUpdate - 每次更新时的回调
   * @param {number} duration - 动画持续时间（毫秒）
   * @returns {Promise<*>} 最终选中的元素
   */
  static async selectWithAnimation(items, onUpdate, duration = 2000) {
    if (!items || items.length === 0) {
      return null;
    }

    const startTime = Date.now();
    const finalItem = this.selectRandom(items);
    
    return new Promise((resolve) => {
      const interval = setInterval(() => {
        const elapsed = Date.now() - startTime;
        
        if (elapsed >= duration) {
          clearInterval(interval);
          onUpdate(finalItem, true);
          resolve(finalItem);
        } else {
          // 随机显示一个项目
          const randomItem = this.selectRandom(items);
          onUpdate(randomItem, false);
        }
      }, 100);
    });
  }
}

export default RandomSelector;
