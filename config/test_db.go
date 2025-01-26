package config

import (
	"log"
	"ticketing-system/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

// SetupTestDB sets up a temporary SQLite database for testing
func SetupTestDB() {
	var err error
	TestDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	// Migrate database schema
	if err := TestDB.AutoMigrate(&entity.Event{}, &entity.Ticket{}, &entity.User{}); err != nil {
		log.Fatalf("failed to migrate test database: %v", err)
	}
}

// TeardownTestDB cleans up the test database
func TeardownTestDB() {
	sqlDB, err := TestDB.DB()
	if err != nil {
		log.Printf("error closing test database: %v", err)
		return
	}
	sqlDB.Close()
}
