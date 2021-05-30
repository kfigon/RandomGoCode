package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testCases := []struct {
		n Note
		exp string
	}{
		{A, "A"},
		{Ab, "Ab"},
		{G, "G"},
		{Note(123), ""},
	}
	for _, tC := range testCases {
		t.Run(tC.exp, func(t *testing.T) {
			assert.Equal(t, tC.exp, (&tC.n).String())
		})
	}
}