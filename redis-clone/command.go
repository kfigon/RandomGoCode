package main
import (
	"fmt"
)

type command []byte
func (c command) validate() error {
	if len(c) < 3 {
		return fmt.Errorf("too short")
	} else if !c.isStringCmd() && !c.isBulk() && !c.isArray() {
		return fmt.Errorf("invalid first character: %v", c[0])
	} else if !equal(c.termination(), []byte{0x0d, 0x0a}) {
		return fmt.Errorf("invalid termination: %q", string(c.termination()))
	}
	return nil
}

func (c command) isStringCmd() bool {
	return c[0] == '+'
}

func (c command) isArray() bool {
	return c[0] == '*'
}

func (c command) isBulk() bool {
	return c[0] == '$'
}

func (c command) termination() []byte {
	return c[len(c)-2:]
}

func (c command) simpleString() string {
	return string(c[1:len(c)-2])
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

type bulkCommand struct {
	command
}

func newBulkString(c command) (*bulkCommand,error) {
	
	return &bulkCommand{c}, nil
}

func (b *bulkCommand) simpleString() string {
	return ""
}

func (b *bulkCommand) byteLen() int {
	return 0
}