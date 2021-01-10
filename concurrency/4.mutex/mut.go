package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// raceCondition()
	// atomicMethod()
	mutexMethod()
}

func atomicMethod()  {
	for x := 0; x < 100; x++ {
		var wg sync.WaitGroup
		var i int32 = 0

		inc := func(){
			atomic.AddInt32(&i, 1)
			wg.Done()
		}

		increments := 1000
		wg.Add(increments)
		for x := 0; x < increments; x++ {
			go inc()
		}
		wg.Wait()

		if i != int32(increments) {
			panic(fmt.Sprint("RACE in atomic", i, " != ", increments))
		}
	}
}

func mutexMethod()  {
	for x := 0; x < 100; x++ {
		var wg sync.WaitGroup
		var mutex sync.Mutex

		i := 0
		inc := func(){
			mutex.Lock()
			i = i + 1
			mutex.Unlock()
			wg.Done()
		}

		increments := 1000
		wg.Add(increments)
		for x := 0; x < increments; x++ {
			go inc()
		}
		wg.Wait()

		if i != increments {
			panic(fmt.Sprint("RACE in mutexed ", i, " != ", increments))
		}
	}
}


func raceCondition()  {
	for x := 0; x < 100; x++ {
		var wg sync.WaitGroup
		i := 0

		inc := func(){
			i++ // 3 operations, might go wrong! critical section!
			wg.Done()
		}

		increments := 1000
		wg.Add(increments)
		for x := 0; x < increments; x++ {
			go inc()
		}
		wg.Wait()

		if i != increments {
			panic(fmt.Sprint("RACE ", i, " != ", increments))
		}
	}
}