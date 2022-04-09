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
	Right    ExpressionNode
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Operator
}
func (p *PrefixExpression) evaluateExpression() {}

type InfixExpressionNode struct {
	Operator string
	Left     ExpressionNode
	Right    ExpressionNode
}

func (i *InfixExpressionNode) TokenLiteral() string {
	return i.Operator
}
func (i *InfixExpressionNode) evaluateExpression() {}

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

func (p *parser) parseExpression(predescense int) ExpressionNode {
	tok := p.currentToken

	var left ExpressionNode
	if bang(tok) || minus(tok) {
		 left = p.parsePrefixExpression()
	} else if isNumberLiteral(tok){
		left = p.parseIntegerLiteralExpression()
	} else if isIdentifier(tok) {
		left = p.parseIdentifierExpression()
	}


	// infix - left associated
	for !isSemicolon(p.nextToken) && predescense < tokensPredescense(p.nextToken) {

		if equals(p.nextToken) || 
			notEquals(p.nextToken) || 
			lessThan(p.nextToken) || 
			lessEqThan(p.nextToken) || 
			greaterEqThan(p.nextToken) || 
			greaterThan(p.nextToken) || 
			plus(p.nextToken) || 
			minus(p.nextToken) || 
			product(p.nextToken) || 
			divide(p.nextToken) {
			
			p.advanceToken()
			left = p.parseInfixExpression(left)
		} else {
			return left
		}
	}

	return left
}

func tokensPredescense(tok lexer.Token) int {
	switch {
	case equals(tok):
		return EQUALS
	case notEquals(tok):
		return EQUALS

	case lessThan(tok):
		return LESSGREATER
	case lessEqThan(tok):
		return LESSGREATER
	case greaterEqThan(tok):
		return LESSGREATER
	case greaterThan(tok):
		return LESSGREATER

	case plus(tok):
		return SUM
	case minus(tok):
		return SUM

	case product(tok):
		return PRODUCT
	case divide(tok):
		return PRODUCT
	default:
		return LOWEST
	}
}

func (p *parser) parseIntegerLiteralExpression() ExpressionNode {
	tok := p.currentToken
	v, err := strconv.Atoi(tok.Lexeme)
	if err != nil {
		p.addError(fmt.Errorf("int literal expression error - error in parsing integer literal in: %v", tok.Lexeme))
		return nil
	}
	return &IntegerLiteralExpression{Value: v}
}

func (p *parser) parseIdentifierExpression() ExpressionNode {
	identifierToken := p.currentToken
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
		Left:     left,
	}
	pred := tokensPredescense(p.currentToken)
	p.advanceToken()
	out.Right = p.parseExpression(pred)
	return out
}
