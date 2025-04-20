package repository

import (
	"context"

	"github.com/undb/undb-go/internal/field/model"
	"gorm.io/gorm"
)

// FieldRepository 定义字段仓库接口
type FieldRepository interface {
	// Create 创建新字段
	Create(ctx context.Context, field *model.Field) error

	// GetByID 根据ID获取字段
	GetByID(ctx context.Context, id uint) (*model.Field, error)

	// GetByTableID 获取表的所有字段
	GetByTableID(ctx context.Context, tableID uint) ([]*model.Field, error)

	// Update 更新字段
	Update(ctx context.Context, field *model.Field) error

	// Delete 删除字段
	Delete(ctx context.Context, id uint) error
}

// fieldRepository 实现字段仓库接口
type fieldRepository struct {
	db *gorm.DB
}

// NewFieldRepository 创建新的字段仓库实例
func NewFieldRepository(db *gorm.DB) FieldRepository {
	return &fieldRepository{db: db}
}

func (r *fieldRepository) Create(ctx context.Context, field *model.Field) error {
	return r.db.WithContext(ctx).Create(field).Error
}

func (r *fieldRepository) GetByID(ctx context.Context, id uint) (*model.Field, error) {
	var field model.Field
	err := r.db.WithContext(ctx).First(&field, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrFieldNotFound
		}
		return nil, err
	}
	return &field, nil
}

func (r *fieldRepository) GetByTableID(ctx context.Context, tableID uint) ([]*model.Field, error) {
	var fields []*model.Field
	err := r.db.WithContext(ctx).Where("table_id = ?", tableID).Find(&fields).Error
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (r *fieldRepository) Update(ctx context.Context, field *model.Field) error {
	return r.db.WithContext(ctx).Save(field).Error
}

func (r *fieldRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Field{}, id).Error
}
