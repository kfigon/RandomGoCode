package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decode(in string) (bencodeObj, error) {
	if len(in) < 3 {
		return nil, fmt.Errorf("input too short: %v", in)
	}

	firstChar := in[0]
	lastChar := in[len(in)-1]
	switch {
	case firstChar == 'i' && lastChar == 'e':
		v, err := strconv.Atoi(in[1:len(in)-1])
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %v", err)
		}
		return intObj(v), nil
	case unicode.IsDigit(rune(firstChar)):
		idx := strings.Index(in, ":")
		if idx == -1 {
			return nil, fmt.Errorf("invalid string, no ':'")
		}
		v, err := strconv.Atoi(in[0:idx])
		if err != nil {
			return nil, fmt.Errorf("invalid string length: %v", err)
		} else if idx+v >= len(in){
			return nil, fmt.Errorf("too long string, expected %v, got %v", (idx+v), len(in))
		}
		str := in[idx+1:idx+1+v]
		return stringObj(str), nil
	case firstChar == 'l' && lastChar == 'e':
	case firstChar == 'd' && lastChar == 'e':
	}
	return nil, fmt.Errorf("unknown input: %q", in[0:3])
}

// lack of sumtypes
type bencodeObj interface {
	dummy()
}

type stringObj string
func (_ stringObj) dummy(){}

type intObj int
func (_ intObj) dummy() {}

type listObj []any
func (_ listObj) dummy() {}

type dictObj map[string]any
func (_ dictObj) dummy() {}