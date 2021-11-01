package main

import (
	"fmt"
	"strconv"
)

func main() {
	board := newBoard()
	fmt.Println(board)
}

func newBoard() *board {
	ar := [][]cell{}
	idx := 0
	for w := 0; w < 8; w++ {
		row := []cell{}
		for h := 0; h < 8; h++ {
			c := cellIdx(idx)
			row = append(row, cell{c,figureFromCell(c)})
			idx++
		}
		ar = append(ar, row)
	}
	return &board{area: ar}
}

func figureFromCell(c cellIdx) figure {
	rowIdx := c.row()
	colIdx := c.col()

	if rowIdx == 1 || rowIdx == 6 {
		isBlack := rowIdx == 6
		return figure{isBlack,pion}
	} else if (rowIdx == 0 || rowIdx == 7) {
		isBlack := rowIdx == 7

		if (colIdx == 0 || colIdx == 7) {
			return figure{isBlack,wieza}
		} else if (colIdx == 1 || colIdx == 6) {
			return figure{isBlack,kon}
		} else if (colIdx == 2 || colIdx == 5) {
			return figure{isBlack,goniec}
		} else if colIdx == 3 {
			return figure{isBlack, krol}
		}
		return figure{isBlack, hetman}
	}
	return figure{false, empty}
}

type cellIdx int
func (c cellIdx) isBlack() bool {
	evenRow := c.row() % 2 == 0
	evenCol := c.col() % 2 == 0
	if !evenRow {
		return !evenCol
	} 
	return evenCol
}

func (c cellIdx) row() int {
	return int(c) / 8
}

func (c cellIdx) col() int {
	return int(c) % 8
}

func (c cellIdx) coord() (rune, int) {
	return rune(c.col()+int('A')),c.row()+1
}

func (c cellIdx) coordStr() string {
	char,num := c.coord()
	return fmt.Sprintf("%c%v",char,num)
}

func (c cellIdx) String() string {
	if c.isBlack() {
		return "#"
	}
	return " "
}

type figure struct {
	black bool
	figureCode
}

type figureCode int
const (
	empty figureCode = iota
	pion
	kon
	goniec
	wieza
	hetman
	krol
)

type cell struct {
	cellIdx
	figure
}

func (c cell) String() string {
	if c.figure.figureCode == empty {
		return c.cellIdx.String()
	}
	return strconv.Itoa(int(c.figure.figureCode)) // todo - better idxs
}

type board struct {
	area [][]cell
}

func (b *board) String() string {
	out := "  ABCDEFGH\n"
	// reverse order on rows, but not columns - it's ok
	for i := len(b.area)-1; i >= 0; i-- {
		out += strconv.Itoa(i+1)+"|"
		for j := 0; j < len(b.area[0]); j++ {
			out += b.area[i][j].String()
		}
		out += "|\n"
	}
	return out
}