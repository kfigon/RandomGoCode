package main

import (
	"net/http"
	"io/ioutil"
	"testing"
	"net/http/httptest"
	"strings"
)

func serve(inputRequest *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var d hotdog
	d.ServeHTTP(w, inputRequest)
	return w
}
func assertStatus200(t *testing.T, w *httptest.ResponseRecorder) {
	receivedStatus := w.Result().StatusCode
	if receivedStatus != 200 {
		t.Error("Invalid status code received: ", receivedStatus)
	}
	headers := w.Result().Header
	if content := headers.Get("Content-Type"); content != "text/html; charset=utf-8" {
		t.Error("Invalid contentType: ", content)
	}
}

func readAllBody(w *httptest.ResponseRecorder) string {
	body, _ := ioutil.ReadAll(w.Result().Body)
	return string(body)
}
func TestGet(t *testing.T) {
	w := serve(httptest.NewRequest("GET", "http://localhost:8080", nil))

	assertStatus200(t, w)
	if strings.Contains(readAllBody(w), "You have filled this thing") {
		t.Error("Non expected string received")
	}
}

func TestGetWithInvalidQuery(t *testing.T) {
	w := serve(httptest.NewRequest("GET", "http://localhost:8080?asd=foo", nil))

	assertStatus200(t, w)
	if strings.Contains(readAllBody(w), "You have filled this thing") {
		t.Error("Non expected string received")
	}
}

func TestGetWithRightQuery(t *testing.T) {
	w := serve(httptest.NewRequest("GET", "http://localhost:8080?myName=foo", nil))

	assertStatus200(t, w)
	if !strings.Contains(readAllBody(w), "You have filled this thing: foo") {
		t.Error("Expected string not found")
	}
}

func TestPost(t *testing.T) {
	reader := strings.NewReader("myName=abc")
	req := httptest.NewRequest("POST", "http://localhost:8080", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := serve(req)

	assertStatus200(t, w)
	if !strings.Contains(readAllBody(w), "You have filled this thing: abc") {
		t.Error("Expected string not found")
	}
}