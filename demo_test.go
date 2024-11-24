package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	HelloHandler(rec, req)

	// Check HTTP status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rec.Code)
	}

	// Check response body
	expected := `{"message":"Hello, World!"}`
	actual := strings.TrimSpace(rec.Body.String())
	if actual != expected {
		t.Errorf("Expected body '%s'; got '%s'", expected, actual)
	}
}
