package main

import "fmt"

func main() {
	makeCall(false)(
		func(v string) {
			fmt.Println("call succeed:", v)
		},
		func(err error) {
			fmt.Println("call failed:", err)
		},
	)
}

type Result[T any] func(
	onOk func(v T),
	onErr func(err error),
)

func Ok[T any](v T) Result[T] {
	return func(onOk func(v T), onErr func(err error)) {
		onOk(v)
	}
}

func Err[T any](err error) Result[T] {
	return func(onOk func(v T), onErr func(err error)) {
		onErr(err)
	}
}

func makeCall(shouldFail bool) Result[string] {
	return func(onOk func(v string), onErr func(err error)) {
		if shouldFail {
			onErr(fmt.Errorf("oopsie"))
		} else {
			onOk("secret code")
		}
	}
}
