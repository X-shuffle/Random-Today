package dto

import (
	"random_today/internal/domain/entity"
	"time"
)

// CreatePreferenceRequest 创建偏好请求
type CreatePreferenceRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
}

// CreatePreferenceResponse 创建偏好响应
type CreatePreferenceResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// GetRandomPreferenceRequest 获取随机偏好请求
type GetRandomPreferenceRequest struct {
	Type string `json:"type" validate:"required"`
}

// GetRandomPreferenceResponse 获取随机偏好响应
type GetRandomPreferenceResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// GetPreferencesByTypeRequest 根据类型获取偏好请求
type GetPreferencesByTypeRequest struct {
	Type string `json:"type" validate:"required"`
}

// GetPreferencesByTypeResponse 根据类型获取偏好响应
type GetPreferencesByTypeResponse struct {
	Preferences []PreferenceInfo `json:"preferences"`
	Count       int              `json:"count"`
}

// GetAllPreferencesResponse 获取所有偏好响应
type GetAllPreferencesResponse struct {
	Preferences []PreferenceInfo `json:"preferences"`
	Count       int              `json:"count"`
}

// PreferenceInfo 偏好信息
type PreferenceInfo struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DeletePreferenceRequest 删除偏好请求
type DeletePreferenceRequest struct {
	ID string `json:"id" validate:"required"`
}

// DeletePreferenceResponse 删除偏好响应
type DeletePreferenceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ToPreferenceInfo 将实体转换为DTO
func ToPreferenceInfo(preference *entity.Preference) PreferenceInfo {
	return PreferenceInfo{
		ID:          preference.ID,
		Name:        preference.Name,
		Type:        string(preference.Type),
		Description: preference.Description,
		CreatedAt:   preference.CreatedAt,
		UpdatedAt:   preference.UpdatedAt,
	}
}

// ToPreferenceInfoList 将实体列表转换为DTO列表
func ToPreferenceInfoList(preferences []*entity.Preference) []PreferenceInfo {
	result := make([]PreferenceInfo, len(preferences))
	for i, preference := range preferences {
		result[i] = ToPreferenceInfo(preference)
	}
	return result
} 