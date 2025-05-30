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
		{
			desc: "infix",
			input: "3 + 1",
			expectedInstr: []byte{
				byte(OpConst), 0,0,
				byte(OpConst), 0,1,
				byte(OpAdd),
			},
			expConstants: []objects.Object{
				&objects.PrimitiveObj[int]{3},
				&objects.PrimitiveObj[int]{1},
			},
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
	t.Run("two constants and add", func(t *testing.T) {
		ins := Instructions{}

		c, err := MakeCommand(OpConst, 2)
		require.NoError(t, err)
		ins = append(ins, c...)

		c,err = MakeCommand(OpAdd)
		require.NoError(t, err)
		ins = append(ins, c...)

		c, err = MakeCommand(OpConst, 15)
		require.NoError(t, err)
		ins = append(ins, c...)

		got := slices.Collect(ins.Iter())
		exp := [][]byte{
			{0, 0, 2},
			{1},
			{0, 0, 0xf},
		}
		assert.Equal(t, exp, got)
	})	

}


func TestStack(t *testing.T) {
	t.Run("push twice, pop twice", func(t *testing.T) {
		s := NewStack[int]()	
		s.Push(2)
		s.Push(333)

		assert.Equal(t, 333, s.Pop())
		assert.Equal(t, 2, s.Pop())
		assert.True(t, s.Empty())
	})

	t.Run("drain", func(t *testing.T) {
		s := NewStack[int]()	
		s.Push(3)
		s.Push(2)
		assert.Equal(t, 2, s.Pop())
		s.Push(1)

		exp := []int{1,3}
		got := []int{}
		for !s.Empty() {
			got = append(got, s.Pop())
		}

		assert.Equal(t, exp, got)
	})
}