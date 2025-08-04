package main

import (
	"log"

	"random_today/internal/application/handler"
	"random_today/internal/infrastructure/database"
)

func main() {
	// 初始化数据库连接
	if err := database.Init(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 初始化路由
	router := handler.SetupRoutes()

	// 启动服务器
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 