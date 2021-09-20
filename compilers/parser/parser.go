package parser

import (
	"programming-lang/lexer"
)

type Grammar int

const (
	Expression Grammar = iota
	BinaryOperator
	Statement
	Assignment
	If
	Loop
	// Function
)

func Parse(tokens []lexer.Token) *Tree {
	return nil
}

type Tree struct {}

type BinaryNode struct {
	Operation string
	Left *ExpressionNode
	Right *ExpressionNode
}

type ExpressionNode struct {
	Assign *AssignmentNode
	Op *BinaryNode
}

type StatementNode struct {
	Exp *ExpressionNode
}

type AssignmentNode struct {
	Identifier string
	Expression *ExpressionNode
}

type IfNode struct {
	Predicate *ExpressionNode
	Statement *StatementNode
}

type LoopNode struct {
	End *ExpressionNode
	Statement *StatementNode
}