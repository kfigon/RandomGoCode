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
			desc: "basic tokens",
			input: `(define somevalue 10)
			(+ 3 (* somevalue somevalue))`,
			expected: []token{
				{ opening, "("},
				{ keyword, "define"},
				{ identifier, "somevalue"},
				{ number, "10"},
				{ closing, ")"},
				{ opening, "("},
				{ operator, "+"},
				{ number, "3"},
				{ opening, "("},
				{ operator, "*"},
				{ identifier, "somevalue"},
				{ identifier, "somevalue"},
				{ closing, ")"},
				{ closing, ")"},
		},
		},

		{
			desc: "identifiers",
			input: ` somevalue definee`,
			expected: []token{
				{ identifier, "somevalue"},
				{ identifier, "definee"},
			},
		},
		{
			desc: "number",
			input: ` 1234`,
			expected: []token{
				{number, "1234"},
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
			desc: "if statement",
			input: `(define apples 5)
			(define oranges 6)
			(if (<= apples oranges)
				(printf "Apples")
				(printf "Oranges"))`,
			expected: []token{
				{ opening, "("},
				{ keyword, "define"},
				{ identifier, "apples"},
				{ number, "5"},
				{ closing, ")"},
				{ opening, "("},
				{ keyword, "define"},
				{ identifier, "oranges"},
				{ number, "6"},
				{ closing, ")"},
				{ opening, "("},
				{ keyword, "if"},
				{ opening, "("},
				{ operator, "<="},
				{ identifier, "apples"},
				{ identifier, "oranges"},
				{ closing, ")"},
				{ opening, "("},
				{ identifier, "printf"},
				{ stringLiteral, "\"Apples\""},
				{ closing, ")"},
				{ opening, "("},
				{ identifier, "printf"},
				{ stringLiteral, "\"Oranges\""},
				{ closing, ")"},
				{ closing, ")"},
			},
		},
		{
			desc: "operators",
			input: `< <= > >= ! !! !=`,
			expected: []token{
				{ operator, "<"},
				{ operator, "<="},
				{ operator, ">"},
				{ operator, ">="},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!="},
			},
		},
		{
			desc: "operators without spaces",
			input: `<<=>>=!!!!=`,
			expected: []token{
				{ operator, "<"},
				{ operator, "<="},
				{ operator, ">"},
				{ operator, ">="},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!"},
				{ operator, "!="},
			},
		},
		{
			desc: "function",
			input: `(define (dbl x)
			(* 2 x))

			(dbl 2)`,
			expected: []token{
				{ opening,"(" },
				{ keyword,"define" },
				{ opening,"(" },
				{ identifier,"dbl" },
				{ identifier,"x" },
				{ closing,")" },
				{ opening,"(" },
				{ operator,"*" },
				{ number, "2"},
				{ identifier,"x" },
				{ closing,")" },
				{ closing,")" },
				{ opening,"(" },
				{ identifier,"dbl" },
				{ number,"2"},
				{ closing,")" },
			},
		},
		{
			desc: "boolean",
			input: `(define x true)
			(define y false)
			(define z(= x y))`,
			expected: []token{
				{ opening, "(" },
				{ keyword, "define" },
				{ identifier, "x" },
				{ boolean, "true"},
				{ closing, ")" },
				{ opening, "(" },
				{ keyword, "define" },
				{ identifier, "y" },
				{ boolean, "false"},
				{ closing, ")" },
				{ opening, "(" },
				{ keyword, "define" },
				{ identifier, "z" } ,
				{ opening, "(" },
				{ operator, "=" },
				{ identifier, "x" },
				{ identifier, "y" },
				{ closing, ")" },
				{ closing, ")" },
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got,err := lex(tC.input)
			require.NoError(t, err)

			assert.Equal(t, tC.expected, got)
		})
	}
}

func TestInvalidInput(t *testing.T) {
	t.Run("invalid string literal on line 2", func(t *testing.T) {
		input := ` (define x 3)
        (define s " hello world `

		_, err := lex(input)
		assert.Error(t, err)
		assert.Equal(t, "Invalid token at line 2: \" hello world ", err.Error())
	}) 

	t.Run("invalid string", func(t *testing.T) {
		input := `" hello world `

		_, err := lex(input)
		assert.Error(t, err)
		assert.Equal(t, "Invalid token at line 1: \" hello world ", err.Error())
	})
}