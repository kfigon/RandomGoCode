package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/")
	if err != nil {
		t.Errorf("Unexpected error during request: %q", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Wanted %v, got %v", http.StatusOK, response.StatusCode)
	}
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	strBody := string(body)	
	if strBody != "Hello World!" {
		t.Error("Invalid body: ", strBody)
	}
}
