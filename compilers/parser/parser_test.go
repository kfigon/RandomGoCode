package parser

import (
	"testing"
	"programming-lang/lexer"
)


func assertNoErrors(t *testing.T, errors []error) {
	if len(errors) > 0 {
		t.Errorf("Got %v errors, expected none", len(errors))
		for _, v := range errors {
			t.Error(v)
		}
	}
}

func assertStatementLength(t *testing.T, statements []StatementNode, expLen int) {
	if len(statements) != expLen {
		t.Fatalf("Invalid statement length, exp %v, got %v", expLen, len(statements))
	}
}

func assertErrorLength(t *testing.T, errors []error, expLen int) {
	if len(errors) != expLen {
		t.Fatalf("Expected %v errors, got %v: %v", expLen, len(errors), errors)
	}
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

func TestVarStatement_Identifiers(t *testing.T) {
	tokens := lexer.Tokenize(`var foo = 123;
	var asd = 3;`)
	
	tree := Parse(tokens)

	assertNoErrors(t, tree.Errors)
	assertStatementLength(t, tree.Statements, 2)

	if tree.Statements[0].TokenLiteral() != "foo" {
		t.Error("Invalid first statement", tree.Statements[0])
	}
	if tree.Statements[1].TokenLiteral() != "asd" {
		t.Error("Invalid second statement", tree.Statements[1])
	}

	assertVarStatementAndIntegerExpression(t, tree.Statements[0], 123)
	assertVarStatementAndIntegerExpression(t, tree.Statements[1], 3)	
}

func TestBasicReturnStatement(t *testing.T) {
	tokens := lexer.Tokenize(`return 123;`)
	
	tree := Parse(tokens)

	assertNoErrors(t, tree.Errors)

	assertStatementLength(t, tree.Statements, 1)
	if tree.Statements[0].TokenLiteral() != "return" {
		t.Error("Invalid literal found", tree.Statements[0])
	}
	
	ret,ok := tree.Statements[0].(*ReturnStatementNode)
	if !ok {
		t.Fatal("Return node note found")
	}

	integer,ok := ret.Value.(*IntegerLiteralExpression)
	if !ok {
		t.Fatal("Integer literal not found")
	}
	if integer.Value != 123 {
		t.Errorf("Invalid integer literal, got %v, exp %v", integer.Value, 123)
	}
}

func TestIdentifierExpression(t *testing.T) {
	tokens := lexer.Tokenize(`var foo = asd;`)
	
	tree := Parse(tokens)

	assertNoErrors(t, tree.Errors)
	assertStatementLength(t, tree.Statements,1)

	if tree.Statements[0].TokenLiteral() != "foo" {
		t.Error("Invalid literal found", tree.Statements[0])
	}
	
	ret,ok := tree.Statements[0].(*VarStatementNode)
	if !ok {
		t.Fatal("Return node note found")
	}

	identifier,ok := ret.Value.(*IdentifierExpression)
	if !ok {
		t.Fatal("Identifier expression not found")
	}
	if identifier.Name != "asd" {
		t.Errorf("Invalid identifier, got %v, exp %v", identifier.Name, "asd")
	}
}

func TestInvalidVarStatements(t *testing.T) {
	input := `var asd 4;
	var = 432;
	var x = foo`

	tree := Parse(lexer.Tokenize(input))

	assertErrorLength(t, tree.Errors, 3)
}

func TestInvalidVarStatementsWithExpressions(t *testing.T) {
	input := `var asd = 4
	var asd = ;`

	tree := Parse(lexer.Tokenize(input))

	assertErrorLength(t,tree.Errors,1)
}

func TestInvalidVarStatementsWithExpressions2(t *testing.T) {
	input := `var asd = ;`

	tree := Parse(lexer.Tokenize(input))

	assertErrorLength(t,tree.Errors,1)
}

func TestVarWithoutAssignment(t *testing.T) {
	tree := Parse(lexer.Tokenize(`var asd;`))
	assertNoErrors(t, tree.Errors)
}

func TestVarWithBinaryOperator(t *testing.T) {
	input := `var asd = 5 + 1;`

	tree := Parse(lexer.Tokenize(input))

	assertNoErrors(t, tree.Errors)
	assertStatementLength(t, tree.Statements, 1)
}
