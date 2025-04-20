package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/undb/undb-go/internal/view/model"
)

// ViewRepository defines the interface for view data access
type ViewRepository interface {
	Create(ctx context.Context, view *model.View) error
	FindByID(ctx context.Context, id string) (*model.View, error)
	FindByTableID(ctx context.Context, tableID string) ([]*model.View, error)
	Update(ctx context.Context, view *model.View) error
	Delete(ctx context.Context, id string) error
}

type viewRepository struct {
	db *gorm.DB
}

func (r *viewRepository) FindByID(ctx context.Context, id string) (*model.View, error) {
	var view model.View
	err := r.db.WithContext(ctx).First(&view, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &view, nil
}

// NewViewRepository creates a new view repository instance
func NewViewRepository(db *gorm.DB) ViewRepository {
	return &viewRepository{db: db}
}

// Create creates a new view
func (r *viewRepository) Create(ctx context.Context, view *model.View) error {
	return r.db.WithContext(ctx).Create(view).Error
}

// FindByID gets a view by ID
func (r *viewRepository) GetByID(ctx context.Context, id string) (*model.View, error) {
	//TODO implement me
	panic("implement me")
}

// FindByTableID gets all views for a table
func (r *viewRepository) FindByTableID(ctx context.Context, tableID string) ([]*model.View, error) {
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
