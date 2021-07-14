package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	srv := httptest.NewServer(newServer(newLogin()))
	defer srv.Close()

	request,err := http.NewRequest(http.MethodGet, srv.URL+"/health", nil)
	assert.NoError(t, err)

	resp, err := srv.Client().Do(request)
	assert.NoError(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	
	content := string(body)
	exp := `"status":"up"`
	if !strings.Contains(content, exp) {
		assert.Failf(t, "Invalid content", "Got %v, exp %v", content, exp)
	}
}