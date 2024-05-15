package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomerInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"id": "123", "first-name": "John", "last-name": "Doe"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetCustomerInfo()

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if resp.Data.ID != "123" {
		t.Errorf("expected %s, got %s", "123", resp.Data.ID)
	}
}

func TestListCustomerAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"account": {"account-number": "123"}}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.ListCustomerAccounts()

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if len(resp.Data.Items) != 1 {
		t.Errorf("expected %d, got %d", 1, len(resp.Data.Items))
	}

	if resp.Data.Items[0].Account.AccountNumber != "123" {
		t.Errorf("expected %s, got %s", "123", resp.Data.Items[0].Account.AccountNumber)
	}
}

func TestGetAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"account-number": "123"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetAccount("123")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if resp.Data.AccountNumber != "123" {
		t.Errorf("expected %s, got %s", "123", resp.Data.AccountNumber)
	}
}
