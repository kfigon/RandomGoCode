package bottles

import (
	"fmt"
	"strings"
)

func generateLyrics() string {
	out := strings.Builder{}
	for i := 99; i >= 0; i-- {
		if i == 2 {
			out.WriteString(fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i))
			out.WriteString(fmt.Sprintf("Take one down and pass it around, %d bottle of beer on the wall.\n\n", i-1))
		} else if i == 1 {
			out.WriteString(fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\n", i, i))
			out.WriteString("Take one down and pass it around, no more bottles of beer on the wall.\n\n")
		} else if i == 0 {
			out.WriteString("No more bottles of beer on the wall, no more bottles of beer.\n")
			out.WriteString("Go to the store and buy some more, 99 bottles of beer on the wall.")
		} else {
			out.WriteString(fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i))
			out.WriteString(fmt.Sprintf("Take one down and pass it around, %d bottles of beer on the wall.\n\n", i-1))
		}
	}
	return out.String()
}
