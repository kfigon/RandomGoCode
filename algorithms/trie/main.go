package main

import (
	"fmt"
)

func main() {
	t := &trie{}

	words := []string{"hi", "hello", "hell", "howdy", "asd", "as"}
	for _, v := range words {
		t.add(v)
	}

	fmt.Println("h", t.suggestions("h"))
	fmt.Println("hi", t.suggestions("hi"))
	fmt.Println("he", t.suggestions("he"))
	fmt.Println("hell", t.suggestions("hell"))
	fmt.Println("hello", t.suggestions("hello"))
	fmt.Println("as", t.suggestions("as"))
}

type node struct {
	children map[rune]*node
	isWord bool
}

func newNode() *node {
	return &node{
		children: map[rune]*node{},
	}
}

type trie struct {
	root *node
}

func (t *trie) add(text string) {
	if t.root == nil {
		t.root = newNode()
	}

	ptr := t.root
	for i, c := range text {
		if v,ok := ptr.children[c]; ok {
			ptr = v
		} else {
			next := newNode()
			ptr.children[c] = next
			ptr = next
		}

		if i == len(text)-1 {
			ptr.isWord = true
		}
	}
}

func (t *trie) suggestions(prefix string) []string {
	var out []string
	var traverse func(prefixStr string, n *node)
	traverse = func(prefixStr string, n *node) {
		if n == nil {
			return
		}
		if n.isWord {
			out = append(out, prefixStr)
		}
		for c, children := range n.children {
			traverse(prefixStr+string(c), children)
		}
	}

	lastCommonNode := t.root
	for _, c := range prefix {
		v, ok := lastCommonNode.children[c]
		if !ok {
			return []string{}
		}
		lastCommonNode = v
	}
	traverse(prefix, lastCommonNode)
	return out
}