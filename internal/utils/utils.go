package utils

import (
	"fmt"
	"os"

	"github.com/omnlgy/RESTHARISGO/internal/models"
	"github.com/omnlgy/RESTHARISGO/internal/seed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	// Database — SQLite file path
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "hris.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(
		&models.Department{},
		&models.Position{},
		&models.Employee{},
		&models.Attendance{},
		&models.Leave{},
		&models.Salary{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	fmt.Println("Database migrated successfully")

	// Seed data when SEED=true
	if os.Getenv("SEED") == "true" {
		seed.SeedAll(db)
	}

	return db, nil
}
