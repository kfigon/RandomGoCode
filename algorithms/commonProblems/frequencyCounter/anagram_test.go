package frequencycounter

import (
	"testing"
	"fmt"
)

// given 2 strings - determine is the second string an anagram of 1st
// anagram - rearrange

func createMap(in string) map[string]int {
	mapa := make(map[string]int)
	for _,c := range in {
		mapa[string(c)]++
	}
	return mapa
}

func isAnagram(in1, in2 string) bool {
	map1 := createMap(in1)
	map2 := createMap(in2)

	for char := range map2 {
		occurenceInFirst := map1[char]
		occurenceInSecond := map2[char]
		if occurenceInSecond > occurenceInFirst {
			return false
		}
	}

	return true
}

func TestAnagram(t *testing.T) {
	testCases := []struct {
		in1 string
		in2 string
		exp bool		
	}{
		{ in1:"", in2:"", exp: true },
		{ in1: "aaz", in2: "zza", exp: false },
		{ in1: "anagram", in2: "nagaram", exp: true },
		{ in1: "cinema", in2: "iceman", exp: true },
		{ in1: "abce", in2: "abcea", exp: false },
		{ in1: "abcea", in2: "abce", exp: true },
		{ in1: "cat", in2: "cat", exp: true },
		{ in1: "rat", in2: "car", exp: false },
		{ in1: "awesome", in2: "awesom", exp: true },
		{ in1: "awesome", in2: "awesme", exp: true },
		{ in1: "awesme", in2:"awesome" , exp: false },
		{ in1: "qwerty", in2: "qeywrt", exp: true },
		{ in1: "texttwisttime", in2: "timetwisttext", exp: true },
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q-%q", tc.in1, tc.in2), func(t *testing.T) {
			if res := isAnagram(tc.in1, tc.in2); res != tc.exp {
				t.Errorf("anagram of %q and %q, res %v != exp %v", tc.in1, tc.in2, res, tc.exp)
			}
		})
	}
}