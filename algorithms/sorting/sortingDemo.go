package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// demoAlgorithm("bubble", 50000, bubbleSort)
	demoAlgorithm("merge", 10000000, mergeSort)
	demoAlgorithm("quick", 10000000, quickSort)
	demoAlgorithm("radix", 10000000, radixSort)
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

