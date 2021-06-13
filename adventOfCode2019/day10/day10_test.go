package day10

import (
	"fmt"
	"io"
	"log"
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
	assert.Equal(t, result{x:23, y:20, visibility:334}, got)
}

// https://adventofcode.com/2019/day/10#part2
func TestPart2File(t *testing.T) {
	theMap := spaceMap(readFile(t))
	station := point{x:23, y:20}
	coord := theMap.orderByVaporization(station)[200-1]
	assert.Equal(t, 1, coord.x)
	assert.Equal(t, 2, coord.y)
	assert.Less(t, 1620, coord.x*100+coord.y)
}

func TestPositionAnalysis(t *testing.T) {
	testCases := []struct {
		x,y int
		exp int
	}{
		{1,0,7},
		{4,0,7},
		
		{0,2,6},
		{1,2,7},
		{2,2,7},
		{3,2,7},
		{4,2,5},
		
		{4,3,7},
		{3,4,8},
		{4,4,7},
	}
	data:=`.#..#.
......
#####.
....#.
...##.`
	s := buildMap(data)
	asteroidSet := s.filterAsteroids()
	assert.Equal(t, 10, len(asteroidSet))
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.x, tc.y), func(t *testing.T) {
			got := s.analyzePosition(point{tc.x,tc.y}, asteroidSet)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestTrigonometry(t *testing.T) {
	testCases := []struct {
		start point
		end point
		expDegree float64
		expLen float64
	}{
		{point{0,0}, point{0,0}, 0,0},
		{point{0,0}, point{1,0}, 0,1},
		{point{0,0}, point{2,0}, 0,2},
		{point{0,0}, point{0,2}, 1.5707963267948966,2},
		{point{1,1}, point{0,3}, 2.0344439357957027,2.23606797749979},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := tc.start.trigonometryVersion(tc.end)
			assert.Equal(t, tc.expDegree, got.degree)
			assert.Equal(t, tc.expLen, got.length)
		})
	}
}

func TestOrderAsteroids(t *testing.T) {
	data := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	s := buildMap(data)
	ordered := s.orderByVaporization(point{11,13})

	for i := 0; i < len(ordered); i++ {
		v:=ordered[i]
		log.Println(i,"---->",v.x,",",v.y)
	}
	assert.Equal(t, point{11,12}, ordered[0])
	assert.Equal(t, point{12,1}, ordered[1])
	assert.Equal(t, point{12,2}, ordered[2])
	assert.Equal(t, point{12,8}, ordered[9])
	assert.Equal(t, point{16,0}, ordered[19])
	assert.Equal(t, point{16,9}, ordered[49])
	assert.Equal(t, point{10,16}, ordered[99])
	assert.Equal(t, point{9,6}, ordered[198])
	assert.Equal(t, point{8,2}, ordered[199])
	assert.Equal(t, point{10,9}, ordered[200])
	assert.Equal(t, point{11,1}, ordered[298])
}