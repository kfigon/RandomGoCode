package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"strings"
)

func createServer() (string, func()) {
	srv := httptest.NewServer(createMux())
	return srv.URL, srv.Close
}

func TestHello(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	response, _ := http.Get(baseURL+"/upload")
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Invalid status, got: %v", response.StatusCode)
	}
}

func TestUploadFile(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	fileData := strings.NewReader("this is my file!")
	response, _ := http.Post(baseURL+"/upload", "multipart/form-data", fileData)
	
	gotStatus := response.StatusCode
	wantedStatus := http.StatusAccepted
	if gotStatus != wantedStatus {
		t.Errorf("Invalid status, got: %v, want: %v", gotStatus, wantedStatus)
	}
}

func TestUploadEmptyFile(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	fileData := strings.NewReader("")
	response, _ := http.Post(baseURL+"/upload", "multipart/form-data", fileData)
	
	gotStatus := response.StatusCode
	wantedStatus := http.StatusBadRequest
	if gotStatus != wantedStatus {
		t.Errorf("Invalid status, got: %v, want: %v", gotStatus, wantedStatus)
	}
}