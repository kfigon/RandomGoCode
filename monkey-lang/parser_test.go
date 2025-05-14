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
		{
			desc: "not grouped expressions",
			input: `1 + 2+3+4;`, // (((1+2)+3)+4)
			expected: []Statement{
				&ExpressionStatement{
					&InfixExpression{
						Operator: Token{Plus, "+"},
						Left: &InfixExpression{
							Operator: Token{Plus, "+"},
							Left: &InfixExpression{
								Operator: Token{Plus, "+"},
								Left: &PrimitiveLiteral[int]{1},
								Right: &PrimitiveLiteral[int]{2},
							},
							Right: &PrimitiveLiteral[int]{3},
						},
						Right: &PrimitiveLiteral[int]{4},
					},
				},
			},
		},
		{
			desc: "grouped expressions",
			input: `1 + (2+3)+4;`,
			expected: []Statement{
				&ExpressionStatement{
					&InfixExpression{
						Operator: Token{Plus, "+"},
						Left: &InfixExpression{
							Operator: Token{Plus, "+"},
							Left: &PrimitiveLiteral[int]{1},
							Right: &InfixExpression{
								Operator: Token{Plus, "+"},
								Left: &PrimitiveLiteral[int]{2},
								Right: &PrimitiveLiteral[int]{3},
							},
						},
						Right: &PrimitiveLiteral[int]{4},
					},
				},
			},
		},
		{
			desc: "if else expression",
			input: `if x < y { x } else { y }`,
			expected: []Statement{
				&ExpressionStatement{
					&IfExpression{
						Predicate: &InfixExpression{
							Operator: Token{LT, "<"},
							Left: &IdentifierExpression{"x"},
							Right: &IdentifierExpression{"y"},
						},
						Consequence: &BlockStatement{[]Statement{&ExpressionStatement{&IdentifierExpression{"x"}}}},
						Alternative: &IfExpression{
							Predicate: nil,
							Consequence: &BlockStatement{[]Statement{&ExpressionStatement{&IdentifierExpression{"y"}}}},
							Alternative: nil,
						},
					},
				},
			},
		},
		{
			desc: "if expression",
			input: `if x < y { x }`,
			expected: []Statement{
				&ExpressionStatement{
					&IfExpression{
						Predicate: &InfixExpression{
							Operator: Token{LT, "<"},
							Left: &IdentifierExpression{"x"},
							Right: &IdentifierExpression{"y"},
						},
						Consequence: &BlockStatement{[]Statement{&ExpressionStatement{&IdentifierExpression{"x"}}}},
						Alternative: nil,
					},
				},
			},
		},
		{
			desc: "if else if expression",
			input: `if x < y { x } else if x > y {y} else {123}`,
			expected: []Statement{
				&ExpressionStatement{
					&IfExpression{
						Predicate: &InfixExpression{
							Operator: Token{LT, "<"},
							Left: &IdentifierExpression{"x"},
							Right: &IdentifierExpression{"y"},
						},
						Consequence: &BlockStatement{[]Statement{&ExpressionStatement{&IdentifierExpression{"x"}}}},
						Alternative: &IfExpression{
							Predicate: &InfixExpression{
								Operator: Token{GT, ">"},
								Left: &IdentifierExpression{"x"},
								Right: &IdentifierExpression{"y"},
							},
							Consequence: &BlockStatement{[]Statement{&ExpressionStatement{&IdentifierExpression{"y"}}}},
							Alternative: &IfExpression{
								Predicate: nil,
								Consequence: &BlockStatement{[]Statement{&ExpressionStatement{&PrimitiveLiteral[int]{123}}}},
								Alternative: nil,
							},
						},
					},
				},
			},
		},
		{
			desc: "function literal",
			input: `fun(x,y){ return x + y; }`,
			expected: []Statement{
				&ExpressionStatement{ &FunctionLiteral{
						Parameters: []*IdentifierExpression{{"x"}, {"y"}},
						Body: &BlockStatement{[]Statement{
							&ReturnStatement{&InfixExpression{
								Operator: Token{Plus,"+"},
								Left: &IdentifierExpression{"x"},
								Right: &IdentifierExpression{"y"},
							}},
						},
					}}},
				},
		},
		{
			desc: "function literal assigned",
			input: `let foo = fun(x,y){ return x + y; };`,
			expected: []Statement{
				&LetStatement{
					&IdentifierExpression{"foo"},
					&FunctionLiteral{
					Parameters: []*IdentifierExpression{{"x"}, {"y"}},
					Body: &BlockStatement{[]Statement{
						&ReturnStatement{&InfixExpression{
							Operator: Token{Plus,"+"},
							Left: &IdentifierExpression{"x"},
							Right: &IdentifierExpression{"y"},
						}}},
					}}},
			},
		},
		{
			desc: "function literal without args",
			input: `fun(){ return x + y; }`,
			expected: []Statement{
				&ExpressionStatement{ &FunctionLiteral{
						Parameters: []*IdentifierExpression{},
						Body: &BlockStatement{[]Statement{
							&ReturnStatement{&InfixExpression{
								Operator: Token{Plus,"+"},
								Left: &IdentifierExpression{"x"},
								Right: &IdentifierExpression{"y"},
							}},
						},
					}}},
				},
		},
		{
			desc: "function call with name",
			input: `foo(a-2, 1+c)`,
			expected: []Statement{
				&ExpressionStatement{
					&FunctionCall{
						&IdentifierExpression{"foo"},
						[]Expression{
							&InfixExpression{
								Operator: Token{Minus, "-"},
								Left: &IdentifierExpression{"a"},
								Right: &PrimitiveLiteral[int]{2},
							},
							&InfixExpression{
								Operator: Token{Plus, "+"},
								Left: &PrimitiveLiteral[int]{1},
								Right: &IdentifierExpression{"c"},
							},
						},
					},
				},
			},
		},
		{
			desc: "function call without args",
			input: `foo()`,
			expected: []Statement{
				&ExpressionStatement{
					&FunctionCall{
						&IdentifierExpression{"foo"},
						[]Expression{}},
					},
			},
		},
		{
			desc: "function call with function literal",
			input: `fun(x,y){ return x + y; }(a, 1+c)`,
			expected: []Statement{
				&ExpressionStatement{
					&FunctionCall{
						&FunctionLiteral{
							Parameters: []*IdentifierExpression{ {"x"}, {"y"} },
							Body: &BlockStatement{
								[]Statement{&ReturnStatement{
									&InfixExpression{
										Token{Plus,"+"},
										&IdentifierExpression{"x"},
										&IdentifierExpression{"y"},
									},
								}},
							},
						},
						[]Expression{
							&IdentifierExpression{"a"},
							&InfixExpression{
								Operator: Token{Plus, "+"},
								Left: &PrimitiveLiteral[int]{1},
								Right: &IdentifierExpression{"c"},
							},
						},
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := Parse(Lex(tC.input))	
			require.NoError(t, err)
			assert.Equal(t, tC.expected, got)
		})
	}
}