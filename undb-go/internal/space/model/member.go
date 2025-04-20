package model

import (
	"time"

	"github.com/undb/undb-go/pkg/utils"
	"gorm.io/gorm"
)

// MemberRole represents the role of a member in a space
type MemberRole string

const (
	MemberRoleOwner  MemberRole = "owner"
	MemberRoleAdmin  MemberRole = "admin"
	MemberRoleMember MemberRole = "member"
)

// Member represents a member in a space
type Member struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	SpaceID   string     `json:"space_id" gorm:"index"`
	UserID    string     `json:"user_id" gorm:"index"`
	Role      MemberRole `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TableName 指定表名
func (Member) TableName() string {
	return "space_members"
}

// BeforeCreate is a GORM hook that runs before creating a new member
func (m *Member) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = utils.GenerateID("mem")
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating a member
func (m *Member) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

// Validate 验证成员数据
func (m *Member) Validate() error {
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
