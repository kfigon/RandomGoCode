package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"strings"
	"io"
)

func createServer() (string, func()) {
	srv := httptest.NewServer(createMux())
	return srv.URL, srv.Close
}

func getAndAssertStatus(t *testing.T, url string, expectedStatus int) {
	response, _ := http.Get(url)
	if response.StatusCode != expectedStatus {
		t.Errorf("Invalid status, got: %v, wanted: %v", response.StatusCode, expectedStatus)
	}
}

func postAndAssertStatus(t *testing.T, url string, expectedStatus int, body io.Reader) {
	response, _ := http.Post(url, "multipart/form-data", body)
	if response.StatusCode != expectedStatus {
		t.Errorf("Invalid status, got: %v, wanted: %v", response.StatusCode, expectedStatus)
	}
}

func TestHello(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	getAndAssertStatus(t, baseURL+"/upload", http.StatusNotFound)
}

func TestUploadFile(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	fileData := strings.NewReader("this is my file!")
	postAndAssertStatus(t, baseURL+"/upload", http.StatusAccepted, fileData)
}

func TestUploadEmptyFile(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	fileData := strings.NewReader("")
	postAndAssertStatus(t, baseURL+"/upload", http.StatusBadRequest, fileData)	
}