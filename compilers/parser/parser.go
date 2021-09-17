package parser

import (
	"programming-lang/lexer"
)

type Grammar int

// * expression (right side of assignment)
// * binary operations (+,-,*,/)
// * statements - list of expressions
// * assignments
// * if-else
// * if 
// * predicate (boolean expression)
// * loop
// * function
const (
	Expression Grammar = iota  // <term> <operator>(+-*/) <term>
	Term
)

func Parse(tokens []lexer.Token) *Tree {
	return nil
}

type Tree struct {

}