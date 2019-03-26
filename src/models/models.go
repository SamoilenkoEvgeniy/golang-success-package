package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	name     string
	email    string
	password string
}

type Row struct {
	gorm.Model
	Code  string
	Price uint
}
