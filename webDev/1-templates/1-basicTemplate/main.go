package main

import (
	"log"
	"text/template"
	"os"
	"time"
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

func processAdHocTemplate(data interface{}, templateString string)  {
	tmpl := template.Must(template.New("").Parse(templateString))
	must(tmpl.ExecuteTemplate(os.Stdout, "", data))
}

// {{.}} - any data at the point of execution
// {{$imie := .}} - assignment to variable in a template
// {{range .}} {{.}} {{end}}
// {{.NazwaPolaWStrukturze}}
// {{funkcja .}} - wola func z argumentem
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
	}{ "Jan", "Kowalski", 15 }
	parseFile("complexTemplate", myStruct)
	

	funcs := template.FuncMap{
		"foo": func(input string) string {
			return input + " AAAAAAAAAAAAAAAA!"
		},
		"fooNoArg": func() string {
			return "func output!"
		},
	}
	tmpl := template.Must(template.New("").Funcs(funcs).Parse(`this is template with function: {{foo .}}, second: {{fooNoArg}}
`))
	must(tmpl.ExecuteTemplate(os.Stdout, "", "my input data"))

	processAdHocTemplate(time.Now(),
	`passed date {{.}}`)

	log.Println("done")
}

func must(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
	}
}