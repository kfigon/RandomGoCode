package main

import (
	"fmt"
	"sort"
)

type data struct {
	name string
	age int
}

func (d data) String() string {
	return fmt.Sprintf("(%v;%v)", d.name, d.age)
}

type byAge []data
func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	values := []data {
		{"asd", 1},
		{"sad", 3},
		{"bar", 2},
		{"foo", 5},
	}

	fmt.Println(values)
	sort.Sort(byAge(values))
	fmt.Println(values)
}