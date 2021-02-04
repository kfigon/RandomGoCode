package main

import (
	"testing"
)

func assertElements(t *testing.T, tr *tree, exp []int) {
	if v := tr.size(); v != len(exp) {
		t.Errorf("Invalid size, exp: %v got: %v", len(exp), v)
	}
	vals := tr.values()
	if ln := len(vals); ln != len(exp) {
		t.Fatalf("Invalid size on values, exp: %v got: %v", len(exp), len(vals))
	}
	for i := range exp {
		e := exp[i]
		got := vals[i]
		if e != got {
			t.Errorf("Invalid element id %v exp %v, got: %v", i, e, got)
		}
	}
}

func TestInserts(t *testing.T) {
	testCases := []struct {
		desc string
		in []int
		exp []int		
	}{
		{"empty", []int{}, []int{}},
		{"single elem", []int{5}, []int{5}},
		{"elements in order", []int{1,2,3,4,5}, []int{1,2,3,4,5}},
		{"elements in not bad order", []int{5,4,3,2,1,}, []int{1,2,3,4,5}},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tr := newTree()
			for _,v := range tc.in {
				tr.insert(v)
			}

			assertElements(t, tr, tc.exp)
		})
	}
}

