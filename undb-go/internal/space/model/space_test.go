package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undb/undb-go/pkg/utils"
)

func TestSpace_Validate(t *testing.T) {
	tests := []struct {
		name    string
		space   *Space
		wantErr error
	}{
		{
			name: "valid space",
			space: &Space{
				ID:          1,
				Name:        "Test Space",
				Description: "Test Description",
				OwnerID:     "user1",
				Visibility:  VisibilityPrivate,
			},
			wantErr: nil,
		},
		{
			name: "empty name",
			space: &Space{
				ID:          1,
				Name:        "",
				Description: "Test Description",
				OwnerID:     "user1",
				Visibility:  VisibilityPrivate,
			},
			wantErr: ErrEmptySpaceName,
		},
		{
			name: "empty owner id",
			space: &Space{
				ID:          1,
				Name:        "Test Space",
				Description: "Test Description",
				OwnerID:     "",
				Visibility:  VisibilityPrivate,
			},
			wantErr: ErrEmptyOwnerID,
		},
		{
			name: "invalid visibility",
			space: &Space{
				ID:          1,
				Name:        "Test Space",
				Description: "Test Description",
				OwnerID:     "user1",
				Visibility:  "invalid",
			},
			wantErr: ErrInvalidVisibility,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.space.Validate()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestSpace_BeforeCreate(t *testing.T) {
	space := &Space{
		ID:          1,
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  VisibilityPrivate,
	}

	err := space.BeforeCreate(nil)
	assert.NoError(t, err)
	assert.NotZero(t, space.CreatedAt)
	assert.NotZero(t, space.UpdatedAt)
}

func TestSpace_BeforeUpdate(t *testing.T) {
	space := &Space{
		ID:          1,
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  VisibilityPrivate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	oldUpdatedAt := space.UpdatedAt
	err := space.BeforeUpdate(nil)
	assert.NoError(t, err)
	assert.NotEqual(t, oldUpdatedAt, space.UpdatedAt)
}
