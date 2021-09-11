package parser

import (
	"testing"
	"programming-lang/lexer"
)

func TestBasicMathTree(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		// 1+2
		tokens := []lexer.Token{
			{lexer.Number, "1"},
			{lexer.Operator, "+"},
			{lexer.Number, "2"},
		}

		tree := Parse(tokens)

		if tree == nil {
			t.Error("Invalid tree built")
		}	
	})

	t.Run("Operation order", func(t *testing.T) {
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
	})
	
}