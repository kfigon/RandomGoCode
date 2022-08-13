package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSimpleString(t *testing.T) {
	go startServer(defaultPort)
	resp, err := sendData(defaultPort, []byte("+some string\r\n"))
	assert.NoError(t, err)

	assert.Equal(t, "+ok\r\n", string(resp))
}