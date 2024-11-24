package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rec.Code)
	}

	expected := `{"status":"Healthy"}`
	if rec.Body.String() != expected {
		t.Errorf("Expected body %s; got %s", expected, rec.Body.String())
	}
}
