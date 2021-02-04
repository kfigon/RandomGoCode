package main

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	tr := newTree()
	assertElements(t, tr, []int{})
}

func assertElements(t *testing.T, tr *tree, exp []int) {
	if v := tr.size(); v != len(exp) {
		t.Errorf("Invalid size, exp: %v got: %v", len(exp), v)
	}
	vals := tr.values()
	if ln := len(vals); ln != len(exp) {
		t.Errorf("Invalid size on values, exp: %v got: %v", len(exp), len(vals))
	}
	for i := range exp {
		e := exp[i]
		got := vals[i]
		if e != got {
			t.Errorf("Invalid element id %v exp %v, got: %v", i, e, got)
		}
	}
}

func TestOneElement(t *testing.T) {
	tr := newTree()
	tr.insert(5)
	assertElements(t, tr, []int{5})
}

