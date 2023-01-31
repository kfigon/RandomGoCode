package main

import "fmt"

// todo: implement a visitor maybe?
type expression interface {
	expr() // dummy method - lack of sum types
}

type literal token
func (literal) expr(){}

type grouping struct {
	expression
}
func (g grouping) expr(){}

type unary struct {
	op token
	ex expression
}
func (unary) expr(){}

type binary struct {
	op token
	left expression
	right expression
}
func (binary) expr(){}

type operatorExpr token
func (operatorExpr) expr(){}


type Parser struct {
	it *iter[token]
	Errors []error
	Expressions []expression
}

func NewParser(toks []token) *Parser {
	return &Parser{it: toIter(toks)}
}

func (p *Parser) Parse() ([]expression, []error) {
	v := p.parseExpression()
	if v == nil {
		p.Errors = append(p.Errors, fmt.Errorf("some error"))
	} else {
		p.Expressions = append(p.Expressions, v)
	}
	return p.Expressions, p.Errors
}

func (p *Parser) parseExpression() expression {
	return p.parseEquality()
}

func (p *Parser) parseEquality() expression {
	ex := p.parseComparison()
	for  {
		current, ok := p.it.current()
		if ok && (checkToken(current, operator, "!=") || checkToken(current, operator, "==")) {
			p.it.consume()
			ex = binary{op: current, left: ex, right: p.parseComparison()}
		} else {
			break
		}
	}
	return ex
}

func (p *Parser) parseComparison() expression {
	ex := p.parseTerm()
	for  {
		current, ok := p.it.current()
		if ok && (checkToken(current, operator, ">") || 
			checkToken(current, operator, ">=") || 
			checkToken(current, operator, "<") ||
			checkToken(current, operator, "<=")) {
			p.it.consume()
			ex = binary{op: current, left: ex, right: p.parseTerm()}
		} else {
			break
		}
	}
	return ex
}

func (p *Parser) parseTerm() expression {
	ex := p.parseFactor()
	for  {
		current, ok := p.it.current()
		if ok && (checkToken(current, operator, "-") || checkToken(current, operator, "+")) {
			p.it.consume()
			ex = binary{op: current, left: ex, right: p.parseFactor()}
		} else {
			break
		}
	}
	return ex
}

func (p *Parser) parseFactor() expression {
	ex := p.parseUnary()
	for  {
		current, ok := p.it.current()
		if ok && (checkToken(current, operator, "/") || checkToken(current, operator, "*")) {
			p.it.consume()
			ex = binary{op: current, left: ex, right: p.parseUnary()}
		} else {
			break
		}
	}
	return ex
}

func (p *Parser) parseUnary() expression {
	current, ok := p.it.current()
	if ok && (checkToken(current, operator, "!") || checkToken(current, operator, "-")) {
		op := current
		return unary{op: op, ex: p.parseUnary()}
	} 
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() expression {
	current, ok := p.it.current()
	if !ok {
		return nil // todo
	}
	if checkToken(current, opening, "(") {
		p.it.consume()
		ex := p.parseExpression()
		next, ok := p.it.current()
		if ok && checkToken(next, closing, ")") {
			p.it.consume()
			return ex
		} else {
			// todo: unmatched paren, error
		}
	} else if checkTokenType(current, number) || checkTokenType(current, boolean) || checkTokenType(current, stringLiteral) {
		p.it.consume()
		return literal(current)
	}
	return nil // error
}

func checkToken(tok token, tokType tokenType, lexeme string) bool {
	return checkTokenType(tok, tokType) && tok.lexeme == lexeme
}

func checkTokenType(tok token, tokType tokenType) bool {
	return tok.tokType == tokType
}