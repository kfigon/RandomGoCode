package main

import (
	"fmt"
)

func main()  {
	fmt.Println("FIBO:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v -> %v\n", i, fibo(i))
	}

	fmt.Println("\nFIZZBUZZ:")
	for i := 0; i < 17; i++ {
		fmt.Printf("%v -> %v\n", i, fizzbuzz(i))
	}
	fmt.Println("\nCHAR COUNT:")
	fmt.Printf("%v -> %v\n", "Kamil", charCount("Kamil"))
	fmt.Printf("%v -> %v\n", "golang", charCount("golang"))
	fmt.Printf("%v -> %v\n", "missisipi", charCount("missisipi"))

	fmt.Println("\nGENERATE EVEN:")
	fmt.Printf("%v -> %v\n", 10, generateEven(10))
}

func fibo(a int) int {
	if a < 2{
		return 1
	}
	return fibo(a-1) + fibo(a-2)
}

func fizzbuzz(a int) string {
	out := ""
	if a % 3 == 0{
		out += "fizz"
	}
	if a % 5 == 0 {
		out += "buzz"
	}

	if out == "" {
		return fmt.Sprint(a)
	}
	return out
}

func charCount(text string) map[string]int {
	dict := make(map[string]int)
	for _, c := range text{
		dict[string(c)]++
	}
	return dict
}

func generateEven(ln int) []int {
	var out []int
	// out := make([]int, 0) // initial len

	for i := 0; i < ln; i++ {
		if i % 2 == 0{
			out = append(out, i)
		}
	}
	return out
}
	