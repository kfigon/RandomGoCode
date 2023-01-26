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
	semicolon
)

type token struct {
	tokType tokenType
	lexeme  string
}

func (t token) String() string {
	tok := [...]string{
		"opening",
		"closing",
		"operator",
		"number",
		"boolean",
		"keyword",
		"identifier",
		"stringLiteral",
		"semicolon",
	}
	return fmt.Sprintf("(%v, %v)", tok[t.tokType], t.lexeme)
}

func isKeyword(word string) bool {
	return word == "let" || word == "for" || word == "if"
}

func lex(input string) ([]token, error) {
	out := []token{}
	idx := 0
	lineNumer := 1
	peek := func() (rune, bool) {
		if idx+1 >= len(input) {
			return 0, false
		}
		return rune(input[idx+1]), true
	}

	currentChar := func() (rune, bool) {
		if idx >= len(input) {
			return 0, false
		}
		return rune(input[idx]), true
	}

	addTok := func(t token) {
		out = append(out, t)
	}

	for current, ok := currentChar(); ok; current,ok = currentChar() {
		if unicode.IsSpace(current) {
			if current == '\n' {
				lineNumer++
			}
		} else if current == ')' || current == '}' {
			addTok(token{closing, string(current)})
		} else if current == ';' {
			addTok(token{semicolon, string(current)})
		} else if current == '(' || current == '{' {
			addTok(token{closing, string(current)})
		} else if current == '+' || current == '-' || current == '*' || current == '/' {
			addTok(token{operator, string(current)})
		} else if current == '!' || current == '<' || current == '>' || current == '=' {
			if next, ok := peek(); ok && next == '=' {
				idx++
				addTok(token{operator, string(current)+"="})
			} else {
				addTok(token{operator, string(current)})
			}
		} else if current == '|' || current == '&' {
			if next, ok := peek(); ok && next == current {
				idx++
				addTok(token{operator, string(current) + string(next)})
			} else {
				return nil, fmt.Errorf("invalid boolean operator on line %d", lineNumer)
			}
		} else if current == '"' {
			word := readUntil(input, &idx, func(r rune) bool {return r != '"'})
			if next, ok := peek(); ok && next == '"' {
				idx++
				addTok(token{stringLiteral, word+"\""})
			} else {
				return nil, fmt.Errorf("Invalid token at line %d: %s", lineNumer, word)
			}
		} else if unicode.IsDigit(current) {
			num := readUntil(input, &idx, unicode.IsDigit)
			addTok(token{number, num})
		} else {
			word := readUntil(input, &idx, unicode.IsLetter)
			if isKeyword(word) {
				addTok(token{keyword, word})
			} else if word == "true" || word == "false" {
				addTok(token{boolean, word})
			} else {
				addTok(token{identifier, word})
			}
		}
		idx++
	}
	return out, nil
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