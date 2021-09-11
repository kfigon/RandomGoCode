package lexer

import (
	"fmt"
	"log"
	"regexp"
)

var enableLogs bool = false


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

type tokenizerEntry struct {
	pattern string
	class TokenClass
}

var tokenizerEntries []tokenizerEntry = []tokenizerEntry{
	{`^(\s+)`, Whitespace},

	{`^(if)($|\s)`, Keyword},
	{`^(else)($|\s)`, Keyword},

	{`^(==)($|\s?)`, Operator},
	{`^(\+)($|\s?)`, Operator},
	{`^(\-)($|\s?)`, Operator},
	{`^(\*)($|\s?)`, Operator},
	{`^(\/)($|\s?)`, Operator},
	{`^(=)($|\s?)`, Assignment},

	{`^(;)`, Semicolon},
	{`^(\))`, CloseParam},
	{`^(\()`, OpenParam},

	{`^([0-9]+\.[0-9]+)`, Number},
	{`^([0-9]+)`, Number},

	{`^(\w+)`, Identifier},
}

func (t Token) String() string {
	return fmt.Sprintf("{%v %q}", t.Class, t.Lexeme)
}

func (t TokenClass) String() string {
	return classesStrings[t]
}

func Tokenize(input string) []Token {
	var tokens []Token
	var idx uint64

	ln := uint64(len(input))
	for idx < ln {
		rest := input[idx:]

		logLine(idx, ln, rest)
		

		found, deltaIdx, token := processAvailableTokens(rest)

		if found {
			idx += uint64(deltaIdx)
			tokens = append(tokens, token)
		} else {
			log.Println("Unknown token at idx", idx)
			idx++
		} 
	}
	return tokens
}

func logLine(idx, ln uint64, rest string) {
	if !enableLogs {
		return
	}
	
	toPrint := 20
	if idx + uint64(toPrint) > ln {
		toPrint = int(ln-idx)
	}
	log.Printf("parsing %q...\n", rest[:toPrint])
}

func processAvailableTokens(input string) (bool, int, Token) {
	for _, entry := range tokenizerEntries {
		substr, ok := findPattern(input, entry.pattern)
		if !ok {
			continue
		}

		if enableLogs {
			log.Printf("found %q -> %v, moving up to %v\n", substr, entry.class, len(substr))
		}

		return true, len(substr), Token{Class: entry.class, Lexeme: substr}
	}

	return false, 0, Token{}
}

func findPattern(input string, pattern string) (string, bool) {
	reg := regexp.MustCompile(pattern)
	res := reg.FindStringSubmatch(input)
	if len(res) < 2 {
		return "", false
	}
	return res[1],true
}