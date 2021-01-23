package main

import (
	"runtime/pprof"
	"fmt"
	"os"
)

// run this program
// go tool pprof -http=":8000" pprofbin ./cpu.pprof


type betterFib struct {
	d map[uint64]uint64
}

func (b *betterFib) calc(n uint64) uint64 {
	var fibBetter func(uint64) uint64

	fibBetter = func(n uint64) uint64 {
		if v,ok := b.d[n]; ok {
			return v
		} else if n <= 1 {
			return n
		}
		return fibBetter(n-1) + fibBetter(n-2)
	}

	res := fibBetter(n)
	b.d[n] = res
	return res
}

type poor struct{}
func (p *poor) calc(n uint64) uint64 {
	var fib func(uint64)uint64
	fib = func(n uint64) uint64 {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	return fib(n)
}

type ifibo interface {
	calc(uint64) uint64
}

func newPoor() ifibo {
	return &poor{}
}
func newFast() ifibo {
	return &betterFib{
		d: make(map[uint64]uint64),
	}
}

func main() {
	f, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// alg := newPoor()
	alg := newFast()
	for i := 0; i < 45; i++ {
		res := alg.calc(uint64(i))
		fmt.Printf("%v -> %v\n", i, res)
	}
}