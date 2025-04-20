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
	"github.com/undb/undb-go/internal/user/model"
)

// MockUserService is a mock implementation of UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	args := m.Called(ctx, email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) Update(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func setupTest() (*gin.Engine, *MockUserService) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	group := r.Group("/api")
	group.POST("/users/register", handler.Register)
	group.POST("/users/login", handler.Login)
	group.GET("/users/:id", handler.GetUser)
	group.PUT("/users/:id", handler.UpdateUser)
	group.DELETE("/users/:id", handler.DeleteUser)

	return r, mockService
}

func TestRegister(t *testing.T) {
	r, mockService := setupTest()

	req := RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	mockService.On("Register", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/users/register", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestRegister_InvalidRequest(t *testing.T) {
	r, _ := setupTest()

	// Test with invalid email
	req := RegisterRequest{
		Email:    "invalid-email",
		Password: "password123",
		Name:     "Test User",
	}

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/users/register", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin(t *testing.T) {
	r, mockService := setupTest()

	req := LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := &model.User{
		ID:    "user1",
		Email: req.Email,
		Name:  "Test User",
	}

	mockService.On("Login", mock.Anything, req.Email, req.Password).Return(user, nil)

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/users/login", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	r, mockService := setupTest()

	req := LoginRequest{
		Email:    "test@example.com",
		Password: "wrong-password",
	}

	mockService.On("Login", mock.Anything, req.Email, req.Password).Return(nil, model.ErrInvalidPassword)

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/users/login", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	r, mockService := setupTest()

	user := &model.User{
		ID:    "user1",
		Email: "test@example.com",
		Name:  "Test User",
	}

	mockService.On("GetByID", mock.Anything, "user1").Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/users/user1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("GetByID", mock.Anything, "nonexistent").Return(nil, model.ErrUserNotFound)

	req := httptest.NewRequest(http.MethodGet, "/api/users/nonexistent", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	r, mockService := setupTest()

	user := &model.User{
		ID:    "user1",
		Email: "test@example.com",
		Name:  "Updated User",
	}

	mockService.On("Update", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/api/users/user1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateUser_NotFound(t *testing.T) {
	r, mockService := setupTest()

	user := &model.User{
		ID:    "nonexistent",
		Email: "test@example.com",
		Name:  "Updated User",
	}

	mockService.On("Update", mock.Anything, mock.AnythingOfType("*model.User")).Return(model.ErrUserNotFound)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/api/users/nonexistent", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("Delete", mock.Anything, "user1").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/users/user1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	r, mockService := setupTest()

	mockService.On("Delete", mock.Anything, "nonexistent").Return(model.ErrUserNotFound)

	req := httptest.NewRequest(http.MethodDelete, "/api/users/nonexistent", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}
