package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	testCases := []struct {
		input	string
		exp		[]string
	}{
		{"H", []string{}},
		{"asd", []string{"asd"}},
		{"as", []string{"asd", "as"}},
		{"a", []string{"asd", "as"}},
		{"", []string{"hi", "hello", "hell", "howdy", "asd", "as"}},
		{" ", []string{}},
		{"asdasd", []string{}},
		{"asdd", []string{}},
		{"h", []string{"hi", "hello", "hell", "howdy"}},
		{"hell", []string{"hello", "hell"}},
		{"ho", []string{"howdy"}},
		{"hi", []string{"hi"}},
	}
	words := []string{"hi", "hello", "hell", "howdy", "asd", "as"}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			tr := create(words)
			assert.ElementsMatch(t, tr.suggestions(tC.input), tC.exp)
		})
	}
}

func TestOrderOfAddingDoesNotMatter(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		tr := create([]string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("he"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hell"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hello"), []string{"hello"})
	})
	
	t.Run("Seconds", func(t *testing.T) {
		tr := create([]string{"hell", "hello"})
		assert.ElementsMatch(t, tr.suggestions("he"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hell"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hello"), []string{"hello"})
	})
}

func create(words []string) *trie {
	t := &trie{}
	for _, v := range words {
		t.add(v)
	}
	return t
}