package countinginversions

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountInversions(t *testing.T) {
	testCases := []struct {
		in []int
		exp []int
	}{
		{[]int{}, []int{}},
		{[]int{1,2,3}, []int{1,2,3}},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v",i), func(t *testing.T) {
			got := countInversions(tc.in)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func countInversions(in []int) []int {
	return in
}