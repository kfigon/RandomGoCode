package parser

import (
	"fmt"
	"programming-lang/lexer"
)

type parser struct {
	iter *tokenIterator
	errors []error
	statements []StatementNode
}

func Parse(tokens []lexer.Token) *Program {
	p := &parser{iter: newIterator(tokens)}
	for {
		tok, ok := p.iter.next()
		if !ok {
			break
		} else if isVarKeyword(tok) {
			p.parseVarStatement()
		} else if isReturnKeyword(tok) {
			p.parseReturnStatement()
		}
	}

	return &Program{p.statements, p.errors}
}

type Program struct {
	Statements []StatementNode
	Errors []error
}

// Node is an interface mostly for debugging and testing
type Node interface {
	TokenLiteral() string
}

type StatementNode interface {
	Node
	evaluateStatement()
}

type ExpressionNode interface {
	Node
	evaluateExpression()
}

func (p *parser) addError(err error) {
	if err != nil {
		p.errors = append(p.errors, err)
	}
}

func (p *parser) addStatement(st StatementNode) {
	if st != nil {
		p.statements = append(p.statements, st)
	}
}

func (p *parser) parseExpression() ExpressionNode {
	for {
		tok, ok := p.iter.next()
		if !ok || isSemicolon(tok){
			break
		}
	}
	return nil
}


func (p *parser) parseVarStatement() {
	tok, ok := p.iter.next()
	out := VarStatementNode{}
	if !ok {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after var"))
		return
	} else if !isIdentifier(tok) {
		p.addError(fmt.Errorf("var error - expected identifier, got %v", tok.Class))
		return
	}

	out.Name = tok.Lexeme

	tok, ok = p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after identifier"))
		return
	} else if !isAssignmentOperator(tok) && !isSemicolon(tok) {
		p.addError(fmt.Errorf("var error - expected assignment or semicolon after identifier, got %v", tok.Class))
		return
	}

	// todo
	out.Value = p.parseExpression()
	p.addStatement(&out)
}

func (p *parser) parseReturnStatement() {
	// todo
	exp := p.parseExpression()
	p.addStatement(&ReturnStatementNode{exp})
}