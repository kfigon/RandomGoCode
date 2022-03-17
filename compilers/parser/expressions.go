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

type infixFn func(leftSide ExpressionNode) ExpressionNode
type prefixFn func() ExpressionNode

type mapPair[V any] struct {
	key lexer.TokenClass
	val V
}

func registerParsingFns[T any](fns ...mapPair[T]) map[lexer.TokenClass]T {
	out := map[lexer.TokenClass]T{}
	for _, v := range fns {
		out[v.key] = v.val
	}

	return out
}

func prefixFns() map[lexer.TokenClass]prefixFn {
	return registerParsingFns[prefixFn]()
}

func infixFns() map[lexer.TokenClass]infixFn {
	return registerParsingFns[infixFn]()
}

func (p *parser) parseExpression(predescense int) ExpressionNode {
	p.advanceToken()
	tok := p.currentToken
	switch {
	case p.eof() || isSemicolon(tok): {
		p.addError(fmt.Errorf("expression error - no expresion found, got %v", tok.Lexeme))
		return nil
	}
	case isNumberLiteral(tok): return p.parseIntegerLiteralExpression()
	case isIdentifier(tok): return p.parseIdentifierExpression()
	default: return nil
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