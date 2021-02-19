package binarySearchTree

import (
	"testing"
	"fmt"
)

func assertValues(t *testing.T, tree *tree, expected []int) {
	vals := tree.values()
	if len(vals) != len(expected) {
		t.Errorf("Value len should be %v, got %v", len(expected), len(vals))
	}

	for i := 0; i < len(vals); i++ {
		exp := expected[i]
		got := vals[i]
		if exp != got {
			t.Errorf("Invalid element in idx %v, got %v, exp %v", i, got, exp)
		}
	}
}
func TestCreateTree(t *testing.T) {
	tree := createTree()
	assertValues(t, tree, []int{})
}

func TestInsertAndTraverse(t *testing.T) {
	tdt := [] struct {
		elements []int
		expected []int
	} {
		{ []int{15}, []int{15}},
		{ []int{15, 16}, []int{15, 16}},
		{ []int{15, 14}, []int{14, 15}},
		{ []int{15, 1, 14}, []int{1, 14, 15}},
		{ []int{-1, 15, 1, 14}, []int{-1, 1, 14, 15}},
		{ []int{-1, 15, 1, 14,18, 3}, []int{-1, 1, 3, 14, 15, 18}},
	}
	recursiveInsert := map[bool]string {
		false: "iterative",
		true: "recursive",
	}
	for recursiveAlg := range recursiveInsert {
		for _, tc := range tdt {
			runName := fmt.Sprintf("%v %v", recursiveInsert[recursiveAlg], tc.elements)
			t.Run(runName, func(t *testing.T) {
				tree := createTree()
				for _, v := range tc.elements {
					
					if recursiveAlg {
						tree.insertRecursive(v)	
					} else {
						tree.insert(v)
					}
				}
				assertValues(t, tree, tc.expected)
			})
		}
	}
}

func searchInTree(t *testing.T, tree *tree, val int, expectToFind bool) {
	var errorMsg string
	if expectToFind {
		errorMsg = fmt.Sprintf("%v expected to be found, but it was not", val)
	} else {
		errorMsg = fmt.Sprintf("%v expected to be not found, but it was", val)
	}

	foundIter := tree.contains(val)
	if (expectToFind && !foundIter) || (!expectToFind && foundIter) {
		t.Error(errorMsg)
	}
	foundRecursive := tree.containsRecursive(val)
	if foundIter != foundRecursive {
		t.Errorf("Difference in result in recursive (%v) and iterative (%v) contains", foundRecursive, foundIter)
	}
	
}

func TestSearchWhenEmpty(t *testing.T) {
	tree := createTree()
	searchInTree(t, tree, 4, false)
}

func TestSearch(t *testing.T) {
	tree := createTree()
	tree.insert(4)
	tree.insert(3)
	tree.insert(2)
	tree.insert(6)
	
	searchInTree(t, tree, 4, true)
	searchInTree(t, tree, 3, true)
	searchInTree(t, tree, 2, true)
	searchInTree(t, tree, 6, true)
	
	searchInTree(t, tree, -1, false)
	searchInTree(t, tree, 5, false)
	searchInTree(t, tree, -7, false)
	searchInTree(t, tree, 8, false)
	searchInTree(t, tree, -8, false)
}

