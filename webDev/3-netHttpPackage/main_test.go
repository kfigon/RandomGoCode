package main

import (
	"net/http"
	"io/ioutil"
	"testing"
	"net/http/httptest"
	"strings"
)

func serve(inputRequest *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	var d hotdog
	d.ServeHTTP(response, inputRequest)
	return response
}
func assertStatus200AndContentType(t *testing.T, response *httptest.ResponseRecorder) {
	receivedStatus := response.Result().StatusCode
	if receivedStatus != 200 {
		t.Error("Invalid status code received: ", receivedStatus)
	}
	headers := response.Result().Header
	if content := headers.Get("Content-Type"); content != "text/html; charset=utf-8" {
		t.Error("Invalid contentType: ", content)
	}
}

func readAllBody(response *httptest.ResponseRecorder) string {
	body, _ := ioutil.ReadAll(response.Result().Body)
	return string(body)
}
func TestGet(t *testing.T) {
	response := serve(httptest.NewRequest("GET", "http://localhost:8080", nil))

	assertStatus200AndContentType(t, response)
	if strings.Contains(readAllBody(response), "You have filled this thing") {
		t.Error("Non expected string received")
	}
}

func TestGetWithInvalidQuery(t *testing.T) {
	response := serve(httptest.NewRequest("GET", "http://localhost:8080?asd=foo", nil))

	assertStatus200AndContentType(t, response)
	if strings.Contains(readAllBody(response), "You have filled this thing") {
		t.Error("Non expected string received")
	}
}

func TestGetWithRightQuery(t *testing.T) {
	response := serve(httptest.NewRequest("GET", "http://localhost:8080?myName=foo", nil))

	assertStatus200AndContentType(t, response)
	if !strings.Contains(readAllBody(response), "You have filled this thing: foo") {
		t.Error("Expected string not found")
	}
}

func TestPost(t *testing.T) {
	reader := strings.NewReader("myName=abc")
	req := httptest.NewRequest("POST", "http://localhost:8080", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := serve(req)

	assertStatus200AndContentType(t, response)
	if !strings.Contains(readAllBody(response), "You have filled this thing: abc") {
		t.Error("Expected string not found")
	}
}