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
		{"", []string{}},
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
			tr := createTrie(words)
			assert.ElementsMatch(t, tr.suggestions(tC.input), tC.exp)
		})
	}
}

func TestOrderOfAddingDoesNotMatter(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		tr := createTrie([]string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("he"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hell"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hello"), []string{"hello"})
	})
	
	t.Run("Seconds", func(t *testing.T) {
		tr := createTrie([]string{"hell", "hello"})
		assert.ElementsMatch(t, tr.suggestions("he"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hell"), []string{"hello", "hell"})
		assert.ElementsMatch(t, tr.suggestions("hello"), []string{"hello"})
	})
}

func createTrie(words []string) *trie {
	t := &trie{}
	for _, v := range words {
		t.add(v)
	}
	return t
}