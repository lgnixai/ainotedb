package repository

import (
	"context"

	"github.com/undb/undb-go/internal/space/model"
	"gorm.io/gorm"
)

// SpaceRepository 空间仓库接口
type SpaceRepository interface {
	Create(ctx context.Context, space *model.Space) error
	GetByID(ctx context.Context, id string) (*model.Space, error)
	GetByOwnerID(ctx context.Context, ownerID string) ([]*model.Space, error)
	List(ctx context.Context) ([]*model.Space, error)
	Update(ctx context.Context, space *model.Space) error
	Delete(ctx context.Context, id string) error
}

// spaceRepository 空间仓库实现
type spaceRepository struct {
	db *gorm.DB
}

// NewSpaceRepository 创建空间仓库实例
func NewSpaceRepository(db *gorm.DB) SpaceRepository {
	return &spaceRepository{db: db}
}

// Create 创建空间
func (r *spaceRepository) Create(ctx context.Context, space *model.Space) error {
	return r.db.WithContext(ctx).Create(space).Error
}

// GetByID 根据ID获取空间
func (r *spaceRepository) GetByID(ctx context.Context, id string) (*model.Space, error) {
	var space model.Space
	if err := r.db.WithContext(ctx).First(&space, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrSpaceNotFound
		}
		return nil, err
	}
	return &space, nil
}

// GetByOwnerID 根据所有者ID获取空间
func (r *spaceRepository) GetByOwnerID(ctx context.Context, ownerID string) ([]*model.Space, error) {
	var spaces []*model.Space
	if err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&spaces).Error; err != nil {
		return nil, err
	}
	return spaces, nil
}

// List 获取空间列表
func (r *spaceRepository) List(ctx context.Context) ([]*model.Space, error) {
	var spaces []*model.Space
	if err := r.db.WithContext(ctx).Find(&spaces).Error; err != nil {
		return nil, err
	}
	return spaces, nil
}

// Update 更新空间
func (r *spaceRepository) Update(ctx context.Context, space *model.Space) error {
	return r.db.WithContext(ctx).Save(space).Error
}

// Delete 删除空间
func (r *spaceRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Space{}, "id = ?", id).Error
}
