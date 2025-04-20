package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undb/undb-go/internal/view/model"
)

type MockViewService struct {
	mock.Mock
}

func (m *MockViewService) CreateView(ctx *gin.Context, view *model.View) error {
	args := m.Called(ctx, view)
	return args.Error(0)
}

func (m *MockViewService) GetView(ctx *gin.Context, id string) (*model.View, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.View), args.Error(1)
}

func (m *MockViewService) GetViews(ctx *gin.Context, tableID string) ([]*model.View, error) {
	args := m.Called(ctx, tableID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.View), args.Error(1)
}

func (m *MockViewService) UpdateView(ctx *gin.Context, view *model.View) error {
	args := m.Called(ctx, view)
	return args.Error(0)
}

func (m *MockViewService) DeleteView(ctx *gin.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockViewService) UpdateViewConfig(ctx *gin.Context, id string, config interface{}) error {
	args := m.Called(ctx, id, config)
	return args.Error(0)
}

func setupTest() (*gin.Engine, *MockViewService, *ViewHandler) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockViewService)
	handler := NewViewHandler(mockService)
	return r, mockService, handler
}

func TestCreateView(t *testing.T) {
	r, mockService, handler := setupTest()
	r.POST("/views", handler.CreateView)

	view := &model.View{
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	mockService.On("CreateView", mock.Anything, view).Return(nil)

	body, _ := json.Marshal(view)
	req, _ := http.NewRequest(http.MethodPost, "/views", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetView(t *testing.T) {
	r, mockService, handler := setupTest()
	r.GET("/views/:id", handler.GetView)

	view := &model.View{
		ID:      "view1",
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	mockService.On("GetView", mock.Anything, "view1").Return(view, nil)

	req, _ := http.NewRequest(http.MethodGet, "/views/view1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetViews(t *testing.T) {
	r, mockService, handler := setupTest()
	r.GET("/views/table/:tableId", handler.GetViews)

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

	mockService.On("GetViews", mock.Anything, "table1").Return(views, nil)

	req, _ := http.NewRequest(http.MethodGet, "/views/table/table1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateView(t *testing.T) {
	r, mockService, handler := setupTest()
	r.PUT("/views/:id", handler.UpdateView)

	view := &model.View{
		ID:      "view1",
		Name:    "Updated View",
		TableID: "table1",
		Type:    model.ViewTypeKanban,
	}

	mockService.On("UpdateView", mock.Anything, view).Return(nil)

	body, _ := json.Marshal(view)
	req, _ := http.NewRequest(http.MethodPut, "/views/view1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteView(t *testing.T) {
	r, mockService, handler := setupTest()
	r.DELETE("/views/:id", handler.DeleteView)

	mockService.On("DeleteView", mock.Anything, "view1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/views/view1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateViewConfig(t *testing.T) {
	r, mockService, handler := setupTest()
	r.PUT("/views/:id/config", handler.UpdateViewConfig)

	config := map[string]interface{}{
		"sort": []map[string]interface{}{
			{
				"field": "name",
				"order": "asc",
			},
		},
	}

	mockService.On("UpdateViewConfig", mock.Anything, "view1", config).Return(nil)

	body, _ := json.Marshal(config)
	req, _ := http.NewRequest(http.MethodPut, "/views/view1/config", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
