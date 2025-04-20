package service

import (
	"context"

	"github.com/undb/undb-go/internal/user/model"
)

// UserService defines the interface for user business logic
type UserService interface {
	// Register registers a new user
	Register(ctx context.Context, user *model.User) error
	// Login logs in a user
	Login(ctx context.Context, email, password string) (*model.User, error)
	// GetByID gets a user by ID
	GetByID(ctx context.Context, id string) (*model.User, error)
	// Update updates a user
	Update(ctx context.Context, user *model.User) error
	// Delete deletes a user by ID
	Delete(ctx context.Context, id string) error
}
