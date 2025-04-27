package main

import (
	"slices"
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
			desc: "some keywords",
			input:`let    true;
			for false{}
			if asdf return true false` ,
			exp: []Token{
				{Let, "let"},
				{True, "true"},
				{Semicolon, ";"},
				{For, "for"},
				{False, "false"},
				{LBrace, "{"},
				{RBrace, "}"},
				{If, "if"},
				{Identifier, "asdf"},
				{Return, "return"},
				{True, "true"},
				{False, "false"},
				{EOF, ""},
			},
		},
		{
			desc: "control chars",
			input: "=,(){} == !=,fun",
			exp: []Token{
				{Assign, "="},
				{Comma, ","},
				{LParen, "("},
				{RParen, ")"},
				{LBrace, "{"},
				{RBrace, "}"},
				{EQ, "=="},
				{NEQ, "!="},
				{Comma, ","},
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
		{
			desc: "operators",
			input: "<> asdf /!*x12",
			exp: []Token{
				{LT, "<"},
				{GT, ">"},
				{Identifier, "asdf"},
				{Slash, "/"},
				{Bang, "!"},
				{Asterisk, "*"},
				{Identifier, "x12"},
				{EOF, ""},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Lex(tC.input)
			assert.Equal(t, tC.exp, slices.Collect(got))
		})
	}
}