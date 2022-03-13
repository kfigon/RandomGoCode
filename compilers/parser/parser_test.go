package parser

import (
	"programming-lang/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertVarStatementAndIntegerExpression(t *testing.T, st StatementNode, exp int) {
	varSt, ok := st.(*VarStatementNode)
	assert.True(t, ok, "expected var statement")

	integer, ok := varSt.Value.(*IntegerLiteralExpression)
	assert.True(t, ok, "expected integer literal")
	assert.Equal(t, exp, integer.Value)
}

func TestVarStatement_Identifiers(t *testing.T) {
	tokens := lexer.Tokenize(`var foo = 123;
	var asd = 3;`)
	
	tree := Parse(tokens)

	assert.Nil(t, tree.Errors)
	assert.Len(t, tree.Statements, 2)

	assert.Equal(t, "foo", tree.Statements[0].TokenLiteral(), "invalid first literal")
	assert.Equal(t, "asd", tree.Statements[0].TokenLiteral(), "invalid second literal")

	assertVarStatementAndIntegerExpression(t, tree.Statements[0], 123)
	assertVarStatementAndIntegerExpression(t, tree.Statements[1], 3)	
}

func TestBasicReturnStatement(t *testing.T) {
	tokens := lexer.Tokenize(`return 123;`)
	
	tree := Parse(tokens)

	assert.Nil(t, tree.Errors)

	assert.Len(t, tree.Statements, 1)
	assert.Equal(t, "return", tree.Statements[0].TokenLiteral())
	
	ret,ok := tree.Statements[0].(*ReturnStatementNode)
	assert.True(t, ok, "return node not found")

	integer,ok := ret.Value.(*IntegerLiteralExpression)
	assert.True(t, ok, "integer literal not found")
	assert.Equal(t, 123, integer.Value)
}

func TestIdentifierExpression(t *testing.T) {
	tokens := lexer.Tokenize(`var foo = asd;`)
	
	tree := Parse(tokens)

	assert.Nil(t, tree.Errors)
	assert.Len(t, tree.Statements, 1)
	assert.Equal(t, "foo", tree.Statements[0].TokenLiteral())
	
	ret,ok := tree.Statements[0].(*VarStatementNode)
	assert.True(t, ok)

	identifier,ok := ret.Value.(*IdentifierExpression)
	assert.True(t, ok)

	assert.Equal(t, "asd", identifier.Name)
}

func TestInvalidVarStatements(t *testing.T) {
	input := `var asd 4;
	var = 432;
	var x = foo`

	tree := Parse(lexer.Tokenize(input))

	assert.Len(t, tree.Errors, 3)
}

func TestInvalidVarStatementsWithExpressions(t *testing.T) {
	input := `var asd = 4
	var asd = ;`

	tree := Parse(lexer.Tokenize(input))
	assert.Len(t, tree.Errors, 1)
}

func TestInvalidVarStatementsWithExpressions2(t *testing.T) {
	input := `var asd = ;`

	tree := Parse(lexer.Tokenize(input))
	assert.Len(t,tree.Errors,1)
}

func TestVarWithoutAssignment(t *testing.T) {
	tree := Parse(lexer.Tokenize(`var asd;`))
	assert.Nil(t, tree.Errors)
}

func TestVarWithBinaryOperator(t *testing.T) {
	input := `var asd = 5 + 1;`

	tree := Parse(lexer.Tokenize(input))

	assert.Nil(t, tree.Errors)	
	assert.Len(t, tree.Statements, 1)
}
