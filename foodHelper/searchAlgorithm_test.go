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

func TestIngredients(t *testing.T) {
	testCases := []struct {
		desc	string
		ingredients *set
		expected []food
	}{
		{ "WhenEmptyIngredients_thenEmptyResult", newSet(), []food{}},
		{ "WhenInvalidIngredients_thenEmptyResult", newSet(noodle,bread), []food{}},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			alg := newSearch(createMockDb())
		
			results := alg.findFoods(tc.ingredients)

			if ln := len(results); ln != 0 {
				t.Fatalf("Invalid len got: %v, exp %v", ln, len(tc.expected))
			}

			for i := range results {
				got := results[i]
				exp := tc.expected[i]
				if exp.name != got.name {
					t.Errorf("Got invalid food (%v), got: %v, exp: %v", i, got.name, exp.name)
				}
			}
		})
	}
}

