package recommendfood

import (
	"fmt"
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
	controller := NewRecommendationController(createMockIngredientsDb(), NewSearch(createMockFoodDb()))
	result := controller.FindFoods(nil)
	assert.Empty(t, result)
}

func TestFindFoodsInvalidInput2(t *testing.T) {
	controller := NewRecommendationController(createMockIngredientsDb(), NewSearch(createMockFoodDb()))
	result := controller.FindFoods(createEmptyTabWithLen(100))
	assert.Empty(t, result)
}

func TestFindFoods(t *testing.T) {
	controller := NewRecommendationController(createMockIngredientsDb(), NewSearch(createMockFoodDb()))
	result := controller.FindFoods([]string{"egg"})
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "Scrambled eggs", result[0].Name)
	assert.Equal(t, []string{"egg"}, result[0].IngredientNames)
}

func TestFindFoods2(t *testing.T) {
	controller := NewRecommendationController(createMockIngredientsDb(), NewSearch(createMockFoodDb()))
	result := controller.FindFoods([]string{"egg", "chicken", "salmon"})
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "first", result[0].Name)
	assert.Equal(t, []string{"egg", "chicken", "salmon"}, result[0].IngredientNames)

	assert.Equal(t, "Scrambled eggs", result[1].Name)
	assert.Equal(t, []string{"egg"}, result[1].IngredientNames)
}

func TestFindFoodsNameAdjuster(t *testing.T) {
	tdt := []struct {
		input              string
		expected           string
		expectedId         int
		expectedFoundState foundState
	}{
		{"egg", "egg", 0, FOUND},
		{"eg", "egg", 0, ADJUSTED},
		{"eggg", "egg", 0, ADJUSTED},
		{"kegg", "egg", 0, ADJUSTED},
		{"chicken", "chicken", 1, FOUND},
		{"chcken", "chicken", 1, ADJUSTED},
		{"chckn", "chicken", 1, ADJUSTED},
		{"chickn", "chicken", 1, ADJUSTED},
		{"cihcken", "chicken", 1, ADJUSTED},
		{"chickens", "chicken", 1, ADJUSTED},
		{"chicknes", "chicken", 1, ADJUSTED},
		{"asdjkhadfs", "", 0, NOT_FOUND},
	}

	controller := NewRecommendationController(createMockIngredientsDb(), NewSearch(createMockFoodDb()))
	for _, tc := range tdt {
		t.Run(tc.input, func(t *testing.T) {
			result := controller.adjustName(tc.input)
			assert.Equal(t, tc.expected, result.foundName)
			assert.Equal(t, tc.expectedId, result.foundId)
			assert.Equal(t, tc.expectedFoundState, result.matchResult)
		})
	}
}

func TestHammingDistance(t *testing.T) {
	tdt := []struct {
		a        string
		b        string
		expected int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"", "ab", 2},
		{"ab", "", 2},
		{"abc", "abc", 0},
		{"abc", "abd", 1},
		{"babc", "abc", 1},
		{"abc", "babc", 1},
		{"abcd", "abc", 1},
		{"abc", "abcd", 1},
	}

	for _, tc := range tdt {
		t.Run(fmt.Sprintf("%v_%v", tc.a, tc.b), func(t *testing.T) {
			res := calcHammingDistance(tc.a, tc.b)
			assert.Equal(t, tc.expected, res)
		})
	}
}
