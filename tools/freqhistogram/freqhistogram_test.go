package freqhistogram

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// Exercise 1-14. Write a program to print a histogram of the frequencies of different characters in its input.
// It is easy to draw the histogram with the bars horizontal; a vertical orientation is more challenging.

func TestHistogram(t *testing.T) {
	input := "asd abc xxx "
	exp := map[rune]int{
		'a': 2,
		'b': 1,
		'c': 1,
		'd': 1,
		's': 1,
		'x': 3,
	}

	assert.Equal(t, exp, histogram(input))
}

func TestPrintHistogram(t *testing.T) {
	input := "aaa"
	assert.Equal(t, "a|***", print(histogram(input)))
}

func histogram(data string) map[rune]int {
	out := map[rune]int{}
	for _, c := range data {
			if unicode.IsSpace(c) {
				continue
			}
			out[c]++
	}
	return out
}

func print(hist map[rune]int) string {
	out := []string{}
	for r, cnt := range hist {
		stars := ""
		for i := 0; i < cnt; i++ {
			stars += "*"
		}
		out = append(out, string(r) + "|" + stars)
	}
	return strings.Join(out, "\n")
}