package main

import (
	"time"
	"fmt"
)

func main() {
	// unbufferedChannel()
	// bufferedChannel()
	selectStatement()
}

func unbufferedChannel()  {
	c := make(chan int)
	multiply := func (a, b int, x chan int)  {
		x <- a*b
	}

	go multiply(1,2,c)
	go multiply(1,3,c)

	// order not quaranteed
	fmt.Println(<- c) // thread blocks here, waiting for data from goroutine
	fmt.Println(<- c)
	// fmt.Println(<- c) fatal error: all goroutines are asleep - deadlock!
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

// when want first available data from any channel
func selectStatement() {
	c1 := make(chan int)
	c2 := make(chan int)

	foo := func (a int, x chan int)  {
		x <-a*2
	}

	go foo(1, c1)
	go foo(2, c2)

	select {
	case a := <- c1:
		 fmt.Println("chan1", a)
	case b := <- c2:
		 fmt.Println("chan2", b)
	
	//  we can even block on sending in select
	// select will execute first available
	// case outChan <- 213: fmt.Println("sending data")
	}
	
	// sometimes we do additional channel to abort the proces
	// or a default to not block
	// for {
	// 	select {
	// 	case <- myAbortChannel:return
	// 	default: fmt.Println("nop")
	// 	}
	// }
	
}
