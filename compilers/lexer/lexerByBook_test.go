package lexer

import (
	"testing"
)

func TestLlexer(t *testing.T) {
	input := `let asd = 123;
for(i=1;i<3; i++){}`

	expected := []llexToken{
		{"KEYWORD", "let"},
		{"IDENTIFIER", "asd"},
		{"ASSIGNMENT", "="},
		{"NUMBER", "123"},
		{"SEMICOLON", ";"},
		{"KEYWORD", "for"},
		{"OPEN_BRACE", "("},
		{"IDENTIFIER", "i"},
		{"ASSIGNMENT", "="},
		{"NUMBER", "1"},
		{"SEMICOLON", ";"},
		{"IDENTIFIER", "i"},
		{"OPERATOR", "<"},
		{"NUMBER", "3"},
		{"SEMICOLON", ";"},
		{"IDENTIFIER", "i"},
		{"OPERATOR", "++"},
		{"CLOSE_BRACE", ")"},
		{"OPEN_CBRACE", "{"},
		{"CLOSE_CBRACE", "}"},
	}

	llex := newLlexer(input)
	var got []llexToken
	for token := llex.emit(); token.token != "EOF"; token = llex.emit() {
		got = append(got, token)
	}

	if len(expected) != len(got) {
		t.Error("GOT", got)
		t.Error("EXP", expected)
		t.Fatal("Invalid lengths got", len(got), "exp", len(expected))
	}

	for i := 0; i < len(got); i++ {
		g := got[i]
		e := expected[i]
		if g != e {
			t.Errorf("Error in %v, got %v, exp %v", i, g, e)
		}
	}
}

func TestSimpleIdentifier(t *testing.T) {
	l := newLlexer(`let;`)
	to := l.emit()
	semi := l.emit()

	if next := l.emit(); next.token != "EOF" {
		t.Error("Found not terminating token", next)
	}
	if to.token != "KEYWORD" || to.lexeme != "let" {
		t.Error("First token invalid", to)
	}
	if semi.token != "SEMICOLON" || semi.lexeme != ";" {
		t.Error("Second token invalid", semi)
	}
}