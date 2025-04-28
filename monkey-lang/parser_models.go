package main

type Statement interface {
	statementTag()
}

type Expression interface {
	expressionTag()
}

type LetStatement struct {
	Ident *IdentifierExpression
	Value Expression
}

func (l *LetStatement) statementTag() {}

type IdentifierExpression struct {
	Name string
}

func (*IdentifierExpression) expressionTag() {}

type ReturnStatement struct {
	Exp Expression
}

func (*ReturnStatement) statementTag() {}
