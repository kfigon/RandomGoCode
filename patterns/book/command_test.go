package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// encapsulating invocation
// undo stack, logging, storing commands, queues...

type command interface {
	execute()
}

type executeFn func()
func (e executeFn) execute() {
	e()
}

func TestCommandPattern(t *testing.T) {
	cmdLog := []string{}

	var cmd1 executeFn = func() {
		cmdLog = append(cmdLog, "1")
	}

	var cmd2 executeFn = func() {
		cmdLog = append(cmdLog, "2")
	}
	var noop executeFn = func() {} // null object pattern

	invoker := func(c command) {
		c.execute()
	}

	t.Run("invoke some cmds", func(t *testing.T) {
		cmdLog = []string{}

		invoker(cmd1)
		invoker(cmd2)
		invoker(noop)

		assert.Equal(t, []string{"1", "2"}, cmdLog)
	})

	t.Run("compose commands", func(t *testing.T) {
		cmdLog = []string{}

		var macroCmds executeFn = func() {
			cmd1()
			cmd2()
			cmd2()
			cmd2()
		}
		invoker(macroCmds)

		assert.Equal(t, []string{"1", "2","2","2"}, cmdLog)
	})
}