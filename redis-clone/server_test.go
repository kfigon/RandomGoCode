package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSimpleString(t *testing.T) {
	testPort := 6666
	go startServer(testPort)
	resp, err := sendData(testPort, []byte("+some string\r\n"))
	assert.NoError(t, err)

	assert.Equal(t, "+ok\r\n", string(resp))
}