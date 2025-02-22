package concurency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	numbers := make(chan int)
	squared := make(chan int)
	limit := 5
	counter := func() {
		for i := 1; i <= limit; i++ {
			numbers<-i
		}
		close(numbers)
	}

	squarer := func() {
		// for { // !ok - closed channel
		// 	v, ok := <- numbers
		// }
		for v := range numbers {
			squared <- v*v
		}
		close(squared)
	}

	go counter()
	go squarer()
	data := []int{}
	for v := range squared {
		data = append(data, v)
	}
	assert.Equal(t, []int{1,4,9,16,25}, data)
}

func TestPipelineWithoutClosure(t *testing.T) {
	numbers := make(chan int)
	squared := make(chan int)
	
	counter := func(out chan<- int) {
		limit := 5
		for i := 1; i <= limit; i++ {
			out<-i
		}
		close(out)
	}

	squarer := func(in <-chan int, out chan<- int) {
		for v := range in {
			out <- v*v
		}
		close(out)
	}

	go counter(numbers)
	go squarer(numbers, squared)
	
	data := []int{}
	for v := range squared {
		data = append(data, v)
	}
	assert.Equal(t, []int{1,4,9,16,25}, data)
}
