package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName returns the table name
func (User) TableName() string {
	return "users"
}

// BeforeCreate is called before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate is called before updating a user
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

// Validate validates the user data
func (u *User) Validate() error {
	if u.Email == "" {
		return ErrEmptyEmail
	}
	if u.Password == "" {
		return ErrEmptyPassword
	}
	if u.Name == "" {
		return ErrEmptyName
	}
	return nil
}
