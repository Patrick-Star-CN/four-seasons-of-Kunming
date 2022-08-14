package database

import (
	"four-seasons-of-Kunming/app/models"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Post{})
}
