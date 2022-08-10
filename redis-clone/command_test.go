package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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