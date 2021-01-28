package main

import (
	"io/ioutil"
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

func getAndAssertStatus(t *testing.T, url string, expectedStatus int) *http.Response {
	response, _ := http.Get(url)
	if response.StatusCode != expectedStatus {
		t.Errorf("Invalid status, got: %v, wanted: %v", response.StatusCode, expectedStatus)
	}
	return response
}

func postAndAssertStatus(t *testing.T, url string, expectedStatus int, body io.Reader) *http.Response {
	response, _ := http.Post(url, "multipart/form-data", body)
	if response.StatusCode != expectedStatus {
		t.Errorf("Invalid status, got: %v, wanted: %v", response.StatusCode, expectedStatus)
	}
	return response
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

func TestGetForm(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()

	resp := getAndAssertStatus(t, baseURL+"/form", http.StatusOK)	
	if head := resp.Header.Get("Content-Type"); head != "text/html" {
		t.Error("invalid content header :" + head)
	}
}

func TestPostFileThroughForm(t *testing.T) {
	baseURL, closeFun := createServer()
	defer closeFun()
	
	fileContent := "this is my file asd"
	fileData := strings.NewReader(fileContent)

	resp := postAndAssertStatus(t, baseURL+"/form", http.StatusOK, fileData)
	if head := resp.Header.Get("Content-Type"); head != "text/html" {
		t.Error("invalid content header :" + head)
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("got error during reading response: ", err)
	}
	responseString := string(responseBody)
	if !strings.Contains(responseString, fileContent) {
		t.Error("Output does not contain uploaded file: ", responseString)
	}
}