package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	t.Fatal("todo")
}

type unionFind struct{}

func (u *unionFind) union(p,q int){}
func (u *unionFind) find(p int) int{
	return -1
}
func (u *unionFind) connected(p,q int) bool{
	return false
}
func (u *unionFind) count(p int) int{
	return -1
}