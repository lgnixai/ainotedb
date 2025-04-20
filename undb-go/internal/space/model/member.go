
package model

import (
	"time"

	"github.com/undb/undb-go/pkg/utils"
	"gorm.io/gorm"
)

// MemberRole 定义成员角色
type MemberRole string

const (
	RoleOwner  MemberRole = "owner"  // 所有者
	RoleAdmin  MemberRole = "admin"  // 管理员
	RoleMember MemberRole = "member" // 普通成员
	RoleEditor MemberRole = "editor" // 编辑者
)

// SpaceMember 表示空间成员
type SpaceMember struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	SpaceID   string     `gorm:"index" json:"space_id"`
	UserID    string     `gorm:"index" json:"user_id"`
	Role      MemberRole `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TableName 指定表名
func (SpaceMember) TableName() string {
	return "space_members"
}

// BeforeCreate 创建前的钩子
func (m *SpaceMember) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = utils.GenerateID("mem")
	}
	m.CreatedAt = utils.Now()
	m.UpdatedAt = utils.Now()
	return nil
}

// BeforeUpdate 更新前的钩子
func (m *SpaceMember) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = utils.Now()
	return nil
}

// Validate 验证成员数据
func (m *SpaceMember) Validate() error {
	if m.SpaceID == "" {
		return ErrEmptySpaceID
	}
	if m.UserID == "" {
		return ErrEmptyUserID
	}
	if m.Role == "" {
		return ErrEmptyRole
	}
	return nil
}
