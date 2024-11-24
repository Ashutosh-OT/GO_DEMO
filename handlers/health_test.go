package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
	actual := strings.TrimSpace(rec.Body.String())
	if actual != expected {
		t.Errorf("Expected body '%s'; got '%s'", expected, actual)
	}
}
