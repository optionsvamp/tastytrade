package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEquityData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"symbol": "AAPL"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetEquityData("AAPL")

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

func TestListEquities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL"}, {"symbol": "GOOG"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListEquities(nil)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 2 {
		t.Errorf("expected %d, got %d", 2, len(resp.Data.Items))
	}
}

func TestListActiveEquities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL"}, {"symbol": "GOOG"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListActiveEquities(nil)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 2 {
		t.Errorf("expected %d, got %d", 2, len(resp.Data.Items))
	}
}
