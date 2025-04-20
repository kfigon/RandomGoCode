package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// todo: hint file support

func main() {
	fmt.Println("hello")

	storage := newStore(NewMemBuffer())

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		s = strings.TrimSpace(s)

		if err != nil {
			fmt.Println("error reading line", err)
			return
		} else if s == "quit" || s == "exit" {
			fmt.Println("bye bye")
			return
		}

		// todo: better parsing
		fields := strings.Fields(s)
		if len(fields) >= 2 && fields[0] == "get" {
			got, err := storage.Get(fields[1])
			if err != nil {
				fmt.Printf("get %q error: %v\n", fields[1], err)
			} else {
				fmt.Printf("get %q: %v\n", fields[1], string(got))
			}
		} else if len(fields) >= 3 && fields[0] == "put" {
			key := fields[1]
			rest := strings.TrimPrefix(s, fields[0]+ " " +fields[1]+" ")
			storage.Put(key, []byte(rest))
			fmt.Printf("put %q done\n", key)
		} else if len(fields) == 1 && fields[0] == "keys" {
			keys := storage.Keys()
			fmt.Println("list of keys:", keys)
		} else {
			fmt.Println("invalid command:", s)
		}
	}

}
