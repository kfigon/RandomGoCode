package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		input string
		exp []token
	}{
		{
			input: "1",
			exp: []token{{number, "1"}},
		},
		{
			input: "1,2,3",
			exp: []token{{number, "1"}, {comma,","},{number,"2"}, {comma,","}, {number,"3"}},
		},
		{
			input: "*",
			exp: []token{{wildcard, "*"}},
		},
		{
			input: "/",
			exp: []token{{div, "/"}},
		},
		{
			input: "12-34",
			exp: []token{{number, "12"}, {dash,"-"},{number,"34"}},
		},
		{
			input: "12 - 34",
			exp: []token{{number, "12"}, {dash,"-"},{number,"34"}},
		},
		{
			input: "1 3",
			exp: []token{{number, "1"},{number,"3"}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got, err := tokenize(tC.input)
			assert.NoError(t, err)

			assert.Equal(t, tC.exp, got)
		})
	}
}

func TestInvalidInput(t *testing.T) {
	testCases := []string {
		"123a",
		"a123",
		"a",
		"1a3",
		"*/3a",
	}
	for _, tC := range testCases {
		t.Run(tC, func(t *testing.T) {
			_, err := tokenize(tC)
			assert.Error(t, err)
		})
	}
}