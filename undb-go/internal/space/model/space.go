package model

import (
	"time"

	"github.com/undb/undb-go/pkg/utils"
	"gorm.io/gorm"
)

// SpaceVisibility 定义空间的可见性
type SpaceVisibility string

const (
	VisibilityPrivate SpaceVisibility = "private" // 私有空间
	VisibilityPublic  SpaceVisibility = "public"  // 公开空间
)

// Space 表示一个空间
type Space struct {
	ID          string          `gorm:"primaryKey" json:"id"`
	Name        string          `gorm:"not null" json:"name"`
	Description string          `json:"description"`
	OwnerID     string          `json:"owner_id" gorm:"index"`
	Visibility  SpaceVisibility `json:"visibility"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// TableName 指定表名
func (Space) TableName() string {
	return "spaces"
}

// BeforeCreate 创建前的钩子
func (s *Space) BeforeCreate(tx *gorm.DB) error {
	s.ID = utils.GenerateID("space")
	s.CreatedAt = utils.Now()
	s.UpdatedAt = utils.Now()
	return nil
}

// BeforeUpdate 更新前的钩子
func (s *Space) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = utils.Now()
	return nil
}

// Validate 验证空间数据
func (s *Space) Validate() error {
	if s.Name == "" {
		return ErrEmptySpaceName
	}
	if s.OwnerID == "" {
		return ErrEmptyOwnerID
	}
	if s.Visibility != VisibilityPrivate && s.Visibility != VisibilityPublic {
		return ErrInvalidVisibility
	}
	return nil
}
