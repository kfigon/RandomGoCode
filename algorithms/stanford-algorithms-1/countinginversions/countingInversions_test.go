package countinginversions

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	in []int
	exp int
}

func data() []testData{
	// max number - (n(n-1))/2
	return []testData{
		{[]int{}, 0},
		{[]int{1,2,3}, 0},
		{[]int{1,3,5,2,4,6}, 3}, // 32,52,54
		{[]int{3,2,1}, 3},
		{[]int{4,3,2,1}, 6},
		{[]int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 190},
		{[]int{50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 
		1225},
	}
}

func TestCountInversions(t *testing.T) {
	testCases := data()
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v",i), func(t *testing.T) {
			got := countInversions(tc.in)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestCountInversionsBrute(t *testing.T) {
	testCases := data()
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v",i), func(t *testing.T) {
			got := countInversionsBrute(tc.in)
			assert.Equal(t, tc.exp, got)
		})
	}
}
// O(nlogn)
func countInversions(in []int) int {
	if len(in) <= 1 {
		return 0
	}
	return 0
	// left := countInversions(in[:len(in)/2])
	// right := countInversions(in[len(in)/2:])
	// split := countSplitInversions(in)
	// return left+right
}

// O(n^2)
func countInversionsBrute(in []int) int {
	invs := 0
	for i := 0; i < len(in); i++ {
		for j := i; j < len(in); j++ {
			if i < j && in[i] > in[j] {
				invs += 1
			}
		}
	}
	return invs
}