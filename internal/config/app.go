package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Connect() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}
