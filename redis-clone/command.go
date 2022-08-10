package main
import (
	"fmt"
)

type command []byte
func (c command) validate() error {
	if len(c) < 3 {
		return fmt.Errorf("too short")
	} else if !c.isStringCmd() && c[0] != '*' && c[0] != '$' {
		return fmt.Errorf("invalid first character: %v", c[0])
	} else if !equal(c.termination(), []byte{0x0d, 0x0a}) {
		return fmt.Errorf("invalid termination: %q", string(c.termination()))
	}
	return nil
}

func (c command) isStringCmd() bool {
	return c[0] == '+'
}

func (c command) termination() []byte {
	x := c[len(c)-2:]
	return x
}

func equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}