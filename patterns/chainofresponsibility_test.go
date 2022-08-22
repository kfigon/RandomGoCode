package patterns

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// functional chain of responsibility - middleware http handlers in go

type handler interface {
	setNext(h handler)
	handle(string)
}

type validatorHandler struct{
	next handler
}
func (a *validatorHandler) setNext(h handler) {
	a.next = h
}
func (a *validatorHandler) handle(v string) {
	if v == "" {
		return
	}
	a.next.handle(v)
}

type processorHandler struct{
	next handler
}
func (a *processorHandler) setNext(h handler) {
	a.next = h
}
func (a *processorHandler) handle(v string) {
	a.next.handle(v + "123")
}

type toIntHandler struct{
	next handler
	result *int
}
func (a *toIntHandler) setNext(h handler) {
	a.next = h
}
func (a *toIntHandler) handle(v string) {
	i, err := strconv.Atoi(v)
	if err != nil {
		return
	}
	a.result = &i
}

func TestChainOfResponsibilityPattern(t *testing.T) {
	init := func() (handler, *toIntHandler) {
		val := &validatorHandler{}
		proc := &processorHandler{}
		result := &toIntHandler{}
	
		val.setNext(proc)
		proc.setNext(result)

		return val, result
	}

	t.Run("Empty string", func(t *testing.T) {
		chainStart, result := init()
		chainStart.handle("")
		assert.Nil(t, result.result)
	})

	t.Run("Not digit", func(t *testing.T) {
		chainStart, result := init()

		chainStart.handle("foobar")
		assert.Nil(t, result.result)
	})

	t.Run("valid", func(t *testing.T) {
		chainStart, result := init()

		chainStart.handle("5")
		assert.NotNil(t, result.result)
		assert.Equal(t, 5123, *result.result)
	})
}