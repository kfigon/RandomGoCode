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
