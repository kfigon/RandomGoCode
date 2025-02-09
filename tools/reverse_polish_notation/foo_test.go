package reversepolishnotation

import (
	"fmt"
	"strconv"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	t.Run("1 2 - 4 5 + *", func(t *testing.T) {
		// (1 - 2) * (4 + 5)
		v, err := rpn("1 2 - 4 5 + *")
		assert.NoError(t, err)
		assert.Equal(t, -9, v)
	})
	
	t.Run("100 2 -", func(t *testing.T) {
		v, err := rpn("100 2 -")
		assert.NoError(t, err)
		assert.Equal(t, 98, v)
	})
}

func rpn(in string) (int, error) {
	stack := []int{}
	pop := func() (int, bool) {
		if len(stack) == 0 {
			return 0, false
		}
		out := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return out, true
	}

	push := func(v int) {
		stack = append(stack, v)
	}

	pendingStr := ""
	for _, c := range in {
		if unicode.IsSpace(c) && pendingStr != "" {
			i, err := strconv.Atoi(pendingStr)
			if err != nil {
				return 0, fmt.Errorf("non number found: %w", err)
			}
			push(i)
			pendingStr = ""
		} else if unicode.IsNumber(c) {
			pendingStr += string(c)
		} else if c == '+' || c == '-' || c =='/' || c =='*' {
			a, _ := pop()
			b, _ := pop()
			switch c {
			case '+':
				push(a+b)
			case '-':
				push(b-a)
			case '*':
				push(b*a)
			case '/':
				if a == 0 {
					return 0, fmt.Errorf("divide by 0")
				}
				push(b/a)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("unterminated operators %v", stack)
	}
	out, _ := pop()
	return out, nil
}