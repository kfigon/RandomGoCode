package graphs

import "strings"

type directedGraph map[node]set

func (g directedGraph) String() string {
	out := []string{}
	for k,nodes := range g {
		str := ""
		str += string(k) + " -> ["
		var ns = []string{}
		for n := range nodes {
			ns = append(ns, string(n))
		}
		str += strings.Join(ns, " ")
		str += "]"
		out = append(out, str)
	}
	return strings.Join(out, " ")
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

func (g directedGraph) topologicalSort() ([]node, error) {
	return nil,nil
}

func (g directedGraph) cycle() []node {
	visited := set{}
	onStack := set{}
	pathTo := map[node]node{}
	cycleNodes := []node{}

	var dfs func(node)
	dfs = func(n node){
		if visited.present(n) {
			return
		}

		onStack.add(n)
		visited.add(n)
		for child := range g[n] {
			if len(cycleNodes) != 0 {
				return
			} else if !visited.present(child) {
				pathTo[child] = n
				dfs(child)
			} else if onStack.present(child) {
				for next := n; next != child; next = pathTo[next] {
					cycleNodes = append(cycleNodes, next)
				}
				cycleNodes = append(cycleNodes, child)
				cycleNodes = append(cycleNodes, n)
				return
			}
		}
		delete(onStack, n)
	}

	for n := range g {
		dfs(n)
	}
	return cycleNodes
}

// only when there're no cycles
func (g directedGraph) topology() []node {
	if len(g.cycle()) != 0 {
		return []node{}
	}

	out := []node{}
	visited := set{}
	var dfs func(node)
	dfs = func(n node) {
		if visited.present(n) {
			return
		}

		visited.add(n)
		for child := range g[n] {
			dfs(child)
		}
		out = append(out, n) // postorder
	}

	for k := range g {
		dfs(k)
	}

	// reverse
	for i := 0; i < len(out)/2; i++ {
		tmp := out[i]
		out[i] = out[len(out)-1-i]
		out[len(out)-1-i] = tmp
	}
	return out
}