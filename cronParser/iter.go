package main

type iter[T any] struct {
	vs []T
	idx int
}

func toIter[T any](vals []T) *iter[T] {
	return &iter[T]{vals,0}
}

func (i *iter[T]) current() (T, bool) {
	if i.idx >= len(i.vs) {
		var out T
		return out, false
	}
	return i.vs[i.idx], true
}

func (i *iter[T]) next() {
	i.idx++
}

func (i *iter[T]) peek() (T, bool) {
	if i.idx+1 >= len(i.vs) {
		var out T
		return out, false
	}
	return i.vs[i.idx+1], true
}