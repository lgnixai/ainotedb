package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undb/undb-go/internal/view/model"
)

type MockViewRepository struct {
	mock.Mock
}

func (m *MockViewRepository) Create(ctx context.Context, view *model.View) error {
	args := m.Called(ctx, view)
	return args.Error(0)
}

func (m *MockViewRepository) GetByID(ctx context.Context, id string) (*model.View, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.View), args.Error(1)
}

func (m *MockViewRepository) GetByTableID(ctx context.Context, tableID string) ([]*model.View, error) {
	args := m.Called(ctx, tableID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.View), args.Error(1)
}

func (m *MockViewRepository) Update(ctx context.Context, view *model.View) error {
	args := m.Called(ctx, view)
	return args.Error(0)
}

func (m *MockViewRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestViewService_CreateView(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	view := &model.View{
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	mockRepo.On("Create", mock.Anything, view).Return(nil)

	err := service.CreateView(context.Background(), view)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestViewService_GetView(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	view := &model.View{
		ID:      "view1",
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	mockRepo.On("GetByID", mock.Anything, "view1").Return(view, nil)

	result, err := service.GetView(context.Background(), "view1")
	assert.NoError(t, err)
	assert.Equal(t, view, result)
	mockRepo.AssertExpectations(t)
}

func TestViewService_GetViews(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	views := []*model.View{
		{
			ID:      "view1",
			Name:    "View 1",
			TableID: "table1",
			Type:    model.ViewTypeGrid,
		},
		{
			ID:      "view2",
			Name:    "View 2",
			TableID: "table1",
			Type:    model.ViewTypeKanban,
		},
	}

	mockRepo.On("GetByTableID", mock.Anything, "table1").Return(views, nil)

	result, err := service.GetViews(context.Background(), "table1")
	assert.NoError(t, err)
	assert.Equal(t, views, result)
	mockRepo.AssertExpectations(t)
}

func TestViewService_UpdateView(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	view := &model.View{
		ID:      "view1",
		Name:    "Updated View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	mockRepo.On("Update", mock.Anything, view).Return(nil)

	err := service.UpdateView(context.Background(), view)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestViewService_DeleteView(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	mockRepo.On("Delete", mock.Anything, "view1").Return(nil)

	err := service.DeleteView(context.Background(), "view1")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestViewService_UpdateViewConfig(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	view := &model.View{
		ID:      "view1",
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	config := &model.GridViewConfig{
		ViewConfig: model.ViewConfig{Type: model.ViewTypeGrid},
		Sort: []model.SortConfig{
			{Field: "name", Order: "asc"},
		},
		Filter: []model.FilterConfig{
			{Field: "status", Operator: "eq", Value: "active"},
		},
		Fields: []string{"field1", "field2"},
	}

	mockRepo.On("GetByID", mock.Anything, "view1").Return(view, nil)
	mockRepo.On("Update", mock.Anything, view).Return(nil)

	err := service.UpdateViewConfig(context.Background(), "view1", config)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestViewService_ApplyView(t *testing.T) {
	mockRepo := new(MockViewRepository)
	service := NewViewService(mockRepo)

	view := &model.View{
		ID:      "view1",
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
		Config: `{
			"type": "grid",
			"sort": [{"field": "name", "order": "asc"}],
			"filter": [{"field": "status", "operator": "eq", "value": "active"}],
			"fields": ["field1", "field2"]
		}`,
	}

	records := []map[string]interface{}{
		{
			"field1": "value1",
			"field2": "value2",
			"name":   "record1",
			"status": "active",
		},
		{
			"field1": "value3",
			"field2": "value4",
			"name":   "record2",
			"status": "inactive",
		},
	}

	mockRepo.On("GetByID", mock.Anything, "view1").Return(view, nil)

	result, err := service.ApplyView(context.Background(), "view1", records)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}
