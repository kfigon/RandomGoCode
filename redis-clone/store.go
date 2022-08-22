package main

import "sync"

type datastore struct {
	d map[string]string
	mut sync.RWMutex
}

func newDataStore() *datastore {
	return &datastore{
		d: map[string]string{},
	}
}

func (d *datastore) get(key string) (string, bool) {
	d.mut.Lock()
	defer d.mut.Unlock()
	v, ok := d.d[key]
	return v, ok
}

func (d *datastore) store(key string, value string) {
	d.mut.Lock()
	defer d.mut.Unlock()
	d.d[key] = value
}

func (d *datastore) delete(key string) {
	d.mut.Lock()
	defer d.mut.Unlock()
	delete(d.d, key)
}

