package main

type tree struct {}

func (t *tree) size() int {
	return 0
}

func (t *tree) values() []int {
	return []int{}
}

func newTree() *tree {
	return &tree{}
}