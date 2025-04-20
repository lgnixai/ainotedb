package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	Email         string    `json:"email" gorm:"unique"`
	Username      string    `json:"username"`
	Password      string    `json:"-"` // 不序列化密码
	EmailVerified bool      `json:"email_verified"`
	Avatar        string    `json:"avatar"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// TableName returns the table name
func (User) TableName() string {
	return "users"
}

// BeforeCreate is called before creating a user
func generateUUID() string {
	return uuid.New().String()
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate called for user:", u.Email)
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	if u.ID == "" {
		u.ID = "usr_" + generateUUID()
	}
	return u.HashPassword() //Hash password before saving
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
	if u.Username == "" {
		return ErrEmptyUsername // Added validation for username
	}
	return nil
}

// Error definitions (assuming these are defined elsewhere, adjust as needed)
var (
	//ErrEmptyEmail     = &Error{Code: "emptyEmail", Message: "Email cannot be empty"}
	//ErrEmptyPassword  = &Error{Code: "emptyPassword", Message: "Password cannot be empty"}
	ErrEmptyUsername = &Error{Code: "emptyUsername", Message: "Username cannot be empty"}
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
