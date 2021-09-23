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
	iter := newIterator(tokens)
	tree := &Tree{}
	for {
		tok, ok := iter.next()
		if !ok {
			break
		}
		if isVarKeyword(tok) {
		}
	}

	return tree
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

func isVarKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "var"
}

func isEqualOperator(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "="
}

func isSemicolon(token lexer.Token) bool {
	return token.Class == lexer.Semicolon && token.Lexeme == ";"
}

func isNumberLiteral(token lexer.Token) bool {
	return token.Class == lexer.Number
}