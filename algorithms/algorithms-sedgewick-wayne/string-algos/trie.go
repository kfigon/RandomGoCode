package stringalgos

type node struct {
	children map[rune]*node
	isWord bool
}

func newNode() *node {
	return &node{
		children: map[rune]*node{},
	}
}

type trie struct {
	root *node
}

func (t *trie) add(text string) {
	if t.root == nil {
		t.root = newNode()
	}

	ptr := t.root
	for i, c := range text {
		if v,ok := ptr.children[c]; ok {
			ptr = v
		} else {
			next := newNode()
			ptr.children[c] = next
			ptr = next
		}

		if i == len(text)-1 {
			ptr.isWord = true
		}
	}
}

func (t *trie) suggestions(prefix string) []string {
	var out []string
	var traverse func(prefixStr string, n *node)
	traverse = func(prefixStr string, n *node) {
		if n == nil {
			return
		}
		if n.isWord {
			out = append(out, prefixStr)
		}
		for c, children := range n.children {
			traverse(prefixStr+string(c), children)
		}
	}

	lastCommonNode := t.root
	for _, c := range prefix {
		v, ok := lastCommonNode.children[c]
		if !ok {
			return []string{}
		}
		lastCommonNode = v
	}
	traverse(prefix, lastCommonNode)
	return out
}