package day3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// https://adventofcode.com/2019/day/3
func TestPointParser(t *testing.T) {
	testCases := []struct {
		in string
		exp []point
	}{
		{"U7,R6,D4,L4", []point{newPoint(0,0), newPoint(0,7),newPoint(6,7),newPoint(6,3),newPoint(2,3)}},
		{"R8,U5,L5,D3", []point{newPoint(0,0), newPoint(8,0),newPoint(8,5),newPoint(3,5),newPoint(3,2)}},
	}
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got := parsePointsToVectors(tc.in)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestDistance(t *testing.T) {
	testCases := []struct {
		p point
		exp int
	}{
		{newPoint(0,0), 0},
		{newPoint(0,1), 1},
		{newPoint(1,1), 2},
		{newPoint(3,3), 6},
	}
	for _, tc := range testCases {
		t.Run(tc.p.String(), func(t *testing.T) {
			assert.Equal(t, tc.exp, tc.p.calcDistance())
		})
	}
}

func TestTodo(t *testing.T) {
	assert.Fail(t, "todo")
}
