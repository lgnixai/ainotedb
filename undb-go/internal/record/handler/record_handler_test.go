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
	"github.com/undb/undb-go/internal/record/model"
)

// MockRecordService 是 RecordService 的模拟实现
type MockRecordService struct {
	mock.Mock
}

func (m *MockRecordService) Create(ctx context.Context, record *model.Record) error {
	args := m.Called(ctx, record)
	return args.Error(0)
}

func (m *MockRecordService) GetByID(ctx context.Context, id string) (*model.Record, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Record), args.Error(1)
}

func (m *MockRecordService) GetByTableID(ctx context.Context, tableID string) ([]*model.Record, error) {
	args := m.Called(ctx, tableID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Record), args.Error(1)
}

func (m *MockRecordService) Update(ctx context.Context, record *model.Record) error {
	args := m.Called(ctx, record)
	return args.Error(0)
}

func (m *MockRecordService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func setupTest() (*gin.Engine, *MockRecordService) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockRecordService)
	handler := NewRecordHandler(mockService)

	group := r.Group("/api")
	group.POST("/records", handler.CreateRecord)
	group.GET("/records/:id", handler.GetRecord)
	group.GET("/records/table/:table_id", handler.GetRecordsByTable)
	group.PUT("/records/:id", handler.UpdateRecord)
	group.DELETE("/records/:id", handler.DeleteRecord)

	return r, mockService
}

func TestCreateRecord(t *testing.T) {
	r, mockService := setupTest()

	record := &model.Record{
		TableID: 1,
		Data: map[string]interface{}{
			"name": "Test Record",
			"age":  25,
		},
	}

	mockService.On("Create", mock.Anything, mock.AnythingOfType("*model.Record")).Return(nil)

	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/api/records", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetRecord(t *testing.T) {
	r, mockService := setupTest()

	record := &model.Record{
		ID:      "record1",
		TableID: 1,
		Data: map[string]interface{}{
			"name": "Test Record",
			"age":  25,
		},
	}

	mockService.On("GetByID", mock.Anything, 1).Return(record, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/records/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetRecordsByTable(t *testing.T) {
	r, mockService := setupTest()

	records := []*model.Record{
		{
			ID:      1,
			TableID: 1,
			Data: map[string]interface{}{
				"name": "Test Record 1",
				"age":  25,
			},
		},
		{
			ID:      2,
			TableID: 1,
			Data: map[string]interface{}{
				"name": "Test Record 2",
				"age":  30,
			},
		},
	}

	mockService.On("GetByTableID", mock.Anything, 1).Return(records, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/records/table/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateRecord(t *testing.T) {
	r, mockService := setupTest()

	record := &model.Record{
		ID:      1,
		TableID: 1,
		Data: map[string]interface{}{
			"name": "Updated Record",
			"age":  35,
		},
	}

	mockService.On("Update", mock.Anything, mock.AnythingOfType("*model.Record")).Return(nil)

	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPut, "/api/records/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteRecord(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/records/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}
