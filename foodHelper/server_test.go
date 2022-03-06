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
			"ingredients": ["b","myyyyk","c","foo","a"]
		}`

	req := httptest.NewRequest("POST", "/suggestions", strings.NewReader(inputJson))
	rec := httptest.NewRecorder()

	handleSuggestions(testData())(rec, req)
	resp := rec.Result()
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
	type result struct {
		Name string `json:"Name"`
		Description string `json:"Description"`
		Ingredients []string `json:"Ingredients"`
	}
	type response struct {
		Results []result `json:"results"`
	}

	var out response
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&out))

	assert.Equal(t, 1, len(out.Results))
	assert.Equal(t, "Food a", out.Results[0].Name)
	assert.Equal(t, "delicious", out.Results[0].Description)
}

type testDataProvider func() []suggestion
func (t testDataProvider) allSuggestions() []suggestion {
	return t()
}

func testData() dataProvider {
	results := []suggestion {
		{
			name: "Food a",
			description: "delicious",
			ingredients: []string{"a", "b", "c", "d"},
		},
		{
			name: "Food B",
			description: "tasty",
			ingredients: []string{"b", "c", "d"},
		},
		{
			name: "Food C",
			description: "wow",
			ingredients: []string{"a", "b", "c", "e", "f"},
		},
	}
	f := func() []suggestion {
		return results
	}
	return testDataProvider(f)
}