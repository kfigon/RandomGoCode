package main

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
