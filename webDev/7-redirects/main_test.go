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
	// 301 - moved permanently - browser will always use new address
	// 302 - found - just redirect to other URL, preserves method. "Legacy"
	// 303 - see other - like 302, but always will be GET. Nice for redirect to status page after POST
	// 307 - temporary redirected - preserves method

	// will redirect and return 200
	if resp.StatusCode != http.StatusOK {
		t.Error("Invalid status. Exp 200, Got: ", resp.StatusCode)
	}
	r, _ := ioutil.ReadAll(resp.Body)
	if string(r) != "Hi" {
		t.Error("invalid body received")
	}
}