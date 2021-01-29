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

	resp, _ := http.Get(srv.URL +"/login")
	expStatus := http.StatusOK
	if gotStatus := resp.StatusCode; gotStatus != expStatus {
		t.Errorf("Wrong status, got %v, exp %v", gotStatus, expStatus)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	if !strings.Contains(strBody, "Hi!") {
		t.Errorf("Invalid body received: %q", strBody)
	}

	cookies := resp.Cookies()
	if cookies == nil || len(cookies) == 0 {
		t.Fatal("No cookies present in response")
	}
	if len(cookies) > 1 {
		t.Error("Expected only 1 cookie. There are others in response: ", cookies)
	}
	if cookies[0].Name != "ziomCookie" {
		t.Error("Invalid cookie received: ", cookies[0])
	}
	if len(cookies[0].Name) == 0 {
		t.Error("Empty cookie received")
	}
}

func TestLoginWithoutCookie(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/")
	expStatus := http.StatusForbidden
	if gotStatus := resp.StatusCode; gotStatus != expStatus {
		t.Errorf("Wrong status, got %v, exp %v", gotStatus, expStatus)
	}
}


func TestLoginWithCookie(t *testing.T) {
	srv := httptest.NewServer(createMux())
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/")
	expStatus := http.StatusOK
	if gotStatus := resp.StatusCode; gotStatus != expStatus {
		t.Errorf("Wrong status, got %v, exp %v", gotStatus, expStatus)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	if !strings.Contains(strBody, "This is secret data") {
		t.Errorf("Invalid body received: %q", strBody)
	}
}
