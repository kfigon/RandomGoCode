package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPort = 6666

func TestSendSimpleString(t *testing.T) {
	go startServer(testPort)
	resp, err := sendData(testPort, []byte("+some string\r\n"))
	assert.NoError(t, err)

	assert.Equal(t, "+ok\r\n", string(resp))
}

func TestRespCommands(t *testing.T) {
	testCases := []struct {
		desc	string
		input   []byte
		expectedOut   []byte
	}{
		{
			desc: "Ping",
			input: []byte("*1\r\n$4\r\nPING\r\n"),
			expectedOut: []byte("+PONG\r\n"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			go startServer(testPort)
			resp, err := sendData(testPort, tC.input)
			assert.NoError(t, err)
		
			assert.Equal(t, tC.expectedOut, resp)
		})
	}
}