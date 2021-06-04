package intcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		in []int
		exp []int
	}{
		{[]int{1,0,0,3,99}, []int{1,0,0,2,99}},
		{[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
		{[]int{1,9,10,3,2,3,11,0,99,30,40,50}, []int{3500,9,10,70,2,3,11,0,99,30,40,50}},
		{[]int{2,3,0,3,99}, []int{2,3,0,6,99}},
		{[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}},
		{[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			got := NewComputer(tc.in).Calc()
			assert.Equal(t, tc.exp, got)
		})
	}
}