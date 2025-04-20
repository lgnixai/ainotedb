package service

import (
	"context"

	"github.com/undb/undb-go/internal/table/model"
	"github.com/undb/undb-go/internal/table/repository"
)

// TableService 定义表格服务接口
type TableService interface {
	// Create 创建新表格
	Create(ctx context.Context, table *model.Table) error

	// GetByID 根据ID获取表格
	GetByID(ctx context.Context, id string) (*model.Table, error)

	// GetBySpaceID 获取空间的所有表格
	GetBySpaceID(ctx context.Context, spaceID string) ([]*model.Table, error)

	// Update 更新表格
	Update(ctx context.Context, table *model.Table) error

	// Delete 删除表格
	Delete(ctx context.Context, id string) error
}

// tableService 实现表格服务接口
type tableService struct {
	repo repository.TableRepository
}

// NewTableService 创建新的表格服务实例
func NewTableService(repo repository.TableRepository) TableService {
	return &tableService{repo: repo}
}

func (s *tableService) Create(ctx context.Context, table *model.Table) error {
	if err := table.Validate(); err != nil {
		return err
	}
	return s.repo.Create(ctx, table)
}

func (s *tableService) GetByID(ctx context.Context, id string) (*model.Table, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *tableService) GetBySpaceID(ctx context.Context, spaceID string) ([]*model.Table, error) {
	return s.repo.GetBySpaceID(ctx, spaceID)
}

func (s *tableService) Update(ctx context.Context, table *model.Table) error {
	if err := table.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, table)
}

func (s *tableService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
