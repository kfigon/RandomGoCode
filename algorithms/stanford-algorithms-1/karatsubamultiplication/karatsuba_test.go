package karatsubamultiplication

import (
	"strconv"
	"fmt"
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
		{"12","2","24"},
		{"2","12","24"},
		{"1234","5678","7006652"},
		{"5678","1234","7006652"},
	}
}

func (m multiplicationInput) desc() string { return m.x + "*" + m.y }

func TestKaratsuba(t *testing.T) {
	testCases := getInput()
	for _, tc := range testCases {
		t.Run(tc.desc(), func(t *testing.T) {
			assert.Equal(t, tc.exp, karatsuba(tc.x,tc.y))
		})
	}
}

func TestRecursive(t *testing.T) {
	testCases := getInput()
	for _, tc := range testCases {
		t.Run(tc.desc(), func(t *testing.T) {
			assert.Equal(t, tc.exp, recursive(tc.x,tc.y))
		})
	}
}

func karatsuba(x string, y string) string {
	return "0"
}

func recursive(x string, y string) string {
	if len(x) == 1 || len(x) == 1{
		xi,_ := strconv.Atoi(x)
		yi,_ := strconv.Atoi(y)
		return fmt.Sprintf("%v", xi*yi)
	}
	return "0"
}