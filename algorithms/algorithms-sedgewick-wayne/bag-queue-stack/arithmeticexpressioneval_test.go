package bagqueuestack

import (
	"strconv"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// usage of stacks - Dijkstra's two-stack algorithm for expression eval

func TestEvaluator(t *testing.T) {
	input := "(1+((2+3)*(4*5)))"
	exp := 101
	got := evaluateExpression(input)
	assert.Equal(t, exp, got)
}

func evaluateExpression(input string) int {
	ops := &stack[rune]{}
	vals := &stack[int]{}

	for _, char := range input {
		if char == '+' || char == '*' || char == '-'{
			ops.push(char)
		}else if unicode.IsDigit(char) {
			v, _ := strconv.Atoi(string(char))
			vals.push(v)
		} else if char == ')' {
			op, ok := ops.pop()
			if !ok {
				return -888
			}
			left, _ := vals.pop()
			right, _ := vals.pop()

			switch op {
			case '+': vals.push(left+right)
			case '-': vals.push(left-right)
			case '*': vals.push(left*right)
			}
		}
	}

	if vals.len() == 1 {
		v, _ := vals.pop()
		return v
	}
	return -888
}