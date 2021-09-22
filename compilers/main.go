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
	defer func() {
		fmt.Println("\nDone", time.Since(start))
	}()

	fmt.Println("Welcome to my bad compiler")
	
	cfg := parseCliArgsToConfig()
	if err := validate(cfg); err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open(cfg.filePath)
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
	if cfg.lex {
		tokens := lexer.Tokenize(fileContent)
		fmt.Println("Tokens:")
		for _, t := range tokens {
			fmt.Println(t)
		}
		return
	} else if cfg.runRepl {
		fmt.Println("Running repl...")
		// todo
		return
	}
}

type config struct {
	filePath string
	lex bool
	runRepl bool
}

func parseCliArgsToConfig() config {
	var cfg config
	flag.StringVar(&cfg.filePath, "file", "", "path to file with code")
	flag.BoolVar(&cfg.lex, "lex", false, "lexes file and prints lexer output")
	flag.BoolVar(&cfg.runRepl, "repl", false, "run REPL, ignores all other params")
	flag.Parse()

	return cfg
}

func validate(c config) error {
	if c.lex && c.filePath == "" {
		return fmt.Errorf("filepath not provided")
	} 
	return nil
}