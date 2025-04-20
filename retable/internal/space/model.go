package space

import (
	"time"
	"retable/internal/auth"
)

type Space struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SpaceMember struct {
	ID        string    `json:"id" gorm:"primarykey"`
	SpaceID   string    `json:"space_id"`
	UserID    string    `json:"user_id"`
	Role      auth.Role `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}