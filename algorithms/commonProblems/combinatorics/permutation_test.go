package combinatorics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	testCases := []struct {
		input []int
		expected [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1,2}, [][]int{{1,2}, {2,1}}},
		{[]int{1,2,3}, [][]int{{1,2,3}, {1,3,2}, {3,2,1}, {3,1,2}, {2,1,3}, {2,3,1} }},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v", tC.input), func(t *testing.T) {
			out := permutation(tC.input)
			assert.ElementsMatch(t, out, tC.expected)
		})
	}
}

func permutation[T any](input []T) [][]T {
	var out [][]T
	var fn func([]T, []T)
	fn = func(remaining, pending []T) {
		if len(remaining) == 0 {
			out = append(out, pending)
			return
		}
		first := remaining[0]
		rest := remaining[1:]

		for i := 0; i < len(pending); i++ {
			var asdf []T
			asdf = append(asdf, pending[:i]...)
			asdf = append(asdf, first)
			asdf = append(asdf, pending[i:]...)
			fn(rest, asdf)
		}
		fn(rest, append(pending, first))
	}

	fn(input, []T{})
	return out
}