/**
 * GitHub API 封装
 * 用于读写 GitHub 仓库中的数据文件
 */

// GitHub API 配置
const config = {
  token: import.meta.env.VITE_GITHUB_TOKEN,
  owner: import.meta.env.VITE_GITHUB_OWNER,
  repo: import.meta.env.VITE_GITHUB_REPO,
  branch: import.meta.env.VITE_GITHUB_BRANCH || 'main',
  filePath: 'data/foods.json'
};

// 验证配置
function validateConfig() {
  if (!config.token || !config.owner || !config.repo) {
    throw new Error('GitHub 配置不完整，请检查环境变量');
  }
}

/**
 * GitHub API 类
 */
export class GitHubAPI {
  constructor() {
    validateConfig();
    this.baseURL = 'https://api.github.com';
    this.currentSHA = null;
  }

  /**
   * 获取文件内容
   * @returns {Promise<{content: Object, sha: string}>}
   */
  async getFile() {
    const url = `${this.baseURL}/repos/${config.owner}/${config.repo}/contents/${config.filePath}?ref=${config.branch}`;
    
    try {
      const response = await fetch(url, {
        headers: {
          'Authorization': `token ${config.token}`,
          'Accept': 'application/vnd.github.v3+json'
        }
      });

      if (!response.ok) {
        if (response.status === 404) {
          // 文件不存在，返回空数据
          return { content: { version: '1.0.0', foods: [] }, sha: null };
        }
        throw new Error(`GitHub API 错误: ${response.status} ${response.statusText}`);
      }

      const data = await response.json();
      
      // Base64 解码
      const content = JSON.parse(decodeURIComponent(escape(atob(data.content))));
      this.currentSHA = data.sha;
      
      return { content, sha: data.sha };
    } catch (error) {
      throw error;
    }
  }

  /**
   * 更新文件内容
   * @param {Object} content - 文件内容
   * @param {string} message - 提交信息
   * @returns {Promise<Object>}
   */
  async updateFile(content, message = 'Update foods data') {
    const url = `${this.baseURL}/repos/${config.owner}/${config.repo}/contents/${config.filePath}`;
    
    try {
      // 准备内容
      const jsonContent = JSON.stringify(content, null, 2);
      const encodedContent = btoa(unescape(encodeURIComponent(jsonContent)));

      const body = {
        message,
        content: encodedContent,
        branch: config.branch
      };

      // 如果有 SHA，添加到请求中（更新现有文件）
      if (this.currentSHA) {
        body.sha = this.currentSHA;
      }

      const response = await fetch(url, {
        method: 'PUT',
        headers: {
          'Authorization': `token ${config.token}`,
          'Accept': 'application/vnd.github.v3+json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      });

      if (!response.ok) {
        throw new Error(`GitHub API 错误: ${response.status} ${response.statusText}`);
      }

      const result = await response.json();
      this.currentSHA = result.content.sha;
      
      return result;
    } catch (error) {
      throw error;
    }
  }

  /**
   * 创建文件（首次）
   * @param {Object} content - 文件内容
   * @returns {Promise<Object>}
   */
  async createFile(content) {
    this.currentSHA = null;
    return this.updateFile(content, 'Initial commit: Create foods data file');
  }
}

export default new GitHubAPI();
