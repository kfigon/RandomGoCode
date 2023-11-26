package main

import "sync"

type Protected[T any] struct {
	data T
	mut sync.Mutex
}

func (p *Protected[T]) Access(fn func(*T)) {
	p.mut.Lock()
	fn(&p.data)
	p.mut.Unlock()
}