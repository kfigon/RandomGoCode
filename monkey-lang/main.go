package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	fmt.Println("hello")

	it := slices.Values([]int{1,2,3,4,5})
	next, stop := iter.Pull(it)
	defer stop()
	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
	fmt.Println("all done")
}