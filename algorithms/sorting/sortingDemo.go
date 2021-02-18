package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// demoAlgorithm("bubble", 50000, bubbleSort)
	demoAlgorithm("merge", 1000000, mergeSort)
}

func demoAlgorithm(name string, tabLen int, alg func([]int)[]int) {
	randomTab := make([]int, tabLen)
	for i := 0; i < len(randomTab); i++ {
		randomTab[i] = rand.Intn(tabLen)
	}
	startTime := time.Now()
	// sorted := alg(randomTab)
	alg(randomTab)
	duration := time.Since(startTime)

	fmt.Printf("%v took %v, tabLen %v\n", name, duration, tabLen)
	// fmt.Println(randomTab)
	// fmt.Println(sorted)
}

