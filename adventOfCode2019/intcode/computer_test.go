package intcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{[]int{1002,5,3,5,99,33}, []int{1002,5,3,5,99,99}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			got := NewComputer(tc.in).Calc()
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestOpcodeExtraction(t *testing.T) {
	testCases := []struct {
		val int
		exp int		
	}{
		{1002, 2},
		{2, 2},
		{1112, 12},
		{1199, 99},
		{99, 99},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.val), func(t *testing.T) {
			got := opcode(tc.val).extractOpcode()
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestModeExtraction(t *testing.T) {
	op := opcode(12398)
	assert.Equal(t, 98, op.extractOpcode())
	assert.Equal(t, 3, op.modeForParam(0))
	assert.Equal(t, 2, op.modeForParam(1))
	assert.Equal(t, 1, op.modeForParam(2))
	
	assert.Equal(t, 0, op.modeForParam(3))
	assert.Equal(t, 0, op.modeForParam(4))
	assert.Equal(t, 0, op.modeForParam(5))
	assert.Equal(t, 0, op.modeForParam(6))
	assert.Equal(t, 0, op.modeForParam(7))
}

func TestHandleInput(t *testing.T) {
	in := []int{3,0,4,1,99}
	c := NewComputer(in)
	c.SetUserInput(15)
	c.Calc()
	
	require.NotNil(t, c.userInput)
	require.NotNil(t, c.userOutput)

	assert.Equal(t, 0, c.userOutput)
	assert.Equal(t, 15, c.userInput)
	assert.Equal(t, []int{15,0,4,1,99}, c.instructions)
}