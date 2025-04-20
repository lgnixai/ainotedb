package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/undb/undb-go/internal/user/model"
	"github.com/undb/undb-go/internal/user/repository"
)

// userService implements UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Register registers a new user
func (s *userService) Register(ctx context.Context, user *model.User) error {

	// Check if user already exists
	_, err := s.repo.GetByEmail(ctx, user.Email)
	if err == nil {
		return model.ErrEmailAlreadyExists
	}
	if !errors.Is(err, model.ErrUserNotFound) {
		return err
	}
	// 不要再手动 hash 密码了，直接保存
	return s.repo.Create(ctx, user)

}

// Login logs in a user
func (s *userService) Login(ctx context.Context, email, password string) (*model.User, error) {
	// Get user by email
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, model.ErrInvalidPassword
	}

	return user, nil
}

// GetByID gets a user by ID
func (s *userService) GetByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

// Update updates a user
func (s *userService) Update(ctx context.Context, user *model.User) error {
	return s.repo.Update(ctx, user)
}

// Delete deletes a user by ID
func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// GetUserByEmail gets a user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
