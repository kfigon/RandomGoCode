package sorting

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