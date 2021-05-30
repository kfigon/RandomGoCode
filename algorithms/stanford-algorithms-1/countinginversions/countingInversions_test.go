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
	return []testData{
		{[]int{}, 0},
		{[]int{1,2,3}, 0},
		{[]int{1,3,5,2,4,6}, 3}, // 32,52,54
		{[]int{3,2,1}, 3},
		{[]int{4,3,2,1}, 6},
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

func countInversions(in []int) int {
	return 0
}

func countInversionsBrute(in []int) int {
	invs := 0
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in); j++ {
			if i < j && in[i] > in[j] {
				invs += 1
			}
		}
	}
	return invs
}