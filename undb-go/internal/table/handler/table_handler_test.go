package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undb/undb-go/internal/table/model"
)

// MockTableService 是 TableService 的模拟实现
type MockTableService struct {
	mock.Mock
}

func (m *MockTableService) Create(ctx context.Context, table *model.Table) error {
	args := m.Called(ctx, table)
	return args.Error(0)
}

func (m *MockTableService) GetByID(ctx context.Context, id string) (*model.Table, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Table), args.Error(1)
}

func (m *MockTableService) GetBySpaceID(ctx context.Context, spaceID string) ([]*model.Table, error) {
	args := m.Called(ctx, spaceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Table), args.Error(1)
}

func (m *MockTableService) Update(ctx context.Context, table *model.Table) error {
	args := m.Called(ctx, table)
	return args.Error(0)
}

func (m *MockTableService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func setupTest() (*gin.Engine, *MockTableService) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockTableService)
	handler := NewTableHandler(mockService)

	group := r.Group("/api")
	group.POST("/tables", handler.CreateTable)
	group.GET("/tables/:id", handler.GetTable)
	group.GET("/tables/space/:space_id", handler.GetTablesBySpace)
	group.PUT("/tables/:id", handler.UpdateTable)
	group.DELETE("/tables/:id", handler.DeleteTable)

	return r, mockService
}

func TestCreateTable(t *testing.T) {
	r, mockService := setupTest()

	table := &model.Table{
		Name:        "Test Table",
		Description: "Test Description",
		SpaceID:     "space1",
	}

	mockService.On("Create", mock.Anything, mock.AnythingOfType("*model.Table")).Return(nil)

	body, _ := json.Marshal(table)
	req := httptest.NewRequest(http.MethodPost, "/api/tables", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetTable(t *testing.T) {
	r, mockService := setupTest()

	table := &model.Table{
		ID:          "table1",
		Name:        "Test Table",
		Description: "Test Description",
		SpaceID:     "space1",
	}

	mockService.On("GetByID", mock.Anything, "table1").Return(table, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/tables/table1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetTablesBySpace(t *testing.T) {
	r, mockService := setupTest()

	tables := []*model.Table{
		{
			ID:          "table1",
			Name:        "Test Table 1",
			Description: "Test Description 1",
			SpaceID:     "space1",
		},
		{
			ID:          "table2",
			Name:        "Test Table 2",
			Description: "Test Description 2",
			SpaceID:     "space1",
		},
	}

	mockService.On("GetBySpaceID", mock.Anything, "space1").Return(tables, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/tables/space/space1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateTable(t *testing.T) {
	r, mockService := setupTest()

	table := &model.Table{
		ID:          "table1",
		Name:        "Updated Table",
		Description: "Updated Description",
		SpaceID:     "space1",
	}

	mockService.On("Update", mock.Anything, mock.AnythingOfType("*model.Table")).Return(nil)

	body, _ := json.Marshal(table)
	req := httptest.NewRequest(http.MethodPut, "/api/tables/table1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteTable(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("Delete", mock.Anything, "table1").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/tables/table1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}
