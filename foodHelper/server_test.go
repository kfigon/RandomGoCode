package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthcheck", nil)
	rec := httptest.NewRecorder()

	healthcheck(rec, req)
	resp := rec.Result()
	defer resp.Body.Close()

	body := map[string]string{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "ok", body["status"])
}