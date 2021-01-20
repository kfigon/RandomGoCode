package memory

import (
	"testing"
	// "fmt"
)

func TestInitGame(t *testing.T) {
	game := newField()
	for i := 0; i < game.size(); i++ {
		if game.isRevealed(i) {
			t.Errorf("Field %v shouldn't be revealed at the start", i)
		}
	}
}

func TestGenerateAndShuffle(t *testing.T)  {
	maxValue := 8
	vals := generatePairsAndShuffle(maxValue)
	if len(vals) != 8 {
		t.Fatal("Wrong number of values", len(vals))
	}

	occurences := make(map[int]int)
	for _,v := range vals {
		occurences[v]++
	}
	if len(occurences) != 4 {
		t.Fatalf("wrong number of unique values: %v", len(occurences))
	}
	for key := range occurences {
		if occurences[key] != 2 {
			t.Errorf("Key %q, has invalid occurences %v", key, occurences[key])
		}
	}
}