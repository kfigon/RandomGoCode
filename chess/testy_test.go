package main
import (
	"testing"
)

func TestBoardGeneration(t *testing.T) {
	expected := `  ABCDEFGH
8|42365324|
7|11111111|
6| # # # #|
5|# # # # |
4| # # # #|
3|# # # # |
2|11111111|
1|42365324|
`

	b := newBoard(true)
	got := b.String()

	if got != expected {
		t.Errorf("Invalid board generated exp: \n%v\ngot \n%v", expected, got)
	}
}

func TestEmptyBoardGeneration(t *testing.T) {
	expected := `  ABCDEFGH
8| # # # #|
7|# # # # |
6| # # # #|
5|# # # # |
4| # # # #|
3|# # # # |
2| # # # #|
1|# # # # |
`

	b := newBoard(false)
	got := b.String()

	if got != expected {
		t.Errorf("Invalid board generated exp: \n%v\ngot \n%v", expected, got)
	}
}

func TestCoorString(t *testing.T) {
	testCases := []struct {
		exp	string
		input int
	}{
		{"A1", 0 },
		{"B1", 1 },
		{"C1", 2 },
		{"D1", 3 },
		{"E1", 4 },
		{"F1", 5 },
		{"G1", 6 },
		{"H1", 7 },
		{"A2", 8 },
		{"F8", 61 },
		{"G8", 62 },
		{"H8", 63 },
	}
	for _, tC := range testCases {
		t.Run(tC.exp, func(t *testing.T) {
			got := cellIdx(tC.input).coordStr()
			if got != tC.exp {
				t.Errorf("Got %v, exp %v for %v", got, tC.exp, tC.input)
			}
		})
	}	
}
