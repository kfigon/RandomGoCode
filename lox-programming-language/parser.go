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

func (p *Parser) parseExpression() expression {
	current,_ := p.it.current()
	t := current.tokType

	if t == number || t == stringLiteral || t == boolean {
		return literal(current)
	} else if t == operator && (current.lexeme == "!" || current.lexeme == "-") {
		p.it.consume()
		e := p.parseExpression()
		// todo nil
		return unary{op: current, ex: e}
	} else if t == operator {
		return operatorExpr(current)
	} else if t == opening && current.lexeme == "(" {
		p.it.consume()
		e := p.parseExpression()
		p.it.consume()
		next, ok := p.it.current() 
		if ok && next.tokType == closing && next.lexeme == ")" {
			return nil // todo
		}
		return grouping{e}
	}
	return nil
}