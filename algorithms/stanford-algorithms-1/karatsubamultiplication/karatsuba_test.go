package karatsubamultiplication

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type multiplicationInput struct {
	x string
	y string
	exp string
}

func getInput() []multiplicationInput {
	return []multiplicationInput{
		{"1","2","2"},
		{"2","3","6"},
		{"3","2","6"},
		{"1234","5678","7006652"},
		{"5678","1234","7006652"},
	}
}

func (m multiplicationInput) testName() string {
	return m.x + "*" + m.y
}

func TestKaratsuba(t *testing.T) {
	testCases := getInput()
	for _, tc := range testCases {
		t.Run(tc.testName(), func(t *testing.T) {
			got := karatsuba(tc.x,tc.y)
			assert.Equal(t, tc.exp, got)
		})
	}
}


func TestRecursive(t *testing.T) {
	testCases := getInput()
	for _, tc := range testCases {
		t.Run(tc.testName(), func(t *testing.T) {
			got := recursive(tc.x,tc.y)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func karatsuba(x string, y string) string {
	return "0"
}

func recursive(x string, y string) string {
	return "0"
}