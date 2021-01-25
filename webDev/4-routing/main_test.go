package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"io/ioutil"
)

func assertStatusGetBody(t *testing.T, expStatus int, response *httptest.ResponseRecorder) string {
	status := response.Result().StatusCode
	if status != expStatus {
		t.Errorf("Wanted %v, got %v", expStatus, status)
	}
	body, err := ioutil.ReadAll(response.Result().Body)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	return string(body)
}

func TestRoot(t *testing.T) {
	// request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	response := httptest.NewRecorder()

	strBody := assertStatusGetBody(t, http.StatusOK, response)
	if strBody != "Hello World" {
		t.Error("Invalid body", strBody)
	}
}