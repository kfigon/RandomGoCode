package parser

import (
	"fmt"
	"programming-lang/lexer"
	"strconv"
)

type Program struct {
	Statements []StatementNode
	Errors []error
}

func Parse(tokens []lexer.Token) *Program {
	p := &parser{
		tokens: tokens,
		prefixParsingFns: prefixFns(),
		infixParsingFns: infixFns(),
	}

	// populate current and next
	p.advanceToken()

	for !p.eof(){
		if isVarKeyword(p.currentToken) {
			p.addStatement(p.parseVarStatement())
		} else if isReturnKeyword(p.currentToken) {
			p.addStatement(p.parseReturnStatement())
		} else {
			p.addStatement(p.parseExpressionStatement())
		}
	}

	return &Program{p.statements, p.errors}
}

type parser struct {
	tokens []lexer.Token
	idx int
	currentToken lexer.Token
	nextToken lexer.Token

	prefixParsingFns map[lexer.TokenClass]prefixFn
	infixParsingFns map[lexer.TokenClass]infixFn

	errors []error
	statements []StatementNode
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

func (p *parser) advanceToken() {
	currentToken := func() lexer.Token {
		if p.idx >= len(p.tokens) {
			return lexer.Token{Class: lexer.EOF}
		}
		return p.tokens[p.idx]
	}

	p.currentToken = currentToken()
	p.idx++
	p.nextToken = currentToken()
}

func (p *parser) eof() bool {
	return eof(p.currentToken)
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

	exp := p.parseExpression()
	if exp == nil {
		return nil
	}
	return &VarStatementNode{Name: identifierTok.Lexeme, Value: exp}
}

func (p *parser) parseReturnStatement() StatementNode {
	exp := p.parseExpression()
	if exp == nil {
		return nil
	}
	return &ReturnStatementNode{exp}
}

func (p *parser) parseExpressionStatement() StatementNode {
	p.advanceToken()
	return nil
}