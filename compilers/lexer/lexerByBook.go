package lexer

// dont know if this actually works, it's just a mechanism for a reference

type llex struct {
	input string
	idx int
}

func newLlexer(input string) *llex {
	return &llex{input: input}
}

type llexToken struct {
	token string
	lexeme string
}

func (l *llex) readChar() rune {
	toRet := l.input[l.idx]
	l.idx++
	return rune(toRet)
}

func (l *llex) peekChar() rune {
	return rune(l.input[l.idx+1])
}

func (l *llex) eof() bool {
	return l.idx >= len(l.input)
}

func (l *llex) emit() llexToken {
	if l.eof() {
		return llexToken{token: "EOF", lexeme: ""}
	}
	char := l.readChar()
	switch char{
	case ';': return llexToken{"SEMICOLON", string(char)}
	case '=': {
		if !l.eof() && l.peekChar() == '=' {
			l.readChar()
			return llexToken{"EQUALS", "=="}
		}
		return llexToken{"ASSIGNMENT", "="}
	}
	case '!': {
		if !l.eof() && l.peekChar() == '=' {
			l.readChar()
			return llexToken{"NOT_EQUALS", "!="}
		}
		return llexToken{"NEGATION", "!"}
	}
	default: {
		 if isAlpha(char) {
			consumedString := string(char)+l.eatIdentifier()
			if isKeyword(consumedString) {
				return llexToken{"KEYWORD", consumedString}
			}
			return llexToken{"IDENTIFIER", consumedString}
		 } else if isDigit(char) {
			consumedString := string(char)+l.eatNumber()
			return llexToken{"NUMBER", consumedString}
		 } else if isWhiteSpace(char) {
			 return llexToken{"WHITESPACE",string(char)}
		 }
	}
	}
	return llexToken{"UNKNOWN_TOKEN", string(char)}
}

func isWhiteSpace(char rune) bool {
	return char==' ' || char == '\t' || char == '\n'
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func (l *llex) eatIdentifier() string {
	var out string
	for !l.eof() {
		c := l.readChar()
		if !isAlpha(c) {
			break
		}
		out += string(c)
	}
	return out
}

func (l *llex) eatNumber() string {
	var out string
	for !l.eof() {
		c := l.readChar()
		if !isDigit(c) {
			break
		}
		out += string(c)
	}
	return out
}

func isKeyword(input string) bool {
	keywords := map[string]bool{
		"let": true,
		"func": true,
		"for": true,
	}
	return keywords[input]
}
