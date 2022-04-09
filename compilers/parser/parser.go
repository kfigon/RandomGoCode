package parser

import (
	"programming-lang/lexer"
)

type Program struct {
	Statements []StatementNode
	Errors []error
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

func Parse(tokens []lexer.Token) *Program {
	p := &parser{tokens: tokens}

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
		
		p.advanceToken()
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
	String() string
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
