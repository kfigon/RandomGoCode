package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"programming-lang/lexer"
	"time"
)

func main() {
	start := time.Now()
	defer func(){
		fmt.Println("\nDone", time.Since(start))
	}()

	fmt.Println("Welcome to my bad compiler\n")

	filePath := flag.String("file", "", "path to file with code")
	lex := flag.Bool("lex", false, "print lexer output")
	flag.Parse()

	if filePath == nil || *filePath == "" {
		fmt.Println("Filepath not provided")
		return
	}
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error when reading file:", err)
		return
	}
	defer file.Close()

	fileByteContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error in reading file:", err)
		return
	}
	fileContent := string(fileByteContent)
	if *lex {
		tokens := lexer.Tokenize(fileContent)
		fmt.Println("Tokens:")
		for _, t := range tokens {
			fmt.Println(t)
		}
		return
	}
}