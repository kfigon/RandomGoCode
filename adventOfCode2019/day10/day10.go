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

func (s spaceMap) findVaporizedCoordinate(startingPoint point, nthAsteroid int) point {
	asteroids := s.filterAsteroids()
	pointsToOrder := orderByAngle{asteroids, startingPoint}
	sort.Sort(pointsToOrder)
	return pointsToOrder.points[nthAsteroid]
}


type orderByAngle struct {
	points	[]point
	startingPoint point
}
func (a orderByAngle) Len() int           { return len(a.points) }
func (a orderByAngle) Swap(i, j int)      { 
	a.points[i], a.points[j] = a.points[j], a.points[i] 
}

func (a orderByAngle) Less(i, j int) bool { 
	trigI := a.startingPoint.trigonometryVersion(a.points[i])
	trigJ := a.startingPoint.trigonometryVersion(a.points[j])
	
	if trigI.degree < trigJ.degree {
		return true
	}
	return trigI.length < trigJ.length
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

	radians := math.Atan2((float64(end.y) - float64(p.y)), (float64(end.x) - float64(p.x)));
	// angle := radians * (180 / math.Pi);
	return trigonometryPoint{
		degree: radians,
		length: distance,
	}
}