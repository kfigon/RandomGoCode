package concurency

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelProcessor(t *testing.T) {
	strs := []string{"aa", "bb", "cc"}
	c := make(chan string)
	var wg sync.WaitGroup

	for _, v := range strs {
		wg.Add(1)
		go func() {
			c <- processor(v)
			wg.Done()
		}()
	}

	go func () {
		wg.Wait()
		close(c)
	}()

	out := []string{}
	for v := range c {
		out = append(out, v)
	}
	assert.ElementsMatch(t, []string{"aa foo!", "bb foo!", "cc foo!"}, out)
}

func processor(i string) string {
	return i + " foo!"
}