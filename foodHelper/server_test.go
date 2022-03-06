package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
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


func TestHandleIngredients(t *testing.T) {
	testCases := []struct {
		input	string
		expected []string
	}{
		{"h", []string{"hi", "hello", "hell", "howdy"}},
		{"hell", []string{"hello", "hell"}},
		{"ho", []string{"howdy"}},
		{"asd", []string{}},
		{"", []string{}},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			words := []string{"hi", "hello", "hell", "howdy"}
			tr := createTrie(words)

			req := httptest.NewRequest("GET", "/ing?prefix="+tC.input, nil)
			rec := httptest.NewRecorder()

			handleIngredients(tr)(rec, req)
			resp := rec.Result()
			defer resp.Body.Close()

			body := map[string][]string{}
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
		
			assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
			assert.Equal(t, 200, resp.StatusCode)
			assert.ElementsMatch(t, tC.expected, body["ingredients"])
		})
	}
}

func TestSuggestions(t *testing.T) {
	inputJson := `
		{
			"ingredients": ["first","second","foo"]
		}`

	req := httptest.NewRequest("POST", "/suggestions", strings.NewReader(inputJson))
	rec := httptest.NewRecorder()

	handleSuggestions(testData())(rec, req)
	resp := rec.Result()
	defer resp.Body.Close()

	body := map[string][]string{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "foo", body["results"])
}

func testData() dataProvider {
	return nil
}