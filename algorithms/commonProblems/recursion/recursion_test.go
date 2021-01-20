package recursion

import (
	"fmt"
	"testing"
)

// # check if any of number is odd
func isAnyOdd(tab []int) bool {
	if len(tab) == 0 {
		return false
	}
	return (tab[0] % 2 == 1) || isAnyOdd(tab[1:])
}

func TestIsAnyOdd(t *testing.T) {
	testCases := []struct {
		in []int
		exp bool
		
	}{
		{in: []int{1,2,3,4,5}, exp: true},
		{in: []int{2,4,6,0}, exp: false},
		{in: []int{22,2,4,4,1}, exp: true},
		{in: []int{22,2,1,4,4}, exp: true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			if res := isAnyOdd(tc.in); tc.exp != res {
				t.Errorf("Expected %v, actual %v", tc.exp, res)
			}
		})
	}
}

func sumRange(rang int) int {
	if rang <= 0 {
		return 0
	}
	return rang + sumRange(rang-1)
}

func TestSumRange(t *testing.T) {
	testCases := []struct {
		in int
		exp int
		
	}{
		{in: 6, exp: 21},
		{in: 10, exp: 55},
		{in: 5, exp: 15},
		{in: 4, exp: 10},
		{in: 3, exp: 6},
		{in: 30, exp: 465},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			if res := sumRange(tc.in); tc.exp != res {
				t.Errorf("Expected %v, actual %v", tc.exp, res)
			}
		})
	}
}


func factorial(in int) int {
	if in == 0 || in == 1{
		return 1
	}
	return in * factorial(in-1)
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		in int
		exp int
		
	}{
		{in: 0, exp: 1},
		{in: 1, exp: 1},
		{in: 2, exp: 2},
		{in: 3, exp: 6},
		{in: 4, exp: 24},
		{in: 5, exp: 120},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			if res := factorial(tc.in); tc.exp != res {
				t.Errorf("Expected %v, actual %v", tc.exp, res)
			}
		})
	}
}

func collectOdd(tab []int) []int {
	out := make([]int,0)

	var run func(in []int)
	run = func(in []int) {
		if len(in) == 0 {
			return
		} else if in[0] % 2 == 1 {
			out = append(out, in[0])
		}
		run(in[1:])
	}

	run(tab)
	return out
}
func TestCollectOdd(t *testing.T) {
	testCases := []struct {
		in []int
		exp []int
		
	}{
		{in: []int{1,2,3,4,5,6,7}, exp: []int{1,3,5,7}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			res := collectOdd(tc.in)
			if len(res) != len(tc.exp) {
				t.Fatalf("Invalid len. Exp %v, actual %v", len(tc.exp), len(res))
			}

			for i := 0; i < len(res); i++ {
				if res[i] != tc.exp[i] {
					t.Errorf("Error on idx %v, actual %v, exp %v", i, res[i], tc.exp[i])
				}
			}
		})
	}
}