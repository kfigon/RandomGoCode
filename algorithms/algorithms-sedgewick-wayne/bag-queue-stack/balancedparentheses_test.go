package bagqueuestack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// we can check if parenthesis are balanced
func TestBalanced(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		input := "[()]{}{[()()]()}"
		assert.True(t, validateParenthesis(input))
	})

	t.Run("not ok", func(t *testing.T) {
		input := "[()]{}{()()]()}"
		assert.False(t, validateParenthesis(input))
	})

	t.Run("not ok2", func(t *testing.T) {
		input := "[(])"
		assert.False(t, validateParenthesis(input))
	})

	t.Run("not ok3", func(t *testing.T) {
		input := "[()"
		assert.False(t, validateParenthesis(input))
	})
}

func validateParenthesis(input string) bool {
	stack := &stack[rune]{}

	openingBrace := map[rune]bool {
		'{': true,
		'(': true,
		'[': true,
	}

	complementaryOpeningBrace := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range input {
		if openingBrace[v] {
			stack.push(v)
		} else {
			lastOpen, _ := stack.pop()
			if lastOpen != complementaryOpeningBrace[v] {
				return false
			}
		}
	}

	return stack.empty()
}