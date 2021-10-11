package parser

import (
	"strconv"
)

type IntegerLiteralExpression struct {
	Value int
}
func (ile *IntegerLiteralExpression) TokenLiteral() string {
	return strconv.Itoa(ile.Value)
}
func (ile *IntegerLiteralExpression) evaluateExpression() {}
