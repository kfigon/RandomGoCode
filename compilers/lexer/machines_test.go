package lexer

import (
	"testing"
)

func TestLessSimpleAutomaton(t *testing.T) {
	type stateMachine func() (string, bool)
	drain := func(machine stateMachine) string {
		state, more := machine()
		for more {
			state, more = machine()
		}
		return state
	}

	t.Run("Simple run, invalid", func(t *testing.T) {
		machine := anyNumberOfOnes("11")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got", state)
		}
	})

	t.Run("Simple run, no ones, valid", func(t *testing.T) {
		machine := anyNumberOfOnes("0")
		state := drain(machine)
		if state != "accept" {
			t.Error("Invalid state, expected accept, got", state)
		}
	})

	t.Run("Simple run, invalid2", func(t *testing.T) {
		machine := anyNumberOfOnes("100")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got", state)
		}
	})

	t.Run("Simple run - ok", func(t *testing.T) {
		machine := anyNumberOfOnes("110")
		state := drain(machine)
		if state != "accept" {
			t.Error("Invalid state, expected accept, got", state)
		}
	})
}

// go test ./lexer -run TestSimpleAutomaton
func TestSimpleAutomaton(t *testing.T) {
	type stateMachine func() (string, bool)
	drain := func(machine stateMachine) string {
		state, more := machine()
		for more {
			state, more = machine()
		}
		return state
	}

	t.Run("Simple run", func(t *testing.T) {
		machine := simpleState("1")
		state := drain(machine)
		if state != "accept" {
			t.Error("Invalid state, expected accept, got", state)
		}
	})

	t.Run("Simple run - invalid", func(t *testing.T) {
		machine := simpleState("0")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got", state)
		}
	})

	t.Run("Simple run - too long string", func(t *testing.T) {
		machine := simpleState("10")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got", state)
		}
	})
}
