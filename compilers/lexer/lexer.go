package lexer

type Token struct {
	Class TokenClass
	Content string
}

type TokenClass int
const (
	Whitespace TokenClass = iota
	Keyword
	Identifier // variables
	Number
	Operator

	OpenParam
	CloseParam
	Semicolon
	Assignment
)

func Tokenize(input string) []Token {
	return []Token{}
}