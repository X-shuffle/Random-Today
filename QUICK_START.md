# 🚀 快速启动指南

## 第一步：安装依赖

```bash
npm install
```

## 第二步：配置 GitHub

### 1. 获取 GitHub Token

访问：https://github.com/settings/tokens

点击 "Generate new token (classic)"，勾选 `repo` 权限，生成 token。

### 2. 创建数据仓库

在 GitHub 创建一个新仓库（建议私有），例如：`random-food-data`

在仓库中创建文件 `data/foods.json`，内容如下：

```json
{
  "version": "1.0.0",
  "foods": []
}
```

### 3. 配置环境变量

复制环境变量模板：

```bash
cp .env.example .env
```

编辑 `.env` 文件：

```env
VITE_GITHUB_TOKEN=ghp_你的token
VITE_GITHUB_OWNER=你的GitHub用户名
VITE_GITHUB_REPO=random-food-data
VITE_GITHUB_BRANCH=main
VITE_ADMIN_PASSWORD=你的管理员密码
```

## 第三步：启动开发服务器

```bash
npm run dev
```

浏览器会自动打开 http://localhost:3000

## 第四步：开始使用

### 首页使用
1. 访问 http://localhost:3000
2. 选择分类
3. 点击"开始选择"按钮
4. 享受随机选择的乐趣！

### 管理后台
1. 点击首页底部的"管理后台"链接
2. 输入你在 `.env` 中设置的管理员密码
3. 添加、管理食物数据
4. 查看统计信息

## 常见问题

### Q: 提示 GitHub 配置不完整？

A: 检查 `.env` 文件是否正确配置，确保所有变量都已填写。

### Q: 保存失败？

A: 检查 GitHub Token 是否有 `repo` 权限，仓库是否存在。

### Q: 数据没有同步？

A: 点击右上角的"手动同步"按钮，或检查网络连接。

### Q: 如何备份数据？

A: 点击"导出数据"按钮，将数据下载到本地。

## 部署到生产环境

### Cloudflare Pages

1. 推送代码到 GitHub
2. 登录 Cloudflare Pages
3. 连接仓库
4. 配置构建：
   - 构建命令：`npm run build`
   - 输出目录：`dist`
5. 添加环境变量（与 .env 相同，包括管理员密码）
6. 部署完成！

**重要**: 部署后记得测试管理后台的密码是否正确配置。

## 需要帮助？

查看完整文档：[README.md](./README.md)

提交问题：[GitHub Issues](https://github.com/your-repo/issues)
