package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sample-go-app/handlers"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", handlers.HelloHandler).Methods("GET")
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
