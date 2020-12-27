package main

import (
	"fmt"
	"kamil.com/greetings"
)

func main()  {
	fmt.Println("this is my hello world program")
	fmt.Println("msg from other go module:", greetings.Hello("Adam"))
}