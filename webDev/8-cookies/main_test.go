package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
)

func TestHello(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/")
	expStatus := http.StatusOK
	if gotStatus := resp.StatusCode; gotStatus != expStatus {
		t.Errorf("Wrong status, got %v, exp %v", gotStatus, expStatus)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	if !strings.Contains(strBody, "Hi!") {
		t.Errorf("Invalid body received: %q", strBody)
	}
}