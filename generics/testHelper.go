package main

import "testing"

func assertEqual[T comparable](t *testing.T, exp, got T) {
	if got != exp {
		t.Errorf("Exp %v, got %v", exp, got)
	}
}

func assertEqualArr[T comparable](t *testing.T, exp, got []T) {
	if len(exp) != len(got) {
		t.Fatalf("Len exp %v, len got %v", len(exp), len(got))
	}

	for i := 0; i < len(got); i++ {
		if exp[i] != got[i] {
			t.Errorf("Error on idx %v, exp %v, got %v", i, exp[i], got[i])
		}
	}
}