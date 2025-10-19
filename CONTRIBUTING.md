# 贡献指南

感谢你对本项目的关注！我们欢迎任何形式的贡献。

## 如何贡献

### 报告 Bug

如果你发现了 Bug，请：

1. 检查 [Issues](https://github.com/your-repo/issues) 是否已有相关报告
2. 如果没有，创建一个新的 Issue
3. 详细描述问题，包括：
   - 复现步骤
   - 预期行为
   - 实际行为
   - 截图（如果适用）
   - 环境信息（浏览器、操作系统等）

### 提出新功能

如果你有新功能的想法：

1. 创建一个 Issue，标记为 `enhancement`
2. 详细描述功能需求和使用场景
3. 等待讨论和反馈

### 提交代码

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建一个 Pull Request

### 代码规范

- 使用 ES6+ 语法
- 保持代码简洁清晰
- 添加必要的注释
- 遵循现有的代码风格

### 提交信息规范

使用清晰的提交信息：

- `feat: 添加新功能`
- `fix: 修复 Bug`
- `docs: 更新文档`
- `style: 代码格式调整`
- `refactor: 代码重构`
- `test: 添加测试`
- `chore: 构建/工具链更新`

## 开发流程

### 1. 设置开发环境

```bash
# 克隆仓库
git clone https://github.com/your-repo/random-today.git
cd random-today

# 安装依赖
npm install

# 配置环境变量
cp .env.example .env
# 编辑 .env 文件

# 启动开发服务器
npm run dev
```

### 2. 进行开发

- 在 `js/` 目录下修改 JavaScript 代码
- 在 `css/` 目录下修改样式
- 在 `index.html` 中修改页面结构

### 3. 测试

- 手动测试所有功能
- 确保在不同浏览器中正常工作
- 测试移动端响应式布局

### 4. 提交 Pull Request

- 确保代码通过所有检查
- 提供清晰的 PR 描述
- 关联相关的 Issue

## 项目结构

```
random-today/
├── index.html          # 主页面
├── js/                 # JavaScript 模块
│   ├── app.js         # 主应用逻辑
│   ├── storage.js     # 数据存储
│   ├── github-api.js  # GitHub API
│   ├── random.js      # 随机算法
│   └── utils.js       # 工具函数
├── css/               # 样式文件
└── data/              # 数据示例
```

## 需要帮助？

如有任何问题，欢迎：

- 创建 Issue
- 发起 Discussion
- 联系维护者

再次感谢你的贡献！🎉
