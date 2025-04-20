
package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("Database connected successfully")
	return &Database{db}, nil
}

func (db *Database) AutoMigrate() error {
	// Register all models for auto migration
	models := []interface{}{
		&auth.User{},
		&space.Space{},
		&space.Member{},
		&table.Table{},
		&table.Field{},
		&table.Option{},
		&record.Record{},
		&view.View{},
		&view.ViewConfiguration{},
	}

	// Run migrations
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
