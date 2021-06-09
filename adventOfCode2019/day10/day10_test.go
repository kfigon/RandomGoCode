package day10

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/10

func readFile(t *testing.T) []string {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	return split(string(content), "\r\n")
}

func split(data, separator string) []string {
	splitted := strings.Split(data, separator)
	out := []string{}
	for i := 0; i < len(splitted); i++ {
		out = append(out, splitted[i])
	}
	return out
}

func TestUtils(t *testing.T) {
	data := `.#..#.
......
#####.
....#.
...##.`
	s := spaceMap(split(data,"\n"))
	assert.Equal(t, 5, s.rows())
	assert.Equal(t, 6, s.cols())
}

func TestCharAt(t *testing.T) {
	testCases := []struct {
		x int
		y int
		exp rune
	}{
		{0,0,'.'},
		{1,0,'#'},
		{2,0,'.'},
		{3,0,'.'},
		{4,0,'#'},
		{0,1,'.'},
		{0,2,'#'},
		{4,3,'#'},
	}
	data := `.#..#
.....
#####
....#`
	s := spaceMap(split(data,"\n"))
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v",tc.x,tc.y), func(t *testing.T) {
			assert.Equal(t, tc.exp, s.charAt(tc.x, tc.y))
		})
	}
}

func TestPart1(t *testing.T) {
	data := `.#..#
.....
#####
....#
...##`
	s := spaceMap(split(data, "\n"))
	res := s.findBestPlace()
	assert.Equal(t, 3, res.x)
	assert.Equal(t, 4, res.y)
	assert.Equal(t, 8, res.visibility)
}

func isAsteroid(c rune) bool { return c =='#' }
type spaceMap []string

func (s spaceMap) charAt(x,y int) rune {
	return rune(s[y][x])
}

func (s spaceMap) rows() int { return len(s) }
func (s spaceMap) cols() int { return len(s[0]) }

type result struct {
	x,y,visibility int
}
func (s spaceMap) findBestPlace() result { 
	return result{}
}


