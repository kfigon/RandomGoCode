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

func (vsn *VarStatementNode) String() string {
	str := "var " + vsn.Name
	if vsn.Value != nil {
		str += vsn.String()
	}
	return str
}

func (vsn *VarStatementNode) evaluateStatement() {}

type ReturnStatementNode struct {
	Value ExpressionNode
}

func (r *ReturnStatementNode) TokenLiteral() string {
	return "return"
}
func (r *ReturnStatementNode) String() string {
	str := "return"
	if r.Value != nil {
		str += r.String()
	}
	return str
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

func (e *ExpressionStatementNode) String() string {
	return e.Value.String()
}

func (e *ExpressionStatementNode) evaluateStatement() {}



func (p *parser) parseVarStatement() StatementNode {	
	if !isIdentifier(p.nextToken) {
		p.addError(fmt.Errorf("var error - expected identifier, got %v", p.nextToken.Class))
		return nil
	}
	p.advanceToken()
	identifierTok := p.currentToken

	if isSemicolon(p.nextToken) {
		p.advanceToken()
		return &VarStatementNode{Name: identifierTok.Lexeme}
	}

	if !isAssignmentOperator(p.nextToken) {
		p.addError(fmt.Errorf("var error - expected assignment after identifier, got %v", p.nextToken.Class))
		return nil
	}

	p.advanceToken() // assignment
	p.advanceToken() // expression

	exp := p.parseExpression(LOWEST)
	out := &VarStatementNode{Name: identifierTok.Lexeme, Value: exp}
	if !isSemicolon(p.nextToken) {
		p.addError(fmt.Errorf("var error - expected semicolon after expression, got %v", p.nextToken.Class))
		return nil
	}
	p.advanceToken()
	return out
}

func (p *parser) parseReturnStatement() StatementNode {
	p.advanceToken()

	exp := p.parseExpression(LOWEST)
	out := &ReturnStatementNode{exp}
	if !isSemicolon(p.nextToken) {
		p.addError(fmt.Errorf("var error - expected semicolon after expression, got %v", p.nextToken.Class))
		return nil
	}
	p.advanceToken()
	return out
}

func (p *parser) parseExpressionStatement() StatementNode {
	tok := p.currentToken

	exp := p.parseExpression(LOWEST)

	if !isSemicolon(p.nextToken) {
		p.addError(fmt.Errorf("var error - expected semicolon after expression, got %v", p.nextToken.Class))
		return nil
	}
	p.advanceToken()
	return &ExpressionStatementNode{
		Token: tok,
		Value: exp,
	}
}