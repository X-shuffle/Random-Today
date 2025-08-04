package service

import (
	"context"
	"random_today/internal/application/dto"
	"random_today/internal/domain/entity"
	"random_today/internal/domain/service"
)

// PreferenceAppService 偏好应用服务
type PreferenceAppService struct {
	domainService *service.PreferenceService
}

// NewPreferenceAppService 创建偏好应用服务实例
func NewPreferenceAppService(domainService *service.PreferenceService) *PreferenceAppService {
	return &PreferenceAppService{
		domainService: domainService,
	}
}

// CreatePreference 创建偏好
func (s *PreferenceAppService) CreatePreference(ctx context.Context, req dto.CreatePreferenceRequest) (*dto.CreatePreferenceResponse, error) {
	// 转换类型
	preferenceType := entity.PreferenceType(req.Type)
	
	// 调用领域服务
	preference, err := s.domainService.AddPreference(ctx, req.Name, preferenceType, req.Description)
	if err != nil {
		return nil, err
	}
	
	// 转换为响应DTO
	response := &dto.CreatePreferenceResponse{
		ID:          preference.ID,
		Name:        preference.Name,
		Type:        string(preference.Type),
		Description: preference.Description,
		CreatedAt:   preference.CreatedAt,
	}
	
	return response, nil
}

// GetRandomPreference 获取随机偏好
func (s *PreferenceAppService) GetRandomPreference(ctx context.Context, req dto.GetRandomPreferenceRequest) (*dto.GetRandomPreferenceResponse, error) {
	// 转换类型
	preferenceType := entity.PreferenceType(req.Type)
	
	// 调用领域服务
	preference, err := s.domainService.GetRandomPreference(ctx, preferenceType)
	if err != nil {
		return nil, err
	}
	
	// 转换为响应DTO
	response := &dto.GetRandomPreferenceResponse{
		ID:          preference.ID,
		Name:        preference.Name,
		Type:        string(preference.Type),
		Description: preference.Description,
		CreatedAt:   preference.CreatedAt,
	}
	
	return response, nil
}

// GetPreferencesByType 根据类型获取偏好列表
func (s *PreferenceAppService) GetPreferencesByType(ctx context.Context, req dto.GetPreferencesByTypeRequest) (*dto.GetPreferencesByTypeResponse, error) {
	// 转换类型
	preferenceType := entity.PreferenceType(req.Type)
	
	// 调用领域服务
	preferences, err := s.domainService.GetPreferencesByType(ctx, preferenceType)
	if err != nil {
		return nil, err
	}
	
	// 转换为响应DTO
	response := &dto.GetPreferencesByTypeResponse{
		Preferences: dto.ToPreferenceInfoList(preferences),
		Count:       len(preferences),
	}
	
	return response, nil
}

// GetAllPreferences 获取所有偏好
func (s *PreferenceAppService) GetAllPreferences(ctx context.Context) (*dto.GetAllPreferencesResponse, error) {
	// 调用领域服务
	preferences, err := s.domainService.GetAllPreferences(ctx)
	if err != nil {
		return nil, err
	}
	
	// 转换为响应DTO
	response := &dto.GetAllPreferencesResponse{
		Preferences: dto.ToPreferenceInfoList(preferences),
		Count:       len(preferences),
	}
	
	return response, nil
}

// DeletePreference 删除偏好
func (s *PreferenceAppService) DeletePreference(ctx context.Context, req dto.DeletePreferenceRequest) (*dto.DeletePreferenceResponse, error) {
	// 调用领域服务
	err := s.domainService.DeletePreference(ctx, req.ID)
	if err != nil {
		return &dto.DeletePreferenceResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	
	return &dto.DeletePreferenceResponse{
		Success: true,
		Message: "删除成功",
	}, nil
} 