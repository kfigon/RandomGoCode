package parser

import (
	"testing"
	"programming-lang/lexer"
)

func TestVarStatement(t *testing.T) {
	tokens := lexer.Tokenize(`var foo = 123;
	var asd = 3;`)
	
	tree := Parse(tokens)
	if len(tree.Statements) != 2 {
		t.Fatal("Invalid tree built, should be 2, got",len(tree.Statements))
	}
	if tree.Statements[0].Name == nil || 
		*tree.Statements[0].Name != "foo" || 
		tree.Statements[0].Value == nil {
		t.Error("Invalid first statement", tree.Statements[0])
	}
	if tree.Statements[1].Name == nil || 
		*tree.Statements[1].Name != "asd" || 
		tree.Statements[1].Value == nil {
		t.Error("Invalid second statement", tree.Statements[1])
	}
}