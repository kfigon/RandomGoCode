package interpreter

import (
	"monkey-lang/lexer"
	"monkey-lang/parser"
	. "monkey-lang/objects"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	testCases := []struct {
		desc	string
		code string
		exp Object
	}{
		{
			desc: "int literal",
			code: "14",
			exp: &PrimitiveObj[int]{14},
		},
		{
			desc: "false literal",
			code: "false",
			exp: FALSE,
		},
		{
			desc: "true literal",
			code: "true",
			exp: TRUE,
		},
		{
			desc: "null literal",
			code: "null",
			exp: NULL,
		},
		{
			desc: "prefix1",
			code: "!true",
			exp: FALSE,
		},
		{
			desc: "prefix2",
			code: "!!true",
			exp: TRUE,
		},
		{
			desc: "prefix3",
			code: "!false",
			exp: TRUE,
		},
		{
			desc: "prefix4",
			code: "-123",
			exp: &PrimitiveObj[int]{-123},
		},
		{
			desc: "infix plus",
			code: "1 + 3",
			exp: &PrimitiveObj[int]{4},
		},
		{
			desc: "infix plus 2",
			code: "1 + 3 + 1 + 4 + 5",
			exp: &PrimitiveObj[int]{1+3+1+4+5},
		},
		{
			desc: "infix with mixed predescence",
			code: "1 + 3 * 2",
			exp: &PrimitiveObj[int]{7},
		},
		{
			desc: "infix with mixed predescence with grouping",
			code: "(1 + 3) * 2",
			exp: &PrimitiveObj[int]{8},
		},
		{
			desc: "many operators",
			code: "7 + (3 - 2) * (8 + 1) - 4",
			exp: &PrimitiveObj[int]{12},
		},
		{
			desc: "simple if",
			code: "if true { 1 }",
			exp: &PrimitiveObj[int]{1},
		},
		{
			desc: "simple false if",
			code: "if false { 1 }",
			exp: NULL,
		},
		{
			desc: "if else when true",
			code: "if true { 1 } else { 2 }",
			exp: &PrimitiveObj[int]{1},
		},
		{
			desc: "if else when false",
			code: "if false { 1 } else { 2 }",
			exp: &PrimitiveObj[int]{2},
		},
		{
			desc: "if else if when else true",
			code: "if false { 1 } else if true { 2 } else { 3 }",
			exp: &PrimitiveObj[int]{2},
		},
		{
			desc: "evaluated if",
			code: "if 2 > 3 { 1 } else if 1 == 1 { 2 }",
			exp: &PrimitiveObj[int]{2},
		},
		{
			desc: "if else if when false",
			code: "if false { 1 } else if false { 2 } else { 3 }",
			exp: &PrimitiveObj[int]{3},
		},
		{
			desc: "simple return",
			code: `1+1;
			return 3;
			3+13;`,
			exp: &PrimitiveObj[int]{3},
		},
		{
			desc: "block return",
			code: 
			`if true {
				if true {
					return 3;
				}	
				return 1;
			} `,
			exp: &PrimitiveObj[int]{3},
		},
		{
			desc: "let statetements 1",
			code: `let a = 5;
			a;`,
			exp: &PrimitiveObj[int]{5},
		},
		{
			desc: "let statetements 2",
			code: `let a = 5;
			let x = 5 + a;
			x;`,
			exp: &PrimitiveObj[int]{10},
		},
		{
			desc: "let statetements when invalid",
			code: `let a = 5;
			x;`,
			exp: NULL,
		},
		{
			desc: "function",
			code: `let max = fun(x,y) { if x > y { x } else {y}};
			max(2+1,2);`,
			exp: &PrimitiveObj[int]{3},
		},
		{
			desc: "function without args and inner env",
			code: `let foo = fun() {
				let x = 123;
				return x + 3;
			};
			foo();`,
			exp: &PrimitiveObj[int]{126},
		},
		{
			desc: "function literal",
			code: `fun(y) {
				let x = 123;
				return x + y;
			}(3);`,
			exp: &PrimitiveObj[int]{126},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			parsed, err := parser.Parse(lexer.Lex(tC.code))
			require.NoError(t, err, "parsing error")

			got, err := Eval(parsed)

			require.NoError(t, err, "eval error")
			assert.Equal(t, tC.exp, got)
		})
	}
}