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
	}{
		{
			desc: "short string",
			input: []byte("$3\r\nHEY\r\n"),
			expected: "HEY",
		},
		{
			desc: "longer string",
			input: []byte("$18\r\nHELLO WORLD my man\r\n"),
			expected: "HELLO WORLD my man",
		},
		{
			desc: "string with delimiters inside",
			input: []byte("$30\r\nTHIS CONTAINS A \r\n INSIDE IT\r\n"),
			expected: "THIS CONTAINS A \r\n INSIDE IT",
		},
		{
			desc: "Invalid length",
			input: []byte("$11\r\nHELLO WORLD my man\r\n"),
			expected: "HELLO WORLD",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd,err := newBulkString(tC.input)
			require.NoError(t, err)
			
			assert.Equal(t, tC.expected, cmd.simpleString())
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
			desc: "missing bytes",
			input: []byte("$\r\nHEY\r\n"),
			expectedError: "missing bytes",
		},
		{
			desc: "missing delimiter",
			input: []byte("$3HEY\r\n"),
			expectedError: "missing delimiter",
		},
		{
			desc: "too big size",
			input: []byte("$15\r\nHEY\r\n"),
			expectedError: "too big size",
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