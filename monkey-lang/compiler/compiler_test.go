package compiler

import (
	"monkey-lang/lexer"
	"monkey-lang/objects"
	"monkey-lang/parser"
	"slices"
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

func TestInstructionIter(t *testing.T) {
	t.Run("two constants", func(t *testing.T) {
		ins := Instructions{}

		c, err := MakeCommand(OpConst, 2)
		require.NoError(t, err)
		ins = append(ins, c...)

		c, err = MakeCommand(OpConst, 15)
		require.NoError(t, err)
		ins = append(ins, c...)

		got := slices.Collect(ins.Iter())
		exp := [][]byte{
			{0, 0, 2},
			{0, 0, 0xf},
		}
		assert.Equal(t, exp, got)
	})	
}


func TestStack(t *testing.T) {
	s := NewStack[int]()	
	s.Push(2)
	s.Push(333)

	assert.Equal(t, 333, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.True(t, s.Empty())
}