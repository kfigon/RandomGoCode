package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// union find - algorithms for dynamic connectivity problems
// check if 2 elements of a set are connected

// 3 algorithms: quick find, quick union and weighted quick union
func TestQuickFind(t *testing.T) {
	data := []struct{
		a,b int
	}{
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
	qf := newQuickFind(10)

	for _, v := range data {
		qf.union(v.a, v.b)
	}

	assert.Equal(t, 2, qf.count())
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

}

func (qf *quickFind) find(p int) int{
	return -1
}

func (qf *quickFind) connected(p,q int) bool{
	return false
}

func (qf *quickFind) count() int{
	return -1
}