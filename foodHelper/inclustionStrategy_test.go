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

func Test50HitStrat(t *testing.T)  {
	strat := fitnessInclusionStrategy{50}
	required := newSet(int(1), int(2), int(3))
	users := newSet(int(2), int(3))
	
	result := strat.shouldBeIncluded(users, required)
	expected := true
	if result != expected {
		t.Errorf("Got %v, exp: %v", result, expected)
	}
}

func TestFitness(t *testing.T)  {
	strat := fitnessInclusionStrategy{50}
	required := newSet(int(1), int(2), int(3))
	users := newSet(int(2), int(3))
	
	result := strat.calcFitness(users, required)
	expected := 66
	if result != expected {
		t.Errorf("Got %v, exp: %v", result, expected)
	}
}

func TestBetterTestcasesForAlgorithm(t *testing.T) {
	t.Fail()
}