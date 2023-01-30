package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	string
		expected string
	}{
		{
			desc: "",
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			toks,err := lex(tC.input)
			require.NoError(t, err, "got lexer error")

			p := NewParser(toks)
			p.Parse()
			got, errs := p.Parse()
			require.Empty(t, errs, "got parser errors")

			assert.Equal(t, got, "todo")
		})
	}
}

func TestParserErrors(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	string
		expected string
	}{
		{
			desc: "",
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			toks,err := lex(tC.input)
			require.NoError(t, err, "got lexer error")

			p := NewParser(toks)
			p.Parse()
			_, errs := p.Parse()
			require.NotEmpty(t, errs, "expected parser errors")
		})
	}
}