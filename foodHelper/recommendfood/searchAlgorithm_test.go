package recommendfood

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredients(t *testing.T) {
	testCases := []struct {
		desc        string
		ingredients *set
		expected    []food
	}{
		{"default_EmptyIngredients_thenEmptyResult", newSet(), []food{}},
		{"default_InvalidIngredients_thenEmptyResult", newSet(noodle, bread), []food{}},
		{"default_IdealHit", newSet(salad, cheese, apple), []food{mockedFoods[2]}},
		{"default_IdealHit_differentOrder", newSet(apple, cheese, salad), []food{mockedFoods[2]}},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			alg := NewSearch(createMockDb())
			strategy := fitnessInclusionStrategy{100}
			results := alg.findFoods(tc.ingredients, strategy)

			assert.Equal(t, len(tc.expected), len(results))
			for i := range results {
				got := results[i]
				exp := tc.expected[i]

				assert.Equal(t, exp.Name, got.Name)
				assert.Equal(t, 100, got.FitnessLevel)
			}
		})
	}
}
