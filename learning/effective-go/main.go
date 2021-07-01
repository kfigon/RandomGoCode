package main

import "fmt"

func main() {

	arrayFun()
}

func switchFun(c byte) byte {
	// no expression in switch - it's triggered on true
	switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }

	// multiple cases
    switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return 15
	}
	return 0
}

func typeSwitch(v interface{}) {

	// type assertion
	// str, ok := v.(string)

	switch t := v.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t)             // t has type bool
	case int:
		fmt.Printf("integer %d\n", t)             // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	default:
		fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
	}
}

func deferFun() {
	getStr := func() string {
		fmt.Println("getStr called")
		return "asd"
	}
	foo := func() {fmt.Println("foo")}
	bar := func(text string) {fmt.Println("bar",text)}
	
	defer bar(getStr())
	defer foo()
	
	// calls:
	// getStr called
	// foo
	// bar asd
}

func arrayFun() {
	modSlice := func(tab []int) { tab[1] = 99999 }
	modArray := func(tab [5]int) { tab[1] = 99999 }

	slice := []int{1,2,3,4}
	slice2 := make([]int, 10)
	array := [5]int{1,2,3,4,5}

	fmt.Println("slice before",slice)
	modSlice(slice)
	fmt.Println("slice after",slice) // modified

	fmt.Println("slice2 before",slice2)
	modSlice(slice2)
	fmt.Println("slice2 after",slice2) // modified

	fmt.Println("arr before",array)
	modArray(array)
	fmt.Println("arr after",array) // NOT CHANGED, arrays are passed as VALUE!!!1
	// pass a pointer if want a change
}