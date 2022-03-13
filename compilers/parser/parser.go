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
	p := &parser{tokens: tokens}
	// populate current and next
	p.advanceToken()

	for !p.eof(){
		if isVarKeyword(p.currentToken) {
			p.parseVarStatement()
		} else if isReturnKeyword(p.currentToken) {
			p.parseReturnStatement()
		} else {
			p.advanceToken()
		}
	}

	return &Program{p.statements, p.errors}
}

type parser struct {
	tokens []lexer.Token
	idx int
	currentToken lexer.Token
	nextToken lexer.Token

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
	nextToken := func() lexer.Token {
		if p.idx >= len(p.tokens) {
			return lexer.Token{Class: lexer.EOF}
		}
		toRet := p.tokens[p.idx]
		p.idx++
		return toRet
	}
	if p.idx >= len(p.tokens) {
		p.currentToken = nextToken()
		p.nextToken = nextToken()
		return
	}

	p.currentToken = p.tokens[p.idx]
	p.nextToken = nextToken()
}

func (p *parser) eof() bool {
	return p.currentToken.Class == lexer.EOF
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
	case p.eof(): return nil
	case isSemicolon(tok): {
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
	token := p.currentToken
	out := &IdentifierExpression{Name: token.Lexeme}
	p.advanceToken()
	if p.eof() {
		p.addError(fmt.Errorf("identifier expression error - unexpected end of tokens"))
		return nil
	} else if !isSemicolon(p.currentToken) {
		p.addError(fmt.Errorf("identifier expression error - semicolon not found, got %v", p.currentToken))
		return nil
	}
	p.advanceToken()
	return out
}

func (p *parser) parseVarStatement() {
	p.advanceToken()
	identifierTok := p.currentToken
	if p.eof() {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after var"))
		return
	} else if !isIdentifier(identifierTok) {
		p.addError(fmt.Errorf("var error - expected identifier, got %v", identifierTok.Class))
		return
	}

	p.advanceToken()
	if p.eof() {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after identifier"))
		return
	} else if isSemicolon(p.currentToken) {
		out := VarStatementNode{Name: identifierTok.Lexeme}
		p.addStatement(&out)
		return
	} else if !isAssignmentOperator(p.currentToken) {
		p.addError(fmt.Errorf("var error - expected assignment after identifier, got %v", p.currentToken.Class))
		return
	}

	out := VarStatementNode{Name: identifierTok.Lexeme}
	exp := p.parseExpression()
	if exp == nil {
		return
	}
	out.Value = exp
	p.addStatement(&out)
}

func (p *parser) parseReturnStatement() {
	exp := p.parseExpression()
	if exp == nil {
		return
	}
	p.addStatement(&ReturnStatementNode{exp})
}