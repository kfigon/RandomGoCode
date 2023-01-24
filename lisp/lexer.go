package main

import (
	"fmt"
	"unicode"
)

type tokenType int

const (
	opening tokenType = iota
	closing
	operator
	number
	boolean
	keyword
	identifier
	stringLiteral
	invalid
)

type token struct {
	tokType tokenType
	lexeme  string
}

func lex(input string) ([]token, error) {
	out := []token{}
	idx := 0
	lineNumer := 1

	for idx < len(input) {
		current := rune(input[idx])
		if unicode.IsSpace(current) {
			if current == '\n' {
				lineNumer++
			}
		} else if current == ')' {
			out = append(out, token{closing, ")"})
		} else if current == '(' {
			out = append(out, token{opening, "("})
		} else if current == '+' || current == '-' || current == '*' || current == '/' || current == '=' {
			out = append(out, token{operator, string(current)})
		} else if current == '!' || current == '<' || current == '>' {
			if idx +1 < len(input) && input[idx+1] == '=' {
				idx++
				out = append(out, token{operator, string(current)+"="})
			} else {
				out = append(out, token{operator, string(current)})
			}
		} else if current == '"' {
			word := readUntil(input, &idx, func(r rune) bool {return r != '"'})
			if idx +1 < len(input) && input[idx+1] == '"' {
				idx++
				out = append(out, token{stringLiteral, word+"\""})
			} else {
				return nil, fmt.Errorf("Invalid token at line %d: %s", lineNumer, word)
			}
		} else if unicode.IsDigit(current) {
			num := readUntil(input, &idx, unicode.IsDigit)
			out = append(out, token{number, num})
		} else {
			word := readUntil(input, &idx, unicode.IsLetter)
			if isKeyword(word) {
				out = append(out, token{keyword, word})
			} else if word == "true" || word == "false" {
				out = append(out, token{boolean, word})
			} else {
				out = append(out, token{identifier, word})
			}
		}
		idx++
	}
	return out, nil
}

func isKeyword(word string) bool {
	return word == "define" || word == "if"
}

func readUntil(input string, idx *int, fn func(rune)bool) string {
	out := ""
	out += string(input[*idx])
	for *idx +1 < len(input) {
		next := input[*idx+1]
		if fn(rune(next)) {
			*idx++
			out += string(next)
		} else {
			break
		}
	}
	return out
}