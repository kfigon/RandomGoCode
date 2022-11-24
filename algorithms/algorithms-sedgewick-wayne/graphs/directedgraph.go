package graphs

import "strings"

type directedGraph map[node]set

func (g directedGraph) String() string {
	out := ""
	for k,nodes := range g {
		out += string(k) + " -> ["
		var ns = []string{}
		for n := range nodes {
			ns = append(ns, string(n))
		}
		out += strings.Join(ns, " ")
		out += "]"
		out += " "
	}
	return out
}

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

func(g directedGraph) connected(start, end node) bool {
	visited := set{}
	var foo func(node) bool
	foo = func(a node) bool {
		if visited.present(a) {
			return false
		}
		visited.add(a)
		for n := range g[a] {
			if n == end || foo(n) {
				return true
			}
		}
		return false
	}
	return foo(start)
}