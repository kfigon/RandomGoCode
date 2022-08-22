package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type command func()string

type lamp struct {
	state string
	onCommand command
	offCommand command
}

func newLamp(onCommand, offCommand command) *lamp {
	invoker := lamp{onCommand: onCommand, offCommand: offCommand}
	invoker.state = "off"
	return &invoker
}

func (i *lamp) on() {
	i.state = i.onCommand()
}

func (i *lamp) off() {
	i.state = i.offCommand()
}

func TestCommandPattern(t *testing.T) {
	on := func() string {
		return "on"
	}

	off := func() string {
		return "off"
	}

	invoker := newLamp(on, off)
	assert.Equal(t, "off", invoker.state)
	
	invoker.on()
	assert.Equal(t, "on", invoker.state)
	
	invoker.off()
	assert.Equal(t, "off", invoker.state)
}
