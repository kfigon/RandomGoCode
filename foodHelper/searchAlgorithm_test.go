package main

import (
	"testing"
)

type mockDb struct{
	findFoodFun func() []food
}

func (m mockDb) findFoods() []food {
	return m.findFoodFun()
}

type ingredientsId int
const (
	egg ingredientsId = 1
	chicken = 2
	beef = 3
	salmon = 4
	salad = 5
	cheese = 6
	apple = 7
	noodle = 80
	bread = 90
)

func createMockDb() mockDb {
	return mockDb{
		findFoodFun: func() []food {
			return []food{
				food{"first", newSet(int(egg),int(chicken),int(salmon))},
				food{"second", newSet(int(egg),int(chicken),int(salad))},
				food{"third", newSet(int(salad),int(cheese),int(apple))},
			}
		},
	}
}

func TestWhenEmptyIngredients_thenEmptyResult(t *testing.T) {
	ingredients := newSet()
	alg := newSearch(createMockDb())

	results := alg.findFoods(ingredients)

	if ln := len(results); ln != 0 {
		t.Error("Expected empty result, got: ", ln)
	}
}

func TestWhenInvalidIngredients_thenEmptyResult(t *testing.T) {
	ingredients := newSet(noodle,bread)
	alg := newSearch(createMockDb())

	results := alg.findFoods(ingredients)

	if ln := len(results); ln != 0 {
		t.Error("Expected empty result, got: ", ln)
	}
}
