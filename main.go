package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // other necessary imports
)

func main() {
	// Setup database connection
	err := ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ping the database to check the connection
	err = PingDatabase()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Define your routes
	router := setupRouter()

	// Log that the server is starting
	log.Println("Starting server on :8080")

	// Start the server
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Setup routes
	configureRoutes(router)

	return router
}
