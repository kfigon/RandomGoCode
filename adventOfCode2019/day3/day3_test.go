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

func TestPart1(t *testing.T) {
	testCases := []struct {
		wire1	string
		wire2	string
		exp int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4",6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83",159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",135},
		{"", "",135},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := findMinDistance(tc.wire1, tc.wire2)			
			assert.Equal(t, tc.exp, got)
		})
	}
}
