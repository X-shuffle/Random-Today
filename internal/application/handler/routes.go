package handler

import (
	appService "random_today/internal/application/service"
	domainService "random_today/internal/domain/service"
	infraRepo "random_today/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)

	// 创建Gin引擎
	router := gin.Default()

	// 添加CORS中间件
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 创建依赖注入
	// 注意：这里需要完善依赖注入，暂时使用空实例
	preferenceRepo := &infraRepo.MongoDBPreferenceRepository{}
	domainService := domainService.NewPreferenceService(preferenceRepo)
	appService := appService.NewPreferenceAppService(domainService)

	// 创建处理器
	preferenceHandler := NewPreferenceHandler(appService)
	healthHandler := NewHealthHandler()

	// 健康检查路由
	router.GET("/health", healthHandler.HealthCheck)

	// API路由组
	api := router.Group("/api")
	{
		// 偏好管理路由
		preferences := api.Group("/preferences")
		{
			preferences.POST("", preferenceHandler.CreatePreference)
			preferences.GET("", preferenceHandler.GetAllPreferences)
			preferences.GET("/random", preferenceHandler.GetRandomPreference)
			preferences.GET("/type", preferenceHandler.GetPreferencesByType)
			preferences.DELETE("/:id", preferenceHandler.DeletePreference)
		}
	}

	return router
}
