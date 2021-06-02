package day2

import (
	// "bufio"
	// "os"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		in []int
		exp []int
	}{
		{[]int{1,0,0,3,99}, []int{1,0,0,3,99}},
		{[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
		{[]int{1,9,10,3,2,3,11,0,99,30,40,50}, []int{3500,9,10,70,2,3,11,0,99,30,40,50}},
		{[]int{2,3,0,3,99}, []int{2,3,0,6,99}},
		{[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}},
		{[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			got := computer(tc.in).calc()
			assert.Equal(t, tc.exp, got)
		})
	}
}

type computer []int

func (c computer) calc() []int {
	var i int 
	for i < len(c) {
		i = c.handleCommand(i)
	}
	return c
}

const (
	ADD = 1
	MULT = 2
	TERMINATE = 99
)
func (c computer) handleCommand(idx int) int {
	val := c[idx]
	switch val {
	case ADD:	return c.handleAdd(idx)
	case MULT:	return c.handleMult(idx)
	case TERMINATE: return idx+len(c)
	}
	// should not happen
	return idx
}

func (c computer) handleAdd(idx int) int {
	return idx+4
}

func (c computer) handleMult(idx int) int {
	return idx+4
}