package main

import (
	"testing"
)

type mockDb struct{

}

func TestWhenEmptyIngredients_thenEmptyResult(t *testing.T) {
	ingredients := newSet()
	alg := newSearch(mockDb{})

	results := alg.findFoods(ingredients)

	if ln := len(results); ln != 0 {
		t.Error("Expected empty result, got: ", ln)
	}
}

