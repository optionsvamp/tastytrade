package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountTradingStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"account-number": "123456", "day-trade-count": 1}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetAccountTradingStatus("123456")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp.Context != "test" {
		t.Errorf("expected %s, got %s", "test", resp.Context)
	}

	if resp.Data.AccountNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", resp.Data.AccountNumber)
	}

	if resp.Data.DayTradeCount != 1 {
		t.Errorf("expected %d, got %d", 1, resp.Data.DayTradeCount)
	}
}
