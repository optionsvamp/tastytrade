package tastytrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"data": "test"}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	resp, err := api.fetchData(server.URL)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if resp["data"] != "test" {
		t.Errorf("expected %s, got %s", "test", resp["data"])
	}
}

func TestFetchDataAndUnmarshal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"data": "test"}`))
	}))
	defer server.Close()

	api := NewTastytradeAPI(server.URL)
	var v map[string]interface{}
	err := api.fetchDataAndUnmarshal(server.URL, &v)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if v["data"] != "test" {
		t.Errorf("expected %s, got %s", "test", v["data"])
	}
}
