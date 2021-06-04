package day6

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	testInput := testData()

	for _, tc := range testCases {
		t.Run(tc.node, func(t *testing.T) {
			got := buildOrbits(testInput,"\n").calcOrbits(tc.node)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func testData() string {
	return `COM)B
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
}

func testData2() string {
	return testData() + `
K)YOU
I)SAN`
}

func TestAllOrbits(t *testing.T) {
	orbits := buildOrbits(testData(), "\n")
	got := orbits.calcAllOrbits()
	assert.Equal(t, 42, got)
}

func TestAllOrbitsPart1(t *testing.T) {
	orbits := buildOrbits(readFile(t), "\r\n")
	got := orbits.calcAllOrbits()
	assert.Equal(t, 261306, got)
}

func readFile(t *testing.T) string {
	file, err := os.Open("data.txt")
	require.NoError(t,err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t,err)
	return string(content)
}

func TestPart2Example(t *testing.T) {
	orbits := buildOrbits(testData2(), "\n")
	got := orbits.findPath("YOU", "SAN")
	assert.Equal(t, 4, got)
}

func TestPart2File(t *testing.T) {
	orbits := buildOrbits(readFile(t), "\r\n")
	got := orbits.findPath("YOU", "SAN")
	assert.Equal(t, 382, got)
}

type graphNode struct {
	parent string
	children []string
}

func (g *graphNode) addChild(children string) {
	if g.children == nil {
		g.children = make([]string, 0)
	}
	g.children = append(g.children, children)
}

type orbitGraph struct {
	m map[string]*graphNode
}

func buildOrbits(input string, lineSep string) *orbitGraph {
	o := &orbitGraph{
		m: map[string]*graphNode{},
	}
		
	vals := strings.Split(input, lineSep)
	for _, v := range vals {
		splitted := strings.Split(v, ")")
		node := splitted[0]
		orbittingObject := splitted[1]
		
		o.addNode(node,orbittingObject)
		o.addParent(orbittingObject, node)
	}

	return o
}

func (o *orbitGraph) addNode(node string, children string) {
	val, ok := o.m[node]
	if !ok {
		g := &graphNode{}
		g.addChild(children)
		o.m[node] = g
	} else {
		val.addChild(children)
		o.m[node] = val
	}
}

func (o *orbitGraph) addParent(node string, parent string) {
	val, ok := o.m[node]
	if !ok {
		g := &graphNode{}
		g.parent = parent
		o.m[node] = g
	} else {
		val.parent = parent
		o.m[node] = val
	}
}

func (o *orbitGraph) calcOrbits(startingNode string) int {
	orbit, ok := o.m[startingNode]
	if !ok {
		return 0
	} else if orbit.parent == "" {
		return 0
	}

	return 1 + o.calcOrbits(orbit.parent)
}

func (o *orbitGraph) calcAllOrbits() int {
	visitedNodes := make(map[string]struct{})

	sum := 0
	for key := range o.m {
		if _, visited := visitedNodes[key]; visited{
			continue
		}
		visitedNodes[key] = struct{}{}

		sum += o.calcOrbits(key)
	}
	return sum
}

func (o *orbitGraph) findPath(start, end string) int {
	startingNode := o.m[start].parent
	targetNode := o.m[end].parent

	buildHistory := func(node string) map[string]int {
		parentsOfStartingNode := map[string]int{}
		steps := 0
		for node != "" {
			parentsOfStartingNode[node]=steps
			steps++
			node = o.m[node].parent
		}
		return parentsOfStartingNode
	}
	fromStart := buildHistory(startingNode)
	fromEnd := buildHistory(targetNode)

	var minNumberOfSteps *int
	for key,stepsFromStart := range fromStart {
		stepsFromEnd, ok := fromEnd[key]
		if !ok {
			continue
		}

		steps := stepsFromStart + stepsFromEnd
		if minNumberOfSteps == nil || steps < *minNumberOfSteps {
			minNumberOfSteps = &steps
		}
	}
	if minNumberOfSteps == nil {
		return -1
	}
	return *minNumberOfSteps
}