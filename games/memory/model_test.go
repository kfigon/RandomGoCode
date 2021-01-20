package memory

import (
	"testing"
	"fmt"
)

func TestInitGame(t *testing.T) {
	game := newField()
	for i := 0; i < game.size(); i++ {
		if game.isRevealed(i) {
			t.Errorf("Error %v shouldn't be revealed at the start")
		}
	}
}

func TestGenerateAndShuffle(t *testing.T)  {
	
}