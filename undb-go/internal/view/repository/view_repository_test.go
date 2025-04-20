package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/undb/undb-go/internal/view/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.View{})
	assert.NoError(t, err)

	return db
}

func TestViewRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewViewRepository(db)

	view := &model.View{
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}

	err := repo.Create(context.Background(), view)
	assert.NoError(t, err)
	assert.NotEmpty(t, view.ID)
}

func TestViewRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewViewRepository(db)

	// Create a test view
	view := &model.View{
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}
	err := repo.Create(context.Background(), view)
	assert.NoError(t, err)

	// Test getting the view
	found, err := repo.GetByID(context.Background(), view.ID)
	assert.NoError(t, err)
	assert.Equal(t, view.ID, found.ID)
	assert.Equal(t, view.Name, found.Name)
	assert.Equal(t, view.TableID, found.TableID)
	assert.Equal(t, view.Type, found.Type)

	// Test getting non-existent view
	_, err = repo.GetByID(context.Background(), "non-existent")
	assert.Error(t, err)
}

func TestViewRepository_GetByTableID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewViewRepository(db)

	// Create test views
	views := []*model.View{
		{
			Name:    "View 1",
			TableID: "table1",
			Type:    model.ViewTypeGrid,
		},
		{
			Name:    "View 2",
			TableID: "table1",
			Type:    model.ViewTypeKanban,
		},
		{
			Name:    "View 3",
			TableID: "table2",
			Type:    model.ViewTypeGrid,
		},
	}

	for _, view := range views {
		err := repo.Create(context.Background(), view)
		assert.NoError(t, err)
	}

	// Test getting views for table1
	found, err := repo.GetByTableID(context.Background(), "table1")
	assert.NoError(t, err)
	assert.Len(t, found, 2)

	// Test getting views for table2
	found, err = repo.GetByTableID(context.Background(), "table2")
	assert.NoError(t, err)
	assert.Len(t, found, 1)

	// Test getting views for non-existent table
	found, err = repo.GetByTableID(context.Background(), "non-existent")
	assert.NoError(t, err)
	assert.Empty(t, found)
}

func TestViewRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewViewRepository(db)

	// Create a test view
	view := &model.View{
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}
	err := repo.Create(context.Background(), view)
	assert.NoError(t, err)

	// Update the view
	view.Name = "Updated View"
	view.Type = model.ViewTypeKanban
	err = repo.Update(context.Background(), view)
	assert.NoError(t, err)

	// Verify the update
	found, err := repo.GetByID(context.Background(), view.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated View", found.Name)
	assert.Equal(t, model.ViewTypeKanban, found.Type)
}

func TestViewRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewViewRepository(db)

	// Create a test view
	view := &model.View{
		Name:    "Test View",
		TableID: "table1",
		Type:    model.ViewTypeGrid,
	}
	err := repo.Create(context.Background(), view)
	assert.NoError(t, err)

	// Delete the view
	err = repo.Delete(context.Background(), view.ID)
	assert.NoError(t, err)

	// Verify the view is deleted
	_, err = repo.GetByID(context.Background(), view.ID)
	assert.Error(t, err)
}
