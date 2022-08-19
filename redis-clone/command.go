package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

const DELIMITER_LENGTH = 2

type commandBase interface{
	dummy()
}

func parseCommand(b []byte) (commandBase,error) {
	cmd := command(b)
	if err := cmd.basicValidation(); err != nil {
		return nil, err
	}
	if cmd.isStringCmd() {
		return newSimpleString(cmd)
	} else if cmd.isBulk() {
		return newBulkString(cmd)
	} else if cmd.isArray() {
		return newArrayString(cmd)
	}
	return nil, fmt.Errorf("invalid command")
}

type command []byte

func (c command) basicValidation() error {
	if len(c) < 3 {
		return fmt.Errorf("too short")
	} else if !c.isStringCmd() && !c.isBulk() && !c.isArray() {
		return fmt.Errorf("invalid first character: %v", c[0])
	}
	return nil
}

func terminationError(termination []byte) error {
	return fmt.Errorf("invalid termination: %q", string(termination))
}

func stringTerminated(termination []byte) bool {
	return equal(termination, []byte{0x0d, 0x0a})
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

func parseLengthToken(c command) (int, error) {
	byteLenStr := ""
	for i := 1; i < len(c)-1; i++ {
		this := c[i]
		next := c[i+1]
		if unicode.IsDigit(rune(this)) {
			byteLenStr += string(this)
		} else if byteLenStr == "" && this == '\r' && next == '\n' {
			return 0, fmt.Errorf("missing length")
		} else if byteLenStr != "" && this == '\r' && next == '\n' {
			break
		} else {
			return 0, fmt.Errorf("missing delimiter")
		}
	}
	byteLen, err := strconv.Atoi(byteLenStr)
	if err != nil {
		return 0, fmt.Errorf("invalid len %q: %v", byteLenStr, err)
	}
	return byteLen, nil
}

type simpleStringCommand struct {
	command
}

func newSimpleString(c command) (*simpleStringCommand, error) {
	if !c.isStringCmd() {
		return nil, fmt.Errorf("invalid first byte: %q", c[0])
	}
	ln := 1
	for ln < len(c)-1 {
		this := c[ln]
		next := c[ln+1]
		if this == '\r' && next == '\n' {
			break
		}
		ln++
	}
	if !stringTerminated(c[ln:ln+2]) {
		return nil, terminationError(c[ln:ln+2])
	}
	return &simpleStringCommand{c[0:ln+2]}, nil
}

func (s *simpleStringCommand) termination() []byte {
	c := s.command
	return c[len(c)-DELIMITER_LENGTH:]
}

func (s *simpleStringCommand) simpleString() string {
	c := s.command
	return string(c[1 : len(c)-DELIMITER_LENGTH])
}

func (_ *simpleStringCommand) dummy(){}

type bulkCommand struct {
	command
	byteLen int
}

func newBulkString(c command) (*bulkCommand, error) {
	if !c.isBulk() {
		return nil, fmt.Errorf("invalid first byte: %q", c[0])
	}

	byteLen, err := parseLengthToken(c)
	if err != nil {
		return nil, err
	} else if expectedBulkLen(byteLen) > len(c) {
		return nil, fmt.Errorf("invalid length")
	} else if !stringTerminated(c[expectedBulkLen(byteLen)-2 : expectedBulkLen(byteLen)]) {
		return nil, terminationError(c[expectedBulkLen(byteLen)-2 : expectedBulkLen(byteLen)])
	}
	return &bulkCommand{c[0:expectedBulkLen(byteLen)], byteLen}, nil
}

func (b *bulkCommand) bulkString() string {
	charsOfByteLen := charLenOfNum(b.byteLen)
	start := 1 + charsOfByteLen + 2
	end := start + b.byteLen
	return string(b.command)[start:end]
}
func (_ *bulkCommand) dummy(){}

func expectedBulkLen(ln int) int {
	byteLenStr := charLenOfNum(ln)
	return 1 + byteLenStr + DELIMITER_LENGTH + ln + DELIMITER_LENGTH
}

func charLenOfNum(ln int) int {
	return int(math.Log10(float64(ln))) + 1
}

func (b *bulkCommand) len() int {
	return expectedBulkLen(b.byteLen)
}

type arrayCommand struct {
	cmds []command
}

func newArrayString(c command) (*arrayCommand, error) {
	if !c.isArray() {
		return nil, fmt.Errorf("invalid first byte: %q", c[0])
	}
	arrayLen, err := parseLengthToken(c)
	if err != nil {
		return nil, err
	}
	cmds := []command{}

	i := 1 + charLenOfNum(arrayLen) + DELIMITER_LENGTH
	for i < len(c) {
		subCmd := command(c[i:])

		if err := subCmd.basicValidation(); err != nil {
			return nil, err
		}

		switch {
		case subCmd.isArray():
			_, err := newArrayString(subCmd)
			if err != nil {
				return nil, err
			}
			// todo: store it somehow... type switch probably
			i++
			break
		case subCmd.isStringCmd():
			s, err := newSimpleString(subCmd)
			if err != nil {
				return nil, err
			}
			cmds = append(cmds, s.command)
			i += len(s.command)
		case subCmd.isBulk():
			b, err := newBulkString(subCmd)
			if err != nil {
				return nil, err
			}
			cmds = append(cmds, b.command)
			i += len(b.command)
		default:
			return nil, fmt.Errorf("invalid fist byte %q", subCmd[0])
		}
	}

	return &arrayCommand{cmds}, nil
}

func (_ *arrayCommand) dummy(){}

func (a *arrayCommand) commands() []string {
	var out []string
	for _, c := range a.cmds {
		if c.isStringCmd() {
			s, _ := newSimpleString(c)
			out = append(out, s.simpleString())
		} else if c.isBulk() {
			b, _ := newBulkString(c)
			out = append(out, b.bulkString())
		} else if c.isArray() {
			// todo
		}
	}
	return out
}
