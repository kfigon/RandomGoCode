package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFitnessStrategy(t *testing.T) {
	testCases := []struct {
		desc              string
		threshold         int
		in                []int
		required          []int
		expectedInclusion bool
	}{
		{"100HitStrat", 100, []int{2, 3, 4}, []int{2, 3, 4}, true},
		{"100_rearranged", 100, []int{4, 3, 2}, []int{2, 3, 4}, true},
		{"100_emptyUserInput", 100, []int{}, []int{2, 3, 4}, false},
		{"100_notMet", 100, []int{1}, []int{2, 3, 4}, false},
		{"100_notMet2", 100, []int{1, 2}, []int{2, 3, 4}, false},
		{"100_notMet3", 100, []int{2, 4}, []int{2, 3, 4}, false},
		{"50_notIncluded", 50, []int{2}, []int{2, 3, 4}, false},
		{"50_included", 50, []int{2, 4}, []int{2, 3, 4}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			strat := fitnessInclusionStrategy{tc.threshold}

			result := strat.shouldBeIncluded(newSet(tc.in...), newSet(tc.required...))
			assert.Equal(t, tc.expectedInclusion, result)
		})
	}
}
func TestFitness(t *testing.T) {
	tt := []struct {
		desc     string
		in       []int
		required []int
		exp      int
	}{
		{"66%", []int{2, 3}, []int{1, 2, 3}, 66},
		{"33%", []int{2}, []int{1, 2, 3}, 33},
		{"0%", []int{}, []int{1, 2, 3}, 0},
		{"100%", []int{3, 2, 1}, []int{1, 2, 3}, 100},
		{"75%", []int{3, 2, 1}, []int{1, 2, 3, 4}, 75},
		{"75%_additionalVals", []int{3, 2, 1, 10, 11, 12}, []int{1, 2, 3, 4}, 75},
		{"100%_additionalVals", []int{3, 2, 1, 4, 10, 11, 12}, []int{1, 2, 3, 4}, 100},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			strat := fitnessInclusionStrategy{100}

			required := newSet(tc.required...)
			users := newSet(tc.in...)

			result := strat.calcFitness(users, required)

			assert.Equal(t, tc.exp, result)
		})
	}
}
