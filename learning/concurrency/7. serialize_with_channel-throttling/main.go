package main

import "fmt"

func main()  {
	// unbuffered channels can't do this!
	sem := make(chan bool,1)
	data := []int{1,2,3}

	done := make(chan struct{}, len(data))

	for i := 0; i < len(data); i++ {
		// concurrent, but semaphore limits calls to process - throttling
		go func(x int) {
			sem<-true
			process(x)
			<-sem
			
			done<-struct{}{}
		}(data[i])
	}

	for i := 0; i < len(data); i++ {
		<-done
	}

}

func process(i int) {
	fmt.Println("Processing",i)
	fmt.Println("Processing",i,"done")
}