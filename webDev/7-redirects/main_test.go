package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func TestRoot(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	resp, _ := http.Get(srv.URL+"/")
	if resp.StatusCode != http.StatusNotFound {
		t.Error("Invalid status. Got: ", resp.StatusCode)
	}
}

func TestResource(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	resp, _ := http.Get(srv.URL+"/resource")
	if resp.StatusCode != http.StatusOK {
		t.Error("Invalid status. Got: ", resp.StatusCode)
	}
	r, _ := ioutil.ReadAll(resp.Body)
	if string(r) != "Hi" {
		t.Error("invalid body received")
	}
}

func TestRedirect(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	resp, _ := http.Get(srv.URL+"/redirect")
	if resp.StatusCode != http.StatusFound { // 302
		t.Error("Invalid status. Exp 302, Got: ", resp.StatusCode)
	}
	r, _ := ioutil.ReadAll(resp.Body)
	if string(r) != "Hi" {
		t.Error("invalid body received")
	}
}