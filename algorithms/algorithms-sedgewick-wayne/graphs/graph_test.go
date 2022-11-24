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
		assert.Equal(t, 0, g.connectedComponents())
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
		assert.Equal(t, 2, g.connectedComponents())
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

func TestCyclic(t *testing.T) {
	t.Run("non cyclic", func(t *testing.T) {
		g := initGraph([]pair[node, node]{
			{"a","b"}, {"b","c"}, 
			{"c","d"},
		})
		assert.False(t, g.hasCycle())
	})

	t.Run("cyclic0", func(t *testing.T) {
		g := initGraph([]pair[node, node]{
			{"a","b"}, {"b","c"}, 
			{"c","d"}, {"d","a"}, 
		})
		assert.True(t, g.hasCycle())
	})

	t.Run("cyclic1", func(t *testing.T) {
		g := initGraph([]pair[node, node]{
			{"a","b"}, {"c","a"}, 
			{"c","d"}, {"c","f"},
			{"d","e"}, {"d","f"}, 
			{"f","e"},{"b","e"},
		})
		assert.True(t, g.hasCycle())
	})

	t.Run("cyclic2", func(t *testing.T) {
		g := searchGraph()
		assert.True(t, g.hasCycle())
	})
}

func initGraph(connections []pair[node,node]) undirectedGraph {
	g := newGraph()
	for _, v := range connections {
		g.connect(v.a, v.b)	
	}
	return g
}

func searchGraph() undirectedGraph {
	return initGraph([]pair[node,node]{
		{"0","1"}, {"0", "2"}, {"0", "5"},
		{"0","6"}, {"3","5"}, {"5", "4"},
		{"3", "4"}, {"4","6"},

		{"7","8"},

		{"9", "10"}, {"9","11"},{"9","12"},{"11","12"},
	})
}

func TestDfs(t *testing.T) {
	 g := searchGraph()

	assert.ElementsMatch(t, []node{"0","1","2","6","5","3","4"}, g.collectDfs("0"))
	assert.ElementsMatch(t, []node{"7","8"}, g.collectDfs("7"))
	assert.ElementsMatch(t, []node{"9","10","11","12"}, g.collectDfs("9"))
}

func TestIterDfs(t *testing.T) {
	g := searchGraph()

   assert.ElementsMatch(t, []node{"0","1","2","6","5","3","4"}, g.collectIterDfs("0"))
   assert.ElementsMatch(t, []node{"7","8"}, g.collectIterDfs("7"))
   assert.ElementsMatch(t, []node{"9","10","11","12"}, g.collectIterDfs("9"))
}

func TestBfs(t *testing.T) {
	g := searchGraph()

	assert.ElementsMatch(t, []node{"0", "6", "4", "3", "5", "2", "1"}, g.collectBfs("0"))
	assert.ElementsMatch(t, []node{"7","8"}, g.collectBfs("7"))
	assert.ElementsMatch(t, []node{"9","10","11","12"}, g.collectBfs("9"))
}

func TestConnectedComponents(t *testing.T) {
	g := searchGraph()

	assert.Equal(t, 3, g.connectedComponents())
}

type pair[T any, V any] struct {
	a T
	b V
}
