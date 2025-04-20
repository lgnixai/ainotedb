
package space

import (
	"time"
)

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
)

type Space struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SpaceMember struct {
	ID        string    `json:"id" gorm:"primarykey"`
	SpaceID   string    `json:"space_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Role      Role      `json:"role" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Space) Validate() error {
	if s.Name == "" {
		return errors.New("space name is required")
	}
	if s.OwnerID == "" {
		return errors.New("owner ID is required")
	}
	return nil
}
