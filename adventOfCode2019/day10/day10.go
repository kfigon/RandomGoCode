package day10

import "math"

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
	// todo
	// iter through all points
	// build a map[degree]=length
	// if theres already such degree - check length and choose min()
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