package main

import (
	"testing"
)

type randerFn func () bool
func (r randerFn) pass() bool {
	return r()
}

func Test(t *testing.T) {
	testCases := []struct {
		input		string
		expected	string
		
	}{
		{
			input: "foo raz dwa",
			expected: "foo raz dwa",
		},
		{
			input: `asd
dwa trzy
   
			tab
`,
			expected: `asd
dwa trzy
   
			tab
`,
		},
		{
			input: `cześć zależność napis`,
			expected: `cześć zależność napis`,
		},
		{
			input: ` cześć zależność napis   `,
			expected: ` cześć zależność napis   `,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			var r randerFn
			r = func()bool {
				return false
			}
			got := processInput(tC.input, r)

			if got != tC.expected {
				t.Errorf("Got %q, exp %q\n", got, tC.expected)
			}
		})
	}
}