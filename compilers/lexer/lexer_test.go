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
		desc           string
		input          string
		expectedTokens []Token
	}{
		{
			desc:  "simple case with whitespace",
			input: ` i;`,
			expectedTokens: []Token{
				{Identifier, "i"},
				{Semicolon, ";"},
			},
		},
		{
			desc: "complex case 1",
			input: `if (i==j) els = 654.1;
	else els=123;`,
			expectedTokens: []Token{
				{Keyword, "if"},
				{OpenParam, "("},
				{Identifier, "i"},
				{Operator, "=="},
				{Identifier, "j"},
				{CloseParam, ")"},
				{Identifier, "els"},
				{Assignment, "="},
				{Number, "654.1"},
				{Semicolon, ";"},
				{Keyword, "else"},
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
				{Operator, "/"},
				{Number, "2"},
				{Operator, "*"},
				{Number, "1"},
				{Operator, "-"},
				{Number, "12"},
				{Identifier, "x"},
				{Assignment, "="},
				{Number, "3"},
				{Operator, "+"},
				{Number, "2"},
				{Semicolon, ";"},
			},
		},
		{
			desc:  "for loop",
			input: `for(i=0;i<3;i++){`,
			expectedTokens: []Token{
				{Keyword, "for"},
				{OpenParam, "("},
				{Identifier, "i"},
				{Assignment, "="},
				{Number, "0"},
				{Semicolon, ";"},
				{Identifier, "i"},
				{Operator, "<"},
				{Number, "3"},
				{Semicolon, ";"},
				{Identifier, "i"},
				{Operator, "++"},
				{CloseParam, ")"},
				{OpenParam, "{"},
			},
		},
		{
			desc:  "var statement",
			input: `var foo = 123 != 3;`,
			expectedTokens: []Token{
				{Class: Keyword, Lexeme: "var"},
				{Class: Identifier, Lexeme: "foo"},
				{Class: Assignment, Lexeme: "="},
				{Class: Number, Lexeme: "123"},
				{Class: Operator, Lexeme: "!="},
				{Class: Number, Lexeme: "3"},
				{Class: Semicolon, Lexeme: ";"},
			},
		},
		{
			desc:  "function declaration",
			input: `fn asd(){}`,
			expectedTokens: []Token{
				{Class: Keyword, Lexeme: "fn"},
				{Class: Identifier, Lexeme: "asd"},
				{Class: OpenParam, Lexeme: "("},
				{Class: CloseParam, Lexeme: ")"},
				{Class: OpenParam, Lexeme: "{"},
				{Class: CloseParam, Lexeme: "}"},
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
