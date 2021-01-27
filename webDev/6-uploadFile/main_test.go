package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestHello(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	response, _ := http.Get(srv.URL+"/")
	if response.StatusCode != http.StatusOK {
		t.Errorf("Invalid status, got: %v", response.StatusCode)
	}
}