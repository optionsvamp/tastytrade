package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCryptocurrencies(t *testing.T) {
	// Create a test server that returns a dummy response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data": {"items": []}, "context": "/instruments/cryptocurrencies"}`))
	}))
	defer ts.Close()

	// Create a new TastytradeAPI instance with the test server URL
	api := NewTastytradeAPI(ts.URL)

	// Call the ListCryptocurrencies function
	result, err := api.ListCryptocurrencies()

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	if result.Context != "/instruments/cryptocurrencies" {
		t.Errorf("Expected context to be /instruments/cryptocurrencies, got %v", result.Context)
	}
}

func TestGetCryptocurrency(t *testing.T) {
	// Create a test server that returns a dummy response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data": {"symbol": "BTC/USD"}, "context": "/instruments/cryptocurrencies/BTC%2FUSD"}`))
	}))
	defer ts.Close()

	// Create a new TastytradeAPI instance with the test server URL
	api := NewTastytradeAPI(ts.URL)

	// Call the GetCryptocurrency function
	result, err := api.GetCryptocurrency("BTC/USD")

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	if result.Data.Symbol != "BTC/USD" {
		t.Errorf("Expected symbol to be BTC/USD, got %v", result.Data.Symbol)
	}
}
