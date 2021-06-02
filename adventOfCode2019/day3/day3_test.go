package day2

import (
	"strconv"
	"strings"
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
func TestTodo(t *testing.T) {
	assert.Fail(t, "todo")
}

type point struct {
	x int
	y int
}
func (p point) String() string {
	return strconv.Itoa(p.x) +","+strconv.Itoa(p.y)
}

func newPoint(x,y int) point { return point{x,y} }

func parsePointsToVectors(in string) []point {
	splitted := strings.Split(in, ",")
	out := make([]point, len(splitted)+1)
	out[0] = newPoint(0,0)
	for i := 0; i < len(splitted); i++ {
		toParse := splitted[i]
		dir := rune(toParse[0])
		val,_ := strconv.Atoi(toParse[1:])
		
		x,y := out[i].x,out[i].y
		switch dir {
		case 'U': y += val
		case 'D': y -= val
		case 'L': x-=val
		case 'R': x+=val
		}
		out[i+1] = newPoint(x,y)
	}
	return out
}

type void struct{}
type set struct {
	vals map[string]void
}
func newSet() *set {
	return &set{
		vals: make(map[string]void),
	}
}
func (s *set) add(p point) {
	var v void
	s.vals[p.String()] = v
}

func fromKeyToPoint(key string) point {
	vals := strings.Split(key,",")
	x,_:=strconv.Atoi(vals[0])
	y,_:=strconv.Atoi(vals[1])
	return newPoint(x,y)
}

func populateSet(points []point) *set {
	s := newSet()
	for i := 0; i < len(points)-1; i++ {
		current := points[i]
		next := points[i+1]
		s.add(current)
		s.add(next)
		intermediate := getIntermediate(current, next)
		for i := 0; i < len(intermediate); i++ {
			s.add(intermediate[i])
		}
	}
	return s
}

func parseAndCalc(in string) int {
	vals := strings.Split(in,"\r\n")
	if vals == nil || len(vals) != 2 {
		return -1
	}
	return calcStuff(parsePointsToVectors(vals[0]), parsePointsToVectors(vals[1]))
}

func calcStuff(p1 []point, p2 []point) int {
	s1 := populateSet(p1)
	s2 := populateSet(p2)
	
	var minDistance *int
	for key := range s1.vals {
		if _, ok := s2.vals[key]; ok {
			dist := calcDistance(fromKeyToPoint(key))
			if minDistance == nil || dist < *minDistance {
				*minDistance = dist
			}
		}
	}

	if minDistance == nil {
		return -1
	}
	return *minDistance
}

func getIntermediate(current, next point) []point {
	// todo
	return nil
}

func calcDistance(p point) int {
	// todo
	return -1
}