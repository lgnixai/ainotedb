package service

import (
	"context"
	"fmt"
	"log"

	`github.com/undb/undb-go/internal/view/model`
	`github.com/undb/undb-go/internal/view/repository`
)

// ViewService defines operations for managing views
type ViewService interface {
	CreateView(ctx context.Context, view *model.View) error
	GetView(ctx context.Context, id string) (*model.View, error)
	GetViews(ctx context.Context, tableID string) ([]*model.View, error)
	UpdateView(ctx context.Context, view *model.View) error
	DeleteView(ctx context.Context, id string) error
	UpdateViewConfig(ctx context.Context, id string, config interface{}) error
}

type viewServiceImpl struct {
	repo repository.ViewRepository
}

// NewViewService creates a new view service instance
func NewViewService(repo repository.ViewRepository) ViewService {
	return &viewServiceImpl{
		repo: repo,
	}
}

func (s *viewServiceImpl) CreateView(ctx context.Context, view *model.View) error {
	if err := view.Validate(); err != nil {
		return fmt.Errorf("invalid view: %w", err)
	}

	if err := s.repo.Create(ctx, view); err != nil {
		log.Printf("Failed to create view: %v", err)
		return fmt.Errorf("failed to create view: %w", err)
	}

	return nil
}

func (s *viewServiceImpl) GetView(ctx context.Context, id string) (*model.View, error) {
	view, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Printf("Failed to get view: %v", err)
		return nil, fmt.Errorf("failed to get view: %w", err)
	}

	if view == nil {
		return nil, model.ErrViewNotFound
	}

	return view, nil
}

func (s *viewServiceImpl) GetViews(ctx context.Context, tableID string) ([]*model.View, error) {
	views, err := s.repo.FindByTableID(ctx, tableID)
	if err != nil {
		log.Printf("Failed to get views: %v", err)
		return nil, fmt.Errorf("failed to get views: %w", err)
	}

	return views, nil
}

func (s *viewServiceImpl) UpdateView(ctx context.Context, view *model.View) error {
	if err := view.Validate(); err != nil {
		return fmt.Errorf("invalid view: %w", err)
	}

	if err := s.repo.Update(ctx, view); err != nil {
		log.Printf("Failed to update view: %v", err)
		return fmt.Errorf("failed to update view: %w", err)
	}

	return nil
}

func (s *viewServiceImpl) DeleteView(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Printf("Failed to delete view: %v", err)
		return fmt.Errorf("failed to delete view: %w", err)
	}

	return nil
}

func (s *viewServiceImpl) UpdateViewConfig(ctx context.Context, id string, config interface{}) error {
	view, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Printf("Failed to get view for config update: %v", err)
		return fmt.Errorf("failed to get view: %w", err)
	}

	if view == nil {
		return model.ErrViewNotFound
	}

	if err := view.UpdateConfig(config); err != nil {
		log.Printf("Failed to update view config: %v", err)
		return fmt.Errorf("failed to update view config: %w", err)
	}

	if err := s.repo.Update(ctx, view); err != nil {
		log.Printf("Failed to save view with updated config: %v", err)
		return fmt.Errorf("failed to save view: %w", err)
	}

	return nil
}
