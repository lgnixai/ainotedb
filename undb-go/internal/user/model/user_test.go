package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr error
	}{
		{
			name: "valid user",
			user: &User{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "Test User",
			},
			wantErr: nil,
		},
		{
			name: "empty email",
			user: &User{
				Email:    "",
				Password: "password123",
				Name:     "Test User",
			},
			wantErr: ErrEmptyEmail,
		},
		{
			name: "empty password",
			user: &User{
				Email:    "test@example.com",
				Password: "",
				Name:     "Test User",
			},
			wantErr: ErrEmptyPassword,
		},
		{
			name: "empty name",
			user: &User{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "",
			},
			wantErr: ErrEmptyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	user := &User{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	err := user.BeforeCreate(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}

func TestUser_BeforeUpdate(t *testing.T) {
	user := &User{
		ID:        "user1",
		Email:     "test@example.com",
		Password:  "password123",
		Name:      "Test User",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now().Add(-24 * time.Hour),
	}

	oldUpdatedAt := user.UpdatedAt
	err := user.BeforeUpdate(nil)
	assert.NoError(t, err)
	assert.True(t, user.UpdatedAt.After(oldUpdatedAt))
	assert.Equal(t, user.CreatedAt, user.CreatedAt) // CreatedAt should not change
}
