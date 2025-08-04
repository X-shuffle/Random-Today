package repository

import (
	"context"
	"random_today/internal/domain/entity"
)

// PreferenceRepository 偏好仓储接口
type PreferenceRepository interface {
	// Create 创建偏好
	Create(ctx context.Context, preference *entity.Preference) error
	
	// GetByID 根据ID获取偏好
	GetByID(ctx context.Context, id string) (*entity.Preference, error)
	
	// GetByType 根据类型获取偏好列表
	GetByType(ctx context.Context, preferenceType entity.PreferenceType) ([]*entity.Preference, error)
	
	// GetAll 获取所有偏好
	GetAll(ctx context.Context) ([]*entity.Preference, error)
	
	// Update 更新偏好
	Update(ctx context.Context, preference *entity.Preference) error
	
	// Delete 删除偏好
	Delete(ctx context.Context, id string) error
	
	// DeleteByType 根据类型删除偏好
	DeleteByType(ctx context.Context, preferenceType entity.PreferenceType) error
	
	// GetRandomByType 根据类型随机获取一个偏好
	GetRandomByType(ctx context.Context, preferenceType entity.PreferenceType) (*entity.Preference, error)
} 