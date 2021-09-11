package parser

import (
	"testing"
	"programming-lang/lexer"
)

func TestBasicMathTree(t *testing.T) {
	// 1+2*3
	tokens := []lexer.Token{
		{lexer.Number, "1"},
		{lexer.Operator, "+"},
		{lexer.Number, "2"},
		{lexer.Operator, "*"},
		{lexer.Number, "3"},
	}

	tree := Parse(tokens)

	if tree == nil {
		t.Error("Invalid tree built")
	}
}