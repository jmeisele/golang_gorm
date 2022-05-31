package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id int
	Title string
}

var Books = []Book{
	{Id: 1, Title: "Lord of the rings"},
	{Id: 2, Title: "Star Wars"},
	{Id: 3, Title: "The Hobbit"},
}

var db *gorm.DB

func CreateDB() {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	// Migrate the schema
	db.AutoMigrate(&Book{})
	
	// Create
	db.Create(&Book{Id:1, Title: "Lord of the Rings"})
}

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