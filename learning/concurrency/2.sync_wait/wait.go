package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Main start")
	
	// usually wrapping function with closure and calling done will be better
	foo := func (i int, wg *sync.WaitGroup)  {
		fmt.Println("\thello from",i)
		fmt.Println("\tgoodbye from",i)
		wg.Done() // defer wg.Done() can be used
		// wg.Done() // panic: sync: negative WaitGroup counter
	}

	var wg sync.WaitGroup

	numberOfThreads := 3
	for i := 0; i < numberOfThreads; i++ {
		num := i
		go foo(num, &wg)	
	}
	wg.Add(numberOfThreads) 
	// wg.Add(numberOfThreads+1)  fatal error: all goroutines are asleep - deadlock
	wg.Wait()

	fmt.Println("Main end")
}