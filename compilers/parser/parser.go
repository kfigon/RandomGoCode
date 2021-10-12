package parser

import (
	"fmt"
	"programming-lang/lexer"
	"strconv"
)

type parser struct {
	iter *tokenIterator
	errors []error
	statements []StatementNode
}

func Parse(tokens []lexer.Token) *Program {
	p := &parser{iter: newIterator(tokens)}
	for {
		tok, ok := p.iter.next()
		if !ok {
			break
		} else if isVarKeyword(tok) {
			p.parseVarStatement()
		} else if isReturnKeyword(tok) {
			p.parseReturnStatement()
		} else if next, nextOk := p.iter.peek(); nextOk && isFunctionCall(tok, next) {
			// p.parseFunctionCallStatement()
			// todo
		}
	}

	return &Program{p.statements, p.errors}
}

type Program struct {
	Statements []StatementNode
	Errors []error
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
	tok, ok := p.iter.next()
	if !ok {
		return nil
	} else if isNumberLiteral(tok) {
		return p.parseIntegerLiteralExpression(tok)
	} else if isSemicolon(tok) {
		p.addError(fmt.Errorf("expression error - no expresion found, got %v", tok.Lexeme))
		return nil
	}

	return nil
}

func (p *parser) parseIntegerLiteralExpression(tok lexer.Token) ExpressionNode {
	v, err := strconv.Atoi(tok.Lexeme)
	if err != nil {
		p.addError(fmt.Errorf("expression error - error in parsing integer literal in: %v", tok.Lexeme))
		return nil
	}
	tok, ok := p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("expression error - unexpected end of tokens"))
		return nil
	} else if !isSemicolon(tok) {
		p.addError(fmt.Errorf("expression error - expected semicolon, got %v", tok.Lexeme))
		return nil
	}
	return &IntegerLiteralExpression{Value: v}		
}

func (p *parser) parseVarStatement() {
	tok, ok := p.iter.next()
	out := VarStatementNode{}
	if !ok {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after var"))
		return
	} else if !isIdentifier(tok) {
		p.addError(fmt.Errorf("var error - expected identifier, got %v", tok.Class))
		return
	}

	out.Name = tok.Lexeme

	tok, ok = p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after identifier"))
		return
	} else if !isAssignmentOperator(tok) && !isSemicolon(tok) {
		p.addError(fmt.Errorf("var error - expected assignment or semicolon after identifier, got %v", tok.Class))
		return
	}

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

type parsingFunc func(ex ExpressionNode) ExpressionNode
func (p *parser) prattExpressionParser(tok lexer.Token) parsingFunc {
	return nil
}