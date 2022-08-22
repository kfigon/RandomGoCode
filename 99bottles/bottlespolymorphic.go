package bottles
import (
	"strings"
	"fmt"
)

func generateLyricsPolymorph() string {
	out := strings.Builder{}
	for i := 99; i >= 0; i-- {
		verse := factory(i)
		out.WriteString(verse.songVerse())
	}
	return out.String()
}

type verseMaker interface {
	songVerse() string
}

func factory(verseNum int) verseMaker {
	switch verseNum {
	case 0: return &lastVerse{}
	case 1: return &verseOne{}
	case 2: return &verseTwo{}
	default:  {
			out := regularVerse(verseNum)
			return &out
		}
	}
}

type regularVerse int
func (r *regularVerse) songVerse() string {
	i := *r
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n\n",
		i, i, i-1)
}

type lastVerse struct{}
func (l *lastVerse) songVerse() string {
	return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall."
}

type verseOne struct{}
func (v *verseOne) songVerse() string {
	i := 1
	return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake one down and pass it around, no more bottles of beer on the wall.\n\n",
		i, i)
}

type verseTwo struct{}
func (v *verseTwo) songVerse() string {
	i := 2
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n\n",
		i, i, i-1)
}