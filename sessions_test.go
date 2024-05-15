package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"context": "test", "data": {"user": {"email": "test@example.com", "username": "testuser", "external-id": "123", "is-confirmed": true}, "session-token": "testtoken"}}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	err := api.Authenticate("testuser", "testpassword")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if api.authToken != "testtoken" {
		t.Errorf("expected %s, got %s", "testtoken", api.authToken)
	}
}
