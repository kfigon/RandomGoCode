package recommendfood

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createEmptyTabWithLen(len int) []string {
	out := make([]string, len)
	for i := 0; i < len; i++ {
		out[i] = "abc"
	}
	return out
}

func TestValidateInput(t *testing.T) {
	tdt := []struct {
		desc          string
		input         []string
		expectedError bool
	}{
		{desc: "nil input", input: nil, expectedError: true},
		{desc: "empty input", input: []string{}, expectedError: true},
		{desc: "too big input", input: createEmptyTabWithLen(100), expectedError: true},
		{desc: "too big input2", input: createEmptyTabWithLen(21), expectedError: true},
		{desc: "too long ingredient", input: []string{"asd", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, expectedError: true},
		{desc: "ingredients ok", input: []string{"asd", "sad"}, expectedError: false},
		{desc: "ingredients ok2", input: createEmptyTabWithLen(20), expectedError: false},
	}
	for _, tc := range tdt {
		t.Run(tc.desc, func(t *testing.T) {
			err := validateIngredientNames(tc.input)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindFoodsInvalidInput(t *testing.T) {
	controller := NewRecommendationController(ingredientsMock{}, NewSearch(createMockDb()))
	result := controller.FindFoods(nil)
	assert.Empty(t, result)
}
