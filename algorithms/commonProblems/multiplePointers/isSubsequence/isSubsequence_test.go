package isSubsequence

import (
	"fmt"
	"testing"
)

// # write a function which takes in two strings and checks whether 
// # the characters in the first string form a subsequence 
// # of the characters in the second string. 
// # 
// # In other words, the function should check whether the characters
// # in the first string appear somewhere in the second string, 

func isSubsequence(toFind, bigString string) bool {
	searchingPtr := 0
	stringPtr := 0

	for searchingPtr < len(toFind) && stringPtr < len(bigString) {
		if bigString[stringPtr] == toFind[searchingPtr] {
			searchingPtr++
		}
		stringPtr++
	}
	return searchingPtr >= len(toFind)
}

func Test(t *testing.T) {
	testCases := []struct {
		in1 string
		in2 string
		exp bool
		
	}{
		{in1: "hello", in2: "hello world", exp: true },
		{in1: "sing", in2: "sting", exp: true },
		{in1: "abc", in2: "abracadabra", exp: true},
		{in1: "abc", in2: "acb", exp: false}, //# order matters
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q-%q", tc.in1, tc.in2), func(t *testing.T) {
			if res := isSubsequence(tc.in1, tc.in2); res != tc.exp {
				t.Errorf("Exp %v, got %v", tc.exp, res)
			}
		})
	}
}