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
	return result{}
}


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

