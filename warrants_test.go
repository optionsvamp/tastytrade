package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListWarrants(t *testing.T) {
	// Create a test server that returns a dummy response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data": {"items": []}, "context": "/instruments/warrants"}`))
	}))
	defer ts.Close()

	// Create a new TastytradeAPI instance with the test server URL
	api := NewTastytradeAPI(ts.URL)

	// Call the ListWarrants function
	result, err := api.ListWarrants()

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	if result.Context != "/instruments/warrants" {
		t.Errorf("Expected context to be /instruments/warrants, got %v", result.Context)
	}
}

func TestGetWarrant(t *testing.T) {
	// Create a test server that returns a dummy response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data": {"symbol": "SEPAW"}, "context": "/instruments/warrants/SEPAW"}`))
	}))
	defer ts.Close()

	// Create a new TastytradeAPI instance with the test server URL
	api := NewTastytradeAPI(ts.URL)

	// Call the GetWarrant function
	result, err := api.GetWarrant("SEPAW")

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	if result.Data.Symbol != "SEPAW" {
		t.Errorf("Expected symbol to be SEPAW, got %v", result.Data.Symbol)
	}
}
