package parser

import "programming-lang/lexer"

type VarStatementNode struct {
	Name  string
	Value ExpressionNode
}

func (vsn *VarStatementNode) TokenLiteral() string {
	return vsn.Name
}
func (vsn *VarStatementNode) evaluateStatement() {}

type ReturnStatementNode struct {
	Value ExpressionNode
}

func (r *ReturnStatementNode) TokenLiteral() string {
	return "return"
}
func (r *ReturnStatementNode) evaluateStatement() {}

// Statement wrapper for expressions, required for pratt parsing
type ExpressionStatementNode struct {
	Token lexer.Token //first token
	Value ExpressionNode
}

func (e *ExpressionStatementNode) TokenLiteral() string {
	return e.Token.Lexeme
}
func (e *ExpressionStatementNode) evaluateStatement() {}
