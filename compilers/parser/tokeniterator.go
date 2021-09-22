package parser

import "programming-lang/lexer"

type tokenIterator struct {
	tokens []lexer.Token
	idx int
}

func newIterator(tokens []lexer.Token) *tokenIterator {
	return &tokenIterator{tokens: tokens}
}

func (t *tokenIterator) next() (lexer.Token, bool) {
	if t.idx >= len(t.tokens) {
		return lexer.Token{}, false
	}
	toRet := t.tokens[t.idx]
	t.idx++
	return toRet, true
}

func (t *tokenIterator) peek() (lexer.Token, bool) {
	if t.idx >= len(t.tokens) {
		return lexer.Token{}, false
	}
	return t.tokens[t.idx], true
}