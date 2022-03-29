package parser

import (
	"fmt"
	"programming-lang/lexer"
)

type StatementNode interface {
	Node
	evaluateStatement()
}


type VarStatementNode struct {
	Name  string
	Value ExpressionNode
}

func (vsn *VarStatementNode) TokenLiteral() string {
	return vsn.Name
}
func (vsn *VarStatementNode) evaluateStatement() {}

type ReturnStatementNode struct {
	Value ExpressionNode
}

func (r *ReturnStatementNode) TokenLiteral() string {
	return "return"
}
func (r *ReturnStatementNode) evaluateStatement() {}

// Statement wrapper for expressions, required for pratt parsing
type ExpressionStatementNode struct {
	Token lexer.Token //first token
	Value ExpressionNode
}

func (e *ExpressionStatementNode) TokenLiteral() string {
	return e.Token.Lexeme
}
func (e *ExpressionStatementNode) evaluateStatement() {}



func (p *parser) parseVarStatement() StatementNode {
	p.advanceToken()
	identifierTok := p.currentToken
	
	if p.eof() {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after var"))
		return nil
	} else if !isIdentifier(identifierTok) {
		p.addError(fmt.Errorf("var error - expected identifier, got %v", identifierTok.Class))
		return nil
	}

	p.advanceToken()
	if p.eof() {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after identifier"))
		return nil
	} else if isSemicolon(p.currentToken) {
		out := VarStatementNode{Name: identifierTok.Lexeme}
		p.advanceToken()
		return &out
	} else if !isAssignmentOperator(p.currentToken) {
		p.addError(fmt.Errorf("var error - expected assignment after identifier, got %v", p.currentToken.Class))
		return nil
	}

	p.advanceToken()
	exp := p.parseExpression(LOWEST)
	if exp == nil {
		p.advanceToken()
		return nil
	}
	return &VarStatementNode{Name: identifierTok.Lexeme, Value: exp}
}

func (p *parser) parseReturnStatement() StatementNode {
	p.advanceToken()

	exp := p.parseExpression(LOWEST)
	if exp == nil {
		p.advanceToken()
		return nil
	}
	return &ReturnStatementNode{exp}
}

func (p *parser) parseExpressionStatement() StatementNode {
	tok := p.currentToken

	exp := p.parseExpression(LOWEST)
	if exp == nil {
		p.advanceToken()
		return nil
	}

	return &ExpressionStatementNode{
		Token: tok,
		Value: exp,
	}
}