package lexer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	input := `if (i==j) els = 654.1;
	else els=123;`

	got := Tokenize(input)

	exp := []Token {
		{Keyword, "if"},
		{Whitespace, " "},
		{OpenParam, "("},
		{Identifier, "i"},
		{Operator, "=="},
		{Identifier, "j"},
		{Whitespace, " "},
		{Identifier, "els"},
		{Whitespace, " "},
		{Assignment, "="},
		{Whitespace, " "},
		{Number, "654.1"},
		{Semicolon, ";"},
		{Whitespace, "\n\t"},
		{Keyword, "else"},
		{Whitespace, " "},
		{Identifier, "els"},
		{Assignment, "="},
		{Number, "123"},
		{Semicolon, ";"},
	}

	assertTokens(t, exp, got)
}

func assertTokens(t *testing.T, exp, got []Token) {
	if len(exp) != len(got) {
		t.Errorf("Invalid array lengths, exp %v, got %v", len(exp), len(got))
		t.Fatalf("Got: %v", got)
	}

	for i := 0; i < len(got); i++ {
		g := got[i]
		e := exp[i]

		if g != e {
			t.Errorf("Invalid element on position %v, exp %v, got %v", i, e, g)
		}
	}
}