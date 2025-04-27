package main

import (
	"fmt"
	"iter"
	"unicode"
)

type TokenType string
const (
	EOF TokenType = "EOF"

	// identifiers and literals
	Identifier = "Identifier"
	Number = "Number"

	// operators
	Assign = "Assign"
	Plus = "Plus"
	Minus = "Minus"
	Bang = "Bang"
	Asterisk = "Asterisk"
	Slash = "Slash"
	LT = "LT"
	GT = "GT"
	EQ = "=="
	NEQ = "!="

	// delimiters
	Comma = "Comma"
	Semicolon = "Semicolon"
	LParen = "LParen"
	RParen = "RParen"
	LBrace = "LBrace"
	RBrace

	// keywords
	Function = "Function"
	Let = "Let"
	If = "If"
	Else = "Else"
	Return = "Return"
	True = "true"
	False = "false"
	For = "for"
)

func (t TokenType) String() string {
	return string(t)
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
		'-': Minus,
		'!': Bang,
		'*': Asterisk,
		'/': Slash,
		'>': GT,
		'<': LT,
	}

	keywords := map[string]TokenType {
		"fun": Function,
		"let": Let,
		"if": If,
		"else": Else,
		"return": Return,
		"true": True,
		"false": False,
		"for": For,
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
				if c == '!' && peekNext(input, i, '=') {
					i++
					tok = Token{NEQ, "!="}
				} else if c == '=' && peekNext(input, i, '=') {
					i++
					tok = Token{EQ, "=="}
				} else {
					tok = Token{t, string(c)}
				}
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

func peekNext(in string, i int, exp rune) bool {
	return i+1 < len(in) && rune(in[i+1]) == exp
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