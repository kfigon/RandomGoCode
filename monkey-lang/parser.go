package main

import (
	"fmt"
	"iter"
)

type parser struct {
	nextFn func()(Token, bool)
	current Token
	peek Token
}

func Parse(toks iter.Seq[Token]) ([]Statement, error) {
	nextFn, stop := iter.Pull(toks)
	defer stop()

	p := &parser{
		nextFn: nextFn,
	}
	// populate cur and peek
	p.consume()
	p.consume()

	return p.parse()
}

func (p *parser) consume() {
	newEl, ok := p.nextFn()
	if !ok {
		newEl = Token{EOF, ""}
	}
	p.current = p.peek
	p.peek = newEl
}

func (p *parser) eof() bool {
	return p.current.Typ == EOF
}

func (p *parser) parse() ([]Statement, error)  {
	out := []Statement{}
	for !p.eof() {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		out = append(out, stmt)
	}
	return out, nil
}

func (p *parser) parseStatement() (Statement, error) {
	t := p.current
	if t.Typ == Let {
		if !p.expectPeek(Identifier){
			return nil, expectedTokenErr(Identifier, p.current.Typ) 
		}

		identifier := p.current.Lexeme
		if !p.expectPeek(Assign) {
			return nil, expectedTokenErr(Assign, p.current.Typ) 
		}

		exp, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		return &LetStatement{
			&IdentifierExpression{identifier},
			exp,
		}, nil
	} else if t.Typ == Return {
		p.consume()	// return
		exp, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		return &ReturnStatement{exp},nil
	}

	return nil, fmt.Errorf("unknown token %v", t)
}

func (p *parser) parseExpression() (Expression, error) {
	for !p.currentIs(Semicolon) {
		// todo
		p.consume()
	}
	p.consume()// semicolon
	return nil, nil
}

func (p *parser) currentIs(t TokenType) bool {
	return p.current.Typ == t
}

func (p *parser) peekIs(t TokenType) bool {
	return p.peek.Typ == t
}

func (p *parser) expectPeek(t TokenType) bool {
	if p.peekIs(t) {
		p.consume()
		return true
	}
	return false
}

func expectedTokenErr(exp, got TokenType) error {
	return fmt.Errorf("expected %v, got %v", exp, got)
}