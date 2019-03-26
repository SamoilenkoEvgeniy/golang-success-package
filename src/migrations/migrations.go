package migrations

import "github.com/jinzhu/gorm"
import "../models"

// Migrate this is for test
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Order{})

	db.Model(&models.Order{}).AddForeignKey("managerID", "users(id)", "CASCADE", "RESTRICT")
}
