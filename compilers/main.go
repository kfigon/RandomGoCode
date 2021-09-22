package main

import (
	"bufio"
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
	
	if cfg.runRepl {
		handleRepl(cfg)
	} else if cfg.lex {
		parseTokens(cfg.filePath)
	}
}

func readFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error when reading file: %v", err)
	}
	defer file.Close()

	fileByteContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error in reading file: %v", err)
	}
	return string(fileByteContent), nil
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
	if !c.runRepl && c.filePath == "" {
		return fmt.Errorf("filepath not provided")
	} 
	return nil
}

func handleRepl(cfg config) {
	fmt.Println("Running repl...")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		
		if cfg.lex {
			out := lexer.Tokenize(text)
			fmt.Println(out)
		} 
	}

	fmt.Println("Closing repl")
}

func parseTokens(filePath string) {
	fileContent, err := readFileContent(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	tokens := lexer.Tokenize(fileContent)
	fmt.Println("Tokens:")
	for _, t := range tokens {
		fmt.Println(t)
	}
}