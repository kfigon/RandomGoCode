package main

import (
	"fmt"
	"sync"
)

func run(threads int) {
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
		
		}()	
	}

	wg.Wait()
	fmt.Println("all done")
}