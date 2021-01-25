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
	if response.StatusCode != expStatus {
		t.Errorf("Wanted %v, got %v", expStatus, response.StatusCode)
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
	if strBody != "Hi!" {
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

func TestUnknown(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	strBody := doGet(t, server.URL+"/cat/foo", http.StatusOK)
	if strBody != "Hi!" {
		t.Error("Invalid body: ", strBody)
	}
}

func TestQueried(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	strBody := doGet(t, server.URL+"/queried?foo=bar&asd=123", http.StatusOK)
	if strBody != "Got foo=bar and asd=123" && strBody != "Got asd=123 and foo=bar" {
		t.Error("Invalid body: ", strBody)
	}
}

func TestPathParam(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	strBody := doGet(t, server.URL+"/foo/123", http.StatusOK)
	if strBody != "Id: 123" {
		t.Error("Invalid body: ", strBody)
	}
}

func TestPathParamNotFound(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	tt := []string {"/foo/foo/foo", "/foo/", "/foo/bar"}
	for _,tc := range tt {
		t.Run(tc, func(t *testing.T) {
			strBody := doGet(t, server.URL+tc, http.StatusNotFound)
			if strBody != "" {
				t.Error("Invalid body: ", strBody)
			}
		})
	}
}