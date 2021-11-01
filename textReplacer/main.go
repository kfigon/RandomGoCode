package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// go env to see current settings
// $Env:GOOS = "darwin"; $Env:GOARCH = "amd64"; go build
// restore old settings
func main() {
	pathToFile := flag.String("path", "", "path to file with text")
	randThreshold := flag.Int("rand", 2, "rand threshold <0-10>. Higher - more blured words")
	flag.Parse()

	if *pathToFile == "" {
		fmt.Printf("Invalid path to file: %q\n", *pathToFile)
		return
	} else if *randThreshold < 0 || *randThreshold > 10 {
		fmt.Println("Invalid rand threshold:", *randThreshold)
		return
	}

	rand.Seed(time.Now().Unix())
	content, err := os.ReadFile(*pathToFile)
	if err != nil {
		fmt.Println("Error during opening file:", err)
		return
	}
	stringContent := string(content)
	words := strings.Fields(stringContent)
	out := ""
	for _, w := range words {
		if rand.Intn(10) < *randThreshold {
			out += "......."
		} else {
			out += w
		}
		out += " "
	}
	fmt.Println(out)
}