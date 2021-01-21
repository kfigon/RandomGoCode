package main

import (
	"log"
	"text/template"
	"os"
)

func parseFile(path string, data interface{})  {
	log.Println("parsing", path)
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
	parseFile("letterTemplate", "Asd")
	parseFile("listTemplate", []string{"Gandi","Gates","Kacyznski"})

	// myStruct := struct {
	// 	imie string
	// 	nazwisko string
	// 	wiek int
	// }{
	// 	"Jan", "Kowalski", 15,
	// }
	// parseFile("complexTemplate", myStruct)


	log.Println("\ndone")
}