package main

import (
	"time"
	"fmt"
)

func multiply(a, b int, c chan int)  {
	c <- a*b
}

func unbufferedChannel()  {
	c := make(chan int)
	go multiply(1,2,c)
	go multiply(1,3,c)

	// order not quaranteed
	fmt.Println(<- c) // thread blocks here, waiting for data from goroutine
	fmt.Println(<- c)
	// fmt.Println(<- c) fatal error: all goroutines are asleep - deadlock!
}

func main() {
	// unbufferedChannel()
	bufferedChannel()
}

func bufferedChannel() {
	c := make(chan int, 1)

	foo := func(x chan int) {
		a := <- x // this will wait until main will send data
		b := 4
		fmt.Println("thread got",a,b)
		x <- a*b
	}

	go foo(c)
	time.Sleep(5000000)
	c <- 2 // send the data
	// c <- 0 // now we are blocking!


	// every time when data is available
	// it'll continue to read
	// unless some producer calls close(c)
	// normally not need to close (like a file)
	// but when range is used - we have to

	// for i := range c {
	// 	fmt.Println(i)
	// }

	fmt.Println(<- c)
}

