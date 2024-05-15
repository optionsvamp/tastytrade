package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransactions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"id": 1, "account-number": "123456", "symbol": "AAPL", "transaction-type": "buy"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetTransactions("123456", nil)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].ID != 1 {
		t.Errorf("expected %d, got %d", 1, resp.Data.Items[0].ID)
	}

	if resp.Data.Items[0].AccountNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", resp.Data.Items[0].AccountNumber)
	}

	if resp.Data.Items[0].Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].Symbol)
	}

	if resp.Data.Items[0].TransactionType != "buy" {
		t.Errorf("expected %s, got %s", "buy", resp.Data.Items[0].TransactionType)
	}
}

func TestGetTransaction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"id": 1, "account-number": "123", "symbol": "AAPL", "instrument-type": "equity-option"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetTransaction("123", "1")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if resp.Data.Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Symbol)
	}
}

func TestGetTransactionsPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"api-version": "1.0", "context": "test", "data": {"items": [{"id": 1, "account-number": "123", "symbol": "AAPL", "instrument-type": "equity-option"}]}, "pagination": {"per-page": 1, "total-items": 10}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetTransactions("123", &TransactionQueryParams{Symbol: "AAPL"})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Pagination.PerPage != 1 {
		t.Errorf("expected %d, got %d", 1, resp.Pagination.PerPage)
	}

	if resp.Pagination.TotalItems != 10 {
		t.Errorf("expected %d, got %d", 10, resp.Pagination.TotalItems)
	}
}

func TestGetTransactionsError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	_, err := api.GetTransactions("123", &TransactionQueryParams{Symbol: "AAPL"})

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestGetTransactionError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	_, err := api.GetTransaction("123", "1")

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}
