package constructnote

import (
	"fmt"
	"testing"
)

// # write a fun that accepts 2 strings - message and some letters
// # fun should return true if the message can be built with the letters
// # that you are given

// # assume lowercase, not special chars and no spaces

// # space - O(N)
// # time - O(N+M)

func createMap(in string) map[string]int {
	dict := make(map[string]int)
	for _,c := range in {
		dict[string(c)]++
	}
	return dict
}

func constructNote(message, letters string) bool {
	messageMap := createMap(message)
	availableLetters := createMap(letters)

	for c := range messageMap {
		occurencesInMessage := messageMap[c]
		occurencesInLetters := availableLetters[c]
		if occurencesInMessage > occurencesInLetters {
			return false
		}
	}

	return true
}

func Test(t *testing.T) {
	testCases := []struct {
		in1, in2 string
		exp      bool
	}{
		{in1: "aa", in2: "abc", exp: false},
		{in1: "abc", in2: "dcba", exp: true},
		{in1: "aabbcc", in2: "bcabcaddff", exp: true},
		{in1: "kamil", in2: "kamil", exp: true},
		{in1: "kaamil", in2: "kamil", exp: false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q-%q", tc.in1, tc.in2), func(t *testing.T) {
			if res := constructNote(tc.in1, tc.in2); res != tc.exp {
				t.Errorf("got %v, exp %v", res, tc.exp)
			}
		})
	}
}
