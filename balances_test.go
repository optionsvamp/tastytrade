package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountBalances(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if req.URL.String() != "/accounts/123/balances" {
			t.Errorf("got: %s, want: /accounts/123/balances", req.URL.String())
		}
		// Send response to be tested
		rw.Write([]byte(`{"data": {"account-number": "123", "cash-balance": "1000"}, "context": "test"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Initialize a new TastytradeAPI instance
	api := NewTastytradeAPI(server.URL)

	// Invoke the method to be tested
	resp, err := api.GetAccountBalances("123")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Data.AccountNumber != "123" {
		t.Errorf("expected %s, got %s", "123", resp.Data.AccountNumber)
	}

	if resp.Data.CashBalance != 1000 {
		t.Errorf("expected %f, got %f", 1000.0, resp.Data.CashBalance)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}
}
