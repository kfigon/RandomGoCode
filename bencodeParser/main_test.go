package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestEncoder(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "3:yes", encodeStr("yes"))
		assert.Equal(t, "8:a string", encodeStr("a string"))
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, "i59e", encodeInt(59))
	})

	t.Run("list", func(t *testing.T) {
		assert.Equal(t, "li2ei58ee", encodeList([]any{2,58}))
		assert.Equal(t, "l4:spami42ee", encodeList([]any{"spam",42}))
		
		assert.Equal(t, "l2:hili1ei2eed3:fooi58e4:barz3:asdei42ee", encodeList([]any{
			"hi",
			[]any{1,2},
			map[string]any{"foo": 58, "barz": "asd"},
			42 },
		))
	})

	t.Run("dict", func(t *testing.T) {
		assert.Equal(t, "d3:bar4:spam3:fooi42ee", encodeDict(map[string]any{"bar": "spam", "foo": 42}))
		assert.Equal(t, "d3:bar3:asd3:foo4:spami42ee", encodeDict(map[string]any{"bar": "asd", "foo": []any{"spam",42}}))
	})
}