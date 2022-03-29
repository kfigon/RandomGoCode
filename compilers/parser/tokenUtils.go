package parser

import "programming-lang/lexer"

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

func eof(token lexer.Token) bool {
	return token.Class == lexer.EOF
}

func bang(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "!"
}

func minus(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "-"
}