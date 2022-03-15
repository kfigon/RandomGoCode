package main

import "testing"

func assertEqual[T comparable](t *testing.T, exp, got T) {
	if got != exp {
		t.Errorf("Exp %v, got %v", exp, got)
	}
}