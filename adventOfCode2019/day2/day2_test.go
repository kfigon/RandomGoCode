package day2

import (
	"strings"
	"io"
	"strconv"
	"os"
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
			got := computer(tc.in).calc()
			assert.Equal(t, tc.exp, got)
		})
	}
}
func readFile() []int {
	file, err := os.Open("data.txt")
	if err != nil {
		return []int{}
	}
	defer file.Close()

	out := make([]int,0)
	bytes, err := io.ReadAll(file)
	if err != nil {
		return []int{}
	}
	str := string(bytes)
	for _, v := range strings.Split(str, ",") {
		i, _ := strconv.Atoi(v)
		out = append(out, i)
	}
	return out
}
func TestTask1(t *testing.T) {
	data := readFile()
	data[1] = 12
	data[2] = 2
	out := computer(data).calc()

	assert.Equal(t, 5866714, out[0])
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

func (c computer) getIdxs(idx int) (int,int,int) {
	return c[idx+1],c[idx+2],c[idx+3]
}

func (c computer) handleAdd(idx int) int {
	idxA,idxB,idxC := c.getIdxs(idx)
	c[idxC] = c[idxA] + c[idxB]
	return idx+4
}

func (c computer) handleMult(idx int) int {
	idxA,idxB,idxC := c.getIdxs(idx)
	c[idxC] = c[idxA] * c[idxB]
	return idx+4
}