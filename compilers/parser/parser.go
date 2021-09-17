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
	ElseIf
	Else
	Predicate
	Loop
	Function
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

}