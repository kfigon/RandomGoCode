package main

import (
	"testing"
)

func TestFitnessStrategy(t *testing.T) {
	testCases := []struct {
		desc	string
		threshold int
		in []int
		required []int
		expectedInclusion bool
	}{
		{ "100HitStrat", 100, []int{2,3,4}, []int{2,3,4}, true, },
		{ "100_rearranged", 100, []int{4,3,2}, []int{2,3,4}, true, },
		{ "100_emptyUserInput", 100, []int{}, []int{2,3,4}, false, },
		{ "100_notMet", 100, []int{1}, []int{2,3,4}, false, },
		{ "100_notMet2", 100, []int{1,2}, []int{2,3,4}, false, },
		{ "100_notMet3", 100, []int{2,4}, []int{2,3,4}, false, },
		{ "50_notIncluded", 50, []int{2}, []int{2,3,4}, false, },
		{ "50_included", 50, []int{2,4}, []int{2,3,4}, true, },
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			strat := fitnessInclusionStrategy{tc.threshold}
			
			result := strat.shouldBeIncluded(newSet(tc.in...), newSet(tc.required...))
			if result != tc.expectedInclusion {
				t.Errorf("Got %v, exp: %v", result, tc.expectedInclusion)
			}
		})
	}
}
func TestFitness(t *testing.T)  {
	strat := fitnessInclusionStrategy{-1}
	required := newSet(int(1), int(2), int(3))
	users := newSet(int(2), int(3))
	
	result := strat.calcFitness(users, required)
	expected := 66
	if result != expected {
		t.Errorf("Got %v, exp: %v", result, expected)
	}
}