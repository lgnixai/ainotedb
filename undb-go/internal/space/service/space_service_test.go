package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/undb/undb-go/internal/space/model"
	"github.com/undb/undb-go/internal/space/repository"
	"github.com/undb/undb-go/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 迁移数据库
	err = db.AutoMigrate(&model.Space{}, &model.SpaceMember{})
	assert.NoError(t, err)

	return db
}

func TestSpaceService_Create(t *testing.T) {
	db := setupTestDB(t)
	spaceRepo := repository.NewSpaceRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	service := NewSpaceService(spaceRepo, memberRepo, db)

	space := &model.Space{
		ID:          utils.GenerateID("spc_"),
		Name:        "Test Space",
		Description: "Test Description",
		OwnerID:     "user1",
		Visibility:  model.VisibilityPrivate,
	}

	err := service.Create(context.Background(), space)
	assert.NoError(t, err)

	// 验证空间是否创建成功
	var foundSpace model.Space
	err = db.First(&foundSpace, "id = ?", space.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, space.Name, foundSpace.Name)

	// 验证是否自动添加了所有者成员
	var member model.SpaceMember
	err = db.First(&member, "space_id = ? AND user_id = ?", space.ID, space.OwnerID).Error
	assert.NoError(t, err)
	assert.Equal(t, model.RoleOwner, member.Role)
}

func TestSpaceService_GetByID(t *testing.T) {
	db := setupTestDB(t)
	spaceRepo := repository.NewSpaceRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	service := NewSpaceService(spaceRepo, memberRepo, db)

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
	found, err := service.GetByID(context.Background(), space.ID)
	assert.NoError(t, err)
	assert.Equal(t, space.Name, found.Name)
	assert.Equal(t, space.Description, found.Description)
	assert.Equal(t, space.OwnerID, found.OwnerID)
	assert.Equal(t, space.Visibility, found.Visibility)

	// 测试获取不存在的空间
	_, err = service.GetByID(context.Background(), "non-existent")
	assert.Error(t, err)
}

func TestSpaceService_GetByOwnerID(t *testing.T) {
	db := setupTestDB(t)
	spaceRepo := repository.NewSpaceRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	service := NewSpaceService(spaceRepo, memberRepo, db)

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
	found, err := service.GetByOwnerID(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Len(t, found, 2)

	// 测试获取 user2 的空间
	found, err = service.GetByOwnerID(context.Background(), "user2")
	assert.NoError(t, err)
	assert.Len(t, found, 1)

	// 测试获取不存在的用户的空间
	found, err = service.GetByOwnerID(context.Background(), "non-existent")
	assert.NoError(t, err)
	assert.Empty(t, found)
}

func TestSpaceService_Update(t *testing.T) {
	db := setupTestDB(t)
	spaceRepo := repository.NewSpaceRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	service := NewSpaceService(spaceRepo, memberRepo, db)

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
	err = service.Update(context.Background(), space)
	assert.NoError(t, err)

	// 验证更新
	var found model.Space
	err = db.First(&found, "id = ?", space.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, space.Name, found.Name)
	assert.Equal(t, space.Description, found.Description)
}

func TestSpaceService_Delete(t *testing.T) {
	db := setupTestDB(t)
	spaceRepo := repository.NewSpaceRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	service := NewSpaceService(spaceRepo, memberRepo, db)

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

	// 添加成员
	member := &model.SpaceMember{
		ID:      utils.GenerateID("mem_"),
		SpaceID: space.ID,
		UserID:  "user2",
		Role:    model.RoleEditor,
	}
	err = db.Create(member).Error
	assert.NoError(t, err)

	// 删除空间
	err = service.Delete(context.Background(), space.ID)
	assert.NoError(t, err)

	// 验证空间是否被删除
	var foundSpace model.Space
	err = db.First(&foundSpace, "id = ?", space.ID).Error
	assert.Error(t, err)

	// 验证成员是否被删除
	var foundMember model.SpaceMember
	err = db.First(&foundMember, "space_id = ?", space.ID).Error
	assert.Error(t, err)
}
