package main

import (
	"fmt"
	"time"
)

func run(i int) {
	fmt.Println("running thread", i)
	fmt.Println("done", i)
}

func main() {
	fmt.Println("======== start\t", time.Now())

	for i := 0; i < 5; i++ {
		d := i
		go run(d)
	}

	time.Sleep(5*1000*1000*1000)

	fmt.Println("========= End\t",  time.Now())
}

