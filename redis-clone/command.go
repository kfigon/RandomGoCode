package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type command []byte
func (c command) validate() error {
	if len(c) < 3 {
		return fmt.Errorf("too short")
	} else if !c.isStringCmd() && !c.isBulk() && !c.isArray() {
		return fmt.Errorf("invalid first character: %v", c[0])
	} else if (c.isStringCmd() || c.isBulk()) && !stringTerminated(c) {
		return terminationError(c)
	}
	return nil
}

func terminationError(c command) error {
	return fmt.Errorf("invalid termination: %q", string(c.termination()))
}

func stringTerminated(c command) bool {
	return equal(c.termination(), []byte{0x0d, 0x0a})
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
	byteLen int
}

func newBulkString(c command) (*bulkCommand,error) {
	if !c.isBulk() {
		return nil, fmt.Errorf("invalid first byte: %q", c[0])
	} else if !stringTerminated(c) {
		return nil, terminationError(c)
	}

	byteLenStr := ""
	for i := 1; i < len(c)-1; i++ {
		this := c[i]
		next := c[i+1]
		if unicode.IsDigit(rune(this)) {
			byteLenStr += string(this)
		} else if byteLenStr == "" && this == '\r' && next == '\n' {
			return nil, fmt.Errorf("missing length")
		} else if byteLenStr != "" &&  this == '\r' && next == '\n' {
			break
		} else {
			return nil, fmt.Errorf("missing delimiter")
		}
	}
	byteLen, err := strconv.Atoi(byteLenStr)
	if err != nil {
		return nil, fmt.Errorf("invalid byte len %q: %v", byteLenStr, err)
	}
	if  expectedBulkLen(byteLen) > len(c) {
		return nil, fmt.Errorf("too big size")
	}
	return &bulkCommand{c, byteLen}, nil
}

func (b *bulkCommand) bulkString() string {
	charsOfByteLen := int(math.Log10(float64(b.byteLen)))+1
	start := 1+charsOfByteLen+2
	end := start + b.byteLen
	return string(b.command)[start:end]
}

func expectedBulkLen(ln int) int {
	byteLenStr := int(math.Log10(float64(ln)))+1
	return 1 + byteLenStr + 2 + ln + 2
}

func (b *bulkCommand) len() int {
	return expectedBulkLen(b.byteLen)
}

type arrayCommand struct {
	command
	elements int
}

func newArrayString(c command) (*arrayCommand,error) {
	return nil, nil
}

func (a *arrayCommand) commands() []command {
	return nil
}