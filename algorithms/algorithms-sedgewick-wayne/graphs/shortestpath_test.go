package graphs

import "testing"

func TestDijkstra(t *testing.T) {
	t.Skip("todo one day")
}

type weightedEntry struct {
	n node
	v int
}

type weightedDirected map[node][]weightedEntry

func newWeightedDirected() weightedDirected {
	return map[node][]weightedEntry{}
}

func (g weightedDirected) connect(a, b node, v int) {
	nodes := g[a]
	nodes = append(nodes, weightedEntry{b,v})
	g[a] = nodes
}

// dijkstra
func (g weightedDirected) shortestPath() []node {
	return nil
}