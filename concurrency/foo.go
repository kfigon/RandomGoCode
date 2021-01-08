package main

import (
	"fmt"
	"time"
)

func run(i int) {
	fmt.Println("running thread", i)
	time.Sleep(1000)
	fmt.Println("done", i)
}

func main() {
	fmt.Println("starting")
	for i := 0; i < 10; i++ {
		data := i // golang reuses iterators, copy
		go run(data)
	}
	time.Sleep(20000)
	fmt.Println("End")
}

