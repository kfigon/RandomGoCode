package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterpretExpression(t *testing.T) {
	testCases := []struct {
		desc	string
		input	string
		expected []loxObject
	}{
		{
			desc: "simple numeric literal",
			input: "15",
			expected: []loxObject{{v: toAnyPtr(15)}},
		},
		{
			desc: "simple expression1",
			input: "15+5",
			expected: []loxObject{{v: toAnyPtr(20)}},
		},
		{
			desc: "simple expression2",
			input: "15-5",
			expected: []loxObject{{v: toAnyPtr(10)}},
		},
		{
			desc: "simple expression3",
			input: "5-17",
			expected: []loxObject{{v: toAnyPtr(-12)}},
		},
		{
			desc: "simple expression4",
			input: "3*5",
			expected: []loxObject{{v: toAnyPtr(15)}},
		},
		{
			desc: "simple expression5",
			input: "15/5",
			expected: []loxObject{{v: toAnyPtr(3)}},
		},
		{
			desc: "unary expr1",
			input: "-15",
			expected: []loxObject{{v: toAnyPtr(-15)}},
		},
		{
			desc: "unary expr2",
			input: "!false",
			expected: []loxObject{{v: toAnyPtr(true)}},
		},
		{
			desc: "complicated expr",
			input: "5*3+1",
			expected: []loxObject{{v: toAnyPtr(16)}},
		},
		{
			desc: "complicated expr2",
			input: "1+5*3",
			expected: []loxObject{{v: toAnyPtr(16)}},
		},
		{
			desc: "complicated expr3",
			input: "5*3+1 == 16",
			expected: []loxObject{{v: toAnyPtr(true)}},
		},
		{
			desc: "complicated expr4",
			input: "true == !true",
			expected: []loxObject{{v: toAnyPtr(false)}},
		},
		{
			desc: "complicated expr5",
			input: "!false != !true",
			expected: []loxObject{{v: toAnyPtr(true)}},
		},
		{
			desc: "multiple expressions",
			input: `true;!true`,
			expected: []loxObject{{v: toAnyPtr(true)},{v: toAnyPtr(false)}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			toks,err := lex(tC.input)
			require.NoError(t, err, "got lexer error")

			p := NewParser(toks)
			p.Parse()
			exps, errs := p.Parse()
			require.Empty(t, errs, "got parser errors")

			got, interpreterErrs := interpret(exps)
			require.Empty(t, interpreterErrs, "got intepreter errors")

			assert.Equal(t, tC.expected, got)
		})
	}
}