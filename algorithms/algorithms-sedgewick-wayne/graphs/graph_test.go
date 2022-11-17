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
		g := initGraph([]pair[node, node]{
			{"a","b"}, {"c","a"}, 
			{"z","x"},
		})

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

	t.Run("any path to node", func(t *testing.T) {
		// 		a ----b
		//    /		   \
		//   /          \
		//  c --- d ---- e
		//   \    |    /
		//    \   |   /
		//       f   /
		g := initGraph([]pair[node, node]{
			{"a","b"}, {"c","a"}, 
			{"c","d"}, {"c","f"},
			{"d","e"}, {"d","f"}, 
			{"f","e"},{"b","e"},
		})
		// non deterministic due to maps. It's also bad because it can loop
		assert.Greater(t,  len(g.path("a","f")), 1)
	})
}

func initGraph(connections []pair[node,node]) undirectedGraph {
	g := newGraph()
	for _, v := range connections {
		g.connect(v.a, v.b)	
	}
	return g
}

func TestDfs(t *testing.T) {
	g := initGraph([]pair[node,node]{
		{"0","1"}, {"0", "2"}, {"0", "5"},
		{"0","6"}, {"3","5"}, {"5", "4"},
		{"3", "4"}, {"4","6"},

		{"7","8"},

		{"9", "10"}, {"9","11"},{"9","12"},{"11","12"},

	})

	assert.ElementsMatch(t, []node{"0","1","2","6","5","3","4"}, g.collectDfs("0"))
	assert.ElementsMatch(t, []node{"7","8"}, g.collectDfs("7"))
	assert.ElementsMatch(t, []node{"9","10","11","12"}, g.collectDfs("9"))
}


func TestMissingAlgos(t *testing.T) {
	t.Fatal("iter dfs") // same as bfs, but with stack
	t.Fatal("iter bfs") // with queue
}

type pair[T any, V any] struct {
	a T
	b V
}

// adjacency list is in general good underlying data structure for general graphs
type node string
type void struct{}
type set map[node]void
type undirectedGraph map[node]set

func newGraph() undirectedGraph {
	return undirectedGraph{}
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
	visited := set{}
	var foo func(node) bool
	foo = func(n node) bool {
		if n == b {
			return true
		}
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = void{}
		for k := range g[n] {
			if foo(k) {
				return true
			}
		}
		return false
	}
	return foo(a)
}

func (g undirectedGraph) path(a,b node) []node {
	pathToStart := map[node]node{}

	visited := set{}
	var foo func(node) bool
	foo = func(n node) bool {
		if n == b {
			return true
		}
		visited[n] = void{}
		for k := range g[n] {
			if _, ok := visited[k]; ok {
				continue
			}

			pathToStart[k] = n
			if foo(k) {
				return true
			}
		}
		return false
	}
	out := []node{}
	if !foo(a) {
		return out
	}
	out = append(out, b)
	for next, ok := pathToStart[b]; next != a; next, ok = pathToStart[next] {
		if !ok {
			break
		}
		out = append(out, next)
	}
	out = append(out, a)
	return out
}

// visit all nodes in connected graph (there's a path to every node from any node)
func (g undirectedGraph) dfs(a node, fn func(node)) {
	visited := set{}
	var foo func(node)
	foo = func(n node) {
		if _, ok := visited[n]; ok {
			return
		}
		fn(n)
		visited[n] = void{}
		for k := range g[n] {
			foo(k)
		}
	}
	foo(a)
}

func (g undirectedGraph) collectDfs(a node) []node {
	out := []node{}
	g.dfs(a, func(n node) {
		out = append(out, n)
	})
	return out
}