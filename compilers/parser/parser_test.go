package parser

import (
	"programming-lang/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNoErrors(t *testing.T, errors []error) {
	assert.Len(t, errors, 0, "no errors were expected")
}

func parse(input string) *Program {
	return Parse(lexer.Tokenize(input))
}

func assertVarStatementAndIntegerExpression(t *testing.T, st StatementNode, exp int) {
	varSt, ok := st.(*VarStatementNode)
	assert.True(t, ok, "expected var statement")

	integer, ok := varSt.Value.(*IntegerLiteralExpression)
	assert.True(t, ok, "expected integer literal")
	assert.Equal(t, exp, integer.Value)
}

func TestVarStatement_Identifiers(t *testing.T) {
	tree := parse(`var foo = 123;
	var asd = 3;`)

	assertNoErrors(t, tree.Errors)
	assert.Len(t, tree.Statements, 2)

	assert.Equal(t, "foo", tree.Statements[0].TokenLiteral(), "invalid first literal")
	assert.Equal(t, "asd", tree.Statements[1].TokenLiteral(), "invalid second literal")

	assertVarStatementAndIntegerExpression(t, tree.Statements[0], 123)
	assertVarStatementAndIntegerExpression(t, tree.Statements[1], 3)	
}

func TestBasicReturnStatement(t *testing.T) {
	tree := parse(`return 123;`)

	assertNoErrors(t, tree.Errors)

	assert.Len(t, tree.Statements, 1)
	assert.Equal(t, "return", tree.Statements[0].TokenLiteral())
	
	ret,ok := tree.Statements[0].(*ReturnStatementNode)
	assert.True(t, ok, "return node not found")

	integer,ok := ret.Value.(*IntegerLiteralExpression)
	assert.True(t, ok, "integer literal not found")
	assert.Equal(t, 123, integer.Value)
}

func TestIdentifierExpression(t *testing.T) {
	tree := parse(`var foo = asd;`)

	assertNoErrors(t, tree.Errors)
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

	tree := parse(input)

	assert.Len(t, tree.Errors, 3)
}

func TestInvalidVarStatementsWithExpressions(t *testing.T) {
	input := `var asd = 4
	var asd = ;`

	tree := parse(input)
	assert.Len(t, tree.Errors, 1)
}

func TestInvalidVarStatementsWithExpressions2(t *testing.T) {
	tree := parse(`var asd = ;`)
	assert.Len(t,tree.Errors,1)
}

func TestVarWithoutAssignment(t *testing.T) {
	tree := parse(`var asd;`)
	assertNoErrors(t, tree.Errors)
	assert.Len(t, tree.Statements, 1)

	st,ok := tree.Statements[0].(*VarStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "asd", st.Name)
	assert.Nil(t, st.Value)
}
