package main

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
	for current, ok := p.it.current(); ok; current, ok = p.it.current() {
		_ = current
	}
	return p.Expressions, p.Errors
}