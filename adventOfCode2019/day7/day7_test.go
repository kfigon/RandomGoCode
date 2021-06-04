package day5

import (
	"aoc2019/intcode"
	"strconv"

	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/7
func TestPart1Examples(t *testing.T) {
	testCases := []struct {
		exp int
		code []int
	}{
		{43210, []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}},
		{54321, []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}},
		{65210, []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33, 1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := doPart1(tc.code)	
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestPart1(t *testing.T) {
	file := readFile(t)
	assert.Equal(t, 844468, doPart1(file))
}

func TestPart2(t *testing.T) {
	t.Fail() // todo
}

func readFile(t *testing.T) []int{
	file,err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)
	data := string(content)
	splitted := strings.Split(data, ",")
	out := make([]int,len(splitted))
	
	for i := 0; i < len(splitted); i++ {
		v,err := strconv.Atoi(splitted[i])	
		require.NoError(t,err)
		out[i] = v
	}
	return out
}

func doPart1(code []int) int {
	newCode := func(in []int) []int {
		out := make([]int,len(in))
		for i := 0; i < len(out); i++ {
			out[i]=in[i]
		}
		return out
	}

	var maxSignal int
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
	
						if a==b || a==c || a==d || a==e || 
							b==c || b==d || b==e ||
							c==d || c==e || 
							d == e {
							continue
						}

						comp1:=intcode.NewComputer(newCode(code))
						comp2:=intcode.NewComputer(newCode(code))
						comp3:=intcode.NewComputer(newCode(code))
						comp4:=intcode.NewComputer(newCode(code))
						comp5:=intcode.NewComputer(newCode(code))
				
						comp1.SetUserInput(a)
						comp1.SetUserInput(0)
						comp1.Calc()		

						comp2.SetUserInput(b)
						comp2.SetUserInput(comp1.GetOutput())
						comp2.Calc()

						comp3.SetUserInput(c)
						comp3.SetUserInput(comp2.GetOutput())
						comp3.Calc()

						comp4.SetUserInput(d)
						comp4.SetUserInput(comp3.GetOutput())
						comp4.Calc()

						comp5.SetUserInput(e)
						comp5.SetUserInput(comp4.GetOutput())
						comp5.Calc()

						outSignal := comp5.GetOutput()
						if outSignal > maxSignal {
							maxSignal = outSignal
						}
					}
				}
			}
		}
	}
	return maxSignal
}


// func readFile(t *testing.T) []int {
// 	file, err := os.Open("data.txt")
// 	require.NoError(t, err)
// 	defer file.Close()

// 	content, err := io.ReadAll(file)
// 	require.NoError(t, err)
// 	data := string(content)
// 	splitted := strings.Split(data, ",")
// 	out := make([]int, len(splitted))
// 	for i := 0; i < len(out); i++ {
// 		v, err := strconv.Atoi(splitted[i])
// 		require.NoError(t, err)
// 		out[i] = v
// 	}
// 	return out
// }
