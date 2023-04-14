package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		input string
		desc string
		exp []int
	}{
		{
			input: `5`,
			desc: "just number",
			exp: []int{5},
		},
		{
			input: `*`,
			desc: "wildcard",
			exp: []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59},
		},
		{
			input: `5-10`,
			desc: "range",
			exp: []int{5,6,7,8,9,10},
		},
		{
			input: `5,6,7,8,9,10`,
			desc: "enumerated",
			exp: []int{5,6,7,8,9,10},
		},
		{
			input: `1,2,3,10,11,12`,
			desc: "enumerated with skip",
			exp: []int{1,2,3,10,11,12},
		},
		{
			input: `5-10,20-25`,
			desc: "ranges enumerated",
			exp: []int{5,6,7,8,9,10,20,21,22,23,24,25},
		},
		{
			input: `*/10`,
			desc: "wildcard with step",
			exp: []int{0,10,20,30,40,50},
		},
		{
			input: `0-20/10`,
			desc: "range with step",
			exp: []int{0,10,20},
		},
		{
			input: `0-22/10`,
			desc: "non overlapping range with step",
			exp: []int{0,10,20},
		},
		{
			input: `1,2,3,4,5,6,7,8,9,10/2`,
			desc: "enumeration with step",
			exp: []int{2,4,6,8,10},
		},
		{
			input: `5,20-30/5`,
			desc: "enumeration and ranges with step",
			exp: []int{5,20,25,30},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tokenize(tC.input)
			require.NoError(t, err, "unexpected lexer error")

			result, err := eval(got)
			require.NoError(t, err, "parser error")

			assert.Equal(t, tC.exp, result)
		})
	}
}