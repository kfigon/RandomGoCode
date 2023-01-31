package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		interpreterMode()
	} else if len(os.Args) == 2 {
		fileMode(os.Args[1])
	} else {
		fmt.Println("Invalid number of arguments")
	}
}

func fileMode(fileName string) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cant open file %v: %v\n", fileName, err)
		return
	}
	t, err := lex(string(b))
	if err != nil {
		fmt.Println("Got error:", err)
		return
	}
	fmt.Println(t)
}

func interpreterMode() {
	fmt.Println("Welcome to lox interpreter")
	fmt.Println("type 'quit' to exit")
	for true {
		var line string
		fmt.Print("> ")
		fmt.Scanln(&line)
		
		if line == "quit" {
			fmt.Println("Bye")
			return
		} else if line != "" {
			t, err := lex(line)
			if err != nil {
				fmt.Println("got error: ", err)
				continue
			}
			exp, errs := NewParser(t).Parse()
			if len(errs) > 0 {
				fmt.Println("got errors: ", errs)
			} else {
				fmt.Println(exp)
			}
		}
	}
}
