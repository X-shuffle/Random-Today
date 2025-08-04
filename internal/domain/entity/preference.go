package entity

import (
	"time"
)

// PreferenceType 偏好类型
type PreferenceType string

const (
	PreferenceTypeMainFood PreferenceType = "main_food" // 主食
	PreferenceTypeSnack    PreferenceType = "snack"     // 零食
	PreferenceTypeScenery  PreferenceType = "scenery"   // 风景
	PreferenceTypeMilkTea  PreferenceType = "milk_tea"  // 奶茶
)

// Preference 偏好实体
type Preference struct {
	ID          string         `json:"id" bson:"_id,omitempty"`
	Name        string         `json:"name" bson:"name"`
	Type        PreferenceType `json:"type" bson:"type"`
	Description string         `json:"description" bson:"description,omitempty"`
	CreatedAt   time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" bson:"updated_at"`
}

// NewPreference 创建新的偏好实体
func NewPreference(name string, preferenceType PreferenceType, description string) *Preference {
	now := time.Now()
	return &Preference{
		Name:        name,
		Type:        preferenceType,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateDescription 更新描述
func (p *Preference) UpdateDescription(description string) {
	p.Description = description
	p.UpdatedAt = time.Now()
}

// IsValid 验证偏好实体是否有效
func (p *Preference) IsValid() bool {
	return p.Name != "" && p.Type != ""
} 