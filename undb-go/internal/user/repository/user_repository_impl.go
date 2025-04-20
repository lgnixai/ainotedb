package repository

import (
	"context"

	"github.com/undb/undb-go/internal/user/model"
	"gorm.io/gorm"
)

// userRepository implements UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetByID gets a user by ID
func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail gets a user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete deletes a user by ID
func (r *userRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}
