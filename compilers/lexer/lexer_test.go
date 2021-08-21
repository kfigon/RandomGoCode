package lexer

import "testing"

func TestTokenizer(t *testing.T) {
	input := `if (i==j) z = 0;
	else z=1;`

	got := Tokenize(input)

	exp := []Token {
		{Keyword, "if"},
		{Whitespace, " "},
		{OpenParam, "("},
		{Identifier, "i"},
		{Operator, "=="},
		{Identifier, "j"},
		{Whitespace, " "},
		{Identifier, "z"},
		{Whitespace, " "},
		{Assignment, "="},
		{Whitespace, " "},
		{Number, "0"},
		{Semicolon, ";"},
		{Whitespace, "\n\t"},
		{Keyword, "else"},
		{Whitespace, " "},
		{Identifier, "z"},
		{Assignment, "="},
		{Number, "1"},
		{Semicolon, ";"},
	}

	assertTokens(t, exp, got)
}

func assertTokens(t *testing.T, exp, got []Token) {
	if len(exp) != len(got) {
		t.Fatalf("Invalid array lengths, exp %v, got %v", len(exp), len(got))
	}

	for i := 0; i < len(got); i++ {
		g := got[i]
		e := exp[i]

		if g != e {
			t.Errorf("Invalid element on position %v, exp %v, got %v", i, e, g)
		}
	}
}