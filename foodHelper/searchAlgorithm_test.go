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

var mockedFoods = []food {
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

func TestIngredients(t *testing.T) {
	testCases := []struct {
		desc	string
		algType includeStrategy
		ingredients *set
		expected []food
	}{
		{ "default_EmptyIngredients_thenEmptyResult", defaultStrategy, newSet(), []food{}},
		{ "default_InvalidIngredients_thenEmptyResult", defaultStrategy, newSet(noodle,bread), []food{}},
		{ "default_IdealHit", defaultStrategy, newSet(int(salad),int(cheese),int(apple)), []food{mockedFoods[2]}},
		{ "default_IdealHit_differentOrder", defaultStrategy, newSet(int(apple), int(cheese), int(salad)), []food{mockedFoods[2]}},

		{ "80_EmptyIngredients_thenEmptyResult", eightyPercent, newSet(), []food{}},
		{ "80_InvalidIngredients_thenEmptyResult", eightyPercent, newSet(noodle,bread), []food{}},
		{ "80_IdealHit", eightyPercent, newSet(int(salad),int(cheese),int(apple)), []food{mockedFoods[2]}},
		{ "80_IdealHit_differentOrder", eightyPercent,newSet(int(apple), int(cheese), int(salad)), []food{mockedFoods[2]}},
		{ "80_2hit", eightyPercent,newSet(int(apple), int(chicken), int(salad)), []food{}},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			alg := newSearch(createMockDb())
		
			results := alg.findFoods(tc.ingredients, tc.algType)

			if ln := len(results); ln != len(tc.expected) {
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

func TestBetterTestcasesForAlgorithm(t *testing.T) {
	t.Fail()
}