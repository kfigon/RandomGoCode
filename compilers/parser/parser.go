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
		p.addError(fmt.Errorf("unexpected end of tokens"))
		return
	} else if !isIdentifier(tok) {
		p.addError(fmt.Errorf("expected identifier"))
		return
	}

	out.Name = tok.Lexeme

	tok, ok = p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("unexpected end of tokens"))
		return
	} else if !isAssignmentOperator(tok) {
		p.addError(fmt.Errorf("expected assignment after var"))
		return
	}

	// todo
	exp := p.parseExpression()
	out.Value = exp
	p.addStatement(&out)
}

func (p *parser) parseReturnStatement() {
	// todo
	exp := p.parseExpression()
	p.addStatement(&ReturnStatementNode{exp})
}