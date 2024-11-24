package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the health status.
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler responds with the health status.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "Healthy"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
