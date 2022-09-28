package bottles

import (
	"fmt"
	"strconv"
	"strings"
)

func generateLyricsByBook() string {
	var lines []string
	for i := 99; i >= 0; i-- {
		bot := &bottleNumber{i}
		nextBottle := &bottleNumber{successor(i)}

		line := fmt.Sprintf("%s %s of beer on the wall, %s %s of beer.\n", 
						capitalize(bot.number()), 
						bot.container(), 
						bot.number(), 
						bot.container())

		line += fmt.Sprintf("%s, %s %s of beer on the wall.", 
						bot.action(), 
						nextBottle.number(), 
						nextBottle.container())
		
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n\n")
}

func capitalize(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 2 {
		return strings.Title(fields[0]) + " " + fields[1]
	}
	return s
}

func successor(i int) int {
	if i == 0 {
		return 99
	}
	return i-1
}

type bottleNumber struct {
	num int
}

func (b *bottleNumber) number() string {
	if b.num == 0 {
		return "no more"
	}
	return strconv.Itoa(b.num)
}

func (b *bottleNumber) container() string {
	if b.num == 1 {
		return "bottle"
	}
	return "bottles"
}

func (b *bottleNumber) action() string {
	if b.num == 0 {
		return "Go to the store and buy some more"
	}
	return "Take one down and pass it around"
}