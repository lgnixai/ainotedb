package repository

import (
	"context"

	"github.com/undb/undb-go/internal/user/model"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *model.User) error
	// GetByID gets a user by ID
	GetByID(ctx context.Context, id string) (*model.User, error)
	// GetByEmail gets a user by email
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	// Update updates a user
	Update(ctx context.Context, user *model.User) error
	// Delete deletes a user by ID
	Delete(ctx context.Context, id string) error
}
