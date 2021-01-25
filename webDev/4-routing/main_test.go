package main

import (
	"testing"
	"net/http/httptest"
	"io/ioutil"
)

func TestRoot(t *testing.T) {
	// request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	response := httptest.NewRecorder()

	status := response.Result().StatusCode
	if status != 200 {
		t.Error("Wanted 200, got", status)
	}
	body, err := ioutil.ReadAll(response.Result().Body)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	strBody := string(body)
	if strBody != "Hello World" {
		t.Error("Invalid body", strBody)
	}
}