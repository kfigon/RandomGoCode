package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHealthchecks(t *testing.T) {
	testCases := []struct {
		desc	string
		fn http.HandlerFunc
	}{
		{"map", healthcheckMap },
		{"typed", healthcheckTyped },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
	
			tC.fn(rec, req)
			res := rec.Result()
			
			assertEqual(t, http.StatusOK, res.StatusCode)
			assertEqual(t, "application/json", res.Header.Get("Content-type"))
		})
	}
}