package parser

import (
	"programming-lang/lexer"
)

type Grammar int

const (
	Expression Grammar = iota
	Statement
)

func Parse(tokens []lexer.Token) *Tree {
	return nil
}

type Tree struct {
	Statements []LetStatementNode
}

// should be generalized to a Statement
type LetStatementNode struct {
	Name *string
	Value *ExpressionNode
}

type ExpressionNode struct {

}