package karatsubamultiplication

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestKaratsuba(t *testing.T) {
	testCases := []struct {
		x int
		y int
		exp int
	}{
		{1,2,2},
		{2,3,6},
		{3,2,6},
		{1234,5678,7006652},
		{5678,1234,7006652},
	}
	for _, tc := range testCases {
		title := fmt.Sprintf("%v*%v", tc.x,tc.y)
		t.Run(title, func(t *testing.T) {
			got := karatsuba(tc.x,tc.y)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func karatsuba(x int, y int) int {
	return 0
}