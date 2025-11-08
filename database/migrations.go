package database

import (
	"test/models"

	"gorm.io/gorm"
)

func AutoMigrateAll(db *gorm.DB) error {
	models := []interface{}{
		&models.User{},
		&models.UserContact{},
		&models.Passport{},
		&models.Address{},
		&models.Loans{},
	}

	return db.AutoMigrate(models...)
}
