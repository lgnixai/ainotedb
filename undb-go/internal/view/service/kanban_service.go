package service

import (
	"context"

	repository `github.com/undb/undb-go/internal/infrastructure/db`
	"github.com/undb/undb-go/internal/record/model"
)

type KanbanService interface {
	GetKanbanData(ctx context.Context, viewID string) (map[string][]*model.Record, error)
	UpdateCardPosition(ctx context.Context, cardID string, targetGroup string, position int) error
}

type kanbanService struct {
	viewRepo   repository.ViewRepository
	recordRepo repository.RecordRepository
}

func NewKanbanService(viewRepo repository.ViewRepository, recordRepo repository.RecordRepository) KanbanService {
	return &kanbanService{
		viewRepo:   viewRepo,
		recordRepo: recordRepo,
	}
}

func (s *kanbanService) GetKanbanData(ctx context.Context, viewID string) (map[string][]*model.Record, error) {
	view, err := s.viewRepo.GetByID(ctx, viewID)
	if err != nil {
		return nil, err
	}

	// 这里只能断言为 *model.View，不能断言为 *model.KanbanView
	// kanbanView := view  // 如果有 KanbanView 结构体，需要在 model 中定义
	// TODO: 这里如需特殊 kanbanView 字段，请在 model.View 中定义或做类型扩展

	records, err := s.recordRepo.GetByTableID(ctx, view.TableID)
	if err != nil {
		return nil, err
	}

	// Group records by a group field, 这里假设 group field 名为 "group"
	groups := make(map[string][]*model.Record)
	for _, record := range records {
		groupValue, _ := record.Data["group"].(string)
		groups[groupValue] = append(groups[groupValue], record)
	}

	return groups, nil
}

func (s *kanbanService) UpdateCardPosition(ctx context.Context, cardID string, targetGroup string, position int) error {
	// Implementation for updating card position
	// This would involve updating the record's group field value and possibly a position field
	return nil
}
