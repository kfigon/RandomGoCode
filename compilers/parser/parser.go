package parser

import (
	"fmt"
	"programming-lang/lexer"
)

type Grammar int

const (
	Expression Grammar = iota
	Statement
)

type pars struct {
	iter *tokenIterator
}

func Parse(tokens []lexer.Token) *Tree {
	p := &pars{newIterator(tokens)}
	tree := &Tree{}
	for {
		tok, ok := p.iter.next()
		if !ok {
			break
		}
		if isVarKeyword(tok) {
			st, err := p.parseVarExpression()
			if err != nil {
				tree.Errors = append(tree.Errors, fmt.Errorf("error in var statement: %v", err))
				continue
			}
			tree.Statements = append(tree.Statements, st)
		}
	}

	return tree
}

type Tree struct {
	Statements []LetStatementNode
	Errors []error
}

// should be generalized to a Statement
type LetStatementNode struct {
	Name string
	Value *ExpressionNode
}

type ExpressionNode struct {

}


func (p *pars) parseVarExpression() (LetStatementNode, error) {
	tok, ok := p.iter.next()
	out := LetStatementNode{}
	if !ok {
		return out, fmt.Errorf("unexpected end of tokens")
	} else if !isIdentifier(tok) {
		return out, fmt.Errorf("expected identifier")
	}

	out.Name = tok.Lexeme

	tok, ok = p.iter.next()
	if !ok {
		return out, fmt.Errorf("unexpected end of tokens")
	} else if !isAssignmentOperator(tok) {
		return out, fmt.Errorf("expected assignment after var")
	}

	// todo: parse expression
	for {
		tok, ok = p.iter.next()
		if !ok || isSemicolon(tok){
			break
		}
	}
	return out,nil
}

func isVarKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "var"
}

func isAssignmentOperator(token lexer.Token) bool {
	return token.Class == lexer.Assignment && token.Lexeme == "="
}

func isSemicolon(token lexer.Token) bool {
	return token.Class == lexer.Semicolon && token.Lexeme == ";"
}

func isNumberLiteral(token lexer.Token) bool {
	return token.Class == lexer.Number
}

func isIdentifier(token lexer.Token) bool {
	return token.Class == lexer.Identifier
}