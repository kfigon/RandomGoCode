package averagePair

import (
	"fmt"
	"testing"
)

// # given sorted array of ints and target average
// # determine if there is a pair of values that average equals target average

func averagePair(in []int, avg float32) bool {
	left := 0
	right := len(in) - 1
	for left <= right {
		currentAvg := float32((in[left]) + in[right])/2
		if currentAvg == avg {
			return true
		} else if currentAvg > avg {
			right--
		} else {
			left++
		}
	}
	return false
}

func Test(t *testing.T) {
	testCases := []struct {
		in []int
		avg float32
		exp bool
	}{
		{ in: []int{1,2,3}, avg: 2.5 , exp: true },
		{ in: []int{1,2,3,4,5,6}, avg: 3.5 , exp: true },
		{ in: []int{1,3,3,5,6,7,10,12,19}, avg: 8 , exp: true },
		{ in: []int{-1,0,3,4,5,6}, avg: 4.1, exp: false },
		{ in: []int{1,1,1,1,1,1,1,1,1,1,1,19}, avg:2, exp: false },
		{ in: []int{1,1,1,1,1,1,1,1,1,1,1,19}, avg:1, exp: true },
		{ in: []int{1,1,1,1,1,1,1,1,1,1,1,19}, avg:10, exp: true },
		{ in: []int{}, avg: 4, exp: false },
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			if res := averagePair(tc.in, tc.avg); res != tc.exp {
				t.Errorf("Expected %v, got %v", tc.exp, res)
			}
		})
	}
}