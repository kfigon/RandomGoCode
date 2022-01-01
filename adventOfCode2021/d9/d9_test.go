package d9

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

const exampleData string = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestExampleP1(t *testing.T) {
	if got := solveP1(parse(exampleData, "\n")); got != 15 {
		t.Error("Exp", 15, "got", got)
	}
}

func TestP1(t *testing.T) {
	d,err:=os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	if got := solveP1(parse(string(d), "\r\n")); got != 585 {
		t.Error("Exp", 585, "got", got)
	}
}
func TestP2(t *testing.T) {
	t.Fail()
}
func parse(data string, sep string) [][]int {
	lines := strings.Split(data,sep)
	out := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		out = append(out, row)
	}
	return out
}

func solveP1(data [][]int) int {
	neighbours := func(row,col int) []int {
		out := []int{}
		type pair struct {
			r int
			c int
		}

		candidates := []pair {
				{row-1, col},
			{row, col-1}, {row, col+1},
				{row+1, col},
		}
		for _, c := range candidates {
			if c.r < 0 || c.r >= len(data) || c.c < 0 || c.c >= len(data[0]) {
				continue
			}
			out = append(out, data[c.r][c.c])
		}
		return out
	}

	sum := 0
	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data[0]); c++ {
			current := data[r][c]
			nei := neighbours(r,c)

			allOk := true
			for _, v := range nei {
				if current >= v {
					allOk = false
					break
				}
			}
			if allOk {
				sum+=current+1
			}
		}
	}
	return sum
}