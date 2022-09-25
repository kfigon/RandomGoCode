package main

import (
	"fmt"
	"strconv"
)

func decode(in string) (bencodeObj, error) {
	if len(in) < 3 {
		return nil, fmt.Errorf("input too short: %v", in)
	}

	firstChar := in[0]
	lastChar := in[len(in)-1]
	switch {
	case firstChar == 'i' && lastChar == 'e': {
		v, err := strconv.Atoi(in[1:len(in)-1])
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %v", err)
		}
		return intObj(v), nil
	}
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