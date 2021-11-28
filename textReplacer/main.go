package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
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

	content, err := os.ReadFile(*pathToFile)
	if err != nil {
		fmt.Println("Error during opening file:", err)
		return
	}

	stringContent := []rune(string(content))
	r := newRander(*randThreshold)
	
	out := processInput(string(stringContent), r)

	fmt.Println(out)
}

type rander struct {
	threshold int
}

func newRander(threshold int) *rander{
	rand.Seed(time.Now().Unix())
	return &rander{threshold}
}

func (r *rander) pass() bool {
	return rand.Intn(10) < r.threshold
}


type randerInterface interface {
	pass() bool
}
func processInput(stringContent string, r randerInterface) string {
	out := ""
	word := ""
	for _, char := range stringContent {
		if unicode.IsLetter(char) {
			word += string(char)
		} else if word != "" {
			if r.pass() {
				out += "......."
			} else {
				out += word
			}
			out += string(char)
			word = ""
		} else {
			out += string(char)
		}
	}
	out += word
	return out
}