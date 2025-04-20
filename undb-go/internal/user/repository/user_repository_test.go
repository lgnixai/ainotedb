package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undb/undb-go/internal/user/model"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds)
	return args.Get(0).(*gorm.DB)
}

func TestUserRepository_Create(t *testing.T) {
	mockDB := new(MockDB)
	repo := &userRepository{db: &gorm.DB{}}

	user := &model.User{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	mockDB.On("Create", user).Return(&gorm.DB{Error: nil})

	err := repo.Create(context.Background(), user)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUserRepository_GetByID(t *testing.T) {
	mockDB := new(MockDB)
	repo := &userRepository{db: &gorm.DB{}}

	user := &model.User{
		ID:       "user1",
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	mockDB.On("First", &model.User{}, "id = ?", "user1").Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*model.User)
		*arg = *user
	})

	result, err := repo.GetByID(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockDB.AssertExpectations(t)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	mockDB := new(MockDB)
	repo := &userRepository{db: &gorm.DB{}}

	user := &model.User{
		ID:       "user1",
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	mockDB.On("First", &model.User{}, "email = ?", "test@example.com").Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*model.User)
		*arg = *user
	})

	result, err := repo.GetByEmail(context.Background(), "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockDB.AssertExpectations(t)
}

func TestUserRepository_Update(t *testing.T) {
	mockDB := new(MockDB)
	repo := &userRepository{db: &gorm.DB{}}

	user := &model.User{
		ID:        "user1",
		Email:     "test@example.com",
		Password:  "password123",
		Name:      "Test User",
		UpdatedAt: time.Now(),
	}

	mockDB.On("Save", user).Return(&gorm.DB{Error: nil})

	err := repo.Update(context.Background(), user)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUserRepository_Delete(t *testing.T) {
	mockDB := new(MockDB)
	repo := &userRepository{db: &gorm.DB{}}

	mockDB.On("Delete", &model.User{}, "id = ?", "user1").Return(&gorm.DB{Error: nil})

	err := repo.Delete(context.Background(), "user1")
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUserRepository_ErrorCases(t *testing.T) {
	mockDB := new(MockDB)
	repo := &userRepository{db: &gorm.DB{}}

	// Test GetByID with not found error
	mockDB.On("First", &model.User{}, "id = ?", "nonexistent").Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
	_, err := repo.GetByID(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Equal(t, model.ErrUserNotFound, err)

	// Test GetByEmail with not found error
	mockDB.On("First", &model.User{}, "email = ?", "nonexistent@example.com").Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
	_, err = repo.GetByEmail(context.Background(), "nonexistent@example.com")
	assert.Error(t, err)
	assert.Equal(t, model.ErrUserNotFound, err)

	mockDB.AssertExpectations(t)
}
