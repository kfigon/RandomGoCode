package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		desc	string
		input string
		expected []Statement
		wantErr bool
	}{
		{
			desc: "basic let stmt",
			input: `let foobar = 234;`,
			expected: []Statement{
				&LetStatement{
					&IdentifierExpression{"foobar"},
					nil, 
				},
			},
		},
		{
			desc: "return stmt",
			input: `return 234;`,
			expected: []Statement{
				&ReturnStatement{ nil },
			},
		},
		{
			desc: "return stmt with identifier",
			input: `return foobar;`,
			expected: []Statement{
				&ReturnStatement{ nil },
			},
		},
		{
			desc: "number literal",
			input: `15;`,
			expected: []Statement{
				&ExpressionStatement{&PrimitiveLiteral[int]{15}},
			},
		},
		{
			desc: "boolean literal",
			input: `true;`,
			expected: []Statement{
				&ExpressionStatement{&PrimitiveLiteral[bool]{true}},
			},
		},
		{
			desc: "identifier expression",
			input: `foobar;`,
			expected: []Statement{
				&ExpressionStatement{&IdentifierExpression{"foobar"}},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := Parse(Lex(tC.input))	
			if tC.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.expected, got)
			}
		})
	}
}