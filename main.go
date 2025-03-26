package main

import (
	"fmt"
	"go-postgres-app/config"
	"go-postgres-app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Initialize database tables
	err := config.InitializeDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/api/users", handlers.CreateUserhandler).Methods("POST")
	r.HandleFunc("/api/users", handlers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", handlers.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	// r.HandleFunc("/api/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	// Start the server
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
