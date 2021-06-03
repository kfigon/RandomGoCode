package day3

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestSegmentIntersects(t *testing.T) {
	testCases := []struct {
		s1 segment
		s2 segment
		exp point		
	}{
		{
			newSegment(newPoint(0,3), newPoint(0,4)),
			newSegment(newPoint(0,3), newPoint(0,4)),
			newPoint(0,3),
		},
		{
			newSegment(newPoint(1,1), newPoint(1,5)),
			newSegment(newPoint(1,1), newPoint(3,1)),
			newPoint(1,1),
		},
		{
			newSegment(newPoint(1,5), newPoint(1,1)),
			newSegment(newPoint(1,1), newPoint(3,1)),
			newPoint(1,1),
		},
		{
			newSegment(newPoint(1,1), newPoint(1,5)),
			newSegment(newPoint(3,1), newPoint(1,1)),
			newPoint(1,1),
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got1 := tc.s1.intersection(tc.s2)
			got2 := tc.s2.intersection(tc.s1)
			require.NotNil(t, got1)
			require.NotNil(t, got2)
			assert.Equal(t, tc.exp, *got1)
			assert.Equal(t, tc.exp, *got2)
		})
	}
}

func TestSegmentNotIntersects(t *testing.T) {
	testCases := []struct {
		s1 segment
		s2 segment
	}{
		{
			newSegment(newPoint(1,3), newPoint(1,4)),
			newSegment(newPoint(0,3), newPoint(0,4)),
		},
		{
			newSegment(newPoint(1,4), newPoint(1,3)),
			newSegment(newPoint(0,3), newPoint(0,4)),
		},
		{
			newSegment(newPoint(1,4), newPoint(1,3)),
			newSegment(newPoint(0,4), newPoint(0,3)),
		},
		{
			newSegment(newPoint(1,3), newPoint(5,3)),
			newSegment(newPoint(1,0), newPoint(5,0)),
		},
		{
			newSegment(newPoint(5,3), newPoint(1,3)),
			newSegment(newPoint(1,0), newPoint(5,0)),
		},
		{
			newSegment(newPoint(5,3), newPoint(1,3)),
			newSegment(newPoint(5,0), newPoint(1,0)),
		},
		{
			newSegment(newPoint(1,1), newPoint(1,5)),
			newSegment(newPoint(2,3), newPoint(6,3)),
		},
		{
			newSegment(newPoint(1,1), newPoint(1,5)),
			newSegment(newPoint(6,3), newPoint(2,3)),
		},
		{
			newSegment(newPoint(1,5), newPoint(1,1)),
			newSegment(newPoint(2,3), newPoint(6,3)),
		},
		{
			newSegment(newPoint(1,5), newPoint(1,1)),
			newSegment(newPoint(6,3), newPoint(2,3)),
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Nil(t, tc.s1.intersection(tc.s2))
			assert.Nil(t, tc.s2.intersection(tc.s1))
		})
	}
}
func TestTodo(t *testing.T) {
	assert.Fail(t, "todo")
}
