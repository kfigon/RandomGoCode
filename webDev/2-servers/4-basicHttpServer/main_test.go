package main

import (
	"net/http"
	"testing"
)

type singleServerRunStrategy struct{
	iterationNum int
}
func (s *singleServerRunStrategy) run() bool {
	if s.iterationNum > 0 {
		return false
	}
	s.iterationNum++
	return true
}

func performRequestSync(requestFun func()(*http.Response, error)) (*http.Response, error) {
	ch := make(chan bool)

	strategy := &singleServerRunStrategy{}
	go func() {
		startServer(strategy)
		<-ch
	}()

	res, err := requestFun()
	ch<-true
	return res, err
}

func Test200(t *testing.T) {
	resp, err := performRequestSync(func()(*http.Response, error) {return http.Get("http://localhost:8080/")})

	if err != nil {
		t.Fatal("Error not expected, got:", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("Status code != 200:", resp.StatusCode)
	}
}

func Test404(t *testing.T) {
	resp, err := performRequestSync(func()(*http.Response, error) {return http.Get("http://localhost:8080/asd")})

	if err != nil {
		t.Fatal("Error not expected, got:", err)
	}
	if resp.StatusCode != 404 {
		t.Fatal("Status code != 404:", resp.StatusCode)
	}
}

func Test404_2(t *testing.T) {
	resp, err := performRequestSync(func()(*http.Response, error) {return http.Post("http://localhost:8080/x", "application/json", nil)})

	if err != nil {
		t.Fatal("Error not expected, got:", err)
	}
	if resp.StatusCode != 404 {
		t.Fatal("Status code != 404:", resp.StatusCode)
	}
}