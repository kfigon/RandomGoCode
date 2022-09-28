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
	i := 0
	for i < len(in) {
		firstChar := in[i]
		switch {
		case firstChar == 'i':
			return parseInt(in, &i)
		case unicode.IsDigit(rune(firstChar)):
			return parseString(in, &i)
		case firstChar == 'l':
			return parseList(in, &i)
		default:
			i++
		}
	}

	return nil, fmt.Errorf("unknown input: %q", in[0:3])
}

func parseList(in string, idx *int) (listObj, error) {
	*idx++
	out := listObj{}
	for *idx < len(in) {
		if *idx == len(in)-1 && in[*idx] == 'e' {
			break
		}

		v, err := parseString(in, idx)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func parseInt(in string, idx *int) (intObj, error) {
	*idx++
	numStr := ""
	for *idx < len(in) && in[*idx] != 'e' {
		numStr += string(in[*idx])
		*idx++
	}
	v, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %v", err)
	}
	*idx++ // e
	return intObj(v), nil
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

type listObj []bencodeObj
func (_ listObj) dummy() {}

type dictObj map[string]bencodeObj
func (_ dictObj) dummy() {}