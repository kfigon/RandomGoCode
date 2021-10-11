package parser

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

