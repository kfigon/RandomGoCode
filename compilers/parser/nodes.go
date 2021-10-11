package parser

import (
	"strconv"
)

type VarStatementNode struct {
	Name string
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


type IntegerLiteralExpression struct {
	Value int
}
func (ile *IntegerLiteralExpression) TokenLiteral() string {
	return strconv.Itoa(ile.Value)
}
func (ile *IntegerLiteralExpression) evaluateExpression() {}
