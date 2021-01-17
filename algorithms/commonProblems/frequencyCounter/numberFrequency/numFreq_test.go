package numberFrequency

import (
	"testing"
	"fmt"
	"math"
)

// # given 2 positive integers
// # find 2 numbers with same frequency of digits

func getDigits(data int) []int {
	out := make([]int, 0)
	if data == 0 {
		return []int{0}
	}
	
	howManyDigits := int(math.Log10(float64(data)))+1
	for i := 0; i < howManyDigits; i++ {
		powered := int(math.Pow10(i))
		out = append(out, (data/powered) % 10)
	}
	return out
}

func makeMap(data int) map[int]int {
	dict := make(map[int]int)
	for _,v := range getDigits(data) {
		dict[v]++
	}
	return dict
}

func sameFrequency(in1, in2 int) bool {
	d1 := makeMap(in1)
	d2 := makeMap(in2)
	
	if len(d1) != len(d2) {
		return false
	}

	for k := range d1 {
		v := d1[k]
		secondV := d2[k]
		if v != secondV {
			return false
		}
	}

	return true
}

func TestGetDigits(t *testing.T) {
	testCases := []struct {
		in int		
		exp []int
	}{
		{in: 0, exp: []int{0}},
		{in: 1, exp: []int{1}},
		{in: 12, exp: []int{2,1}},
		{in: 123, exp: []int{3,2,1}},
		{in: 1234, exp: []int{4,3,2,1}},
		{in: 12345, exp: []int{5,4,3,2,1}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v",tc.in), func(t *testing.T) {
			res := getDigits(tc.in)
			if len(res) != len(tc.exp) {
				t.Fatalf("len res %v != len exp %v", len(res), len(tc.exp))
			}

			for i := range res {
				actualV := res[i]
				expectedV := tc.exp[i]
				if actualV != expectedV {
					t.Errorf("Error in idx %v exp %v != actual %v", i, expectedV, actualV)
				}
			}

		})
	}
}

func TestNumberFreq(t *testing.T) {
	testCases := []struct {
		in1 int		
		in2 int
		exp bool
	}{
		{in1: 182, in2: 281, exp: true},
		{in1: 34, in2: 14, exp: false},
		{in1: 3589578, in2:  5879385, exp: true},
		{in1: 22, in2: 222, exp: false},
		{in1: 222, in2: 22, exp: false},
		{in1: 22, in2: 221, exp: false},
		{in1: 221, in2: 22, exp: false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v",tc.in1, tc.in2), func(t *testing.T) {
			if res := sameFrequency(tc.in1, tc.in2); res != tc.exp {
				t.Errorf("Got %v, exp %v", res, tc.exp)
			}
		})
	}
}