package main

import (
	"fmt"
	"strconv"
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
		i := 0
		return parseString(in, &i)
	case firstChar == 'l' && lastChar == 'e':
		out := []any{}

		return listObj(out), nil
	case firstChar == 'd' && lastChar == 'e':
	}
	return nil, fmt.Errorf("unknown input: %q", in[0:3])
}

func parseString(in string, idx *int) (stringObj, error) {
	declaredLen := ""
	for *idx < len(in) && in[*idx] != ':' {
		declaredLen += string(in[*idx])
		*idx++
	}
	// :
	*idx++

	v, err := strconv.Atoi(declaredLen)
	if err != nil {
		return "", fmt.Errorf("invalid string length: %v", err)
	} else if *idx+v > len(in){
		return "", fmt.Errorf("too long string, expected %v, got %v", (*idx+v), len(in))
	}

	data := ""
	strLen := 0
	for *idx < len(in) && strLen < v {
		data += string(in[*idx])
		*idx++
		strLen++
	}
	return stringObj(data), nil
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