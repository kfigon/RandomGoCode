package parser

import "programming-lang/lexer"

type infixFn func(leftSide *ExpressionNode) *ExpressionNode
type prefixFn func() *ExpressionNode

type mapPair[V any] struct {
	key lexer.TokenClass
	val V
}

func registerParsingFns[T any](fns ...mapPair[T]) map[lexer.TokenClass]T {
	out := map[lexer.TokenClass]T{}
	for _, v := range fns {
		out[v.key] = v.val
	}

	return out
}
