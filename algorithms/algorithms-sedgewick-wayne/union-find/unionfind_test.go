package unionfind

import (
	"fmt"
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

func populateQuickUnion(num int, data []pair[int,int]) *quickUnion {
	qu := newQuickUnion(num)
	for _, v := range data {
		qu.union(v.a, v.b)
	}
	return qu
}


func TestUnionFind(t *testing.T) {
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

	tdt := []struct{
		desc string
		algo unionFind
	}{
		{"quick find", populateQuickFind(10, data)},
		{"quick union", populateQuickUnion(10, data)},
	}
	for _, tc := range tdt {
		t.Run(tc.desc, func(t *testing.T) {

			assert.Equal(t, 2, tc.algo.count())
			assert.Equal(t, [][]int{{0, 1, 2, 5, 6, 7}, {3, 4, 8, 9}}, tc.algo.connectedComponents())
			
			connectedPairs := []pair[int,int] {
				{0,1},{1,0},{0,2},{2,0},{0,7}, {5,2}, {6,1},
				{3,8}, {8,4}, {9,3}, {3,9}, {4,9},{0,0},
			}
			for _, p := range connectedPairs {
				t.Run(fmt.Sprintf("connected %v-%v", p.a, p.b), func(t *testing.T) {
					assert.True(t, tc.algo.connected(p.a, p.b))
				})
			}
			
			notconnectedPairs := []pair[int,int] {
				{1,3},{6,9},{8,7},
			}
			for _, p := range notconnectedPairs {
				t.Run(fmt.Sprintf("not connected %v-%v", p.a, p.b), func(t *testing.T) {
					assert.False(t, tc.algo.connected(p.a, p.b))
				})
			}
		})	
	}
}

type pair[T any, V any] struct{
	a T
	b V
}

type unionFind interface{
	union(int,int)
	find(int) int
	count() int
	connected(int, int) bool
	connectedComponents() [][]int
}


// find is very quick O(1), but union is slow O(n)
// it's not good for big data sets - at least if you are going to modify it a lot
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



type quickUnion struct{
	tab []int
}

func newQuickUnion(num int) *quickUnion {
	tab := []int{}
	for i := 0; i < num; i++ {
		tab = append(tab, i)
	}
	return &quickUnion{tab}
}

func (qu *quickUnion) union(p,q int) {
}

func (qu *quickUnion) find(p int) int{
	return -1
}

func (qu *quickUnion) connected(p,q int) bool{
	return false
}

func (qu *quickUnion) count() int{
	return len(qu.connectedComponents())
}

func (qu *quickUnion) connectedComponents() [][]int{
	return nil
}
