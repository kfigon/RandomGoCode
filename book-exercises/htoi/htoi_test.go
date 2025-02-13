package htoi

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Exercise 2-3. Write the function htoi(s), which converts a string of hexadecimal digits
// (including an optional 0x or 0X) into its equivalent integer value.
// The allowable digits are 0 through 9, a through f, and A through F.

func TestHtoi(t *testing.T) {
	testCases := []struct {
		input	string
		expected int
		wantErr bool
	}{
		{ input: "0x0", expected: 0 },
		{ input: "0X0", expected: 0 },
		{ input: "0", expected: 0 },
		{ input: "abc", expected: 0xabc },
		{ input: "0x123abc", expected: 0x123abc },
		{ input: "0x123ABC", expected: 0x123abc },
		{ input: "0x123AbC", expected: 0x123abc },
		{ input: "123AbC", expected: 0x123abc },
		{ input: "0X123abC", expected: 0x123abc },
		{ input: "0X123abCf", expected: 0x123abcf },
		{ input: "10X123abCf", expected: 0, wantErr: true },
		{ input: "0X123abCg", expected: 0, wantErr: true },
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got, err := htoi(tC.input)
			if tC.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.expected, got, "%x != %x", tC.expected, got)
			}
		})
	}
}

func assertHex(t *testing.T, i int, prefix string) {
	t.Helper()

	got, err := htoi(fmt.Sprintf("%s%x", prefix, i))
	assert.NoError(t, err)
	assert.Equal(t, i, got, "%s%x != %s%x", prefix, i , prefix, got)
}

func TestStressHex(t *testing.T) {
	for i := 0; i <= 1_000_000; i++ {
		assertHex(t, i, "")
		assertHex(t, i, "0x")
	}
}

func htoi(s string) (int, error) {
	const (
		parsing int = iota
		xFound
		zeroXFound
	)
	num := 0
	state := parsing
	pow := 0
	for i := len(s)-1; i >= 0; i-- {
		c := s[i]
		switch state {
		case parsing:
			if c == 'x' || c == 'X' {
				state = xFound
			} else {
				v := 0
				switch c {
				case '0': v = 0
				case '1': v = 1
				case '2': v = 2
				case '3': v = 3
				case '4': v = 4
				case '5': v = 5
				case '6': v = 6
				case '7': v = 7
				case '8': v = 8
				case '9': v = 9
				case 'a','A': v = 10
				case 'b','B': v = 11
				case 'c','C': v = 12
				case 'd','D': v = 13
				case 'e','E': v = 14
				case 'f','F': v = 15
				default:
					return 0, fmt.Errorf("invalid character %c", c)
				}
				num += int(math.Pow(16.0, float64(pow)))*v
			}
		case xFound:
			if c == '0' {
				state = zeroXFound
			} else {
				return 0, fmt.Errorf("unrecognized character after x: %c", c)
			}
		case zeroXFound:
			return 0, fmt.Errorf("unrecognized character after 0x: %c", c)
		}
		pow++
	}
	return num, nil
}