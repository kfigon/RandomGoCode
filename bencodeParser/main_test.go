package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncoder(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "0:", encodeStr(""))
		assert.Equal(t, "3:yes", encodeStr("yes"))
		assert.Equal(t, "8:a string", encodeStr("a string"))
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, "i59e", encodeInt(59))
	})

	t.Run("list", func(t *testing.T) {
		assert.Equal(t, "le", encodeList([]any{}))
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
		assert.Equal(t, "de", encodeDict(map[string]any{}))
		assert.Equal(t, "d3:bar4:spam3:fooi42ee", encodeDict(map[string]any{"bar": "spam", "foo": 42}))
		assert.Equal(t, "d3:bar4:spam3:fooi42ee", encodeDict(map[string]any{"foo": 42, "bar": "spam"}))
		assert.Equal(t, "d3:bar3:asd3:fool4:spami42eee", encodeDict(map[string]any{
			"bar": "asd", 
			"foo": []any{"spam",42},
		}))
	})
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		desc	string
		input	string
		expected	bencodeObj
	}{
		{
			desc: "str1",
			input: "3:yes",
			expected: stringObj("yes"),
		},
		{
			desc: "str2",
			input: "8:a string",
			expected: stringObj("a string"),
		},
		{
			desc: "str3",
			input: "18:a very long string",
			expected: stringObj("a very long string"),
		},
		{
			desc: "int",
			input: "i59e",
			expected: intObj(59),
		},
		{
			desc: "string list",
			input: "l3:yes18:a very long stringe",
			expected: listObj{stringObj("yes"), stringObj("a very long string")},
		},
		{
			desc: "int list",
			input: "li2ei58ee",
			expected: listObj{intObj(2),intObj(58)},
		},
		{
			desc: "mixed list",
			input: "l4:spami42ee",
			expected: listObj{stringObj("spam"),intObj(42)},
		},
		{
			desc: "mixed list with int tags in string",
			input: "l4:asdfi42ee",
			expected: listObj{stringObj("asdf"),intObj(42)},
		},
		{
			desc: "mixed list with int tags 2",
			input: "l4:asdii42ee",
			expected: listObj{stringObj("asdi"),intObj(42)},
		},
		{
			desc: "complicated list",
			input: "l2:hili1ei2eed4:barz3:asd3:fooi58eei42ee",
			expected: listObj{
				stringObj("hi"),
				listObj{intObj(1),intObj(2)},
				dictObj(map[string]bencodeObj{"foo": intObj(58), "barz": stringObj("asd")}),
				intObj(42) },
		},
		{
			desc: "simple dict",
			input: "d3:bar4:spam3:fooi42ee",
			expected: dictObj(map[string]bencodeObj{"bar": stringObj("spam"), "foo": intObj(42)}),
		},
		{
			desc: "simple dict with reverse order",
			input: "d3:bar4:spam3:fooi42ee",
			expected: dictObj(map[string]bencodeObj{"foo": intObj(42), "bar": stringObj("spam")}),
		},
		{
			desc: "nested dicts",
			input: "d3:bar3:asd3:fool4:spami42eee",
			expected: dictObj(map[string]bencodeObj{
				"bar": stringObj("asd"), 
				"foo": listObj{stringObj("spam"), intObj(42)},
			}),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assertResult(t, tC.input, tC.expected)
		})
	}
}

func TestDecodeInvalid(t *testing.T) {
	testCases := []struct {
		desc	string
		input	string
	}{
		{
			desc: "string without len",
			input: ":foo",
		},
		{
			desc: "string without colon",
			input: "3foo",
		},
		{
			desc: "too short str",
			input: "4:foo",
		},
		{
			desc: "no string",
			input: "0:",
		},{
			desc: "no string but declared",
			input: "2:",
		},
		{
			desc: "invalid len str",
			input: "4asd:foo",
		},
		// todo
		// {
		// 	desc: "too long str",
		// 	input: "2:foo",
		// },
		{
			desc: "list not terminated",
			input: "i5",
		},
		{
			desc: "int without number",
			input: "ifoobare",
		},
		{
			desc: "list not terminated",
			input: "l2",
		},
		{
			desc: "list with invalid str len",
			input: "l4:asdi4ee",
		},
		{
			desc: "dit not terminated",
			input: "d3:asdi4e",
		},
		{
			desc: "dict without value",
			input: "d3:asde",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := decode(tC.input)
			assert.Error(t, err)
		})
	}
}

func assertResult[T bencodeObj](t *testing.T, input string, expected T) {
	obj,err := decode(input)
	require.NoError(t, err)

	c, ok := obj.(T)
	require.True(t, ok)

	assert.Equal(t, expected, c)
}