package day10

import (
	"math"
	"sort"
)


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
	asteroids := s.filterAsteroids()
	candidate := result{}
	for _,coord := range asteroids {
		v := s.analyzePosition(coord, asteroids)
		if v > candidate.visibility {
			candidate.visibility = v
			candidate.x = coord.x
			candidate.y = coord.y
		}
	}
	return candidate
}

func (s spaceMap) filterAsteroids() []point {
	asteroids := []point{}
	for x := 0; x < s.cols(); x++ {
		for y := 0; y < s.rows(); y++ {
			if !isAsteroid(s.charAt(x,y)) {
				continue
			}

			asteroids = append(asteroids, point{x,y})
		}
	}
	return asteroids
}

// returns number of points visible form the starting point
func (s spaceMap) analyzePosition(startingPoint point, asteroids []point) int {
	visiblePoints := map[float64]float64{}

	for _,v := range asteroids {
		if startingPoint.eq(v) {
			continue
		}

		trig := startingPoint.trigonometryVersion(v)
		angle, ok := visiblePoints[trig.degree]
		if !ok {
			visiblePoints[trig.degree] = trig.length
		} else if trig.degree < angle{
			visiblePoints[trig.degree] = trig.length
		}
	}
	return len(visiblePoints)
}

// veeeeery bad, does not work yet
func (s spaceMap) orderByVaporization(startingPoint point) []point {
	asteroids := s.filterAsteroids()
	sortedByAngle := []orderingData{}
	for i := 0; i < len(asteroids); i++ {
		if startingPoint.eq(asteroids[i]) {
			continue
		}
		sortedByAngle = append(sortedByAngle, newOrderingObj(startingPoint, asteroids[i]))
	}
	sort.Slice(sortedByAngle, func(i, j int) bool {
		return sortedByAngle[i].angle < sortedByAngle[j].angle
	})
	
	groupedByAngle := map[float64][]orderingData{}
	for i := 0; i < len(sortedByAngle); i++ {
		v := sortedByAngle[i]
		if _, ok := groupedByAngle[v.angle]; !ok {
			groupedByAngle[v.angle] = []orderingData{}	
		}
		groupedByAngle[v.angle] = append(groupedByAngle[v.angle], v)
	}
	for key := range groupedByAngle {
		sort.Slice(groupedByAngle[key], func(i, j int) bool {
			return groupedByAngle[key][i].length < groupedByAngle[key][j].length
		})
	}

	out := []point{}
	for len(groupedByAngle) > 0 {
		for i := 0; i < len(sortedByAngle); i++ {
			pt := sortedByAngle[i]
			v := groupedByAngle[pt.angle]
			if len(v) == 0 {
				delete(groupedByAngle, pt.angle)
			} else if len(v) == 1 {
				tmp := groupedByAngle[pt.angle][0]
				out = append(out, tmp.pt)
				delete(groupedByAngle, pt.angle)
			} else {
				tmp := groupedByAngle[pt.angle][0]
				out = append(out, tmp.pt)
				groupedByAngle[pt.angle] = groupedByAngle[pt.angle][1:]
			}
		}
	}

	return out
}

type orderingData struct {
	pt	point
	angle, length float64
}

func newOrderingObj(starting point, pt point) orderingData {
	trig := starting.trigonometryVersion(pt)
	return orderingData{
		pt: pt,
		angle: trig.degree,
		length: trig.length,
	}
}


type point struct { x,y int }
func (p point) eq(other point) bool {
	return p.x == other.x && p.y == other.y
}

type trigonometryPoint struct {
	degree, length float64
}
func (p point) trigonometryVersion(end point) trigonometryPoint {
	deltaX := math.Pow((float64(end.x) - float64(p.x)), 2); 
	deltaY := math.Pow((float64(end.y) - float64(p.y)), 2);

	distance := math.Sqrt(deltaY + deltaX);

	const RADIAN_TO_DEGREE = 180 / math.Pi
	angle := 180 - (RADIAN_TO_DEGREE * math.Atan2(float64(end.y) - float64(p.y), float64(end.x) - float64(p.x)))
	return trigonometryPoint{
		degree: angle,
		length: distance,
	}
}