package lexer

type Token struct {
	Class TokenClass
	Lexeme string
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

var keywords []string = []string{"if", "else", "else if", "for"}
var whitespaces []string = []string{" ", "\t", "\n"}

func makeSet() *set {
	out := newSet()
	for _, v := range keywords {
		out.add(v)
	}
	return out
}

func Tokenize(input string) []Token {
	var tokens []Token
	var idx uint64

	ln := uint64(len(input))
	for idx < ln {
		// char := input[idx]

		// todo: state machine per keyword?
		idx++
	}
	return tokens
}