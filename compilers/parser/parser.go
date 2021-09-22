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

type Tree struct {}