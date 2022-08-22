package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type bevarage interface {
	description() string
}

type espressoCoffee struct{}
func (e *espressoCoffee) description() string {
	return "espresso"
}

type whipDecorator struct{
	bevarage
}
func (w *whipDecorator) description() string {
	return w.bevarage.description() + " with whip"
}

type chocolateDecorator struct{
	bevarage
}
func (c *chocolateDecorator) description() string {
	return c.bevarage.description() + " with chocolate"
}


type bevarageFn func()string
type bevarageDecoratorFn func(bevarageFn) bevarageFn

func TestDecorator(t *testing.T) {
	t.Run("objects", func(t *testing.T) {
		espresso := &espressoCoffee{}
		withWhip := &whipDecorator{espresso}
		withChocolate := &chocolateDecorator{withWhip}

		assert.Equal(t, "espresso with whip with chocolate", withChocolate.description())
	})

	t.Run("functions", func(t *testing.T) {
		var espressoFn bevarageFn = func () string {
			return "espresso"
		}
		var whipDecorator bevarageDecoratorFn = func (fn bevarageFn) bevarageFn {
			return func() string {
				return fn() + " with whip"
			}
		}
		var chocolateDecorator bevarageDecoratorFn = func (fn bevarageFn) bevarageFn {
			return func() string {
				return fn() + " with chocolate"
			}
		}

		espressoWithWhipAndChocolate := chocolateDecorator(whipDecorator(espressoFn))
		assert.Equal(t, "espresso with whip with chocolate", espressoWithWhipAndChocolate())
	})

	t.Run("functions generic", func(t *testing.T) {
		var espressoFn bevarageFn = func () string {
			return "espresso"
		}

		genericDecorator := func(what string) func(bevarageFn)bevarageFn {
			return func(bf bevarageFn) bevarageFn {
				return func() string { return bf() + " " + what}
			}
		}
		withWhip := genericDecorator("with whip")
		withChocolate := genericDecorator("with chocolate")

		espressoAndWhip := withWhip(espressoFn)
		espressoWithWhipAndChocolate := withChocolate(espressoAndWhip)

		assert.Equal(t, "espresso with whip with chocolate", espressoWithWhipAndChocolate())
	})
}
