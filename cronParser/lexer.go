package main

import (
	"fmt"
	"unicode"
)

type tokenType int

const (
	wildcard tokenType = iota
	number
	div
	comma
	dash
)

func (t tokenType) String() string {
	return []string{
		"wildcard",
		"number",
		"div",
		"comma",
		"dash",
	}[t]
}

type token struct {
	tokType tokenType
	lexeme  string
}

func tokenize(input string) ([]token, error) {
	i := 0
	out := []token{}
	for i < len(input) {
		c := rune(input[i])
		if unicode.IsSpace(c) {
			//skip
		} else if c == '*' {
			out = append(out, token{wildcard,"*"})
		} else if c == '-' {
			out = append(out, token{dash,"-"})
		} else if c == ',' {
			out = append(out, token{comma,","})
		} else if c == '/' {
			out = append(out, token{div,"/"})
		} else if unicode.IsDigit(c) {
			num := ""
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				num += string(input[i])
				i++
			}
			out = append(out, token{number,num})
			continue // skip incrementation in the end
		} else {
			return nil, fmt.Errorf("unknown value: %q", string(c))
		}
		i++
	}

	return out, nil
}