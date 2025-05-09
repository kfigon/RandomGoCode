package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
					Ident: &IdentifierExpression{"foobar"},
					Value: &PrimitiveLiteral[int]{234},
				},
			},
		},
		{
			desc: "return stmt",
			input: `return 234;`,
			expected: []Statement{
				&ReturnStatement{
					&PrimitiveLiteral[int]{234},
				 },
			},
		},
		{
			desc: "return stmt with identifier",
			input: `return foobar;`,
			expected: []Statement{
				&ReturnStatement{
					&IdentifierExpression{"foobar"},
				},
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
		{
			desc: "prefix bang with identifier",
			input: `!foobar;`,
			expected: []Statement{
				&ExpressionStatement{
					&PrefixExpression{
						Token{Bang,"!"},
						&IdentifierExpression{"foobar"},
					},
				},
			},
		},
		{
			desc: "prefix bang with boolean",
			input: `!true;`,
			expected: []Statement{
				&ExpressionStatement{
					&PrefixExpression{
						Token{Bang,"!"},
						&PrimitiveLiteral[bool]{true},
					},
				},
			},
		},
		{
			desc: "negative 123",
			input: `-123;`,
			expected: []Statement{
				&ExpressionStatement{
					&PrefixExpression{
						Token{Minus, "-"},
						&PrimitiveLiteral[int]{123},
					},
				},
			},
		},
		{
			desc: "basic infix",
			input: `12 + 34;`,
			expected: []Statement{
				&ExpressionStatement{
					&InfixExpression{
						Operator: Token{Plus,"+"},
						Left: &PrimitiveLiteral[int]{12},
						Right: &PrimitiveLiteral[int]{34},
					},
				},
			},
		},
		{
			desc: "basic infix2",
			input: `12 + 34 + 1;`,
			expected: []Statement{
				&ExpressionStatement{
					&InfixExpression{
						Operator: Token{Plus,"+"},
						Left: &InfixExpression{
							Operator: Token{Plus,"+"},
							Left: &PrimitiveLiteral[int]{12},
							Right: &PrimitiveLiteral[int]{34},
						},
						Right: &PrimitiveLiteral[int]{1},
					},
				},
			},
		},
		{
			desc: "infix with 2 operators with different predescence",
			input: `1 + 2 * 3;`,
			expected: []Statement{
				&ExpressionStatement{
					&InfixExpression{
						Operator: Token{Plus,"+"},
						Left: &PrimitiveLiteral[int]{1},
						Right: &InfixExpression{
							Operator: Token{Asterisk, "*"},
							Left: &PrimitiveLiteral[int]{2},
							Right: &PrimitiveLiteral[int]{3},
						},
					},
				},
			},
		},
		{
			desc: "infix with 2 operators with different predescence 2",
			input: `1 * 2 + 3;`,
			expected: []Statement{
				&ExpressionStatement{
					&InfixExpression{
						Operator: Token{Plus,"+"},
						Left: &InfixExpression{
							Operator: Token{Asterisk, "*"},
							Left: &PrimitiveLiteral[int]{1},
							Right: &PrimitiveLiteral[int]{2},
						},
						Right: &PrimitiveLiteral[int]{3},
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := Parse(Lex(tC.input))	
			if tC.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tC.expected, got)
			}
		})
	}
}