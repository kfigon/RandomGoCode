package karatsubamultiplication

import (
	"strconv"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

type multiplicationInput struct {
	x string
	y string
	exp string
}

func getInput() []multiplicationInput {
	return []multiplicationInput{
		{"","2","2"},
		{"2","","2"},
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

func TestSplit(t *testing.T) {
	testCases := []struct {
		in string
		a string
		b string
	}{
		{"1234", "12", "34"},
		{"12345", "12", "345"},
		{"1", "1", ""},
		{"18", "1", "8"},
	}
	for _, tC := range testCases {
		t.Run(tC.in, func(t *testing.T) {
			a,b := split(tC.in)
			assert.Equal(t, tC.a,a)
			assert.Equal(t, tC.b,b)
		})
	}
}

func karatsuba(x string, y string) string {
	return "0"
}

func recursive(x string, y string) string {
	conv := func(val string) int {
		i,err := strconv.Atoi(val)
		if err != nil {
			return 1
		}
		return i
	}

	if len(x) <= 1 || len(y) <= 1 {
		return fmt.Sprintf("%v", conv(x)*conv(y))
	}
	a,b := split(x)
	c,d := split(y)

	tenN := fmt.Sprintf("%v", math.Pow10(len(x)))
	tenN2 := fmt.Sprintf("%v", math.Pow10(len(x)/2))

	first := recursive(tenN, recursive(a,c))
	second := recursive(tenN2, add(recursive(a,d), recursive(b,c)))
	third := recursive(b,d)
	result := add(add(first, second), third)
	return result
}

// todo: implement addition
func add(x, y string) string {
	conv := func(val string) int {
		i,err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return i
	}
	return fmt.Sprintf("%v", conv(x)+conv(y))
}

func split(x string) (string, string) {
	if len(x) == 1 {
		return x,""
	}
	half := len(x)/2
	return x[:half],x[half:]
}