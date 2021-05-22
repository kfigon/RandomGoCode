package main

import (
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"encoding/json"
)

func TestNoPass(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	response, err := http.Get(srv.URL+"/login")
	
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestInvalidPass(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	request,err := http.NewRequest(http.MethodGet, srv.URL+"/login", nil)
	request.SetBasicAuth("foo","bar")

	assert.NoError(t,err)
	response, err := (&http.Client{}).Do(request)
	
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestValidPass(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	request,err := http.NewRequest(http.MethodGet, srv.URL+"/login", nil)
	request.SetBasicAuth("John","secret")

	assert.NoError(t,err)
	response, err := srv.Client().Do(request)
	
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseBody LoginResponse
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.NotEmpty(t, responseBody.Token)
}

func TestAuthenticateWhenEmpty(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	request,err := http.NewRequest(http.MethodGet, srv.URL+"/auth", nil)
	assert.NoError(t, err)
	
	response,err := srv.Client().Do(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestAuthenticateWhenInvalidToken(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	request,err := http.NewRequest(http.MethodGet, srv.URL+"/auth", nil)
	request.Header.Add("MY_TOKEN", "foobar")
	assert.NoError(t, err)
	
	response,err := srv.Client().Do(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestAuthenticateWhenValidToken(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	request,err := http.NewRequest(http.MethodGet, srv.URL+"/login", nil)
	request.SetBasicAuth("John","secret")

	response, err := srv.Client().Do(request)
	var responseBody LoginResponse
	json.NewDecoder(response.Body).Decode(&responseBody)

	request,err = http.NewRequest(http.MethodGet, srv.URL+"/auth", nil)
	request.Header.Add("MY_TOKEN", responseBody.Token)
	assert.NoError(t, err)
	
	response,err = srv.Client().Do(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

