package day9

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/9

func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	strContent := string(content)
	out := make([]int, 0)
	for i := 0; i < len(strContent); i++ {
		v, err := strconv.Atoi(string(strContent[i]))
		require.NoError(t, err)
		
		out = append(out, v)
	}
	return out
	
}
func Test(t *testing.T) {
	t.Fail()
}