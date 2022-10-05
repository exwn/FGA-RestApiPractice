package controllers

import "github.com/jinzhu/gorm"

type Controllers struct {
	masterDB *gorm.DB
	// gorm.Model
	// First_Name string
	// Last_Name  string
}

func New(db *gorm.DB) *Controllers {
	return &Controllers{
		masterDB: db,
	}
}
