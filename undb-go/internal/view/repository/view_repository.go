package repository

import (
	"context"

	"github.com/undb/undb-go/internal/view/model"
	"gorm.io/gorm"
)

// ViewRepository defines the interface for view data access
type ViewRepository interface {
	// Create creates a new view
	Create(ctx context.Context, view *model.View) error
	// GetByID gets a view by ID
	GetByID(ctx context.Context, id string) (*model.View, error)
	// GetByTableID gets all views for a table
	GetByTableID(ctx context.Context, tableID string) ([]*model.View, error)
	// Update updates a view
	Update(ctx context.Context, view *model.View) error
	// Delete deletes a view by ID
	Delete(ctx context.Context, id string) error
}

// viewRepository implements ViewRepository
type viewRepository struct {
	db *gorm.DB
}

// NewViewRepository creates a new view repository
func NewViewRepository(db *gorm.DB) ViewRepository {
	return &viewRepository{db: db}
}

// Create creates a new view
func (r *viewRepository) Create(ctx context.Context, view *model.View) error {
	return r.db.WithContext(ctx).Create(view).Error
}

// GetByID gets a view by ID
func (r *viewRepository) GetByID(ctx context.Context, id string) (*model.View, error) {
	var view model.View
	err := r.db.WithContext(ctx).First(&view, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &view, nil
}

// GetByTableID gets all views for a table
func (r *viewRepository) GetByTableID(ctx context.Context, tableID string) ([]*model.View, error) {
	var views []*model.View
	err := r.db.WithContext(ctx).Where("table_id = ?", tableID).Find(&views).Error
	if err != nil {
		return nil, err
	}
	return views, nil
}

// Update updates a view
func (r *viewRepository) Update(ctx context.Context, view *model.View) error {
	return r.db.WithContext(ctx).Save(view).Error
}

// Delete deletes a view by ID
func (r *viewRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.View{}, "id = ?", id).Error
}
