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

	assertVarStatementAndIntegerExpression(t, tree.Statements[0], 123)
	assertVarStatementAndIntegerExpression(t, tree.Statements[1], 3)	
}

func assertVarStatementAndIntegerExpression(t *testing.T, st StatementNode, exp int) {
	varSt, ok := st.(*VarStatementNode)
	if !ok {
		t.Error("Expected var node in", st.TokenLiteral())
	}
	integer, ok := varSt.Value.(*IntegerLiteralExpression)
	if !ok {
		t.Error("Expected integer literal in", varSt.TokenLiteral())
	}
	if integer.Value != exp {
		t.Errorf("Got integer %v, exp %v", integer.Value, exp)
	}
}

func TestBasicReturnStatement(t *testing.T) {
	tokens := lexer.Tokenize(`return 123;`)
	
	tree := Parse(tokens)

	assertNoErrors(t, tree.Errors)

	if len(tree.Statements) != 1 {
		t.Fatal("Invalid tree built, should be 1, got",len(tree.Statements))
	}
	if tree.Statements[0].TokenLiteral() != "return" {
		t.Error("Invalid literal found", tree.Statements[0])
	}
	
	ret := tree.Statements[0].(*ReturnStatementNode)
	integer := ret.Value.(*IntegerLiteralExpression)

	if integer.Value != 123 {
		t.Errorf("Invalid integer literal, got %v, exp %v", integer.Value, 123)
	}
}

func assertNoErrors(t *testing.T, errors []error) {
	if len(errors) > 0 {
		t.Errorf("Got %v errors, expected none", len(errors))
		for _, v := range errors {
			t.Error(v)
		}
	}
}

func TestInvalidVarStatements(t *testing.T) {
	input := `var asd 4;
	var = 432;`

	tree := Parse(lexer.Tokenize(input))

	if len(tree.Errors) != 2 {
		t.Error("Expected 2 errors, got", tree.Errors)
	}
}

func TestInvalidVarStatementsWithExpressions(t *testing.T) {
	input := `var asd = 4
	var asd = ;`

	tree := Parse(lexer.Tokenize(input))

	if len(tree.Errors) != 1 {
		t.Error("Expected 1 errors, got", len(tree.Errors), tree.Errors)
	}
}

func TestInvalidVarStatementsWithExpressions2(t *testing.T) {
	input := `var asd = ;`

	tree := Parse(lexer.Tokenize(input))

	if len(tree.Errors) != 1 {
		t.Error("Expected 1 errors, got", len(tree.Errors), tree.Errors)
	}
}

func TestVarWithoutAssignment(t *testing.T) {
	tree := Parse(lexer.Tokenize(`var asd;`))
	assertNoErrors(t, tree.Errors)
}