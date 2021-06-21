package day12

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input string =`<x=-8, y=-18, z=6>
<x=-11, y=-14, z=4>
<x=8, y=-3, z=-10>
<x=-2, y=-16, z=1>`

// https://adventofcode.com/2019/day/12

func TestParsing(t *testing.T) {
	in := parseInput()
	assert.Equal(t, position{-8,-18,6}, in[0])
	assert.Equal(t, position{-11,-14,4}, in[1])
	assert.Equal(t, position{8,-3,-10}, in[2])
	assert.Equal(t, position{-2,-16,1}, in[3])
}

func TestPart1(t *testing.T) {
	t.Fail()
}

type position struct {
	x,y,z int
}

func parseInput() []position {
	out := []position{}
	reg := regexp.MustCompile(`<\w=(\-?\d+), \w=(\-?\d+), \w=(\-?\d+)>`)
	splitted := strings.Split(input, "\n")
	for _, v := range splitted {
		for _, match := range reg.FindAllStringSubmatch(v,-1) {
			x,_ := strconv.Atoi(match[1])
			y,_ := strconv.Atoi(match[2])
			z,_ := strconv.Atoi(match[3])
			out = append(out, position{x,y,z})
		}
	}
	return out
}