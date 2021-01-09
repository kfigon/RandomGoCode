package main

import (
	"fmt"
	"time"
)

func multiply(a, b int, c chan int)  {
	dur := time.Duration(b) * time.Second
	time.Sleep(dur)
	c <- a*b
}

func main() {
	c := make(chan int)
	go multiply(1,2,c)
	go multiply(1,3,c)

	// order not quaranteed
	fmt.Println(<- c)
	fmt.Println(<- c)
	// fmt.Println(<- c) fatal error: all goroutines are asleep - deadlock!
}

