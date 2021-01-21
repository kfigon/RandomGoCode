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
// {{range .}} {{.}} {{end}}
// {{.NazwaPolaWStrukturze}}
func main() {
	parseFile("basicTemplate.gohtml", map[string]string {"userName": "Jacek"})
	parseFile("letterTemplate", "Asd")
	parseFile("listTemplate", []string{"Gandi","Gates","Kacyznski"})
	parseFile("mapTempl", map[string]int{
		"first":1,
		"second":2,
		"third":3,
	})

	myStruct := struct {
		Imie string
		Nazwisko string
		Wiek int
	}{
		"Jan", "Kowalski", 15,
	}
	parseFile("complexTemplate", myStruct)


	log.Println("done")
}