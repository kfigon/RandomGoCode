package lexer

import (
	"testing"
)

func assertTokens(t *testing.T, exp, got []Token) {
	if len(exp) != len(got) {
		t.Errorf("Invalid array lengths, exp %v, got %v", len(exp), len(got))
		t.Errorf("Got: %v", got)
		t.Fatalf("Exp: %v", exp)
	}

	for i := 0; i < len(got); i++ {
		g := got[i]
		e := exp[i]

		if g != e {
			t.Errorf("Invalid element on position %v, exp %v, got %v", i, e, g)
		}
	}
}

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	string
		expectedTokens []Token
	}{
		{
			desc: "simple case with whitespace",
			input: ` i;`,
			expectedTokens: []Token {
				{Whitespace, " "},
				{Identifier, "i"},
				{Semicolon, ";"},
			},
		},
		{
			desc: "complex case 1",
			input: `if (i==j) els = 654.1;
	else els=123;`,
			expectedTokens: []Token {
				{Keyword, "if"},
				{Whitespace, " "},
				{OpenParam, "("},
				{Identifier, "i"},
				{Operator, "=="},
				{Identifier, "j"},
				{CloseParam, ")"},
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
			},
		},
		{
			desc: "case with math operators",
			input: `abc=123 / 2*1-12
	x=3+2;`,
			expectedTokens: []Token{
				{Identifier, "abc"},
				{Assignment, "="},
				{Number, "123"},
				{Whitespace, " "},
				{Operator, "/"},
				{Whitespace, " "},
				{Number, "2"},
				{Operator, "*"},
				{Number, "1"},
				{Operator, "-"},
				{Number, "12"},
				{Whitespace, "\n\t"},
				{Identifier, "x"},
				{Assignment, "="},
				{Number, "3"},
				{Operator, "+"},
				{Number, "2"},
				{Semicolon, ";"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Tokenize(tC.input)		
			assertTokens(t, tC.expectedTokens, got)
		})
	}
}