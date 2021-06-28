package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)



func randStringRunes(n int) string {
    b := make([]rune, n)
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

type unitOfWork string
func (u unitOfWork) doWork() {
	wait := time.Duration(rand.Int31n(5))*time.Second
	fmt.Printf("thread %v start, ETA: %v\n", u, wait)
	time.Sleep(wait)
	fmt.Printf("thread %v done\n", u)
}
const numberOfThreads = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	// runWithWaitGroup()
	runWithChannel()

	fmt.Println("all done in", time.Since(start))
}

func runWithWaitGroup() {
	workers := createWorkers()
	
	var wg sync.WaitGroup
	wg.Add(len(workers))
	for i := 0; i < len(workers); i++ {
		go func(u unitOfWork) {
			u.doWork()
			wg.Done()
		}(workers[i])
	}
	wg.Wait()
}

func createWorkers() []unitOfWork {
	out := make([]unitOfWork, numberOfThreads)
	for i := 0; i < len(out); i++ {
		out[i] = unitOfWork(randStringRunes(5))
	}
	return out
}

func runWithChannel() {
	workers := createWorkers()
	done := make(chan struct{}, len(workers))
	
	for i := 0; i < len(workers); i++ {
		go func(u unitOfWork) {
			u.doWork()
			done<-struct{}{}
		}(workers[i])
	}

	for i := 0; i < len(workers); i++ {
		<-done
	}
}