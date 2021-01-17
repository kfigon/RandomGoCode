package isSubsequence

import (
	"fmt"
	"testing"
)

// # sumZero
// # input - sorted list of ints
// # find first pair when the sum is 0
func notFound() error {
	return fmt.Errorf("Not found")
}
func sumZero(tab []int) (int, int, error) {
	if len(tab) < 2 {
		return 0,0, notFound()
	}

	firstIdx := 0
	secondIdx := len(tab)-1
	for firstIdx <= secondIdx {
		sum := tab[firstIdx] + tab[secondIdx]
		if sum == 0 {
			return tab[firstIdx], tab[secondIdx], nil
		} else if sum < 0 {
			firstIdx++
		} else {
			secondIdx--
		}
	}
	return 0, 0, notFound()
}

type testResult struct {
	resA int
	resB int
	found bool
}

func Test(t *testing.T) {
	testCases := []struct {
		in []int
		exp testResult
	}{
		{in: nil, exp: testResult{found: false} },
		{in: []int{}, exp: testResult{found: false} },
		{in: []int{1,2}, exp: testResult{found: false} },
		{in: []int{-3}, exp: testResult{found: false} },
		{in: []int{-3,-2,-1,0,1,2,3}, exp: testResult{resA: -3, resB: 3, found: true} },
		{in: []int{-2,0,1,3}, exp:testResult{found: false} },
		{in: []int{1,2,3}, exp: testResult{found: false} },
		{in: []int{-4,-3,-2,-1,0,1,2,5}, exp: testResult{resA: -2, resB: 2, found: true} },
		{in: []int{-4,-3,-2,-1,0,1,2,3,10}, exp: testResult{resA: -3, resB: 3, found: true} },
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			a,b,err := sumZero(tc.in)
			shouldBeFound := tc.exp.found
			if shouldBeFound && err != nil {
				t.Fatalf("Should be found, but it's not")
			} else if a != tc.exp.resA || b != tc.exp.resB {
				t.Fatalf("Expected %v, %v, got: %v, %v", tc.exp.resA, tc.exp.resB, a,b)
			}
		})
	}
}