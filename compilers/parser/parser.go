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
	p := &parser{iter: newIterator(tokens)}
	for {
		tok, ok := p.iter.next()
		if !ok {
			break
		} else if isVarKeyword(tok) {
			p.parseVarStatement()
		} else if isReturnKeyword(tok) {
			p.parseReturnStatement()
		}
	}

	return &Program{p.statements, p.errors}
}

type parser struct {
	iter *tokenIterator
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
	switch {
	case !ok: return nil
	case isSemicolon(tok): {
		p.addError(fmt.Errorf("expression error - no expresion found, got %v", tok.Lexeme))
		return nil
	}
	case isNumberLiteral(tok): return p.parseIntegerLiteralExpression(tok)
	case isIdentifier(tok): return p.parseIdentifierExpression(tok)
	default: return nil
	}
}

func (p *parser) parseIntegerLiteralExpression(tok lexer.Token) ExpressionNode {
	v, err := strconv.Atoi(tok.Lexeme)
	if err != nil {
		p.addError(fmt.Errorf("int literal expression error - error in parsing integer literal in: %v", tok.Lexeme))
		return nil
	}
	tok, ok := p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("int literal expression error - unexpected end of tokens"))
		return nil
	} else if !isSemicolon(tok) {
		p.addError(fmt.Errorf("int literal expression error - expected semicolon, got %v", tok.Lexeme))
		return nil
	}
	return &IntegerLiteralExpression{Value: v}		
}

func (p *parser) parseIdentifierExpression(token lexer.Token) ExpressionNode {
	out := &IdentifierExpression{Name: token.Lexeme}
	tok, ok := p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("identifier expression error - unexpected end of tokens"))
		return nil
	} else if !isSemicolon(tok) {
		p.addError(fmt.Errorf("identifier expression error - semicolon not found, got %v", tok))
		return nil
	}
	return out
}

func (p *parser) parseVarStatement() {
	identifierTok, ok := p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after var"))
		return
	} else if !isIdentifier(identifierTok) {
		p.addError(fmt.Errorf("var error - expected identifier, got %v", identifierTok.Class))
		return
	}

	tok, ok := p.iter.next()
	if !ok {
		p.addError(fmt.Errorf("var error - unexpected end of tokens after identifier"))
		return
	} else if isSemicolon(tok) {
		out := VarStatementNode{Name: identifierTok.Lexeme}
		p.addStatement(&out)
		return
	} else if !isAssignmentOperator(tok) {
		p.addError(fmt.Errorf("var error - expected assignment after identifier, got %v", tok.Class))
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