package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("asd")
	fmt.Println(parse(`# dmidecode 3.1
Getting SMBIOS data from sysfs.
SMBIOS 2.6 present.

Handle 0x0001, DMI type 1, 27 bytes
System Information
	Manufacturer: LENOVO
	Product Name: 20042
	Version: Lenovo G560
	Serial Number: 2677240001087
	UUID: CB3E6A50-A77B-E011-88E9-B870F4165734
	Wake-up Type: Power Switch
	SKU Number: Calpella_CRB
	Family: Intel_Mobile`))
}

type dmiOutput map[string]section

type section struct {
	handle string
	title string
	props map[string]property
}

type property struct {
	name string
	value string
	items []string
}

func parse(input string) dmiOutput {
	lines := strings.Split(input, "\n")
	result := dmiOutput{}

	lastSection := ""
	lastProperty := ""

	i := 0
	for i < len(lines) {
		line := lines[i]
		currentIndent := indentantionLevel(line)

		if line == "" {
			i++
			continue
		}

		if currentIndent == 0 && strings.HasPrefix(line, "Handle") {
			if i + 1 >= len(lines) || 
				indentantionLevel(line) != 0 || 
				indentantionLevel(lines[i+1]) != 0 {
				i++
				continue
			}
			s := section{}
			s.props = map[string]property{}
			s.handle = line

			i++
			s.title = lines[i]
			lastSection = s.title
			result[s.title] = s
		} else if currentIndent == 1 && lastSection != "" {
			pair := strings.Split(line, ":")
			if len(pair) != 2 {
				i++
				continue
			}

			lastProperty = strings.TrimSpace(pair[0])
			prop := property{
				name: lastProperty,
				value: strings.TrimSpace(pair[1]),
			}
			result[lastSection].props[lastProperty] = prop
		} else if currentIndent == 2 && lastSection != "" && lastProperty != "" {
			p := result[lastSection].props[lastProperty]
			p.items = append(p.items, strings.TrimSpace(line))
			result[lastSection].props[lastProperty] = p
		}
		i++
	}

	return result
}

func indentantionLevel(line string) int {
	ind := 0
	for _, c := range line {
		if c == '\t' {
			ind++
		} else {
			break
		}
	}
	return ind
}