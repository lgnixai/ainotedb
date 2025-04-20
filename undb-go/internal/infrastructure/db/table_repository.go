package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/undb/undb-go/internal/table/model"
)

// TableRepository 定义表格仓库接口
type TableRepository interface {
	// Create 创建新表格
	Create(ctx context.Context, table *model.Table) error

	// GetByID 根据ID获取表格
	GetByID(ctx context.Context, id uint) (*model.Table, error)

	// GetBySpaceID 获取空间的所有表格
	GetBySpaceID(ctx context.Context, spaceID uint) ([]*model.Table, error)

	// Update 更新表格
	Update(ctx context.Context, table *model.Table) error

	// Delete 删除表格
	Delete(ctx context.Context, id uint) error
}

// tableRepository 实现表格仓库接口
type tableRepository struct {
	db *gorm.DB
}

// NewTableRepository 创建新的表格仓库实例
func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{db: db}
}

func (r *tableRepository) Create(ctx context.Context, table *model.Table) error {
	return r.db.WithContext(ctx).Create(table).Error
}

func (r *tableRepository) GetByID(ctx context.Context, id uint) (*model.Table, error) {
	var table model.Table
	err := r.db.WithContext(ctx).First(&table, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrTableNotFound
		}
		return nil, err
	}
	return &table, nil
}

func (r *tableRepository) GetBySpaceID(ctx context.Context, spaceID uint) ([]*model.Table, error) {
	var tables []*model.Table
	err := r.db.WithContext(ctx).Where("space_id = ?", spaceID).Find(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *tableRepository) Update(ctx context.Context, table *model.Table) error {
	return r.db.WithContext(ctx).Save(table).Error
}

func (r *tableRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Table{}, id).Error
}
