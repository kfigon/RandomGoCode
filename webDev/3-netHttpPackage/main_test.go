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
}
func TestGet(t *testing.T) {
	w := serve(httptest.NewRequest("GET", "http://localhost:8080", nil))

	assertStatus200(t, w)
	body,_ := ioutil.ReadAll(w.Result().Body)
	if strings.Contains(string(body), "You have filled this thing") {
		t.Error("Non expected string received")
	}
}

func TestGetWithQuery(t *testing.T) {
	w := serve(httptest.NewRequest("GET", "http://localhost:8080?asd=foo", nil))

	assertStatus200(t, w)
	body,_ := ioutil.ReadAll(w.Result().Body)
	if strings.Contains(string(body), "You have filled this thing") {
		t.Error("Non expected string received")
	}
}

func TestGetWithRightQuery(t *testing.T) {
	w := serve(httptest.NewRequest("GET", "http://localhost:8080?myName=foo", nil))

	assertStatus200(t, w)
	body,_ := ioutil.ReadAll(w.Result().Body)
	if !strings.Contains(string(body), "You have filled this thing: foo") {
		t.Error("Expected string not found")
	}
}

func TestPost(t *testing.T) {
	reader := strings.NewReader("myName=2")
	req := httptest.NewRequest("POST", "http://localhost:8080", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := serve(req)

	assertStatus200(t, w)
	body,_ := ioutil.ReadAll(w.Result().Body)
	if !strings.Contains(string(body), "You have filled this thing: 2") {
		t.Error("Expected string not found")
	}
}