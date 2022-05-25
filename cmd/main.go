package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmeisele/golang_gorm/internal/config"
	"github.com/jmeisele/golang_gorm/internal/routes"
)

const host = "localhost"
const port = 8081


func main() {
	
	// Creating DB
	log.Println("Creating our initial DB")
	config.CreateDB()
	
	// Registering routes
	fmt.Println("Registering routes")
	router := mux.NewRouter()
	routes.BookStoreRoutes(router)
	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
	if err != nil {
		log.Fatal(err)
	}
}