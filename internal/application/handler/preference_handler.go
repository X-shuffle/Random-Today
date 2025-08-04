package handler

import (
	"encoding/json"
	"net/http"
	"random_today/internal/application/dto"
	"random_today/internal/application/service"
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
func (h *PreferenceHandler) CreatePreference(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePreferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// 简单验证
	if req.Name == "" || req.Type == "" {
		http.Error(w, "Name and Type are required", http.StatusBadRequest)
		return
	}
	
	response, err := h.appService.CreatePreference(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetRandomPreference 获取随机偏好
func (h *PreferenceHandler) GetRandomPreference(w http.ResponseWriter, r *http.Request) {
	preferenceType := r.URL.Query().Get("type")
	if preferenceType == "" {
		http.Error(w, "Type parameter is required", http.StatusBadRequest)
		return
	}
	
	req := dto.GetRandomPreferenceRequest{Type: preferenceType}
	response, err := h.appService.GetRandomPreference(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetPreferencesByType 根据类型获取偏好列表
func (h *PreferenceHandler) GetPreferencesByType(w http.ResponseWriter, r *http.Request) {
	preferenceType := r.URL.Query().Get("type")
	if preferenceType == "" {
		http.Error(w, "Type parameter is required", http.StatusBadRequest)
		return
	}
	
	req := dto.GetPreferencesByTypeRequest{Type: preferenceType}
	response, err := h.appService.GetPreferencesByType(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllPreferences 获取所有偏好
func (h *PreferenceHandler) GetAllPreferences(w http.ResponseWriter, r *http.Request) {
	response, err := h.appService.GetAllPreferences(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeletePreference 删除偏好
func (h *PreferenceHandler) DeletePreference(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中获取ID，暂时使用查询参数
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}
	
	req := dto.DeletePreferenceRequest{ID: id}
	response, err := h.appService.DeletePreference(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	if response.Success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(response)
} 