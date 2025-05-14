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
			exp: &PrimitiveObj[bool]{false},
		},
		{
			desc: "true literal",
			code: "true",
			exp: &PrimitiveObj[bool]{true},
		},
		{
			desc: "null literal",
			code: "null",
			exp: &NullObj{},
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