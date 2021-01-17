package areThereDuplicates

import (
	"fmt"
	"testing"
	"sort"
)

// # check wheter there are duplicates in array
// O(nlogn)
// dont use maps
func areThereDuplicates(tab []int) bool {
	sort.Ints(tab)
	for i := 0; i < len(tab); i++ {
		if i+1 < len(tab) && tab[i] == tab[i+1] {
			return true
		}
	}

	return false
}

func TestDup(t *testing.T) {
	testCases := []struct {
		in []int
		exp bool
	}{
		{in: nil, exp: false },
		{in: []int{}, exp: false },
		{in: []int{1,2,3}, exp: false },
		{in: []int{1,2,2}, exp: true },
		{in: []int{1,2,3,1}, exp: true },
		{in: []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}, exp: false },
		{in: []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,6,15,16,17,18,19,20}, exp: true },
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			if res := areThereDuplicates(tc.in); res != tc.exp {
				t.Errorf("Expected %v, got %v", tc.exp, res)
			}
		})
	}
}