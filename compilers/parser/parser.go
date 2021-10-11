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
			p.statements = append(p.statements, p.parseVarStatement())
		} else if isReturnKeyword(tok) {
			p.statements = append(p.statements, p.parseReturnStatement())
		}
	}

	return &Program{p.statements, p.errors}
}

type Program struct {
	Statements []StatementNode
	Errors []error
}

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


type VarStatementNode struct {
	Name string
	Value ExpressionNode
}
func (vsn *VarStatementNode) TokenLiteral() string {
	return vsn.Name
}
func (vsn *VarStatementNode) evaluateStatement() {}


type IntegerLiteralExpression struct {
	Value int
}
func (ile *IntegerLiteralExpression) TokenLiteral() string {
	return strconv.Itoa(ile.Value)
}
func (ile *IntegerLiteralExpression) evaluateExpression() {}

func (p *parser) addError(err error) {
	p.errors = append(p.errors, err)
}

func (p *parser) parseExpression() ExpressionNode {
	for {
		tok, ok := p.iter.next()
		if !ok || isSemicolon(tok){
			break
		}
	}
	return nil
}


func (p *parser) parseVarStatement() StatementNode {
	tok, ok := p.iter.next()
	out := VarStatementNode{}
	if !ok {
		p.addError(fmt.Errorf("unexpected end of tokens"))
		return nil
	} else if !isIdentifier(tok) {
		p.addError(fmt.Errorf("expected identifier"))
		return nil
	}

	out.Name = tok.Lexeme

	tok, ok = p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("unexpected end of tokens"))
		return nil
	} else if !isAssignmentOperator(tok) {
		p.addError(fmt.Errorf("expected assignment after var"))
		return nil
	}

	// todo
	exp := p.parseExpression()
	out.Value = exp
	return &out
}

func (p *parser) parseReturnStatement() StatementNode {
	// todo
	p.parseExpression()
	return nil
}




func isVarKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "var"
}

func isAssignmentOperator(token lexer.Token) bool {
	return token.Class == lexer.Assignment && token.Lexeme == "="
}

func isSemicolon(token lexer.Token) bool {
	return token.Class == lexer.Semicolon && token.Lexeme == ";"
}

func isNumberLiteral(token lexer.Token) bool {
	return token.Class == lexer.Number
}

func isIdentifier(token lexer.Token) bool {
	return token.Class == lexer.Identifier
}

func isReturnKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "return"
}