package bottles
import (
	"strings"
	"fmt"
)

func generateLyricsFunctional() string {
	out := strings.Builder{}
	for i := 99; i >= 0; i-- {
		verse := factoryFn(i)
		out.WriteString(verse())
	}
	return out.String()
}

type verseFn func() string
func factoryFn(verseNum int) verseFn {
	switch verseNum {
	case 0: return func() string {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall."
	}

	case 1: return func() string {
		return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake one down and pass it around, no more bottles of beer on the wall.\n\n",
		verseNum, verseNum)
	}

	case 2: return func() string {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n\n",
		 verseNum, verseNum, verseNum-1)
	}

	default:  return func() string {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n\n", 
		verseNum, verseNum, verseNum-1)
	}
	}
}
