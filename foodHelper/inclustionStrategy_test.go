package main

import (
	"testing"
)

func Test100HitStrat(t *testing.T)  {
	strat := fitnessInclusionStrategy{100}
	required := newSet(int(1), int(2), int(3))
	users := newSet(int(2), int(3))
	
	result := strat.shouldBeIncluded(users, required)
	expected := false
	if result != expected {
		t.Errorf("Got %v, exp: %v", result, expected)
	}
}

func TestBetterTestcasesForAlgorithm(t *testing.T) {
	t.Fail()
}