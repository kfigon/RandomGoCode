package bagqueuestack

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
		assert.Equal(t, []int{5,6,7}, q.elements())
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
		assert.Equal(t, []int{6,7}, q.elements())

		q.enqueue(88)
		assert.Equal(t, []int{6,7,88}, q.elements())

		dequeuePresent(t,q,6)
		assert.Equal(t, []int{7,88}, q.elements())
		dequeuePresent(t,q,7)
		assert.Equal(t, []int{88}, q.elements())
		dequeuePresent(t,q,88)
		assert.Equal(t, []int{}, q.elements())
	})
}

type queue struct{
	first *listNode
	last *listNode
}

func (q *queue) enqueue(val int){
	node := newListNode(val)
	if q.first == nil {
		q.first = node
		q.last = node
	} else if q.first == q.last {
		q.first.next = node
		q.last = node
	} else {
		q.last.next = node
		q.last = node
	}
}

func (q *queue) dequeue()(int,bool){
	if q.first == nil || q.last == nil{
		return -1, false
	}
	if q.first == q.last {
		toRet := q.first.val
		q.first = nil
		q.last = nil
		return toRet,true
	}
	toRet := q.first.val
	q.first = q.first.next
	return toRet, true
}

func (q *queue) elements() []int{
	out := []int{}
	ptr := q.first
	for ptr != nil {
		out = append(out, ptr.val)
		ptr = ptr.next
	}
	return out
}