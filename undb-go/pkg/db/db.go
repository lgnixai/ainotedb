package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/undb/undb-go/pkg/config"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseURL
	//fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
