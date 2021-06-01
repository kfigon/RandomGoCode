package day1

import (
	"bufio"
	"os"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

// https://adventofcode.com/2019/day/1
func calcFuel(mass int) int {
	return mass/3 - 2
}

func calcSummedFuel(masses []int) int {
	out := 0
	for _,v := range masses	{
		out += calcFuel(v)
	}
	return out
}

func TestCalc(t *testing.T) {
	testCases := []struct {
		in int
		exp int		
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tC := range testCases {
		t.Run(strconv.Itoa(tC.in), func(t *testing.T) {
			assert.Equal(t, tC.exp, calcFuel(tC.in))
		})
	}
}

func parseFile() []int {
	out := make([]int,0)
	file, err := os.Open("data.txt")
	if err != nil {
		return out
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		out = append(out, v)
	}
	return out
}

func TestCalcFromFile(t *testing.T) {
	data := parseFile()
	sum := calcSummedFuel(data)
	assert.Equal(t, 3234871, sum)
}