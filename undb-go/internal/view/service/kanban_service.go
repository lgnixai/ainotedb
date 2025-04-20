
package service

import (
	"context"
	"github.com/undb/undb-go/internal/view/model"
	"github.com/undb/undb-go/internal/record/model"
)

type KanbanService interface {
	GetKanbanData(ctx context.Context, viewID string) (map[string][]*record.Record, error)
	UpdateCardPosition(ctx context.Context, cardID string, targetGroup string, position int) error
}

type kanbanService struct {
	viewRepo   ViewRepository
	recordRepo record.RecordRepository
}

func NewKanbanService(viewRepo ViewRepository, recordRepo record.RecordRepository) KanbanService {
	return &kanbanService{
		viewRepo:   viewRepo,
		recordRepo: recordRepo,
	}
}

func (s *kanbanService) GetKanbanData(ctx context.Context, viewID string) (map[string][]*record.Record, error) {
	view, err := s.viewRepo.GetByID(ctx, viewID)
	if err != nil {
		return nil, err
	}

	kanbanView, ok := view.(*model.KanbanView)
	if !ok {
		return nil, ErrInvalidViewType
	}

	records, err := s.recordRepo.GetByTableID(ctx, view.GetTableID())
	if err != nil {
		return nil, err
	}

	// Group records by the kanban group field
	groups := make(map[string][]*record.Record)
	for _, record := range records {
		groupValue := record.Data[kanbanView.GroupField].(string)
		groups[groupValue] = append(groups[groupValue], record)
	}

	return groups, nil
}

func (s *kanbanService) UpdateCardPosition(ctx context.Context, cardID string, targetGroup string, position int) error {
	// Implementation for updating card position
	// This would involve updating the record's group field value and possibly a position field
	return nil
}
