package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryFutures(t *testing.T) {
	expectedURL := "/instruments/futures"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"data": {"items": [{"symbol": "AAPL"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.QueryFutures(nil)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].Symbol)
	}
}

func TestQueryFuturesError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	_, err := api.QueryFutures(nil)

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestQueryFuturesWithParams(t *testing.T) {
	expectedURL := "/instruments/futures?product-code%5B%5D=123&symbol%5B%5D=AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"data": {"items": [{"symbol": "AAPL", "product-code": "123"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	params := &FuturesQueryParams{
		Symbol:      []string{"AAPL"},
		ProductCode: []string{"123"},
	}
	resp, err := api.QueryFutures(params)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].Symbol)
	}

	if resp.Data.Items[0].ProductCode != "123" {
		t.Errorf("expected %s, got %s", "123", resp.Data.Items[0].ProductCode)
	}
}

func TestGetFuture(t *testing.T) {
	expectedURL := "/instruments/futures/AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"data": {"symbol": "AAPL"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetFuture("AAPL")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Data.Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Symbol)
	}
}

func TestGetFutureOptionProduct(t *testing.T) {
	expectedURL := "/instruments/future-option-products/exchange/AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"context": "test", "data": {"root-symbol": "AAPL"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetFutureOptionProduct("exchange", "AAPL")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if resp.Data.RootSymbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.RootSymbol)
	}
}

func TestListFutureOptions(t *testing.T) {
	expectedURL := "/instruments/future-options?symbol%5B%5D=AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL", "instrument-type": "future-option"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListFutureOptions(&FutureOptionsQueryParams{Symbol: []string{"AAPL"}})

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

func TestListFutureOptionChainsNested(t *testing.T) {
	expectedURL := "/futures-option-chains/AAPL/nested"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"context": "test", "data": {"futures": [{"symbol": "AAPL"}], "option-chains": [{"underlying-symbol": "AAPL"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListFutureOptionChainsNested("AAPL")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Futures) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Futures))
	}

	if resp.Data.Futures[0].Symbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Futures[0].Symbol)
	}

	if len(resp.Data.OptionChains) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.OptionChains))
	}

	if resp.Data.OptionChains[0].UnderlyingSymbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.OptionChains[0].UnderlyingSymbol)
	}
}

func TestListFutureOptionChainsDetailed(t *testing.T) {
	expectedURL := "/futures-option-chains/AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL", "underlying-symbol": "AAPL"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListFutureOptionChainsDetailed("AAPL")

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

func TestGetFutureOption(t *testing.T) {
	expectedURL := "/instruments/future-options/AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"context": "test", "data": {"symbol": "AAPL", "instrument-type": "future-option"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetFutureOption("AAPL")

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

func TestListFutureOptionProducts(t *testing.T) {
	expectedURL := "/instruments/future-option-products"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"root-symbol": "AAPL", "instrument-type": "future-option"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListFutureOptionProducts()

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].RootSymbol != "AAPL" {
		t.Errorf("expected %s, got %s", "AAPL", resp.Data.Items[0].RootSymbol)
	}
}

func TestGetFutureOptionProductError(t *testing.T) {
	expectedURL := "/instruments/future-option-products/exchange/AAPL"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != expectedURL {
			t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
		}
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	_, err := api.GetFutureOptionProduct("exchange", "AAPL")

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}
