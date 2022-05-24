package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jmeisele/golang_gorm/internal/config"
	"github.com/jmeisele/golang_gorm/internal/routes"
)

func main() {
	
	// Creating DB
	fmt.Println("Creating our initial DB")
	config.CreateDB()
	
	// Registering routes
	fmt.Println("Registering routes")
	router := mux.NewRouter()
	routes.BookStoreRoutes(router)
}