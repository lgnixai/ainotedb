package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestView_Validate(t *testing.T) {
	tests := []struct {
		name    string
		view    *View
		wantErr error
	}{
		{
			name: "valid view",
			view: &View{
				Name:    "Test View",
				TableID: "table1",
				Type:    ViewTypeGrid,
			},
			wantErr: nil,
		},
		{
			name: "empty name",
			view: &View{
				Name:    "",
				TableID: "table1",
				Type:    ViewTypeGrid,
			},
			wantErr: ErrEmptyViewName,
		},
		{
			name: "empty table id",
			view: &View{
				Name:    "Test View",
				TableID: "",
				Type:    ViewTypeGrid,
			},
			wantErr: ErrEmptyTableID,
		},
		{
			name: "empty type",
			view: &View{
				Name:    "Test View",
				TableID: "table1",
				Type:    "",
			},
			wantErr: ErrEmptyViewType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.view.Validate()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestView_BeforeCreate(t *testing.T) {
	view := &View{
		Name:    "Test View",
		TableID: "table1",
		Type:    ViewTypeGrid,
	}

	err := view.BeforeCreate(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, view.ID)
	assert.NotZero(t, view.CreatedAt)
	assert.NotZero(t, view.UpdatedAt)
}

func TestView_BeforeUpdate(t *testing.T) {
	view := &View{
		ID:        "view1",
		Name:      "Test View",
		TableID:   "table1",
		Type:      ViewTypeGrid,
		CreatedAt: time.Now().Add(-1 * time.Hour),
		UpdatedAt: time.Now().Add(-1 * time.Hour),
	}

	oldUpdatedAt := view.UpdatedAt
	err := view.BeforeUpdate(nil)
	assert.NoError(t, err)
	assert.True(t, view.UpdatedAt.After(oldUpdatedAt))
}
