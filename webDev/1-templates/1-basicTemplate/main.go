package main

import (
	"log"
	"text/template"
	"os"
)

func parseFile(path string, data interface{})  {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Fatal("Opening file failed: ", err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal("Template processing failed: ", err)
	}
}

// {{.}} - any data at the point of execution
// {{$imie := .}} - assignment to variable in a template

func main() {
	parseFile("basicTemplate.gohtml", map[string]string {"userName": "Jacek"})
	log.Println("\n")
	parseFile("letterTemplate", "Asd")

	log.Println("\ndone")
}