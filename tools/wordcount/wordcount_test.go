package wordcount

import (
	"strconv"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	testCases := []struct {
		input string
		exp int
	}{
		{
			input: "foo bar asd 123",
			exp: 4,
		},
		{
			input: "   foo bar asd 123   ",
			exp: 4,
		},
		{
			input: "1 foo bar asd 123 1",
			exp: 6,
		},
		{
			input: "  foo   bar   asd   123  1  ",
			exp: 5,
		},
	}
	for i, tC := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tC.exp, wordCount(tC.input), "invalid res for %q", tC.input)
		})
	}
}

func wordCount(data string) int {
	cnt := 0
	inWord := false
	for _, c := range data {
		if inWord && unicode.IsSpace(c) {
			cnt++
			inWord = false
		} else if !unicode.IsSpace(c) {
			inWord = true
		}
	}
	if inWord {
		cnt++
	}
	return cnt
}