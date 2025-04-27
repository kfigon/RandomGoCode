package main

import (
	"unicode"
)

type TokenType int
const (
	EOF TokenType = iota
	// identifiers and literals
	Identifier
	Number
	// operators
	Assign
	Plus
	// delimiters
	Comma
	Semicolon
	LParen
	RParen
	LBrace
	RBrace
	// keywords
	Function
	Let
)

func (t TokenType) String() string {
	return [...]string{
		"EOF",
		"Identifier",
		"Number",
		"Assign",
		"Plus",
		"Comma",
		"Semicolon",
		"LParen",
		"RParen",
		"LBrace",
		"RBrace",
		"Function",
		"Let",
	}[t]
}

type Token struct {
	Typ TokenType
	Lexeme string
}

func Lex(input string) []Token {
	out := []Token{}
	
	singleCharTokens := map[rune]TokenType {
		'+': Plus,
		',': Comma,
		';': Semicolon,
		'(': LParen,
		')': RParen,
		'{': LBrace,
		'}': RBrace,
		'=': Assign,
	}

	keywords := map[string]TokenType {
		"fun": Function,
		"let": Let,
	}

	// no utf support
	for i := 0; i < len(input); i++ {
		c := rune(input[i])
		if unicode.IsSpace(c) {
			continue
		} else if t, ok := singleCharTokens[c]; ok {
			out = append(out, Token{t, string(c)})
		} else if unicode.IsDigit(c) {
			candidate := readUntil(input, &i, unicode.IsDigit)
			out = append(out, Token{Number, candidate})
		} else {
			word := readUntil(input, &i, func(r rune) bool {
				return unicode.IsDigit(r) || unicode.IsLetter(r)
			})

			if v, ok := keywords[word]; ok {
				out = append(out, Token{v, word})
			} else {
				out = append(out, Token{Identifier, word})
			}
		}
	}

	out = append(out, Token{EOF,""})
	return out
}

func readUntil(in string, i *int, pred func(rune)bool) string {
	out := string(in[*i])

	for *i + 1 < len(in) {
		if pred(rune(in[*i + 1])) {
			out += string(in[*i + 1])
			*i++
		} else {
			break
		}
	}
	return out
}