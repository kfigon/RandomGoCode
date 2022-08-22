package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// when we have 2 interfaces and we want to use one with another
// wrapper on one interface to satisfy the other

type duckInterface interface {
	fly() string
	quack() string
}

type ducky struct{}
func (d *ducky) fly() string {
	return "flying"
}

func (d *ducky) quack() string {
	return "quacking"
}

type turkeyInterface interface {
	shortFly() string
	makeNoise() string
}

type turkey struct{}
func (t *turkey) shortFly() string {
	return "short fly"
}

func (t *turkey) makeNoise() string {
	return "noise!"
}

type turkeyDuckAdapter struct {
	turkey
}

func (t *turkeyDuckAdapter) fly() string {
	return t.shortFly()
}

func (t *turkeyDuckAdapter) quack() string {
	return t.makeNoise()
}

func TestAdapter(t *testing.T) {
	consumer := func (d duckInterface) string {
		return d.fly() + " " + d.quack()
	}

	t.Run("Duck", func(t *testing.T) {
		res := consumer(&ducky{})
		assert.Equal(t, "flying quacking", res)
	})

	t.Run("Turkey in duck adapter", func(t *testing.T) {
		res := consumer(&turkeyDuckAdapter{turkey{}})
		assert.Equal(t, "short fly noise!", res)
	})
}

