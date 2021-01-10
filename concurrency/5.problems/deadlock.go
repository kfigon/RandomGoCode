package main

import (
	// "fmt"
	"sync"
)

func main() {
	deadlock()
}

func deadlock()  {
	var wg sync.WaitGroup
	
	doStuff := func(c1 chan int, c2 chan int)  {
		<- c1 // waits for receiving data on c1
		c2 <- 15 // sends data
		wg.Done()
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)
	go doStuff(ch1, ch2) //wait for ch1, send to ch2
	go doStuff(ch2, ch1) //wait for ch2, send to ch1
	// both waits: first for ch1, second for ch2, 
	// both are blocked, so no work is done
	wg.Wait()
}