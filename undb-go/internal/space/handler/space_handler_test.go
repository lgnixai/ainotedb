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
	"github.com/undb/undb-go/internal/space/model"
)

// MockSpaceService 是 SpaceService 的模拟实现
type MockSpaceService struct {
	mock.Mock
}

func (m *MockSpaceService) Create(ctx context.Context, space *model.Space) error {
	args := m.Called(ctx, space)
	return args.Error(0)
}

func (m *MockSpaceService) GetByID(ctx context.Context, id string) (*model.Space, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Space), args.Error(1)
}

func (m *MockSpaceService) GetByOwnerID(ctx context.Context, ownerID string) ([]*model.Space, error) {
	args := m.Called(ctx, ownerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Space), args.Error(1)
}

func (m *MockSpaceService) Update(ctx context.Context, space *model.Space) error {
	args := m.Called(ctx, space)
	return args.Error(0)
}

func (m *MockSpaceService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSpaceService) AddMember(ctx context.Context, spaceID, userID string, role model.MemberRole) error {
	args := m.Called(ctx, spaceID, userID, role)
	return args.Error(0)
}

func (m *MockSpaceService) RemoveMember(ctx context.Context, spaceID, userID string) error {
	args := m.Called(ctx, spaceID, userID)
	return args.Error(0)
}

func (m *MockSpaceService) UpdateMemberRole(ctx context.Context, spaceID, userID string, role model.MemberRole) error {
	args := m.Called(ctx, spaceID, userID, role)
	return args.Error(0)
}

func (m *MockSpaceService) GetSpaceMembers(ctx context.Context, spaceID string) ([]*model.SpaceMember, error) {
	args := m.Called(ctx, spaceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.SpaceMember), args.Error(1)
}

func setupTest() (*gin.Engine, *MockSpaceService) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockSpaceService)
	handler := NewSpaceHandler(mockService)

	group := r.Group("/api")
	group.POST("/spaces", handler.CreateSpace)
	group.GET("/spaces/:id", handler.GetSpace)
	group.GET("/spaces", handler.ListSpaces)
	group.PUT("/spaces/:id", handler.UpdateSpace)
	group.DELETE("/spaces/:id", handler.DeleteSpace)
	group.POST("/spaces/:id/members", handler.AddMember)
	group.DELETE("/spaces/:id/members/:user_id", handler.RemoveMember)
	group.PUT("/spaces/:id/members/:user_id/role", handler.UpdateMemberRole)
	group.GET("/spaces/:id/members", handler.GetSpaceMembers)

	return r, mockService
}

func TestCreateSpace(t *testing.T) {
	r, mockService := setupTest()

	space := &model.Space{
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}

	mockService.On("Create", mock.Anything, mock.AnythingOfType("*model.Space")).Return(nil)

	body, _ := json.Marshal(space)
	req := httptest.NewRequest(http.MethodPost, "/api/spaces", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetSpace(t *testing.T) {
	r, mockService := setupTest()

	space := &model.Space{
		ID:         3,
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}

	mockService.On("GetByID", mock.Anything, uint(3)).Return(space, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/spaces/3", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestListSpaces(t *testing.T) {
	r, mockService := setupTest()

	spaces := []*model.Space{
		{
			ID:         3,
			Name:        "Test Space 1",
			Description: "Test Description 1",
			OwnerID:     "user1",
			Visibility:  model.VisibilityPrivate,
		},
		{
			ID:         2,
			Name:        "Test Space 2",
			Description: "Test Description 2",
			OwnerID:     "user1",
			Visibility:  model.VisibilityPublic,
		},
	}

	mockService.On("GetByOwnerID", mock.Anything, "user1").Return(spaces, nil)

	// 设置用户ID到上下文
	r.Use(func(c *gin.Context) {
		c.Set("user_id", "user1")
		c.Next()
	})

	req := httptest.NewRequest(http.MethodGet, "/api/spaces", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateSpace(t *testing.T) {
	r, mockService := setupTest()

	space := &model.Space{
		ID:         3,
		Name:        "Updated Space",
		Description: "Updated Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}

	mockService.On("Update", mock.Anything, mock.AnythingOfType("*model.Space")).Return(nil)

	body, _ := json.Marshal(space)
	req := httptest.NewRequest(http.MethodPut, "/api/spaces/3", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteSpace(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("Delete", mock.Anything, uint(3)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/spaces/3", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

func TestAddMember(t *testing.T) {
	r, mockService := setupTest()

	req := AddMemberRequest{
		UserID: "user2",
		Role:   model.RoleEditor,
	}

	mockService.On("AddMember", mock.Anything, uint(3), "user2", model.RoleEditor).Return(nil)

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/spaces/3/members", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestRemoveMember(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("RemoveMember", mock.Anything, uint(3), "user2").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/spaces/3/members/user2", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateMemberRole(t *testing.T) {
	r, mockService := setupTest()

	req := UpdateMemberRoleRequest{
		Role: model.RoleAdmin,
	}

	mockService.On("UpdateMemberRole", mock.Anything, uint(3), "user2", model.RoleAdmin).Return(nil)

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/api/spaces/3/members/user2/role", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetSpaceMembers(t *testing.T) {
	r, mockService := setupTest()

	members := []*model.SpaceMember{
		{
			ID:      "member1",
			SpaceID: 3,
			UserID:  "user1",
			Role:    model.RoleOwner,
		},
		{
			ID:      "member2",
			SpaceID: 3,
			UserID:  "user2",
			Role:    model.RoleEditor,
		},
	}

	mockService.On("GetSpaceMembers", mock.Anything, uint(3)).Return(members, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/spaces/3/members", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
