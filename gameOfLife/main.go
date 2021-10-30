package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"flag"
)

func main() {

	width := flag.Int("width", 100, "width of plain")
	height := flag.Int("height", 50, "height of plain")
	iter := flag.Int("iter", 1000, "number of iterations")
	delay := flag.Int("delay", 1, "delay in seconds")
	flag.Parse()

	p := createPlain(*width,*height)
	p.applyDiff(p.randomize())

	for i := 0; i < *iter; i++ {
		printPlane(p)
		fmt.Printf("%v / %v\n",i,*iter)
		time.Sleep(time.Duration(*delay)*time.Second)
		p.applyDiff(p.lifecycle())
	}
}

func printPlane(p *plain) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(p)
}

type point struct {
	x,y int
}

type cell struct {
	point
	alive bool
}

type plain struct {
	area [][]cell
}

func createPlain(width, height int) *plain {
	ar := [][]cell{}
	for h := 0; h < height; h++ {
		row := []cell{}
		for w := 0; w < width; w++ {
			row = append(row, cell{point{w,h},false})
		}
		ar = append(ar, row)
	}
	return &plain{ar}
}

func (p *plain) width() int {
	return len(p.area[0])
}

func (p *plain) height() int {
	return len(p.area)
}

func (p *plain) isAlive(pt point) bool {
	return p.area[pt.y][pt.x].alive
}

func (p *plain) kill(pt point) {
	p.area[pt.y][pt.x].alive = false
}

func (p *plain) revive(pt point) {
	p.area[pt.y][pt.x].alive = true
}

func (p *plain) String() string {
	out := ""
	for h := 0; h < p.height(); h++ {
		for w := 0; w < p.width(); w++ {
			r := ' '
			if p.isAlive(point{w,h}) {
				r = '#'
			}
			out += string(r)
		}
		out += "\n"
	}
	return out
}

type diff struct {
	point
	shouldBeAlive bool
}

func (p *plain) randomize() []diff {
out := []diff{}
	rand.Seed(time.Now().UnixNano())
	for h := 0; h < p.height(); h++ {
		for w := 0; w < p.width(); w++ {
			pt := point{w,h}
			r := rand.Intn(2)
			if r == 0 {
				out = append(out, diff{pt, false})
			} else {
				out = append(out, diff{pt, true})
			}
		}
	}
	return out
}

func (p *plain) lifecycle() []diff {
// Martwa komórka, która ma dokładnie 3 żywych sąsiadów, staje się żywa w następnej jednostce czasu (rodzi się)
// Żywa komórka z 2 albo 3 żywymi sąsiadami pozostaje nadal żywa; przy innej liczbie sąsiadów umiera (z „samotności” albo „zatłoczenia”)
	out := []diff{}
	for h := 0; h < p.height(); h++ {
		for w := 0; w < p.width(); w++ {
			pt := point{w,h}
			nei := p.neighbours(pt)
			aliveNei := countAlive(nei)
			if !p.isAlive(pt) && aliveNei == 3 {
				out = append(out, diff{point: pt, shouldBeAlive: true})
			} else if p.isAlive(pt) && (aliveNei != 2 && aliveNei != 3) {
				out = append(out, diff{point: pt, shouldBeAlive: false})
			}
		}
	}
	return out
}

func (p *plain) applyDiff(ds []diff) {
	for _, d := range ds {
		if d.shouldBeAlive {
			p.revive(d.point)
		} else {
			p.kill(d.point)
		}
	}
}

func (p *plain) neighbours(pt point) []cell {
	candidates := []point{
		{pt.x-1,pt.y-1},{pt.x,pt.y-1},{pt.x+1,pt.y-1},
		{pt.x-1,pt.y},  			 {pt.x+1,pt.y},
		{pt.x-1,pt.y+1},{pt.x,pt.y+1},{pt.x+1,pt.y+1},
	}
	out := []cell{}
	for _, c := range candidates {
		if c.x >= 0 && c.x < p.width() && c.y >= 0 && c.y < p.height() {
			out = append(out, cell{c, p.isAlive(c)})
		}
	}
	return out
}

func countAlive(cells []cell) int {
	out := 0
	for _, c := range cells {
		if c.alive {
			out++
		}
	}
	return out
}