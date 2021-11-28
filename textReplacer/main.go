package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
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

	content, err := os.ReadFile(*pathToFile)
	if err != nil {
		fmt.Println("Error during opening file:", err)
		return
	}

	stringContent := []rune(string(content))
	r := newRander(*randThreshold)
	
	out := ""
	i := 0
	for i < len(stringContent) {
		wordResult, found := findWord(stringContent[i:])
		if found {
			if r.pass() {
				out += "......."
			} else {
				out += string(wordResult)
			}
			i += len(wordResult)
		} else {
			out += string(stringContent[i])
			i++
		}
	}

	fmt.Println(out)
}

var wordReg = regexp.MustCompile(`^(\w+)`)
func findWord(content []rune) ([]rune, bool) {
	res := wordReg.FindSubmatch([]byte(string(content)))
	if len(res) < 2 {
		return []rune{}, false
	}
	bytes := res[1]
	out := []rune{}
	for _, b := range bytes {
		out = append(out, rune(b))
	}
	return out,true
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