package day10

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

func buildMap(data string) spaceMap {
	return spaceMap(split(data,"\n"))
}

func TestUtils(t *testing.T) {
	data := `.#..#.
......
#####.
....#.
...##.`
	s := buildMap(data)
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
	s := buildMap(data)
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v",tc.x,tc.y), func(t *testing.T) {
			assert.Equal(t, tc.exp, s.charAt(tc.x, tc.y))
		})
	}
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		data string
		expected result
	}{
		{".#..#\n.....\n#####\n....#\n...##", result{3,4,8}},
		{"......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####", result{5,8,33}},
		{"#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###.", result{1,2,35}},
		{".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#..", result{6,3,41}},
		{".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##", result{11,13,210}},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := buildMap(tc.data)
			got := s.findBestPlace()
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestPart1File(t *testing.T) {
	s := spaceMap(readFile(t))
	got := s.findBestPlace()
	assert.Equal(t, result{1,1,-1}, got)
}

func TestPointOnLine(t *testing.T) {
	p1 := point{0,0}
	p2 := point{3,1}
	p3 := point{6,2}

	f := buildFunction(p1,p2)

	got := f.isPointOnTheLine(p3)
	assert.True(t, got)
}

func TestAsteroidSetContains(t *testing.T) {
	testCases := []struct {
		x,y int
		exp bool		
	}{
		{0,0,false},
		{1,0,true},
		{2,0,false},
		{3,0,false},
		{4,0,true},
		{5,0,false},
		
		{0,1,false},
		{1,1,false},
		{2,1,false},
		{3,1,false},
		{4,1,false},
		{5,1,false},

		{0,2,true},
		{1,2,true},
		{2,2,true},
		{3,2,true},
		{4,2,true},
		{5,2,false},

		{0,3,false},
		{1,3,false},
		{2,3,false},
		{3,3,false},
		{4,3,true},
		{5,3,false},
	}
	data:=`.#..#.
......
#####.
....#.
...##.`
	s := buildMap(data)
	asteroidSet := s.buildAsteroidSet()
	assert.Equal(t, 10, asteroidSet.len())
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.x, tc.y), func(t *testing.T) {
			assert.Equal(t, tc.exp, asteroidSet.contains(asteroidPosition{tc.x,tc.y}))
		})
	}
}
