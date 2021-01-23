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

func processAdHocTemplate(data interface{}, templateString string)  {
	tmpl := template.Must(template.New("").Parse(templateString))
	must(tmpl.ExecuteTemplate(os.Stdout, "", data))
}

// {{.}} - any data at the point of execution
// {{$imie := .}} - assignment to variable in a template
// {{range .}} {{.}} {{end}}
// {{.NazwaPolaWStrukturze}}
// {{funkcja .}} - wola func z argumentem
// {{funkcja}} - wola func bezz argumentu
// {{. | foo1 | foo2}} - pipeline - wez wartosc, zawolaj foo1, wynik do foo2
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

	// using predefined funs
	processAdHocTemplate([]int{1,2},
`passed 2 values, are they equal? {{index . 0}} and {{index . 1}} ?
{{$first := index . 0}} {{$second := index . 1}}-> {{eq $first $second}}
{{$res := eq $first $second}}
{{if $res}} yes, equal!
{{else }} no, not equal
{{end}}
`)

	// nested template

	processAdHocTemplate([]string{"Adam", "Pawel"},
`{{define "myTemplate"}} Hello {{.}} from custom template {{end}}
Nested template usage:
{{range .}}
{{template "myTemplate" .}} {{end}}
`)

	log.Printf("done")
}

func must(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
	}
}