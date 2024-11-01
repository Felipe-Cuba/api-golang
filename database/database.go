package database

import (
	"api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}

func AutoMigrates(db *gorm.DB) {
	db.AutoMigrate(models.Student{})
}