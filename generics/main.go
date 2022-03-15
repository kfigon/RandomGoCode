package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(sumInt([]int{1,2,3,4}))
	fmt.Println(sum([]int{1,2,3,4}))
	fmt.Println(sum([]float32{1,2,3,4}))
}

func sumInt(vals []int) int {
	var out int
	for _, v := range vals {
		out += v
	}
	return out
}

// or use number type:
type Number interface {
    int | float64
}
func sum[T int | float32](vals []T) T {
	var out T
	for _, v := range vals {
		out += v
	}
	return out
}
