package week1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func multiply(a string, b string) string {
	return ""
}

func TestMultiplication(t *testing.T) {
	tdt := []struct {
		a        string
		b        string
		expected string
	}{
		{"", "", ""},
		{"1", "", ""},
		{"", "1", ""},
		{"2", "3", "6"},
		{"12", "3", "26"},
		{"3", "12", "26"},
		{"9", "12", "108"},
		{"12", "9", "108"},
		{"798654", "231456", "184853260224"},
		{"123432798654", "231456", "28569261845260224"},
		{"231456", "123432798654", "28569261845260224"},
	}
	for _, tc := range tdt {
		testName := fmt.Sprintf("%v*%v", tc.a, tc.b)
		t.Run(testName, func(t *testing.T) {
			result := multiply(tc.a, tc.b)
			assert.Equal(t, tc.expected, result)
		})
	}
}
