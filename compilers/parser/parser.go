package parser

import (
	"programming-lang/lexer"
)

type Grammar int

const (
	Expression Grammar = iota 
	Term
)

func Parse(tokens []lexer.Token) *Tree {
	return nil
}

type Tree struct {

}