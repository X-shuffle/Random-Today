#!/bin/bash

# Random Today 项目设置脚本

echo "🚀 开始设置 Random Today 项目..."

# 检查 Go 版本
echo "📋 检查 Go 版本..."
go version

# 检查 MongoDB 是否运行
echo "📋 检查 MongoDB 状态..."
if pgrep -x "mongod" > /dev/null; then
    echo "✅ MongoDB 正在运行"
else
    echo "⚠️  MongoDB 未运行，请启动 MongoDB"
    echo "   启动命令: mongod"
fi

# 安装依赖
echo "📦 安装 Go 依赖..."
go mod tidy

# 创建必要的目录
echo "📁 创建必要的目录..."
mkdir -p logs
mkdir -p tmp

# 设置权限
echo "🔐 设置文件权限..."
chmod +x scripts/*.sh

echo "✅ 项目设置完成！"
echo ""
echo "📝 下一步："
echo "1. 确保 MongoDB 正在运行"
echo "2. 运行项目: go run cmd/main.go"
echo "3. 访问 API: http://localhost:8080/api"
echo ""
echo "📚 查看文档:"
echo "- API 文档: docs/API.md"
echo "- 项目文档: README.md" 