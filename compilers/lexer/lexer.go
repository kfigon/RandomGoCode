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

func (t Token) String() string {
	return fmt.Sprintf("{%v %q}", t.Class, t.Lexeme)
}

func (t TokenClass) String() string {
	return classesStrings[t]
}

func Tokenize(input string) []Token {
	var tokens []Token
	var idx uint64

	update := func(found string, class TokenClass) {
		idx += uint64(len(found))
		tokens = append(tokens, Token{Class: class, Lexeme: found})
	}

	tokenizerEntries := []tokenizerEntry {
		{`^(\s+)`, Whitespace},

		{`^(if)($|\s)`, Keyword},
		{`^(else)($|\s)`, Keyword},

		{`^(==)($|\s?)`, Operator},
		{`^(=)($|\s?)`, Assignment},

		{`^([0-9]+\.[0-9]+)`, Number},
		{`^([0-9]+)`, Number},
		
		{`^(\w+)`, Identifier},

		{`^(;)`, Semicolon},
		{`^(\))`, CloseParam},
		{`^(\()`, OpenParam},
	}

	ln := uint64(len(input))
	for idx < ln {
		rest := input[idx:]
	
		if enableLogs {
			toPrint := 20
			if idx + uint64(toPrint) > ln {
				toPrint = int(ln-idx-1)
			}
			log.Printf("parsing %q...\n", rest[:toPrint])
		}

	
		found := false
		for _, entry := range tokenizerEntries {
			substr, ok := findStr(rest, entry.pattern)
			if !ok {
				continue
			}

			if enableLogs {
				log.Printf("found %q -> %v, moving up to %v\n", substr, entry.class, len(substr))
			}

			update(substr, entry.class)
			found = true
			break
		}

		if !found {
			log.Println("Unknown token at idx", idx)
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