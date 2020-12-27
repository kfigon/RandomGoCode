package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name  == "" {
		return "", errors.New("empty name provided")
	}
	return fmt.Sprintf("Hello from %s", name), nil
}