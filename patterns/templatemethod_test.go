package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func doStuff(in string, fn func() string) string {
	in += " <tag>"
	in += fn()
	in += "</tag>"
	return in
}

func TestTemplateMethod(t *testing.T) {
	result := doStuff("hello", func() string {
		return "ZIOM"
	})

	assert.Equal(t, "hello <tag>ZIOM</tag>", result)
}