package day3

import (
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
func (p point) String() string {
	return "("+strconv.Itoa(p.x)+","+strconv.Itoa(p.y)+")"
}

func (p point) calcDistance() int {
	return p.x+p.y
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

type segment struct {
	start point
	end point
}

func newSegment(start, end point) segment { return segment{start,end} }

func (s segment) intersection(other segment) *point {
	return &s.start
}