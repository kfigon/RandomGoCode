package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		g := newGraph()
		assert.False(t, g.connected("a","b"))
		assert.Equal(t, []node{}, g.adjecent("A"))
	})

	t.Run("more nodes", func(t *testing.T) {
		g := newGraph()
		g.connect("a","b")
		g.connect("c","a")

		g.connect("z","x")

		assert.True(t, g.connected("a","b"))
		assert.True(t, g.connected("b","a"))
		
		assert.True(t, g.connected("a","c"))
		assert.True(t, g.connected("c","a"))

		assert.True(t, g.connected("b","c"))
		assert.True(t, g.connected("c","b"))

		assert.False(t, g.connected("z","a"))
		assert.False(t, g.connected("x","b"))

		assert.ElementsMatch(t, []node{"b","c"}, g.adjecent("a"))
	})
}

// adjacency list is in general good underlying data structure for general graphs
type node string
type void struct{}
type set map[node]void
type undirectedGraph map[node]set

func newGraph() *undirectedGraph {
	return &undirectedGraph{}
}

func (g undirectedGraph) connect(a,b node) {
	add := func(n,m node) {
		nodes, ok := g[n];
		if !ok {
			nodes = set{}
		}
		nodes[m] = void{}
		g[n] = nodes
	}
	add(a,b)
	add(b,a)
}

func (g undirectedGraph) adjecent(n node) []node {
	out := []node{}
	for k := range g[n] {
		out = append(out, k)
	}
	return out
}

func (g undirectedGraph) connected(a,b node) bool {
	return false
}