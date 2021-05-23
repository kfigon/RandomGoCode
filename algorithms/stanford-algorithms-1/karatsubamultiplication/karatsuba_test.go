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
		// {"","2","2"},
		// {"2","","2"},
		// {"1","2","2"},
		// {"2","3","6"},
		// {"3","2","6"},
		// {"12","2","24"},
		// {"2","12","24"},
		// {"56","12","672"},
		// {"12","56","672"},
		{"15","10","150"},
		{"10","15","150"},

		{"15","100","1500"},
		{"100","15","1500"},
		// {"56","34","1904"},
		// {"34","56","1904"},	
		// {"78","12","936"},
		// {"12","78","936"},
		// {"34","78","2652"},
		// {"78","34","2652"},
		// {"1234","5678","7006652"},
		// {"5678","1234","7006652"},
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
	conv := func(val string) int64 {
		i,err := strconv.ParseInt(val,10,64)
		if err != nil {
			return 1
		}
		return i
	}

	if len(x) <= 1 || len(y) <= 1 {
		return fmt.Sprintf("%v", conv(x)*conv(y))
	}

	getTen := func(ln int) string {
		out := ""
		for i := 0; i < ln; i++ {
			out += "0"
		}
		return out
	}

	a,b := split(x)
	c,d := split(y)

	ac:=recursive(a,c)
	ad:=recursive(a,d)
	bc:=recursive(b,c)
	bd:=recursive(b,d)

	first := ac + getTen(len(x)/2+len(y)/2)
	second := ad + getTen(len(x)/2)
	third := bc + getTen(len(y)/2)
	fourth := bd
	return add(add(add(first, second), third), fourth)
}

// todo: implement addition
func add(x, y string) string {
	conv := func(val string) int64 {
		i,err := strconv.ParseInt(val,10,64)
		if err != nil {
			return 0
		}
		return i
	}
	return fmt.Sprintf("%v", (conv(x)+conv(y)))
}

func split(x string) (string, string) {
	if len(x) == 1 {
		return x,""
	}
	half := len(x)/2
	return x[:half],x[half:]
}