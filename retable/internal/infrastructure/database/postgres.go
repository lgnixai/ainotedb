package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
	"gorm.io/gorm/logger"

	"github.com/your-org/your-repo/auth" // Replace with your actual import path
	"github.com/your-org/your-repo/record" // Replace with your actual import path
	"github.com/your-org/your-repo/space" // Replace with your actual import path
	"github.com/your-org/your-repo/table" // Replace with your actual import path
	"github.com/your-org/your-repo/view" // Replace with your actual import path

)

type Database struct {
	*gorm.DB
}

func NewDatabase(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func (db *Database) AutoMigrate() error {
	// Register all models for auto migration.  Note:  Some models from original
        // are omitted due to inconsistencies in the provided changes.  Consider
        // adding them back if needed.
	models := []interface{}{
		&auth.User{},
		&space.Space{},
		&space.SpaceMember{}, // Assuming SpaceMember is the correct model name
		&table.Table{},
		&table.Field{},
		&record.Record{},
		&view.View{},
	}

	if err := db.DB.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to auto migrate models: %w", err)
	}

	// Add any necessary indexes
	if err := db.createIndexes(); err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	return nil
}

func (db *Database) createIndexes() error {
	// Add record table indexes
	if err := db.DB.Exec("CREATE INDEX IF NOT EXISTS idx_records_table_id ON records(table_id)").Error; err != nil {
		return err
	}

	// Add field table indexes  
	if err := db.DB.Exec("CREATE INDEX IF NOT EXISTS idx_fields_table_id ON fields(table_id)").Error; err != nil {
		return err
	}

	// Add view table indexes
	if err := db.DB.Exec("CREATE INDEX IF NOT EXISTS idx_views_table_id ON views(table_id)").Error; err != nil {
		return err
	}

	// Add member table indexes
	if err := db.DB.Exec("CREATE INDEX IF NOT EXISTS idx_members_space_id ON members(space_id)").Error; err != nil {
		return err
	}

	return nil
}