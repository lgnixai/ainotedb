package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/undb/undb-go/internal/space/model"
	"github.com/undb/undb-go/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 迁移数据库
	err = db.AutoMigrate(&model.Space{})
	assert.NoError(t, err)

	return db
}

func TestSpaceRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewSpaceRepository(db)

	space := &model.Space{
		ID:          utils.GenerateID("spc_"),
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}

	err := repo.Create(context.Background(), space)
	assert.NoError(t, err)

	// 验证是否创建成功
	var found model.Space
	err = db.First(&found, "id = ?", space.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, space.Name, found.Name)
	assert.Equal(t, space.Description, found.Description)
	assert.Equal(t, space.OwnerID, found.OwnerID)
	assert.Equal(t, space.Visibility, found.Visibility)
}

func TestSpaceRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewSpaceRepository(db)

	// 创建测试数据
	space := &model.Space{
		ID:          utils.GenerateID("spc_"),
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}
	err := db.Create(space).Error
	assert.NoError(t, err)

	// 测试获取
	found, err := repo.GetByID(context.Background(), space.ID)
	assert.NoError(t, err)
	assert.Equal(t, space.Name, found.Name)
	assert.Equal(t, space.Description, found.Description)
	assert.Equal(t, space.OwnerID, found.OwnerID)
	assert.Equal(t, space.Visibility, found.Visibility)

	// 测试获取不存在的空间
	_, err = repo.GetByID(context.Background(), "non-existent")
	assert.Error(t, err)
	assert.Equal(t, model.ErrSpaceNotFound, err)
}

func TestSpaceRepository_GetByOwnerID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewSpaceRepository(db)

	// 创建测试数据
	spaces := []*model.Space{
		{
			ID:          utils.GenerateID("spc_"),
			Name:        "Space 1",
			Description: "Description 1",
			OwnerID:     "user1",
			Visibility:  model.VisibilityPrivate,
		},
		{
			ID:          utils.GenerateID("spc_"),
			Name:        "Space 2",
			Description: "Description 2",
			OwnerID:     "user1",
			Visibility:  model.VisibilityPrivate,
		},
		{
			ID:          utils.GenerateID("spc_"),
			Name:        "Space 3",
			Description: "Description 3",
			OwnerID:     "user2",
			Visibility:  model.VisibilityPrivate,
		},
	}

	for _, space := range spaces {
		err := db.Create(space).Error
		assert.NoError(t, err)
	}

	// 测试获取 user1 的空间
	found, err := repo.GetByOwnerID(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Len(t, found, 2)

	// 测试获取 user2 的空间
	found, err = repo.GetByOwnerID(context.Background(), "user2")
	assert.NoError(t, err)
	assert.Len(t, found, 1)

	// 测试获取不存在的用户的空间
	found, err = repo.GetByOwnerID(context.Background(), "non-existent")
	assert.NoError(t, err)
	assert.Empty(t, found)
}

func TestSpaceRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewSpaceRepository(db)

	// 创建测试数据
	space := &model.Space{
		ID:          utils.GenerateID("spc_"),
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}
	err := db.Create(space).Error
	assert.NoError(t, err)

	// 更新空间
	space.Name = "Updated Space"
	space.Description = "Updated Description"
	err = repo.Update(context.Background(), space)
	assert.NoError(t, err)

	// 验证更新
	var found model.Space
	err = db.First(&found, "id = ?", space.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, space.Name, found.Name)
	assert.Equal(t, space.Description, found.Description)
}

func TestSpaceRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewSpaceRepository(db)

	// 创建测试数据
	space := &model.Space{
		ID:          utils.GenerateID("spc_"),
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}
	err := db.Create(space).Error
	assert.NoError(t, err)

	// 删除空间
	err = repo.Delete(context.Background(), space.ID)
	assert.NoError(t, err)

	// 验证删除
	var found model.Space
	err = db.First(&found, "id = ?", space.ID).Error
	assert.Error(t, err)
}
