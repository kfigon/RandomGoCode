package findPair

import (
	"testing"
	"fmt"
	"sort"
)

// # given an unsorted array and a number, find if there's a pair
// # whose difference is num

// # a) space O(n), time O(n)
// # b) space O(1), time O(n logn)

// set would be nice...
func makeMap(data []int) map[int]int {
	dict := make(map[int]int)
	for _,v := range data {
		dict[v]++
	}
	return dict
}

func findPair(data []int, difference int) bool {
	dict := makeMap(data)

	for _, v := range data {
		toFind := v - difference
		occurences, found := dict[toFind]
		if found {
			if toFind == v && occurences > 1 {
				return true
			} else if toFind != v {
				return true
			}
		}
	}
	return false
}

func binarySearch(data []int, toFind int) int {
	left := 0
	right := len(data) - 1

	for left <= right {
		middleIdx := (right+left)/2
		value := data[middleIdx]
		if toFind == value {
			return middleIdx
		} else if toFind < value {
			right = middleIdx -1
		} else {
			left = middleIdx + 1
		}
	}

	return -1
}

func isMoreThanOne(foundIdx int, data []int) bool {
	toFind := data[foundIdx]
	if foundIdx -1 >= 0 && data[foundIdx-1] == toFind {
		return true
	} else if foundIdx + 1 < len(data) && data[foundIdx+1] == toFind {
		return true
	}
	return false
}

func findPair2(data []int, difference int) bool {
	sort.Ints(data)
	for _,v := range data {
		toFind := v - difference
		foundIdx := binarySearch(data, toFind)

		if foundIdx != -1 {
			if toFind == v && isMoreThanOne(foundIdx, data) {
				return true
			} else if toFind != v {
				return true
			}
		}	
	}
	return false
}

func TestFirst(t *testing.T) {
	testThings(t, findPair)
}

func TestBinarySearch(t *testing.T)  {
	tt := []struct {
		in []int
		toFind int
		exp int
	}{
		{ in: []int{1,2,3,4,5,6,7}, toFind: 2, exp: 1 },
		{ in: []int{1,2,3,4,5,6,7}, toFind: 1, exp: 0 },
		{ in: []int{1,2,3,4,5,6,7}, toFind: 6, exp: 5 },

		{ in: []int{1,2,3,4,5,6,7}, toFind: 7, exp: 6 },
		{ in: []int{1,2,3,4,5,6,7}, toFind: 4, exp: 3 },
		{ in: []int{1,2,3,4,5,6,7}, toFind: 8, exp: -1 },
		{ in: []int{1,2,3,4,5,6,7}, toFind: 0, exp: -1 },

		{ in: []int{1,2,3,4,5,6}, toFind: 2, exp: 1 },
		{ in: []int{1,2,3,4,5,6}, toFind: 1, exp: 0 },
		{ in: []int{1,2,3,4,5,6}, toFind: 7, exp: -1 },
		{ in: []int{1,2,3,4,5,6}, toFind: 4, exp: 3 },
		{ in: []int{1,2,3,4,5,6}, toFind: 6, exp: 5 },
		{ in: []int{1,2,3,4,5,6}, toFind: 8, exp: -1 },
		{ in: []int{1,2,3,4,5,6}, toFind: 0, exp: -1 },
	}
	for _, tC := range tt {
		t.Run(fmt.Sprintf("%v - %v", tC.in, tC.exp), func(t *testing.T) {
			if res := binarySearch(tC.in, tC.toFind); res != tC.exp {
				t.Errorf("got %v, exp %v", res, tC.exp)
			}
		})
	}
}

func TestSecond(t *testing.T) {
	testThings(t, findPair2)
}

type myFun func(data []int, difference int) bool
// func testThings(t *testing.T, testToFun func(data []int, difference int) bool) {
func testThings(t *testing.T, testToFun myFun) {
	tt := []struct {
		in []int
		difference int
		exp bool
	} {
		// {in: []int{6,1,4,10,2,4}, difference: 2, exp: true},
		// {in: []int{8,6,2,4,1,0,2,5,13}, difference: 1, exp: true},
		// {in: []int{4,-2,3,10}, difference: -6, exp: true},
		{in: []int{4,-2,3,10}, difference: 0, exp: false},
		// {in: []int{5,5}, difference: 0, exp: true},
		// {in: []int{-4,4}, difference: -8, exp: true},
		// {in: []int{-4,4}, difference: 8, exp: true},
		// {in: []int{1}, difference: 1, exp: false},
		// {in: []int{1, 0}, difference: 1, exp: true},
		// {in: []int{1, 1, 3}, difference: 0, exp: true},
		// {in: []int{1,3,4,6}, difference: -2, exp: true},
		// {in: []int{0,1,3,4,6}, difference: -2, exp: true},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprint(tc.in), func (t *testing.T) {
			if res := testToFun(tc.in, tc.difference); res != tc.exp {
				t.Errorf("Exp %v, got %v", tc.exp, res)
			}
		})
	}
}