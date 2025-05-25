package compiler

import (
	"monkey-lang/lexer"
	"monkey-lang/objects"
	"monkey-lang/parser"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompiler(t *testing.T) {
	testCases := []struct {
		desc	string
		input string
		expectedInstr Instructions
		expConstants []objects.Object
	}{
		{
			desc: "simple literal",
			input: "3",
			expectedInstr: []byte{byte(OpConst), 0,0},
			expConstants: []objects.Object{&objects.PrimitiveObj[int]{3}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree, err := parser.Parse(lexer.Lex(tC.input))
			require.NoError(t, err, "parsing err")

			instr, constants, err := Compile(tree)
			require.NoError(t, err, "compiler err")

			assert.Equal(t, tC.expectedInstr, instr)
			assert.Equal(t, tC.expConstants, constants)
		})
	}
}