package parser

import (
	"testing"
	"programming-lang/lexer"
)

func TestVarStatement_Identifiers(t *testing.T) {
	tokens := lexer.Tokenize(`var foo = 123;
	var asd = 3;`)
	
	tree := Parse(tokens)

	assertNoErrors(t, tree.Errors)

	if len(tree.Statements) != 2 {
		t.Fatal("Invalid tree built, should be 2, got", len(tree.Statements))
	}
	if tree.Statements[0].TokenLiteral() != "foo" {
		t.Error("Invalid first statement", tree.Statements[0])
	}
	if tree.Statements[1].TokenLiteral() != "asd" {
		t.Error("Invalid second statement", tree.Statements[1])
	}
	// todo: assert expressions
}

func TestBasicReturnStatement(t *testing.T) {
	tokens := lexer.Tokenize(`return 123;`)
	
	tree := Parse(tokens)

	assertNoErrors(t, tree.Errors)

	if len(tree.Statements) != 1 {
		t.Fatal("Invalid tree built, should be 1, got",len(tree.Statements))
	}
	// todo: assert expressions	
}

func assertNoErrors(t *testing.T, errors []error) {
	if len(errors) > 0 {
		t.Errorf("Got %v errors, expected none", len(errors))
		for _, v := range errors {
			t.Error(v)
		}
	}
}