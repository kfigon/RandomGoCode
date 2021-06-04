package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestOrbits(t *testing.T) {
	testCases := []struct {
		node string
		exp int		
	}{
		{"D",3},
		{"L",7},
		{"COM",0},
		{"I",4},
		{"F",5},
		{"E",4},
	}
	testInput := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	for _, tc := range testCases {
		t.Run(tc.node, func(t *testing.T) {
			got := buildOrbits(testInput).calcOrbits(tc.node)
			assert.Equal(t, tc.exp, got)
		})
	}
}

type graphNode struct {
	parent string
	children []string
}

func (g *graphNode) addChildrenIfPresent(children string) {
	if children == "" {
		return
	}
	if g.children == nil {
		g.children = make([]string, 0)
	}
	
	g.children = append(g.children, children)
}

func (g *graphNode) addParentIfPresent(parent string) {
	if parent != "" {
		g.parent = parent
	}
}

type orbitGraph struct {
	m map[string]*graphNode
}

func buildOrbits(input string) *orbitGraph {
	o := &orbitGraph{
		m: map[string]*graphNode{},
	}
		
	vals := strings.Split(input, "\n")
	for _, v := range vals {
		splitted := strings.Split(v, ")")
		node := splitted[0]
		orbittingObject := splitted[1]
		
		o.addNode(node,orbittingObject)
		o.addChild(orbittingObject, node)
	}

	return o
}

func (o *orbitGraph) addNode(node string, children string) {
	val, ok := o.m[node]
	if !ok {
		g := &graphNode{}
		g.addChildrenIfPresent(children)
		o.m[node] = g
	} else {
		val.addChildrenIfPresent(children)
		o.m[node] = val
	}
}

func (o *orbitGraph) addChild(node string, parent string) {
	val, ok := o.m[node]
	if !ok {
		g := &graphNode{}
		g.addParentIfPresent(parent)
		o.m[node] = g
	} else {
		val.addParentIfPresent(parent)
		o.m[node] = val
	}
}

func (o *orbitGraph) calcOrbits(startingNode string) int {
	// orbit, ok := o.m[startingNode]
	// if !ok {
	// 	return 0
	// }
	return -1
}