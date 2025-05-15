package main

import (
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
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			parsed, err := Parse(Lex(tC.code))
			require.NoError(t, err, "parsing error")

			got, err := Eval(parsed)

			require.NoError(t, err, "eval error")
			assert.Equal(t, tC.exp, got)
		})
	}
}