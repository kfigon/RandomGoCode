package graphs


// adjacency list is in general good underlying data structure for general graphs
type node string
type void struct{}
type set map[node]void

func (s set) present(n node) bool {
	_, ok := s[n]
	return ok
}

func (s set) add(n node) {
	s[n] = void{}
}

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
		if visited.present(n) {
			return false
		}
		visited.add(n)
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
		visited.add(n)
		for k := range g[n] {
			if visited.present(k){
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
// go max left/right then change direction
func (g undirectedGraph) dfs(a node, fn func(node)) {
	visited := set{}
	var foo func(node)
	foo = func(n node) {
		if visited.present(n){
			return
		}
		fn(n)
		visited.add(n)
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

func (g undirectedGraph) collectBfs(a node) []node {
	out := []node{}
	g.bfs(a, func(n node) {
		out = append(out, n)
	})
	return out
}

// first check neighbours. Better for near searches
func (g undirectedGraph) bfs(a node, fn func(node)) {
	visited := set{}
	queue := []node{}
	enqueue := func(n node)	{
		queue = append(queue, n)
	}

	dequeue := func() node {
		toRet := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		return toRet
	}

	enqueue(a)
	for len(queue) > 0 {
		current := dequeue()
		if visited.present(current){
			continue
		}
		visited.add(current)
		fn(current)
		for k := range g[current] {
			enqueue(k)
		}
	}
}

// same as bfs, just with stack instead of queue
func (g undirectedGraph) iterDfs(a node, fn func(node)) {
	visited := set{}
	stack := []node{}
	push := func(n node) {
		stack = append(stack, n)
	}

	pop := func() node {
		toRet := stack[0]
		stack = stack[1:]
		return toRet
	}

	push(a)
	for len(stack) > 0 {
		current := pop()
		if visited.present(current){
			continue
		}
		visited.add(current)
		fn(current)
		for k := range g[current] {
			push(k)
		}
	}

}

func (g undirectedGraph) collectIterDfs(a node) []node {
	out := []node{}
	g.iterDfs(a, func(n node) {
		out = append(out, n)
	})
	return out
}

func (g undirectedGraph) connectedComponents() int {
	visited := set{}

	var foo func(node) bool
	foo = func(n node) bool {
		if visited.present(n){
			return false
		}
		visited.add(n)
		any := false
		for k := range g[n] {
			any = true
			foo(k)
		}
		return any
	}

	connected := 0
	for node := range g {
		if foo(node) {
			connected++
		}
	}
	return connected
}

func (g undirectedGraph) hasCycle() bool {
	visited := set{}
	var foo func(node, node) bool
	foo = func(current, parent node) bool {
		if visited.present(current) {
			return false
		}

		visited.add(current)
		for child := range g[current] {
			if visited.present(child) {
				if child != parent {
					return true
				}
				continue
			} 

			if foo(child, current) {
				return true
			}

		}
		return false
	}

	for k := range g {
		if foo(k, k) {
			return true
		}
	}
	return false
}