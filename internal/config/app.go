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
	{1, "Lord of the rings"},
	{2, "Star Wars"},
	{3, "The Hobbit"},
}

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