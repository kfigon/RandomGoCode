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
	return parseObj(in, &i)
}

func parseObj(in string, i *int) (bencodeObj, error) {
	firstChar := in[*i]
	switch {
	case firstChar == 'i':
		return parseInt(in, i)
	case unicode.IsDigit(rune(firstChar)):
		return parseString(in, i)
	case firstChar == 'l':
		return parseList(in, i)
	case firstChar == 'd':
		return parseDict(in, i)
	default:
		*i++
	}
	return nil, fmt.Errorf("unknown type %v", firstChar)
}

func parseDict(in string, i *int) (dictObj, error) {
	*i++
	out := dictObj{}
	for *i < len(in) && in[*i] != 'e'{
		obj, err := parseObj(in, i)
		if err != nil {
			return nil, err
		}
		strObj, ok := obj.(stringObj)
		if !ok {
			return nil, fmt.Errorf("key is not string %v", strObj)
		}
		
		obj, err = parseObj(in, i)
		if err != nil {
			return nil, err
		}
		out[string(strObj)] = obj
	}
	*i++
	return out, nil
}

func parseList(in string, idx *int) (listObj, error) {
	*idx++
	out := listObj{}
	for *idx < len(in) && in[*idx] != 'e' {
		obj, err := parseObj(in, idx)
		if err != nil {
			return nil, err
		}
		out = append(out, obj)
	}
	*idx++
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