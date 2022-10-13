package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	dequeuePresent := func(t *testing.T, q *queue, exp int) {
		v, ok := q.dequeue()
		assert.True(t, ok)
		assert.Equal(t, exp, v)
	}

	dequeueAbsent := func(t *testing.T, q *queue) {
		_, ok := q.dequeue()
		assert.False(t, ok)
	}

	t.Run("empty", func(t *testing.T) {
		q := &queue{}
		assert.Equal(t, []int{}, q.elements())
	})

	t.Run("dequeue empty", func(t *testing.T) {
		q := &queue{}
		dequeueAbsent(t,q)
		assert.Equal(t, []int{}, q.elements())
	})

	t.Run("enqueue", func(t *testing.T) {
		q := &queue{}
		q.enqueue(5)
		assert.Equal(t, []int{5}, q.elements())
	})

	t.Run("enqueue more", func(t *testing.T) {
		q := &queue{}
		q.enqueue(5)
		q.enqueue(6)
		q.enqueue(7)
		assert.Equal(t, []int{7,6,5}, q.elements())
	})

	t.Run("dequeue single", func(t *testing.T) {
		q := &queue{}
		q.enqueue(5)

		dequeuePresent(t,q,5)
		dequeueAbsent(t,q)
		dequeueAbsent(t,q)

		assert.Equal(t, []int{}, q.elements())
	})

	t.Run("dequeue more", func(t *testing.T) {
		q := &queue{}
		q.enqueue(5)
		q.enqueue(6)
		q.enqueue(7)
		
		dequeuePresent(t,q,5)
		assert.Equal(t, []int{7,6}, q.elements())

		q.enqueue(88)
		assert.Equal(t, []int{88,7,6}, q.elements())

		dequeuePresent(t,q,6)
		assert.Equal(t, []int{88,7}, q.elements())
		dequeuePresent(t,q,7)
		assert.Equal(t, []int{88}, q.elements())
		dequeuePresent(t,q,88)
		assert.Equal(t, []int{}, q.elements())
	})
}

type queue struct{}

func (q *queue) enqueue(val int){}
func (q *queue) dequeue()(int,bool){
	return -1,false
}

func (q *queue) elements() []int{
	return nil
}