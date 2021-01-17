package findAllDuplicates

import (
	"fmt"
	"testing"
)

// # given an array of positive ints, some elements appear
// # twice, others once. Find all elements that appear twice

func makeMap(data []int) map[int]int {
	dict := make(map[int]int)
	for _,v := range data {
		dict[v]++
	}
	return dict
}

func findAllDuplicates(data []int) []int {
	result := make([]int, 0)
	dict := makeMap(data)
	for k := range dict {
		v := dict[k]
		if v == 2 {
			result = append(result, k)
		}
	}
	return result
}

func Test(t *testing.T) {
	testCases := []struct {
		in []int
		exp []int
	}{
		{in :[]int{4,3,2,7,8,2,3,1}, exp: []int{3,2}},
		{in :[]int{4,3,2,1,0}, exp: []int{}},
		{in :[]int{4,3,2,1,0,1,2,3}, exp: []int{3,2,1}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			compare(tc.exp, findAllDuplicates(tc.in), t)
		})
	}
}

func compare(exp []int, result []int, t *testing.T) {
	if len(exp) != len(result) {
		t.Fatalf("Exp len %v, actual len %v", len(exp), len(result))
	}

	contains := func (toFind int) bool {
		for _,v := range result {
			if v == toFind{ 
				return true
			}
		}
		return false
	}

	for i := range exp {
		expectedVal := exp[i]
		if !contains(expectedVal) {
			t.Errorf("%v not found in result", expectedVal)
		}
	}
}