package main

import "testing"

// pattern to emulate sum types and pattern matching
// using only functions

type Option[T any] func(
	onSome func(T) T,
	onEmpty func() T,
) T

// constructors
func None[T any]() Option[T] {
	return func(onSome func(T) T, onEmpty func() T) T {
		return onEmpty()
	}
}

func Some[T any](v T) Option[T] {
	return func(onSome func(T) T, onEmpty func() T) T {
		return onSome(v)
	}
}

func TestOption(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		opt := Some(42)
		val := opt(
			func(v int) int { return v },
			func() int { return 0 },
		)
		if val != 42 {
			t.Error(val, "!= 42")
		}
	})

	t.Run("none", func(t *testing.T) {
		opt := None[int]()
		val := opt(
			func(v int) int { return v },
			func() int { return 0 },
		)
		if val != 0 {
			t.Error(val, "!= 0")
		}
	})
}
