package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		desc	string
		input string
		exp []Token
	}{
		{
			desc: "basic declaration",
			input: "let abc = 123;",
			exp: []Token{
				{Let, "let"},
				{Identifier, "abc"},
				{Assign, "="},
				{Number, "123"},
				{Semicolon, ";"},
				{EOF, ""},
			},
		},
		{
			desc: "control chars",
			input: "=,(){} fun",
			exp: []Token{
				{Assign, "="},
				{Comma, ","},
				{LParen, "("},
				{RParen, ")"},
				{LBrace, "{"},
				{RBrace, "}"},
				{Function, "fun"},
				{EOF, ""},
			},
		},
		{
			desc: "identifier at end",
			input: "let foobar",
			exp: []Token{
				{Let, "let"},
				{Identifier, "foobar"},
				{EOF, ""},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Lex(tC.input)
			assert.Equal(t, tC.exp, got)
		})
	}
}