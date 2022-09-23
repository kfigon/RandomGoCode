package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `
[foo]
key1 = 3
key2=some long name

[empty]

[x]
[another]
key1 = 1

[indented]
	asd = 123  `

	data := parse(input)
	fmt.Println(data)
}

type keys map[string]string
type config map[string]keys

func parse(input string) config {
	lines := strings.Split(input, "\n")
	

	sectionName := ""

	out := config{}
	for _, line := range lines {
		stripped := strings.TrimSpace(line)
		if stripped == "" || stripped == "#" {
			continue
		}

		if strings.HasPrefix(stripped, "[") && strings.HasSuffix(stripped, "]") {
			header := strings.Trim(stripped, "[")
			header = strings.Trim(header, "]")
			sectionName = header
			out[sectionName] = keys{}
		} else if sectionName != "" {
			kvPair := strings.Split(stripped, "=")
			if len(kvPair) != 2 {
				continue
			}
			out[sectionName][strings.TrimSpace(kvPair[0])] = strings.TrimSpace(kvPair[1])
		}
	}
	return out
}

func (c config) hasSection(name string) bool {
	_, ok := c[name]
	return ok
}

func (c config) hasProperty(section, property string) bool {
	_, ok := c.get(section, property)
	return ok
}

func (c config) String() string {
	out := ""
	for k := range c {
		out += fmt.Sprintf("[%v]\n", k)
		for subK, subV := range c[k] {
			out += fmt.Sprintf("%v = %v\n", subK, subV)
		}
		out += fmt.Sprintln()
	}
	return out
}

func (c config) get(section, property string) (string, bool) {
	s, ok := c[section]
	if !ok {
		return "", false
	}
	v, ok := s[property]
	return v, ok
}