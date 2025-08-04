package handler

import (
	"net/http"
	"random_today/internal/application/dto"
	"random_today/internal/application/service"

	"github.com/gin-gonic/gin"
)

// PreferenceHandler 偏好处理器
type PreferenceHandler struct {
	appService *service.PreferenceAppService
}

// NewPreferenceHandler 创建偏好处理器实例
func NewPreferenceHandler(appService *service.PreferenceAppService) *PreferenceHandler {
	return &PreferenceHandler{
		appService: appService,
	}
}

// CreatePreference 创建偏好
func (h *PreferenceHandler) CreatePreference(c *gin.Context) {
	var req dto.CreatePreferenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	
	// 简单验证
	if req.Name == "" || req.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Type are required"})
		return
	}
	
	response, err := h.appService.CreatePreference(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}

// GetRandomPreference 获取随机偏好
func (h *PreferenceHandler) GetRandomPreference(c *gin.Context) {
	preferenceType := c.Query("type")
	if preferenceType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type parameter is required"})
		return
	}
	
	req := dto.GetRandomPreferenceRequest{Type: preferenceType}
	response, err := h.appService.GetRandomPreference(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

// GetPreferencesByType 根据类型获取偏好列表
func (h *PreferenceHandler) GetPreferencesByType(c *gin.Context) {
	preferenceType := c.Query("type")
	if preferenceType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type parameter is required"})
		return
	}
	
	req := dto.GetPreferencesByTypeRequest{Type: preferenceType}
	response, err := h.appService.GetPreferencesByType(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

// GetAllPreferences 获取所有偏好
func (h *PreferenceHandler) GetAllPreferences(c *gin.Context) {
	response, err := h.appService.GetAllPreferences(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

// DeletePreference 删除偏好
func (h *PreferenceHandler) DeletePreference(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}
	
	req := dto.DeletePreferenceRequest{ID: id}
	response, err := h.appService.DeletePreference(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if response.Success {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusBadRequest, response)
	}
} 