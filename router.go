package main

import (
	"net/http"

	"github.com/Smile2578/smiled-go/controllers"
	"github.com/Smile2578/smiled-go/middleware"
	"github.com/gorilla/mux"
)

func configureRoutes(router *mux.Router) {
	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Example of applying middleware to a route
	router.Handle("/api/patients", middleware.RoleCheckMiddleware(http.HandlerFunc(controllers.GetPatients))).Methods("GET")
	// Add other routes here
}
