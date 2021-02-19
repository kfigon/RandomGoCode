package main

import (
	"testing"
	"fmt"
)

type sortingAlg func([]int) []int

func TestSort(t *testing.T) {
	tdt := []struct {
		in []int
		exp []int
	} {
		{[]int{1,2,3,4,5}, []int{1,2,3,4,5} },
		{[]int{1,2,3,4,5,6}, []int{1,2,3,4,5,6} },
		{[]int{6,5,4,3,2,1}, []int{1,2,3,4,5,6} },
		{[]int{1,5,2,4,3,6}, []int{1,2,3,4,5,6} },
		{[]int{1,5,2,4,3}, []int{1,2,3,4,5} },
		{[]int{6,5,3,1,8,7,2,4}, []int{1,2,3,4,5,6,7,8} },
	}
	algorithms := []struct {
		desc string
		alg sortingAlg
	} {
		{"bubble", bubbleSort},
		{"insertion", insertionSort},
		{"selection", selectionSort},
		{"merge", mergeSort},
		{"quick", quickSort},
		{"mergeParallel", mergeSortParallel},
		{"raidx", raidxSort},
	}

	for _,alg := range algorithms {
		for _, tc := range tdt {
			t.Run(fmt.Sprintf("%v->%v", alg.desc, tc.in), func(t *testing.T) {
				got := alg.alg(tc.in)
				assertResult(t, got, tc.exp)
			})
		}
	}
}

func assertResult(t *testing.T, got []int, exp []int) {
	if len(got) != len(exp) {
		t.Fatalf("Invalid len, got %v, exp %v", len(got), len(exp))
	}
	for i := 0; i < len(exp); i++ {
		expEl := exp[i]
		gotEl := got[i]

		if expEl != gotEl {
			t.Errorf("Error at %v, got %v, exp %v", i, gotEl, expEl)
			t.Errorf("%v != %v", got, exp)
			break
		}
	}
}

func TestMergeTabs(t *testing.T) {
	tdt := [] struct {
		desc string
		a []int
		b []int
		exp []int
	} {
		{"empty", []int{}, []int{}, []int{}},
		{"singleLeft", []int{4}, []int{}, []int{4}},
		{"singleRight", []int{}, []int{4}, []int{4}},
		{"singleBoth", []int{4}, []int{4}, []int{4,4}},
		{"singleBoth2", []int{2}, []int{3}, []int{2,3}},
		{"emptyRight", []int{2,3,4}, []int{}, []int{2,3,4}},
		{"emptyLeft", []int{}, []int{2,3,4}, []int{2,3,4}},
		{"uneven1", []int{1,3,4,6,8}, []int{2,5}, []int{1,2,3,4,5,6,8}},
		{"uneven2", []int{2,5}, []int{1,3,4,6,8}, []int{1,2,3,4,5,6,8}},
		{"even1", []int{1,2,3}, []int{4,5,6}, []int{1,2,3,4,5,6}},
		{"even2", []int{1,5,6}, []int{2,3,4}, []int{1,2,3,4,5,6}},
	}

	for _,tc := range tdt {
		t.Run(tc.desc, func(t *testing.T) {
			got := mergeTabs(tc.a, tc.b)
			assertResult(t, got, tc.exp)
		})
	}
}

func TestExtractDigit(t *testing.T) {
	testCases := []struct {
		number int
		digitNum int
		expected int
	}{
		{0,0,0},
		{0,1,0},
		{0,2,0},
		{1,0,1},
		{1,1,0},
		{34,0,3},
		{34,1,4},
		{34,2,0},
		{567893,0,3},
		{567893,1,9},
		{567893,2,8},
		{567893,3,7},
		{567893,4,6},
		{567893,5,5},
		{567893,6,0},
		{567893,7,0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v[%v]", tc.number, tc.digitNum), func(t *testing.T) {
			got := extractDigit(tc.number, tc.digitNum)
			if got != tc.expected {
				t.Errorf("%v[%v], got %v, exp %v", tc.number, tc.digitNum, got, tc.expected)
			}
		})
	}
}