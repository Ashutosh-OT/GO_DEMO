package handlers

import (
	"encoding/json"
	"net/http"
)

// HelloResponse represents a sample response.
type HelloResponse struct {
	Message string `json:"message"`
}

// HelloHandler responds with a greeting.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
