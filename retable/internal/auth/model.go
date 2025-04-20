package auth

import (
	"time"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        string    `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
)