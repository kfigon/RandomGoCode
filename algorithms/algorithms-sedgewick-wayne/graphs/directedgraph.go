package graphs

type directedGraph map[node]set

func newDirectedGraph() directedGraph{
	return directedGraph{}
}

func(g directedGraph) connect(a,b node) {
	vals, ok := g[a]
	if !ok {
		vals = set{}
	}
	vals.add(b)
	g[a]=vals
}

func(g directedGraph) connected(a,b node) bool {
	return false
}