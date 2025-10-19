# 今天吃什么 - 随机选择食物应用

基于 htmx + GitHub API 构建的纯前端随机选择食物应用，帮助你解决"今天吃什么"的困扰。

## ✨ 特性

- 🎲 **随机选择** - 从指定分类中随机选择食物，带动画效果
- 🔐 **管理后台** - 密码保护的管理界面，安全管理食物数据
- 📝 **食物管理** - 添加、删除、分类管理食物（管理后台）
- ☁️ **云端存储** - 数据存储在你自己的 GitHub 仓库
- 🔄 **多设备同步** - 通过 GitHub 实现跨设备数据同步
- 📱 **响应式设计** - 完美支持移动端和桌面端
- 🚀 **极简技术栈** - htmx + 原生 JavaScript + DaisyUI
- 💾 **离线支持** - localStorage 缓存支持离线使用
- 🔒 **数据安全** - 数据完全由你控制，支持私有仓库

## 🚀 快速开始

### 1. 克隆项目

```bash
git clone <your-repo-url>
cd random-today
```

### 2. 安装依赖

```bash
npm install
```

### 3. 配置 GitHub

#### 3.1 创建 GitHub Personal Access Token

1. 访问 [GitHub Settings > Developer settings > Personal access tokens > Tokens (classic)](https://github.com/settings/tokens)
2. 点击 "Generate new token (classic)"
3. 勾选 `repo` 权限（完整的仓库访问权限）
4. 生成并复制 token

#### 3.2 创建数据仓库

1. 在 GitHub 创建一个新仓库（可以是私有仓库）
2. 在仓库中创建 `data/foods.json` 文件
3. 初始化为空数据：

```json
{
  "version": "1.0.0",
  "createdAt": "2024-01-01T00:00:00.000Z",
  "updatedAt": "2024-01-01T00:00:00.000Z",
  "foods": []
}
```

#### 3.3 配置环境变量

复制 `.env.example` 为 `.env`：

```bash
cp .env.example .env
```

编辑 `.env` 文件，填入你的 GitHub 信息和管理员密码：

```env
VITE_GITHUB_TOKEN=ghp_your_token_here
VITE_GITHUB_OWNER=your_github_username
VITE_GITHUB_REPO=your_repo_name
VITE_GITHUB_BRANCH=main
VITE_ADMIN_PASSWORD=your_admin_password
```

### 4. 启动开发服务器

```bash
npm run dev
```

访问 http://localhost:3000

## 📦 部署

### 方式一：部署到 Vercel（推荐）

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone)

1. 点击上方按钮，或访问 [Vercel](https://vercel.com/)
2. 导入你的 Git 仓库
3. 配置项目：
   - **Framework Preset**: Vite
   - **Build Command**: `npm run build`
   - **Output Directory**: `dist`
   - **Install Command**: `npm install`

4. 配置环境变量（在 Vercel 项目设置中）：
   ```
   VITE_GITHUB_TOKEN=your_token
   VITE_GITHUB_OWNER=your_username
   VITE_GITHUB_REPO=your_repo
   VITE_GITHUB_BRANCH=main
   VITE_ADMIN_PASSWORD=your_admin_password
   ```

5. 点击 Deploy，等待部署完成！

### 方式二：部署到 Cloudflare Pages

1. 构建项目：
   ```bash
   npm run build
   ```

2. 登录 [Cloudflare Pages](https://pages.cloudflare.com/)
3. 连接你的 Git 仓库
4. 配置构建设置：
   - **构建命令**: `npm run build`
   - **构建输出目录**: `dist`
   - **根目录**: `/`

5. 配置环境变量（与 Vercel 相同）

6. 部署完成！

## 📖 使用说明

### 首页 - 随机选择

1. 访问首页 `/`
2. 选择食物分类
3. 点击"开始选择"按钮
4. 等待动画结束，查看结果

### 管理后台

1. 点击首页底部的"管理后台"链接，或直接访问 `/admin.html`
2. 输入管理员密码登录
3. 在管理后台可以：
   - **添加食物**: 选择分类，输入食物名称和地点
   - **删除食物**: 在食物列表中删除不需要的食物
   - **导出数据**: 将数据导出为 JSON 文件到本地
   - **手动同步**: 从 GitHub 拉取最新数据
   - **清空数据**: 清空所有食物数据（谨慎操作）
   - **查看统计**: 查看各分类的食物数量

## 🛠️ 技术栈

- **前端框架**: htmx 1.9.10
- **UI 组件**: DaisyUI 4.4.0 + Tailwind CSS
- **构建工具**: Vite 5.0
- **数据存储**: GitHub API + localStorage
- **部署平台**: Vercel / Cloudflare Pages

## 📁 项目结构

```
random-today/
├── index.html              # 主页面（随机选择）
├── admin.html              # 管理后台页面
├── js/
│   ├── app.js             # 主页面逻辑
│   ├── admin.js           # 管理后台逻辑
│   ├── storage.js         # 数据存储管理
│   ├── github-api.js      # GitHub API 封装
│   ├── random.js          # 随机算法
│   └── utils.js           # 工具函数
├── css/
│   └── custom.css         # 自定义样式
├── manifest.json          # PWA 配置
├── package.json           # 项目配置
├── vite.config.js         # Vite 配置
└── .env.example           # 环境变量示例
```

## 🔒 安全提示

- ⚠️ **不要将 `.env` 文件提交到 Git 仓库**
- ⚠️ **GitHub Token 应该具有 `repo` 权限**
- ⚠️ **建议使用私有仓库存储数据**
- ⚠️ **设置强密码保护管理后台**
- ⚠️ **定期备份数据（使用导出功能）**
- ⚠️ **管理员密码存储在环境变量中，部署时需要配置**

## 📝 开发计划

- [ ] 支持食物编辑功能
- [ ] 添加食物图片上传
- [ ] 支持食物评分和备注
- [ ] 添加历史记录功能
- [ ] 支持多语言
- [ ] 添加主题切换

## 📄 License

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📮 联系方式

如有问题或建议，请提交 Issue。
