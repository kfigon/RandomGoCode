package parser

import (
	"programming-lang/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestVarStatement_Identifier(t *testing.T) {
	tree := parse(`var foo = 123;`)

	assertNoErrors(t, tree.Errors)
	assert.Len(t, tree.Statements, 1)

	assert.Equal(t, "foo", tree.Statements[0].TokenLiteral(), "invalid first literal")

	assertVarStatementAndIntegerExpression(t, tree.Statements[0], 123)
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

func TestInvalidVarStatementsWithExpressions(t *testing.T) {
	t.Run("missing expression after assignment", func(t *testing.T) {
		tree := parse(`var asd = ;`)
		assert.Len(t, tree.Errors, 1)
	})

	t.Run("var return", func(t *testing.T) {
		tree := parse(`var return 123;`)
		assert.Len(t, tree.Errors, 1)
	})

	t.Run("unexpected eof", func(t *testing.T) {
		tree := parse(`var foo = `)
		assert.Len(t, tree.Errors, 1)
	})
}

func TestInvalidVarStatements(t *testing.T) {
	input := `var asd 4;
	var = 432;
	var x = foo`

	tree := parse(input)
	assert.Len(t, tree.Errors, 4)
}

func TestFirstVarNotTerminated_SecondExpressionles(t *testing.T) {
	input := `var asd = 4
	var asd = ;`

	tree := parse(input)
	assert.Len(t, tree.Errors, 2)
}

func TestVarWithoutAssignment(t *testing.T) {
	tree := parse(`var asd;`)
	assertNoErrors(t, tree.Errors)
	assert.Len(t, tree.Statements, 1)

	st,ok := tree.Statements[0].(*VarStatementNode)
	require.True(t, ok)
	assert.Equal(t, "asd", st.Name)
	assert.Nil(t, st.Value)
}

func TestExpressionStatement(t *testing.T) {
	tree := parse(`foobar;`)
	assertNoErrors(t, tree.Errors)
	assert.Len(t, tree.Statements, 1)

	exp,ok := tree.Statements[0].(*ExpressionStatementNode)
	require.True(t, ok, "ExpressionStatementNode not found")

	identifier, ok := exp.Value.(*IdentifierExpression)
	require.True(t, ok)
	
	assert.Equal(t, "foobar", identifier.Name)
	assert.Equal(t, "foobar", identifier.TokenLiteral())
}

func TestPrefixExpression(t *testing.T) {
	tree := parse(`-5;`)
	assertNoErrors(t, tree.Errors)
	assert.Len(t, tree.Statements, 1)

	exp,ok := tree.Statements[0].(*ExpressionStatementNode)
	require.True(t, ok, "ExpressionStatementNode not found")

	prefix, ok := exp.Value.(*PrefixExpression)
	require.True(t, ok, "prefix expression expected")
	
	assert.Equal(t, "-", prefix.Operator)
	
	val, ok :=  prefix.Right.(*IntegerLiteralExpression)
	require.True(t, ok, "integer expression expected")
	assert.Equal(t, 5, val.Value)
}

func TestInfixExpression(t *testing.T) {
	t.Skip("todo")
	
	testCases := []struct {
		input	string
		expectedOperator string
		left int
		right int
	}{
		{"1+2", "+", 1, 2},
		{"1 + 2", "+", 1, 2},
		{"1*2", "*", 1, 2},
		{"1 * 2", "*", 1, 2},
		{"4/2", "/",4, 2},
		{"3 == 4", "==",3, 4},
		{"3 != 4", "!=",3, 4},
		{"3 > 4", ">",3, 4},
		{"3 >= 4", ">=",3, 4},
		{"3 < 4", "<",3, 4},
		{"3 <= 4", "<=",3, 4},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			tree := parse(tC.input)
			assertNoErrors(t, tree.Errors)
			require.Len(t, tree.Statements, 1, "statement not found")
		
			exp,ok := tree.Statements[0].(*ExpressionStatementNode)
			require.True(t, ok, "ExpressionStatementNode not found")
		
			infix, ok := exp.Value.(*InfixExpressionNode)
			require.True(t, ok, "infix expression expected")
			
			assert.Equal(t, tC.expectedOperator, infix.Operator)
			
			left, ok :=  infix.Left.(*IntegerLiteralExpression)
			require.True(t, ok, "integer expression expected for left")
			assert.Equal(t, tC.left, left.Value)

			right, ok :=  infix.Right.(*IntegerLiteralExpression)
			require.True(t, ok, "integer expression expected for right")
			assert.Equal(t, tC.right, right.Value)
		})
	}
}