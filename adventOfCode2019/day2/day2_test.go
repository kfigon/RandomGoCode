package day2

import (
	"strings"
	"io"
	"strconv"
	"os"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"aoc2019/intcode"
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
			got := intcode.Computer(tc.in).Calc()
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
func TestPart1(t *testing.T) {
	data := readFile()
	data[1] = 12
	data[2] = 2
	out := intcode.Computer(data).Calc()

	assert.Equal(t, 5866714, out[0])
}

func TestPart2(t *testing.T) {
	exp := 5208
	assert.Equal(t, exp, findPart2Result())
}

func findPart2Result() int {
	const valueToFind int = 19690720
	fileData := readFile()
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			
			data := make([]int, len(fileData))
			copy(data, fileData)

			data[1] = noun
			data[2] = verb
			out := intcode.Computer(data).Calc()
			if out[0] == valueToFind {
				return 100*out[1]+out[2]
			}
		}
	}
	return -1
}
