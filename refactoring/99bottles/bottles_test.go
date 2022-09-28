package bottles

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func read(t *testing.T) string {
	data, err := os.ReadFile("lyrics.txt")
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestLyrics(t *testing.T) {
	t.Run("simple approach", func(t *testing.T) {
		assert.Equal(t, read(t), generateLyrics())	
	})

	t.Run("polymorphic approach", func(t *testing.T) {
		assert.Equal(t, read(t), generateLyricsPolymorph())	
	})

	t.Run("functional polymorphic approach", func(t *testing.T) {
		assert.Equal(t, read(t), generateLyricsFunctional())	
	})

	t.Run("polymorphic by book", func(t *testing.T) {
		assert.Equal(t, read(t), generateLyricsByBook())	
	})
}

func BenchmarkGenerateLyricsSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateLyrics()
	}
}
