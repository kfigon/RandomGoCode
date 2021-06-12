package day10

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
	asteroidSet := s.buildAsteroidSet()
	candidate := result{}
	for coord := range asteroidSet.data {
		v := s.analyzePosition(coord, asteroidSet)
		if v > candidate.visibility {
			candidate.visibility = v
			candidate.x = coord.x
			candidate.y = coord.y
		}
	}
	return candidate
}

func (s spaceMap) buildAsteroidSet() *set {
	asteroidSet := newSet()
	for x := 0; x < s.cols(); x++ {
		for y := 0; y < s.rows(); y++ {
			if !isAsteroid(s.charAt(x,y)) {
				continue
			}

			asteroidSet.add(point{x,y})
		}
	}
	return asteroidSet
}

func (s spaceMap) analyzePosition(a point, asteroidSet *set) int {
	if !asteroidSet.contains(a) {
		return 0
	}

	var out int
	for coord := range asteroidSet.data {
		f := buildFunction(point{a.x, a.y}, 
						   point{coord.x, coord.y})
	
		for tmp := range asteroidSet.data {
			if f.isPointOnTheLine(point{tmp.x, tmp.y}) {
				out++
			}	
		}
		
	}
	return out
}

type point struct { x,y int }
type void struct{}
type set struct{
	data map[point]void
}

func (p point) eq(other point) bool {
	return p.x == other.x && p.y == other.y
}

func newSet() *set {
	return &set{
		data: map[point]void{},
	}
}

func (s *set) add(a point) {
	var v void
	s.data[a] = v
}

func (s *set) contains(a point) bool {
	_, ok := s.data[a]
	return ok
}

func (s *set) len() int { return len(s.data) }

type fun struct { 
	p1,p2 point
}

func buildFunction(p1,p2 point) fun {
	return fun{
		p1: p1,
		p2: p2,
	}
}

func (f fun) isPointOnTheLine(p point) bool {
	isPointOnEndge := f.p1.eq(p) || f.p2.eq(p)
	isSingularity := f.p1.eq(f.p2)

	if isSingularity {
		return isPointOnEndge
	} else if isPointOnEndge {
		return true
	}

	// todo - float?
	a := (f.p1.y - f.p2.y)/(f.p1.x-f.p2.x)
	b := f.p1.y - f.p1.x*a
	return p.y == a*p.x + b
}
