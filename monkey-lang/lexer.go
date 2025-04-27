package main

import (
	"fmt"
	"iter"
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

func (t Token) String() string {
	return fmt.Sprintf("<%v; %s>", t.Typ, t.Lexeme)
}

func Lex(input string) iter.Seq[Token] {
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
	i := 0

	return func(yield func(Token) bool) {
		// no utf support
		for ; i < len(input); i++ {
			c := rune(input[i])
			var tok Token

			if unicode.IsSpace(c) {
				continue
			} else if t, ok := singleCharTokens[c]; ok {
				tok = Token{t, string(c)}
			} else if unicode.IsDigit(c) {
				candidate := readUntil(input, &i, unicode.IsDigit)
				tok = Token{Number, candidate}
			} else {
				word := readUntil(input, &i, func(r rune) bool {
					return unicode.IsDigit(r) || unicode.IsLetter(r)
				})

				if v, ok := keywords[word]; ok {
					tok = Token{v, word}
				} else {
					tok = Token{Identifier, word}
				}
			}


			if !yield(tok) {
				return
			}
		}
		
		if !yield(Token{EOF, ""}) {
			return
		}
	}
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