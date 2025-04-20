package service

import (
	"context"

	"github.com/undb/undb-go/internal/field/model"
	repository `github.com/undb/undb-go/internal/infrastructure/db`
)

// FieldService 定义字段服务接口
type FieldService interface {
	// Create 创建新字段
	Create(ctx context.Context, field *model.Field) error

	// GetByID 根据ID获取字段
	GetByID(ctx context.Context, id uint) (*model.Field, error)

	// GetByTableID 获取表格的所有字段
	GetByTableID(ctx context.Context, tableID uint) ([]*model.Field, error)

	// Update 更新字段
	Update(ctx context.Context, field *model.Field) error

	// Delete 删除字段
	Delete(ctx context.Context, id uint) error
}

// fieldService 实现字段服务接口
type fieldService struct {
	repo repository.FieldRepository
}

// NewFieldService 创建新的字段服务实例
func NewFieldService(repo repository.FieldRepository) FieldService {
	return &fieldService{repo: repo}
}

func (s *fieldService) Create(ctx context.Context, field *model.Field) error {
	if err := field.Validate(); err != nil {
		return err
	}
	return s.repo.Create(ctx, field)
}

func (s *fieldService) GetByID(ctx context.Context, id uint) (*model.Field, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *fieldService) GetByTableID(ctx context.Context, tableID uint) ([]*model.Field, error) {
	return s.repo.GetByTableID(ctx, tableID)
}

func (s *fieldService) Update(ctx context.Context, field *model.Field) error {
	if err := field.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, field)
}

func (s *fieldService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
