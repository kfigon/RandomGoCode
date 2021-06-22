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

func TestApplyGravity(t *testing.T) {
	ganymede := newMoon(position{3,0,0})
	callisto := newMoon(position{6,0,0})

	applyGravity(ganymede, callisto)

	assert.Equal(t, position{4,0,0}, ganymede.position)
	assert.Equal(t, position{5,0,0}, callisto.position)
}

func TestApplyVelocity(t *testing.T) {
	m := newMoon(position{1,2,3})
	m.velocity = position{3,-1,2}
	
	applyVelocity(m)

	assert.Equal(t, position{3,-1,2}, m.velocity)
	assert.Equal(t, position{4,1,5}, m.position)
}

func applyGravity(m1 *moon, m2 *moon) {
	applyPosition := func(coord1 *int, coord2 *int) {
		if *coord1 < *coord2 {
			*coord1++
			*coord2--
		} else if *coord1 > *coord2 {
			*coord1--
			*coord2++
		}
	}
	applyPosition(&m1.position.x, &m2.position.x)
	applyPosition(&m1.position.y, &m2.position.y)
	applyPosition(&m1.position.z, &m2.position.z)
}

func applyVelocity(m *moon) {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
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

type moon struct {
	position position
	velocity position
}

func newMoon(position position) *moon {
	return &moon{position: position}
}


