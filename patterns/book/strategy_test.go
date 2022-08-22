package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// strategy - algorithms decoupled and set using composition
// able to change in runtime
// more codereuse than just interface on subtypes

// can be used also as template method to define details of an algorithm in a general outline

// base class, fully abstract
type duck interface {
	// these 2 are specific to every duck
	swim() string
	display() string

	// new requirements, need code reuse. Change less often
	fly() string
	quack() string
}

// alternative - copy paste code for fly and quack. Strategy pattern is more extensible

// pluggable with reuse
type flyBehavior interface {
	fly() string
}

type quackBehavior interface {
	quack() string
}

// concrete duck types
type mallardDuck struct {
	flyBehavior
	quackBehavior
}

func (m *mallardDuck) display() string {
	return "mallard on the screen "
}

func (m *mallardDuck) swim() string {
	return "mallard swims "
}

type rubberdDuck struct {
	flyBehavior
	quackBehavior
}

func (r *rubberdDuck) display() string {
	return "rubber on the screen "
}

func (r *rubberdDuck) swim() string {
	return "rubber duck swims "
}

type woodenDuck struct {
	flyBehavior
	quackBehavior
}

func (w *woodenDuck) display() string {
	return "wooden duck on the screen "
}

func (w *woodenDuck) swim() string {
	return "sinks "
}

type flying struct{}
func (f *flying) fly() string {
	return "flying high "
}

// todo: isnt this violation of Liskov?!
type notFlying struct{}
func (f *notFlying) fly() string {
	return ""
}

type quacking struct{}
func (q *quacking) quack() string {
	return "quacks like a duck "
}

type squicking struct{}
func (s *squicking) quack() string {
	return "squick "
}

type notQuacking struct{}
func (f *notQuacking) quack() string {
	return ""
}

func TestDucks(t *testing.T) {
	consumeDuck := func(d duck) string {
		out := ""
		out += d.display()
		out += d.fly()
		out += d.swim()
		out += d.quack()
		return out
	}

	t.Run("Mallard duck", func(t *testing.T) {
		d := &mallardDuck{flyBehavior: &flying{}, quackBehavior: &quacking{}}
		res := consumeDuck(d)
		assert.Equal(t, "mallard on the screen flying high mallard swims quacks like a duck ", res)
	})

	t.Run("Rubber duck", func(t *testing.T) {
		d := &rubberdDuck{flyBehavior: &notFlying{}, quackBehavior: &squicking{}}
		res := consumeDuck(d)
		assert.Equal(t, "rubber on the screen rubber duck swims squick ", res)
	})

	t.Run("wooden toy duck", func(t *testing.T) {
		d := &woodenDuck{flyBehavior: &notFlying{}, quackBehavior: &notQuacking{}}
		res := consumeDuck(d)
		assert.Equal(t, "wooden duck on the screen sinks ", res)
	})
}
