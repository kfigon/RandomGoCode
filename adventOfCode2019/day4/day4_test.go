package day4

import (
	// "strconv"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/4
func TestPasswordCheck(t *testing.T) {
	testCases := []struct {
		in string	
		exp bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got := password(tc.in).isOk()
			assert.Equal(t, tc.exp, got)
		})
	}
}

type password string
func (p password) isOk() bool {
	return false
}