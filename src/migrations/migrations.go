package migrations

import "github.com/jinzhu/gorm"
import "../models"

// SomeFunc this is for test
func SomeFunc(db *gorm.DB) string {

	db.AutoMigrate(&models.Row{})

	return "Some String"
}
