package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("asd")
	if len(os.Args) < 2 {
		fmt.Println("no file provided")
		return
	}
	path := os.Args[1]
	fmt.Println(path)
}