package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doGet(t *testing.T, url string, expStatus int) string {
	response, err := http.Get(url)
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
	return string(body)
}

func TestRoot(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	strBody := doGet(t, server.URL+"/", http.StatusOK)
	if strBody != "Hello World!" {
		t.Error("Invalid body: ", strBody)
	}
}

func TestCat(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	strBody := doGet(t, server.URL+"/cat", http.StatusOK)
	if strBody != "Hello Cat!" {
		t.Error("Invalid body: ", strBody)
	}
}
