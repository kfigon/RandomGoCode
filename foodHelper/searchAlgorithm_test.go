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

var mockedFoods = []food{
	food{"first", newSet(int(egg),int(chicken),int(salmon))},
	food{"second", newSet(int(egg),int(chicken),int(salad))},
	food{"third", newSet(int(salad),int(cheese),int(apple))},
}

func createMockDb() mockDb {
	return mockDb{
		findFoodFun: func() []food {
			return mockedFoods
		},
	}
}

func TestIngredientsDefaultStrategy(t *testing.T) {
	testCases := []struct {
		desc	string
		ingredients *set
		expected []food
	}{
		{ "EmptyIngredients_thenEmptyResult", newSet(), []food{}},
		{ "InvalidIngredients_thenEmptyResult", newSet(noodle,bread), []food{}},
		{ "IdealHit", newSet(int(salad),int(cheese),int(apple)), []food{mockedFoods[2]}},
		{ "IdealHit_differentOrder", newSet(int(apple), int(cheese), int(salad)), []food{mockedFoods[2]}},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			alg := newSearch(createMockDb())
		
			results := alg.findFoods(tc.ingredients, defaultStrategy)

			if ln := len(results); ln != len(tc.expected) {
				t.Fatalf("Invalid len got: %v, exp %v", ln, len(tc.expected))
			}

			for i := range results {
				got := results[i]
				exp := tc.expected[i]
				if exp.name != got.f.name {
					t.Errorf("Got invalid food (%v), got: %v, exp: %v", i, got.f.name, exp.name)
				}
				if got.fitnessLevel != 100 {
					t.Errorf("Default strategy requires perfect fitness, got: %v", got.fitnessLevel)
				}
			}
		})
	}
}

func TestRankingStrategy(t *testing.T) {
	t.Fail()
}
