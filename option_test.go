package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListOptionsChainsDetailed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL", "underlying-symbol": "AAPL"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListOptionsChainsDetailed("AAPL")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].Symbol)
	}

	if resp.Data.Items[0].UnderlyingSymbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].UnderlyingSymbol)
	}
}

func TestListOptionChainsNested(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"underlying-symbol": "AAPL", "root-symbol": "AAPL"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListOptionChainsNested("AAPL")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].UnderlyingSymbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].UnderlyingSymbol)
	}
}

func TestGetOptionChainsCompact(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"underlying-symbol": "AAPL", "root-symbol": "AAPL"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetOptionChainsCompact("AAPL")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].UnderlyingSymbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].UnderlyingSymbol)
	}
}

func TestGetEquityOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL", "instrument-type": "equity-option"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetEquityOptions(&EquityOptionsQueryParams{Symbol: []string{"AAPL"}})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].Symbol)
	}
}

func TestGetEquityOption(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"symbol": "AAPL", "instrument-type": "equity-option"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetEquityOption("AAPL")

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
