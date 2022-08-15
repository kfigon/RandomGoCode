package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvalidCommands(t *testing.T) {
	testCases := []struct {
		desc string
		input []byte
	}{
		{
			desc: "too short1",
			input: []byte{},
		},
		{
			desc: "too short2",
			input: []byte("ab"),
		},
		{
			desc: "invalid cmd",
			input: []byte("^asd"),
		},
		{
			desc: "invalid termination",
			input: []byte("+asd\n"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd := command(tC.input)
			assert.Error(t, cmd.validate())
		})
	}
}

func TestValidCommands(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	[]byte
	}{
		{
			desc: "short string",
			input: []byte("+OK\r\n"),
		},
		{
			desc: "short string2",
			input: []byte{'+', 'O', 'K', 0x0D, 0x0A},
		},
		{
			desc: "short string3",
			input: []byte{'+', 'O', 'K', '\r', '\n'},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd := command(tC.input)
			assert.NoError(t, cmd.validate())
		})
	}
}

func TestParseString(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	[]byte
		exp		string
	}{
		{
			desc: "ok1",
			input: []byte("+OK\r\n"),
			exp: "OK",
		},
		{
			desc: "ok2",
			input: []byte{'+', 'O', 'K', 0x0D, 0x0A},
			exp: "OK",
		},
		{
			desc: "ok3",
			input: []byte{'+', 'O', 'K', '\r', '\n'},
			exp: "OK",
		},
		{
			desc: "some long msg",
			input: []byte("+hello world\r\n"),
			exp: "hello world",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd := command(tC.input)
			assert.NoError(t, cmd.validate())
			assert.Equal(t, tC.exp, cmd.simpleString())
		})
	}
}

func TestBulkString(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	[]byte
		expected string
		expectedByteLen int
		expectedLen int
	}{
		{
			desc: "short string",
			input: []byte("$3\r\nHEY\r\n"),
			expected: "HEY",
			expectedByteLen: 3,
			expectedLen: 9,
		},
		{
			desc: "longer string",
			input: []byte("$18\r\nHELLO WORLD my man\r\n"),
			expected: "HELLO WORLD my man",
			expectedByteLen: 18,
			expectedLen: 25,
		},
		{
			desc: "string with delimiters inside",
			input: []byte("$28\r\nTHIS CONTAINS A \r\n INSIDE IT\r\n"),
			expected: "THIS CONTAINS A \r\n INSIDE IT",
			expectedByteLen: 28,
			expectedLen: 35,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd,err := newBulkString(tC.input)
			require.NoError(t, err)
			
			assert.Equal(t, tC.expected, cmd.bulkString(), "invalid string parsed")
			assert.Equal(t, tC.expectedByteLen, cmd.byteLen, "invalid len")
			assert.Equal(t, tC.expectedLen, cmd.len(), "invalid len")
		})
	}
}

func TestInvalidBulkStrings(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	[]byte
		expectedError string
	}{
		{
			desc: "Invalid first byte",
			input: []byte("+5\r\nHEY\r\n"),
			expectedError: "invalid first byte",
		},
		{
			desc: "missing length",
			input: []byte("$\r\nHEY\r\n"),
			expectedError: "missing length",
		},
		{
			desc: "missing delimiter",
			input: []byte("$3HEY\r\n"),
			expectedError: "missing delimiter",
		},
		{
			desc: "too big size",
			input: []byte("$15\r\nHEY\r\n"),
			expectedError: "invalid length",
		},
		{
			desc: "missing termination",
			input: []byte("$3\r\nHEY"),
			expectedError: "invalid length",
		},
		{
			desc: "Too little length",
			input: []byte("$3\r\nHELLO WORLD\r\n"),
			expectedError: "invalid termination",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd,err := newBulkString([]byte(tC.input))
			assert.Nil(t, cmd)
			assert.Error(t, err)
			assert.ErrorContains(t, err, tC.expectedError)
		})
	}
}

func TestArrayCommand(t *testing.T) {
	build := func(t *testing.T, data []byte) *arrayCommand {
		arr, err := newArrayString(data)
		require.NoError(t, err)
		return arr
	}

	t.Run("simple array", func(t *testing.T) {
		data := []byte("*2\r\n" + "+OK\r\n" + "+hello world\r\n")
		arr := build(t, data)	
		
		require.Len(t, arr.commands(), 2)

		assert.True(t, arr.commands()[0].isStringCmd())
		assert.Equal(t, "OK", arr.commands()[0].simpleString())
		
		assert.True(t, arr.commands()[1].isStringCmd())
		assert.Equal(t, "hello world", arr.commands()[1].simpleString())
	})

	t.Run("array with bulk", func(t *testing.T) {
		data := []byte("*2\r\n" + "+OK\r\n" + "$28\r\nTHIS CONTAINS A \r\n INSIDE IT\r\n")
		arr := build(t, data)	
		
		require.Len(t, arr.commands(), 2)

		assert.True(t, arr.commands()[0].isStringCmd())
		assert.Equal(t, "OK", arr.commands()[0].simpleString())
		
		assert.True(t, arr.commands()[1].isBulk())
		blk, err := newBulkString(arr.commands()[1])
		require.NoError(t, err)

		assert.Equal(t, "THIS CONTAINS A \r\n INSIDE IT", blk.bulkString())
	})

	t.Run("array with bulks", func(t *testing.T) {
		t.Fatal("todo")
	})

	t.Run("many arrays", func(t *testing.T) {
		t.Fatal("todo")
	})
}

func TestInvalidArrays(t *testing.T) {
	t.Fatal("todo")
}