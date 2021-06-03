package day4

import (
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

// https://adventofcode.com/2019/day/4
func TestPasswordCheck(t *testing.T) {
	testCases := []struct {
		in string	
		exp bool
	}{
		{"333333", true},
		{"223450", false},
		{"123789", false},
	}
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got := password(tc.in).isOkv1()
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestPart1(t *testing.T) {
	howManyOk := 0
	for i := min; i <= max; i++ {
		in := strconv.Itoa(i)
		if password(in).isOkv1() {
			howManyOk++
		}
	}
	assert.Equal(t, 1099, howManyOk)
}

func TestPasswordCheck2(t *testing.T) {
	testCases := []struct {
		in string	
		exp bool
	}{
		{"333333", false},
		{"412233", true},
		{"523444", false},
		{"333344", true},
	}
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got := password(tc.in).isOkv2()
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestPart2(t *testing.T) {
	howManyOk := 0
	for i := min; i <= max; i++ {
		in := strconv.Itoa(i)
		if password(in).isOkv2() {
			howManyOk++
		}
	}
	assert.Equal(t, 1099, howManyOk)
}

const min = 245182
const max = 790572

type password string

func (p password) isOkv1() bool {
	return p.lengthOk() && 
		   p.withinRange() && 
		   p.adjecentTheSame() && 
		   p.neverDecrease()
}

func (p password) isOkv2() bool {
	return p.isOkv1() && p.adjecentNotPartOfLargerGroup()
}

func (p password) lengthOk() bool { return len(p) == 6 }

func (p password) withinRange() bool { 
	i, _ := strconv.Atoi(string(p))
	return i >= min && i <= max
}

func (p password) adjecentTheSame() bool { 
	for i := 0; i < len(p)-1; i++ {
		current := rune(p[i])
		next := rune(p[i+1])

		if current == next {
			return true
		}
	}
	return false
}

func (p password) neverDecrease() bool { 
	for i := 0; i < len(p)-1; i++ {
		current,_ := strconv.Atoi(string(p[i]))
		next,_ := strconv.Atoi(string(p[i+1]))

		if next < current {
			return false
		}
	}
	return true
}

func (p password) adjecentNotPartOfLargerGroup() bool { 
	return false
}
