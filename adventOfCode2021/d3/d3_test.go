package d3

import (
	"testing"
	"os"
	"strconv"
	"strings"
)

func TestPart1Example(t *testing.T) {
	in := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	got := solveP1(strings.Split(in,"\n"))
	if got != 198 {
		t.Error("Invalid p1", got, "exp", 198)
	}
}
func TestPart1(t *testing.T) {
	got := solveP1(parse(t))
	if got != 2261546 {
		t.Error("Invalid p1", got, "exp", 2261546)
	}
}

func TestPart2(t *testing.T) {
	got := solveP2(parse(t))
	if got != 6775520 {
		t.Error("Invalid p1", got, "exp", 6775520)
	}
}

func parse(t *testing.T) []string {
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("Error in reading file",err)
		return nil
	}
	return strings.Split(string(d), "\r\n")
}

func solveP1(d []string) int {
	out := ""

	for i := len(d[0])-1; i >= 0; i-- {
		occurences := map[byte]int{}
		for _, line := range d {
			c := line[i]
			occurences[c]++
		}
		if occurences['0'] > occurences['1'] {
			out += "0"
		} else {
			out += "1"
		}
	}
	reversed := ""
	for i := len(out)-1; i >= 0; i-- {
		reversed += string(out[i])
	}
	gamma, _ := strconv.ParseInt(reversed, 2, 64)
	epsilon, _ := strconv.ParseInt(negate(reversed), 2, 64)
	return int(gamma)*int(epsilon)
}

func negate(in string) string {
	out := ""
	for _, c := range in {
		switch c {
		case '0': out += "1"
		case '1': out += "0"
		}
	}
	return out
}

func solveP2(d []string) int {
	return 0
}