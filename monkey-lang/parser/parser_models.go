package parser
import (
	"monkey-lang/lexer"
)

type Precedence int
const (
	Lowest Precedence = iota+1
	Equals
	LessGreater
	Sum
	Product
	Prefix
	Call
)

type Statement interface {
	statementTag()
}

type Expression interface {
	expressionTag()
}

// ------------------- statements ------------------
type LetStatement struct {
	Ident *IdentifierExpression
	Value Expression
}

func (*LetStatement) statementTag() {}

type ReturnStatement struct {
	Exp Expression
}

func (*ReturnStatement) statementTag() {}


type ExpressionStatement struct {
	Exp Expression
}

func (*ExpressionStatement) statementTag(){}

type BlockStatement struct {
	Stmts []Statement
}

func (*BlockStatement) statementTag(){}

// ------------------- expression ------------------
// numbers, booleans, strings
type PrimitiveLiteral[T any] struct {
	Val T
}
func (*PrimitiveLiteral[T]) expressionTag(){}

type IdentifierExpression struct {
	Name string
}

func (*IdentifierExpression) expressionTag() {}

type PrefixExpression struct {
	Operator lexer.Token
	Expr Expression
}

func (*PrefixExpression) expressionTag(){}

type InfixExpression struct {
	Operator lexer.Token
	Left Expression
	Right Expression
}

func (*InfixExpression) expressionTag(){}

type IfExpression struct {
	// if <predicate> block
	// optional else if <predicate> block
	// optional else block   <- here predicate will be nil

	Predicate Expression
	Consequence *BlockStatement
	Alternative *IfExpression
}

func (*IfExpression) expressionTag(){}

type FunctionLiteral struct {
	Parameters []*IdentifierExpression
	Body *BlockStatement
}

func (*FunctionLiteral) expressionTag(){}

type FunctionCall struct {
	Func Expression // Identifier or function literal
	Arguments []Expression
}

func(*FunctionCall) expressionTag(){}