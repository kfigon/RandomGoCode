package day11

import (
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/11
func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	splitted := strings.Split(string(content),",")
	out := make([]int, 0)
	for i := 0; i < len(splitted); i++ {
		v, err := strconv.Atoi(string(splitted[i]))
		require.NoError(t, err)
		
		out = append(out, v)
	}
	return out	
}

func TestAsd(t *testing.T) {
	t.Fail()
	// comp := intcode.NewComputer(readFile(t))
	// finished := comp.SingleInstruction()
}