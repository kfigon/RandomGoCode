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
	t.Run("non cyclic1", func(t *testing.T) {
		g := initDirected([]pair[node,node]{
			{"a","b"},{"b","c"},{"c","d"},
			{"d","e"},
		})
		assert.Equal(t, []node{"e","d","c","b","a"}, g.topology())
	})

	t.Run("non cyclic2", func(t *testing.T) {
		g := initDirected([]pair[node,node]{
			{"5","0"},{"0","1"},{"0","2"},
			{"1","3"}, {"3","2"},
		})
		assert.Equal(t, []node{"2","3","1","0","5"}, g.topology())
	})
	
	t.Run("non cyclic3", func(t *testing.T) {
		g := exampleDiGraph2()
		topo := g.topology()
		assert.Greater(t, len(topo), 0)
		// non deterministic order
		// assert.Equal(t, []node{"8", "7", "2", "3", "0", "6", "9",
		// "10", "11", "12", "1", "5", "4"}, g.topology())
	})

	t.Run("non cyclic4", func(t *testing.T) {
		g := initDirected([]pair[node, node]{
			{"deploy", "setup docker"},
			{"deploy", "build binary"},
			{"build binary", "run tests"},
		})
		// non deterministic order
		assert.Greater(t, len(g.topology()), 0)
		// assert.Equal(t, []node{"setup docker", "run tests","build binary", "deploy"}, g.topology())
		// assert.Equal(t, []node{"run tests", "build binary", "setup docker", "deploy"}, g.topology())
	})

	t.Run("cyclic2", func(t *testing.T) {
		g := exampleDiGraph()
		assert.Equal(t, []node{}, g.topology())
	})

	t.Run("cyclic1", func(t *testing.T) {
		g := initDirected([]pair[node,node] {
			{"a","b"},{"b","c"},{"c","d"},{"d","e"},
			{"e","b"},
		})
		assert.Equal(t, []node{}, g.topology())
	})
}

func TestCycle(t *testing.T) {
	t.Run("simple non cyclic", func(t *testing.T) {
		g := initDirected([]pair[node,node]{
			{"a","b"},{"b","c"},{"c","d"},
			{"d","e"},
		})
		assert.Len(t, g.cycle(), 0)
	})
	
	t.Run("non cyclic", func(t *testing.T) {
		g := exampleDiGraph2()
		cycle := g.cycle()
		assert.Len(t, cycle, 0)
	})

	t.Run("cyclic", func(t *testing.T) {
		g := exampleDiGraph()
		cycle := g.cycle()
		assert.NotEqual(t, 0, len(cycle))
		t.Log(cycle)
	})

	t.Run("simple cyclic", func(t *testing.T) {
		g := initDirected([]pair[node,node] {
			{"a","b"},{"b","c"},{"c","d"},{"d","e"},
			{"e","b"},
		})
		cycle := g.cycle()
		assert.NotEqual(t, 0, len(cycle))

		// reversed order
		// assert.Equal(t, []node{"e","d","c","b","e"}, cycle)
	})
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