package main

import (
	"fmt"
	"time"
)


// zawsze main jest w gorutynie automatycznie
func main() {
	// synchronous()
	// async()
	// demo1()
	// demo2()
}

// usually only "main routine" is printed, 
// because main ends before 
// thread to pop and finish
func demo1()  {
	go fmt.Println("new routine")
	fmt.Println("main routine")
}

// hack - sleep main thread so new goroutine can spin. DO NOT RELY ON TIME!
// go runtime scheduler sees that there's useless code - sleep, so
// utilize that time
func demo2()  {
	go fmt.Println("new routine")
	time.Sleep(100 * time.Microsecond)
	fmt.Println("main routine")
}

// gdy uzyjemy go - nie blokujemy glownej funkcji, 
// one ida rownolegle. 
// Poczekamy na nie sleepem
func async()  {
	fmt.Println("======== start\t", time.Now())
	go run(1)
	go run(2)

	time.Sleep(1*1000*1000*1000)
	fmt.Println("========= End\t",  time.Now())
}

func run(i int) {
	fmt.Println("running thread", i)
	fmt.Println("done", i)
}

// synchronicznie, najpierw print, run1, run2, end
func synchronous()  {
	fmt.Println("======== start\t", time.Now())
	run(1)
	run(2)
	fmt.Println("========= End\t",  time.Now())
}

