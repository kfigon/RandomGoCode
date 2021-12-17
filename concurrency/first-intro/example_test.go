package firstintro

import (
	"sync"
	"testing"
	"time"
)

// race solved
// race condition is not - we still don't know the result
func TestCriticalSection(t *testing.T) {
	var mut sync.Mutex
	value := 0

	go func() {
		mut.Lock()
		value++
		mut.Unlock()
	}()

	mut.Lock()
	if value == 0 {
		// some code
	} else {
		// some other code
	}
	mut.Unlock()
}

// 2 threads are waiting for eachother
// and can't move
func TestDeadlock(t *testing.T) {
	t.Skip()
	type val struct {
		mut sync.Mutex
		v int
	}

	var wg sync.WaitGroup
	sum := func(a,b *val) int {
		defer wg.Done() // last defer called
		a.mut.Lock()
		defer a.mut.Unlock() // second defer called

		time.Sleep(1*time.Second) // work...

		b.mut.Lock()
		defer b.mut.Unlock() // first defer called
		return a.v + b.v
	}

	var a,b val
	
	wg.Add(2)
	go sum(&a,&b) // a locks first, wait for b
	go sum(&b,&a) // b lock first, wait for a
	wg.Wait()
}