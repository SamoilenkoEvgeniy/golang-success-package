package models

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	name     string
	email    string
	password string
}

// Order struct
type Order struct {
	gorm.Model
	hash      string
	price     uint
	managerID uint
}
