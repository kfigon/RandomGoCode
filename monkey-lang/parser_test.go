package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		desc	string
		input string
		expected []Statement
		wantErr bool
	}{
		{
			desc: "basic let stmt",
			input: `let foobar = 234;`,
			expected: []Statement{
				&LetStatement{
					&IdentifierExpression{"foobar"},
					nil, 
				},
			},
			wantErr: false,
		},
		{
			desc: "return stmt",
			input: `return 234;`,
			expected: []Statement{
				&ReturnStatement{ nil },
			},
			wantErr: false,
		},
		{
			desc: "return stmt with identifier",
			input: `return foobar;`,
			expected: []Statement{
				&ReturnStatement{ nil },
			},
			wantErr: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := Parse(Lex(tC.input))	
			if tC.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.expected, got)
			}
		})
	}
}