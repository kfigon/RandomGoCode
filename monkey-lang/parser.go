package main

import (
	"fmt"
	"iter"
	"strconv"
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
		p.consume()
	}
	return out, nil
}

func (p *parser) parseStatement() (Statement, error) {
	t := p.current
	if t.Typ == Let {
		return p.parseLetStatement()
	} else if t.Typ == Return {
		return p.parseReturnStatement()
	}
	return p.parseExpressionStatement()
}

func (p *parser) parseLetStatement() (*LetStatement, error) {
	// current Let, peek Identifier
	if !p.expectPeek(Identifier){
		return nil, expectedTokenErr(Identifier, p.peek.Typ) 
	}

	identifier := p.current.Lexeme
	if !p.expectPeek(Assign) {
		return nil, expectedTokenErr(Assign, p.peek.Typ) 
	}

	p.consume() // assign
	exp, err := p.parseExpression(Lowest)
	if err != nil {
		return nil, err
	}

	if !p.expectPeek(Semicolon){
		return nil, expectedTokenErr(Semicolon, p.peek.Typ)
	}

	return &LetStatement{
		&IdentifierExpression{identifier},
		exp,
	}, nil
}

func (p *parser) parseReturnStatement() (*ReturnStatement, error) {
	p.consume()	// return
	exp, err := p.parseExpression(Lowest)
	if err != nil {
		return nil, err
	}

	if !p.expectPeek(Semicolon){
		return nil, expectedTokenErr(Semicolon, p.peek.Typ)
	}

	return &ReturnStatement{exp},nil
}

func (p *parser) parseExpressionStatement() (*ExpressionStatement, error) {
	exp, err := p.parseExpression(Lowest)
	if err != nil {
		return nil, err
	}

	if p.peekIs(Semicolon){
		p.consume()
	}

	return &ExpressionStatement{exp},nil
}

func (p *parser) parseExpression(precedence Precedence) (Expression, error) {
	left, err := p.parsePrefixExpression()
	if err != nil {
		return nil, err
	}

	for !p.eof() && !p.peekIs(Semicolon) && precedence < precedenceForToken(p.peek.Typ) {
		p.consume()
		newExpr, err := p.parseInfixExpression(left)
		if err != nil {
			return nil, err
		} else if newExpr == nil {
			break
		}
		left = newExpr
	}

	return left, nil
}

func (p *parser) parseInfixExpression(left Expression) (Expression, error) {
	switch p.current.Typ {
	case Plus, Minus, Slash, Asterisk, EQ, NEQ, LT, GT: 
		return p.parseLogicalAndArithmeticInfix(left)
	case LParen: 
		return nil, nil
	default: return nil, nil
	}
}

func (p *parser) parseLogicalAndArithmeticInfix(left Expression)(Expression, error) {
	op := p.current

	precedence := precedenceForToken(op.Typ)
	p.consume()
	right, err := p.parseExpression(precedence)
	if err != nil {
		return nil, err
	}

	return &InfixExpression{
		Operator: op,
		Left: left,
		Right: right,
	}, nil
}

func (p *parser) parsePrefixExpression() (Expression, error) {
	switch p.current.Typ {
	case Identifier: return &IdentifierExpression{p.current.Lexeme}, nil
	case Number:
		num, err := strconv.Atoi(p.current.Lexeme)
		if err != nil {
			return nil, err
		}
		return &PrimitiveLiteral[int]{num}, nil
	case True, False:
		b, err := strconv.ParseBool(p.current.Lexeme)
		if err != nil {
			return nil, err
		}
		return &PrimitiveLiteral[bool]{b}, nil
	case Bang, Minus:
		op := p.current
		p.consume()
		left, err := p.parseExpression(Prefix)
		if err != nil {
			return nil, err
		}
		return &PrefixExpression{
			Operator: op,
			Expr: left,
		},nil
	case LParen:
		p.consume() // (
		ex, err := p.parseExpression(Lowest)
		if err != nil {
			return nil, err
		}
		if !p.expectPeek(RParen) {
			return nil, fmt.Errorf("unmatched parens for grouped expressions, got %v", p.peek.Typ)
		}
		return ex, nil
	case If: return p.parseIf()
	case Function: return p.parseFun()
	}


	return nil,fmt.Errorf("invalid token for prefix expression: %v", p.current.Typ)
}

func (p *parser) parseIf() (*IfExpression, error) {
	p.consume() // if
	ex, err := p.parseExpression(Lowest)
	if err != nil {
		return nil, err
	}

	if !p.expectPeek(LBrace) {
		return nil, fmt.Errorf("if expression error, expected open brace, got %v", p.peek.Typ)
	}

	block, err := p.parseBlockStatements()
	if err != nil {
		return nil, err
	}

	ifExp := &IfExpression{
		Predicate: ex,
		Consequence: block,
	}

	if p.peekIs(Else) {
		p.consume() // }
		p.consume() // else

		if p.current.Typ == If { // if-elseif
			next, err := p.parseIf()
			if err != nil {
				return nil, err
			}
			ifExp.Alternative = next
		} else if p.current.Typ == LBrace { // if-else
			alternativeBlock, err := p.parseBlockStatements()
			if err != nil {
				return nil, err
			}
			ifExp.Alternative = &IfExpression{
				Predicate: nil,
				Consequence: alternativeBlock,
				Alternative: nil,
			}
		}
	}
	return ifExp,nil
}

func (p *parser) parseFun() (*FunctionLiteral, error) {
	if !p.expectPeek(LParen) {
		return nil, fmt.Errorf("expected parameter list for function, got %v", p.peek)
	}

	p.consume() // fn, ( 
	params := []*IdentifierExpression{}
	for !p.eof() && p.current.Typ != RParen {
		param, err := p.parsePrefixExpression()
		if err != nil {
			return nil, err
		}
		ident, ok := param.(*IdentifierExpression)
		if !ok {
			return nil, fmt.Errorf("expected identifier for function params, got %T", ident)
		}
		params = append(params, ident)
		p.consume() // ident
		if p.current.Typ == Comma && p.peekIs(Identifier) {
			p.consume() // comma
		} else if !p.eof() && p.current.Typ != RParen {
			return nil, fmt.Errorf("error parsing parameter list, expected identifier separated by commas and end with ')', got %v", p.current.Typ)
		}
	}

	if !p.expectPeek(LBrace) {
		return nil, fmt.Errorf("expected function body, got %v", p.peek.Typ)
	}

	block, err := p.parseBlockStatements()
	if err != nil {
		return nil, err
	}

	return &FunctionLiteral{
		Parameters: params,
		Body: block,
	}, nil
}

func (p *parser) parseBlockStatements() (*BlockStatement, error) {
	p.consume() // {
	var out []Statement
	for !p.eof() && p.current.Typ != RBrace {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		out = append(out, stmt)
		p.consume()
	}

	return &BlockStatement{out}, nil
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

func precedenceForToken(tok TokenType) Precedence {
	switch tok {
	case EQ, NEQ: return Equals
	case LT, GT: return LessGreater
	case Plus, Minus: return Sum
	case Asterisk, Slash: return Product
	default: return Lowest
	}
}
