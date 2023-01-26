package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	testCases := []struct {
		desc	string
		input	string
		expected []token
	}{
		{
			desc: "identifiers",
			input: ` somevalue ife`,
			expected: []token{
				{ identifier, "somevalue"},
				{ identifier, "ife"},
			},
		},
		{
			desc: "number",
			input: ` 1234;`,
			expected: []token{
				{number, "1234"},
				{semicolon, ";"},
			},
		},
		{
			desc: "whitespaces",
			input: " \t \n 123\t",
			expected: []token{
				{number, "123"},
			},
		},
		{
			desc: "whitespaces and string",
			input: ` 	 
" fo
o	"	`,
			expected: []token{
				{stringLiteral, "\" fo\no\t\""},
			},
		},
		{
			desc: "string literal",
			input: "\" hello world if 123\" 123",
			expected: []token{
				{stringLiteral, "\" hello world if 123\""},
				{number, "123"},
			},
		},
		{
			desc: "operators",
			input: `= == < <= > >= ! !! != || &&`,
			expected: []token{
				{ operator, "="},
				{ operator, "=="},
				{ operator, "<"},
				{ operator, "<="},
				{ operator, ">"},
				{ operator, ">="},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!="},
				{ operator, "||"},
				{ operator, "&&"},
			},
		},
		{
			desc: "operators without spaces",
			input: `==<<=>>=||&&!!!!=`,
			expected: []token{
				{ operator, "=="},
				{ operator, "<"},
				{ operator, "<="},
				{ operator, ">"},
				{ operator, ">="},
				{ operator, "||"},
				{ operator, "&&"},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!="},
			},
		},
	}
	for _, tC := range testCases {
		stringify := func(xs []token) []string {
			var out []string
			for _, x := range xs {
				out = append(out, x.String())
			}
			return out
		}
		t.Run(tC.desc, func(t *testing.T) {
			got,err := lex(tC.input)
			require.NoError(t, err)
			
			assert.Equal(t, stringify(tC.expected), stringify(got))
		})
	}
}

func TestInvalidInput(t *testing.T) {
	t.Run("invalid string", func(t *testing.T) {
		input := `" hello world `

		_, err := lex(input)
		assert.Error(t, err)
		assert.Equal(t, "Invalid token at line 1: \" hello world ", err.Error())
	})
}