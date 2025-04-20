package service

import (
	"context"
	"errors"
	"time"

	`github.com/golang-jwt/jwt`
)

type AuthService struct {
	userService UserService
	jwtSecret   []byte
}

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func NewAuthService(userService UserService, jwtSecret string) *AuthService {
	return &AuthService{
		userService: userService,
		jwtSecret:   []byte(jwtSecret),
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userService.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !user.CheckPassword(password) {
		return "", ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidToken = errors.New("invalid token")
