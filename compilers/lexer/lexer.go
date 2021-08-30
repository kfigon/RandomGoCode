package lexer

import "regexp"

type Token struct {
	Class  TokenClass
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
var classesStrings = []string{
	"Whitespace",
	"Keyword",
	"Identifier",
	"Number",
	"Operator",
	"OpenParam",
	"CloseParam",
	"Semicolon",
	"Assignment",
}

func (t TokenClass) String() string {
	return classesStrings[t]
}

func Tokenize(input string) []Token {
	var tokens []Token
	var idx uint64

	ln := uint64(len(input))
	update := func(found string, class TokenClass) {
		idx += uint64(len(found))
		tokens = append(tokens, Token{Class: class, Lexeme: found})
	}

	tokenizerEntries := []tokenizerEntry {
		{`\s`, Whitespace},

		{`(if)\s`, Keyword},
		{`(else)\s`, Keyword},

		{`(==)\s`, Operator},
		{`(=)\s`, Assignment},
		
		{`([0-9]+)\s`, Number},
		{`([0-9]+\.[0-9]+)\s`, Number},
		
		{`(\w+)\s`, Identifier},

		{`(;)\s`, Semicolon},
		{`(\))\s`, OpenParam},
		{`(\()\s`, CloseParam},
	}
	for idx < ln {
		rest := input[idx:]

		found := false
		for _, entry := range tokenizerEntries {
			substr, ok := findStr(rest,entry.pattern)
			if ok {
				update(substr, entry.class)
				found = true
				break
			}
		}

		if !found {
			idx++
		}
	}
	return tokens
}

type tokenizerEntry struct {
	pattern string
	class TokenClass
}

func findStr(input string, pattern string) (string, bool) {
	reg := regexp.MustCompile(pattern)
	res := reg.FindStringSubmatch(input)
	if len(res) < 2 {
		return "", false
	}
	return res[1],true
}