package parser

import (
	"testing"
	"programming-lang/lexer"
)

func TestBasicMathTree(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		// 1+2
		tokens := []lexer.Token{
			{Class: lexer.Number, Lexeme: "1"},
			{Class: lexer.Operator, Lexeme: "+"},
			{Class: lexer.Number, Lexeme: "2"},
		}

		tree := Parse(tokens)

		if tree == nil {
			t.Error("Invalid tree built")
		}	
	})

	t.Run("Operation order", func(t *testing.T) {
		// 1+2*3
		tokens := []lexer.Token{
			{Class: lexer.Number, Lexeme: "1"},
			{Class: lexer.Operator, Lexeme: "+"},
			{Class: lexer.Number, Lexeme: "2"},
			{Class: lexer.Operator, Lexeme: "*"},
			{Class: lexer.Number, Lexeme: "3"},
		}

		tree := Parse(tokens)

		if tree == nil {
			t.Error("Invalid tree built")
		}	
	})
	
}