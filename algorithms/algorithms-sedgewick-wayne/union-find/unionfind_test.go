package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// union find - algorithms for dynamic connectivity problems
// check if 2 elements of a set are connected

// 3 algorithms: quick find, quick union and weighted quick union

func populateQuickFind(num int, data []pair[int,int]) *quickFind {
	qf := newQuickFind(num)
	for _, v := range data {
		qf.union(v.a, v.b)
	}
	return qf
}

func TestQuickFind(t *testing.T) {
	data := []pair[int,int]{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{6, 7},
	}
	qf := populateQuickFind(10, data)

	assert.Equal(t, 2, qf.count())
	assert.Equal(t, [][]int{{0, 1, 2, 5, 6, 7}, {3, 4, 8, 9}}, qf.connectedComponents())
}

type pair[T any, V any] struct{
	a T
	b V
}

type unionFind interface{
	union(int,int)
	find(int)
}

type quickFind struct{
	tab []int
}

func newQuickFind(num int) *quickFind {
	tab := []int{}
	for i := 0; i < num; i++ {
		tab = append(tab, i)
	}
	return &quickFind{tab}
}

func (qf *quickFind) union(p,q int) {
	if qf.connected(p,q) {
		return
	}
	toFind := qf.find(p)
	toSet := qf.find(q)
	for i := 0; i < len(qf.tab); i++ {
		if qf.find(i) == toFind {
			qf.tab[i] = toSet
		}
	}
}

func (qf *quickFind) find(p int) int{
	return qf.tab[p]
}

func (qf *quickFind) connected(p,q int) bool{
	return qf.find(p) == qf.find(q)
}

func (qf *quickFind) count() int{
	return len(qf.connectedComponents())
}

func (qf *quickFind) connectedComponents() [][]int{
	m := map[int][]int{}
	for i, v := range qf.tab {
		ids := m[v]
		ids = append(ids, i)
		m[v] = ids
	}
	out := [][]int{}
	for _, group := range m {
		out = append(out, group)
	}
	return out
}