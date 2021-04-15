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
	values := createSampleValues()
	fmt.Println(values)
	sort.Sort(byAge(values))
	fmt.Println(values)


	fmt.Println("\nother method")
	values2 := createSampleValues()
	fmt.Println(values2)
	sort.Slice(values2, func(i,j int) bool {
		return values2[i].age < values2[j].age
	})
	fmt.Println(values2)
}

func createSampleValues() []data {
	return []data {
		{"asd", 1},
		{"sad", 3},
		{"bar", 2},
		{"foo", 5},
	}
}