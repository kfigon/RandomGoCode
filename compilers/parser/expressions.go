package parser

import (
	"fmt"
	"programming-lang/lexer"
	"strconv"
)

type ExpressionNode interface {
	Node
	evaluateExpression()
}

type IntegerLiteralExpression struct {
	Value int
}
func (ile *IntegerLiteralExpression) TokenLiteral() string {
	return strconv.Itoa(ile.Value)
}
func (ile *IntegerLiteralExpression) evaluateExpression() {}


type IdentifierExpression struct {
	Name string
}
func (ide *IdentifierExpression) TokenLiteral() string {
	return ide.Name
}
func (ide *IdentifierExpression) evaluateExpression() {}

type PrefixExpression struct {
	Operator string
	Right ExpressionNode
}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Operator
}
func (p *PrefixExpression) evaluateExpression() {}

type InfixExpressionNode struct {
	Operator string
	Left ExpressionNode
	Right ExpressionNode
}

func (i *InfixExpressionNode) TokenLiteral() string {
	return i.Operator
}
func (i *InfixExpressionNode) evaluateExpression() {}

const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT // *
	PREFIX // -X or !X
	CALL // myFunction(X)
)

func (p *parser) parseExpression(predescense int) ExpressionNode {
	tok := p.currentToken
	switch {
	case p.eof() || isSemicolon(tok): {
		p.addError(fmt.Errorf("expression error - no expresion found, got %q", tok.Lexeme))
		return nil
	}
	case isNumberLiteral(tok): return p.parseIntegerLiteralExpression()
	case isIdentifier(tok): return p.parseIdentifierExpression()
	case bang(tok) || minus(tok): return p.parsePrefixExpression()
	}

	p.addError(fmt.Errorf("expression error - syntax error %q", tok.Lexeme))
	return nil
}

func predescense(tok lexer.Token) int {
	switch {
	case equals(tok): return EQUALS
	case notEquals(tok): return EQUALS
	
	case lessThan(tok): return LESSGREATER
	case lessEqThan(tok): return LESSGREATER
	case greaterEqThan(tok): return LESSGREATER
	case greaterThan(tok): return LESSGREATER

	case plus(tok): return SUM
	case minus(tok): return SUM
	
	case product(tok): return PRODUCT
	case divide(tok): return PRODUCT
	default: return LOWEST
	}
}

func (p *parser) parseIntegerLiteralExpression() ExpressionNode {
	tok := p.currentToken
	v, err := strconv.Atoi(tok.Lexeme)
	if err != nil {
		p.addError(fmt.Errorf("int literal expression error - error in parsing integer literal in: %v", tok.Lexeme))
		return nil
	}
	p.advanceToken()
	if p.eof() {
		p.addError(fmt.Errorf("int literal expression error - unexpected end of tokens"))
		return nil
	} else if !isSemicolon(p.currentToken) {
		p.addError(fmt.Errorf("int literal expression error - expected semicolon, got %v", p.currentToken.Lexeme))
		return nil
	}
	p.advanceToken()
	return &IntegerLiteralExpression{Value: v}		
}

func (p *parser) parseIdentifierExpression() ExpressionNode {
	identifierToken := p.currentToken
	p.advanceToken()
	if p.eof() {
		p.addError(fmt.Errorf("identifier expression error - unexpected end of tokens"))
		return nil
	} else if !isSemicolon(p.currentToken) {
		p.addError(fmt.Errorf("identifier expression error - semicolon not found, got %v", p.currentToken))
		return nil
	}
	p.advanceToken()
	return &IdentifierExpression{Name: identifierToken.Lexeme}
}

func (p *parser) parsePrefixExpression() ExpressionNode {
	operator := p.currentToken
	p.advanceToken()
	return &PrefixExpression{Operator: operator.Lexeme, Right: p.parseExpression(PREFIX)}
}

func (p *parser) parseInfixExpression(left ExpressionNode) ExpressionNode {
	out := &InfixExpressionNode{
		Operator: p.currentToken.Lexeme,
		Left: left,
	}
	pred := predescense(p.currentToken)
	p.advanceToken()
	out.Right = p.parseExpression(pred)
	return out
}