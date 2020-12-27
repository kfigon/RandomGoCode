package main

import (
	"fmt"
	"log"
	"kamil.com/greetings"
)

func main()  {
	fmt.Println("this is my hello world program")
	greeting, err := greetings.Hello("Adam")
	if err != nil {
		log.Fatal("got error from module: ", err.Error())
	} 
	fmt.Println("msg from other go module:", greeting)
}