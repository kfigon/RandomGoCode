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
		} else if isVarKeyword(tok) {
			st, err := p.parseVarExpression()
			tree.addResult(st,err, "var")
		} else if isReturnKeyword(tok) {
			st, err := p.parseReturnExpression()
			tree.addResult(st,err, "return")
		}
	}

	return tree
}

func (t *Tree) addResult(st LetStatementNode, err error, statementType string) {
	if err != nil {
		t.Errors = append(t.Errors, fmt.Errorf("error in %v statement: %v", statementType, err))
		return
	}
	t.Statements = append(t.Statements, st)
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

type IntegerLiteralExpression struct {
	Value int
}

func (p *pars) parseExpression() (*ExpressionNode, error) {
	for {
		tok, ok := p.iter.next()
		if !ok || isSemicolon(tok){
			break
		}
	}
	return nil,nil
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

	// todo
	_, err := p.parseExpression()
	if err != nil {
		return out, fmt.Errorf("error in parsing expression %v", err)
	}
	return out,nil
}

func (p *pars) parseReturnExpression() (LetStatementNode, error) {
	// todo
	p.parseExpression()
	return LetStatementNode{},nil
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

func isReturnKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "return"
}