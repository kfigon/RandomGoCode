package main

import "fmt"

func main() {
	fmt.Println("hello")

	// opt := None[int]()
	opt := Some(42)

	// this is pattern matching
	got := opt(
		func(v int) int {
			fmt.Println("got", v)
			return v
		},
		func() int {
			fmt.Println("nothing, returning default")
			return -1
		},
	)
	fmt.Println(got)
}

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
