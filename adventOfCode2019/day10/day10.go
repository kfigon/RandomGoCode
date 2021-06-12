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

			asteroidSet.add(asteroidPosition{x,y})
		}
	}
	return asteroidSet
}

func (s spaceMap) analyzePosition(a asteroidPosition, asteroidSet *set) int {
	if !asteroidSet.contains(a) {
		return 0
	}
	var out int
	for coord := range asteroidSet.data {
		f := buildFunction(point{float64(a.x), float64(a.y)}, 
						   point{float64(coord.x), float64(coord.y)})
	
		for tmp := range asteroidSet.data {
			if f.isPointOnTheLine(point{float64(tmp.x), float64(tmp.y)}) {
				out++
			}	
		}
		
	}
	return out
}

type asteroidPosition struct { x,y int }
type void struct{}
type set struct{
	data map[asteroidPosition]void
}

func newSet() *set {
	return &set{
		data: map[asteroidPosition]void{},
	}
}

func (s *set) add(a asteroidPosition) {
	var v void
	s.data[a] = v
}

func (s *set) contains(a asteroidPosition) bool {
	_, ok := s.data[a]
	return ok
}

func (s *set) len() int { return len(s.data) }

type point struct { x,y float64 }
type fun struct { a,b float64 }

func buildFunction(p1,p2 point) fun {
	a := (p1.y - p2.y)/(p1.x-p2.x)
	b := p1.y - p1.x*a
	return fun{
		a: a,
		b: b,
	}
}

func (f fun) isPointOnTheLine(p point) bool {
	return p.y == f.a*p.x + f.b
}

