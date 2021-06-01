package day1

import (
	"bufio"
	"os"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

func calcFuel(mass int) int {
	return mass/3 - 2
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

func TestCalcFromFile(t *testing.T) {
	file, err := os.Open("data.txt")
	assert.NoError(t,err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		assert.NoError(t,err)
		sum += calcFuel(v)
	}
	assert.Equal(t, 123, sum)
}