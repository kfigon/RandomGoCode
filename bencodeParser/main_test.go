package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		
		assert.Equal(t, "l2:hili1ei2eed4:barz3:asd3:fooi58eei42ee", encodeList([]any{
			"hi",
			[]any{1,2},
			map[string]any{"foo": 58, "barz": "asd"},
			42 },
		))
	})

	t.Run("dict", func(t *testing.T) {
		assert.Equal(t, "d3:bar4:spam3:fooi42ee", encodeDict(map[string]any{"bar": "spam", "foo": 42}))
		assert.Equal(t, "d3:bar4:spam3:fooi42ee", encodeDict(map[string]any{"foo": 42, "bar": "spam"}))
		assert.Equal(t, "d3:bar3:asd3:fool4:spami42eee", encodeDict(map[string]any{
			"bar": "asd", 
			"foo": []any{"spam",42},
		}))
	})
}

func TestDecode(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		assertResult[stringObj, string](t, "3:yes", "yes")
		assertResult[stringObj, string](t, "8:a string", "a string")
	})

	t.Run("int", func(t *testing.T) {
		assertResult[intObj, int](t, "i59e", 59)
	})

	t.Run("list", func(t *testing.T) {
		assertResult[listObj, []any](t, "li2ei58ee", []any{2,58})
		assertResult[listObj, []any](t, "l4:spami42ee", []any{"spam",42})
		
		assertResult[listObj, []any](t, "l2:hili1ei2eed4:barz3:asd3:fooi58eei42ee", []any{
			"hi",
			[]any{1,2},
			map[string]any{"foo": 58, "barz": "asd"},
			42 },
		)
	})

	t.Run("dict", func(t *testing.T) {
		assertResult[dictObj, map[string]any](t, "d3:bar4:spam3:fooi42ee", map[string]any{"bar": "spam", "foo": 42})
		assertResult[dictObj, map[string]any](t, "d3:bar4:spam3:fooi42ee", map[string]any{"foo": 42, "bar": "spam"})
		assertResult[dictObj, map[string]any](t, "d3:bar3:asd3:fool4:spami42eee", map[string]any{
			"bar": "asd", 
			"foo": []any{"spam",42},
		})
	})
}

func assertResult[T bencodeObj, K any](t *testing.T, input string, expected K) {
	obj := decode(input)
	c, ok := obj.(T)
	require.True(t, ok)

	assert.Equal(t, expected, c)
}