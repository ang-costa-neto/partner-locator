package db

import (
	"log"
	"partner-locator/internal/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	database, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	if err := database.AutoMigrate(&model.Company{}); err != nil {
		log.Fatal("failed to migrate database")
	}

	return database
}
