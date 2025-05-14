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
			desc: "literal",
			code: "14;",
			exp: nil,
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