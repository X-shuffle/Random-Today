package service

import (
	"context"
	"errors"
	"random_today/internal/domain/entity"
	"random_today/internal/domain/repository"
)

// PreferenceService 偏好领域服务
type PreferenceService struct {
	repo repository.PreferenceRepository
}

// NewPreferenceService 创建偏好服务实例
func NewPreferenceService(repo repository.PreferenceRepository) *PreferenceService {
	return &PreferenceService{
		repo: repo,
	}
}

// AddPreference 添加偏好
func (s *PreferenceService) AddPreference(ctx context.Context, name string, preferenceType entity.PreferenceType, description string) (*entity.Preference, error) {
	// 验证输入
	if name == "" {
		return nil, errors.New("偏好名称不能为空")
	}
	
	if preferenceType == "" {
		return nil, errors.New("偏好类型不能为空")
	}
	
	// 创建偏好实体
	preference := entity.NewPreference(name, preferenceType, description)
	
	// 验证实体
	if !preference.IsValid() {
		return nil, errors.New("偏好实体无效")
	}
	
	// 保存到仓储
	if err := s.repo.Create(ctx, preference); err != nil {
		return nil, err
	}
	
	return preference, nil
}

// GetRandomPreference 随机获取偏好
func (s *PreferenceService) GetRandomPreference(ctx context.Context, preferenceType entity.PreferenceType) (*entity.Preference, error) {
	// 验证类型
	if preferenceType == "" {
		return nil, errors.New("偏好类型不能为空")
	}
	
	// 从仓储获取随机偏好
	preference, err := s.repo.GetRandomByType(ctx, preferenceType)
	if err != nil {
		return nil, err
	}
	
	if preference == nil {
		return nil, errors.New("该类型下没有偏好数据")
	}
	
	return preference, nil
}

// DeletePreference 删除偏好
func (s *PreferenceService) DeletePreference(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("偏好ID不能为空")
	}
	
	return s.repo.Delete(ctx, id)
}

// GetPreferencesByType 根据类型获取偏好列表
func (s *PreferenceService) GetPreferencesByType(ctx context.Context, preferenceType entity.PreferenceType) ([]*entity.Preference, error) {
	if preferenceType == "" {
		return nil, errors.New("偏好类型不能为空")
	}
	
	return s.repo.GetByType(ctx, preferenceType)
}

// GetAllPreferences 获取所有偏好
func (s *PreferenceService) GetAllPreferences(ctx context.Context) ([]*entity.Preference, error) {
	return s.repo.GetAll(ctx)
} 