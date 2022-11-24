package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectivity(t *testing.T) {
	testCases := []struct {
		a,b node
		exp bool
		
	}{
		{"0","1", true},
		{"2","0", true},
		{"0","5", true},
		{"0","4", true},
		{"8","0", true},
		{"6","5", true},
		{"10","3", true},
		{"12","0", true},
		{"5","1", true},
		
		{"0","8", false},
		{"1","0", false},
		{"1","5", false},
		{"6","7", false},
	}

	g := exampleDiGraph()
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v-%v -> %v", tC.a, tC.b, tC.exp), func(t *testing.T) {
			assert.Equal(t, tC.exp, g.connected(tC.a, tC.b))	
		})
	}
}

// aka make build system - order with required steps
// job schedule, course schedule with prerequisites, spreadsheet formulas, symbolic links, inheritance hierarchy etc.
func TestTopologicalSort(t *testing.T) {
	// g := exampleDiGraph2()
	t.Fatal("todo")
}

func initDirected(pairs []pair[node,node]) directedGraph {
	g := newDirectedGraph()
	for _, v := range pairs {
		g.connect(v.a, v.b)
	}
	return g
}

func exampleDiGraph() directedGraph {
	return initDirected([]pair[node,node]{
		{"0","1"}, {"0","5"},
		{"2","0"}, {"2","3"},
		{"3","2"}, {"3","5"},
		{"4","2"}, {"4","3"},
		{"5","4"},
		{"6","0"}, {"6","4"}, {"6","9"}, 
		{"7","6"}, {"7","8"},
		{"8","7"}, {"8","9"},
		{"9","10"}, {"9","11"}, 
		{"10","12"},
		{"11","4"},{"11","12"},
		{"12","9"},
	})
}

func exampleDiGraph2() directedGraph {
	return initDirected([]pair[node,node]{
		{"0","1"}, {"0","5"},{"0","6"},
		{"2","0"},{"2","3"},
		{"3","5"},
		{"5","4"},
		{"6","4"},{"6","9"},
		{"7","6"},
		{"8","7"},
		{"9","10"}, {"9","11"},{"9","12"},
		{"11","12"},
	})
}