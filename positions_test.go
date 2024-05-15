package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPositions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"items": [{"symbol": "AAPL", "quantity": "100"}]}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.GetPositions("123456")

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

	if resp.Data.Items[0].Quantity != "100" {
		t.Errorf("expected %s, got %s", "100", resp.Data.Items[0].Quantity)
	}
}
