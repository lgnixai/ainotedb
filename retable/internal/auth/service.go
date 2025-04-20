package auth

import (
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role string

const (
	RoleOwner   Role = "owner"
	RoleAdmin   Role = "admin"
	RoleMember  Role = "member"
	RoleVisitor Role = "visitor"
)


type UserRepository interface {
	ExistsByEmail(email string) (bool, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	GetUserRole(userID string, spaceID string) (Role, error)
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type AuthService struct {
	jwtSecret []byte
	userRepo  UserRepository
}

func NewAuthService(secret string, userRepo UserRepository) *AuthService {
	return &AuthService{
		jwtSecret: []byte(secret),
		userRepo:  userRepo,
	}
}

func (s *AuthService) Register(email string, password string) (*User, error) {
	// Check if user exists
	exists, err := s.userRepo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &User{
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email string, password string) (string, error) {
	// Get user
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (s *AuthService) GetUserRole(userID string, spaceID string) (Role, error) {
	return s.userRepo.GetUserRole(userID, spaceID)
}

func (s *AuthService) HasPermission(userID string, spaceID string, action string) bool {
	role, err := s.GetUserRole(userID, spaceID)
	if err != nil {
		return false
	}

	switch role {
	case RoleOwner:
		return true
	case RoleAdmin:
		return action != "delete_space"
	case RoleMember:
		return action == "read" || action == "create_record" || action == "update_record"
	default:
		return false
	}
}