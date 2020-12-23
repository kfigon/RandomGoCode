package main

import (
	"time"
	"fmt"
	"math/rand"
	"math"
)

// this wont work in global!
// aaaasd := 1
// var aaaasd int = 1 // this will

func main() {
	fmt.Println("asd", rand.Intn(10), add(1,3))
	fmt.Println(math.Pi)
	
	a,b := "foo", "bar"
	fmt.Println(a,b)
	a,b = swap(a,b)
	fmt.Println(a,b)

	fmt.Println(named(1))

	// variables
	// declare
	var asd int = 1
	// asd = 2
	fmt.Println(asd)

	// type infere, declare define
	foo := "aaaa"
	fmt.Printf("type: %T, val: %v\n", foo, foo) // %v - value

	first,second,last := 1,2,3
	fmt.Println(first, second, last)

	// zero values
	var defaultInt int
	var defaultStr string
	var defaultBool bool
	fmt.Printf("int: %v, str: %q, bool: %v\n", defaultInt, defaultStr, defaultBool)

	casted := float64(defaultInt)
	fmt.Printf("casted type: %T, val: %v\n", casted, casted)

	// := does not work
	const constVal = "asd"
	fmt.Println("constVal", constVal)

	// loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum+=i
	}
	fmt.Println("sum", sum)

	sum = 0
	for sum < 10{
		sum += 1
	}
	fmt.Println("after while", sum)

	// if
	if sum > 5 {
		fmt.Println("sum > 5")
	} else if sum == 5{
		fmt.Println("==5")
	} else {
		fmt.Println("other")
	}

	// if with statement before condition. scope until end of if
	if v := math.Pow(3,2); v > 1{
		fmt.Println("if with assignment", v)
	}
	// fmt.Println(v) // undefined

	
	switch switchCondition := "3"; switchCondition {
	case "asd": fmt.Println("its asd") // does not fallback to rest, no break
	case "sad": fmt.Println("its sad")
	case fmt.Sprint(add(1,2)): fmt.Println("its 3!!!") // will be calculated only if needed
	default : fmt.Println("unknown swicth")
	}

	t := time.Now()
	switch { // instead of if elses
	case t.Hour() < 12: fmt.Println("Good morning!")
	case t.Hour() < 17: fmt.Println("Good afternoon.")
	default: fmt.Println("Good evening.")
	}

	defer fmt.Println("goodbye go!") // executes when function returns. More - it's like a stack

	// pointers!
	//https://tour.golang.org/moretypes/1
	someVal := 1
	p := &someVal
	*p = 2
	fmt.Println("someval", someVal)

	myPoint := point{1,2}
	fmt.Println(myPoint)
	structPointer := &myPoint
	structPointer.X = 2 // no need do dereference (*x).val = 2
	fmt.Println(myPoint)

	v1 := point{1, 2}  // has type Vertex
	v2 := point{X: 1}  // Y:0 is implicit
	v3 := point{}      // X:0 and Y:0
	ptr := &point{1, 2} // has type *Vertex
	fmt.Println(v1, ptr, v2, v3) // {1 2} &{1 2} {1 0} {0 0}

	// constant len
	var arr = [2]string{"hello", "go"}
	fmt.Println(arr)

	slice := []int{1,2,3}
	fmt.Println(slice)

	fmt.Println(ptr.toString())
	
	// https://tour.golang.org/methods/9
}	

func add(a int, b int) int {
	return a + b
}

// more args
func sub(a,b int) int {
	return a-b
}

// multiple return
func swap(a,b string) (string, string) {
	return b,a
}

// named returns. naked return
func named(a int) (x, y int) {
	x = a+1
	y = a-1
	return
}

type point struct {
	X int
	Y int
}

// method
// reference. Can be passed by value, so it's a shallow copy
func (p *point) toString() string {
	return fmt.Sprintf("(%v;%v)", p.X, p.Y)
}