package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undb/undb-go/internal/user/model"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &model.User{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	// Test successful registration
	mockRepo.On("GetByEmail", mock.Anything, user.Email).Return(nil, model.ErrUserNotFound)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

	err := service.Register(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Password) // Password should be hashed
	mockRepo.AssertExpectations(t)

	// Test email already exists
	mockRepo.On("GetByEmail", mock.Anything, user.Email).Return(user, nil)
	err = service.Register(context.Background(), user)
	assert.Error(t, err)
	assert.Equal(t, model.ErrEmailAlreadyExists, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		ID:       "user1",
		Email:    email,
		Password: string(hashedPassword),
		Name:     "Test User",
	}

	// Test successful login
	mockRepo.On("GetByEmail", mock.Anything, email).Return(user, nil)
	result, err := service.Login(context.Background(), email, password)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)

	// Test invalid password
	_, err = service.Login(context.Background(), email, "wrong-password")
	assert.Error(t, err)
	assert.Equal(t, model.ErrInvalidPassword, err)
	mockRepo.AssertExpectations(t)

	// Test user not found
	mockRepo.On("GetByEmail", mock.Anything, "nonexistent@example.com").Return(nil, model.ErrUserNotFound)
	_, err = service.Login(context.Background(), "nonexistent@example.com", password)
	assert.Error(t, err)
	assert.Equal(t, model.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &model.User{
		ID:       "user1",
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	// Test successful get
	mockRepo.On("GetByID", mock.Anything, "user1").Return(user, nil)
	result, err := service.GetByID(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)

	// Test user not found
	mockRepo.On("GetByID", mock.Anything, "nonexistent").Return(nil, model.ErrUserNotFound)
	_, err = service.GetByID(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Equal(t, model.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Update(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &model.User{
		ID:       "user1",
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Updated User",
	}

	// Test successful update
	mockRepo.On("Update", mock.Anything, user).Return(nil)
	err := service.Update(context.Background(), user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Delete(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	// Test successful delete
	mockRepo.On("Delete", mock.Anything, "user1").Return(nil)
	err := service.Delete(context.Background(), "user1")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test user not found
	mockRepo.On("Delete", mock.Anything, "nonexistent").Return(model.ErrUserNotFound)
	err = service.Delete(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Equal(t, model.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}
