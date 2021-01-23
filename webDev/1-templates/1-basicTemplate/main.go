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
	e := tmpl.ExecuteTemplate(os.Stdout, "", data)
	if e != nil {
		log.Fatal("Error: ", e)
	}
}

// {{.}} - any data at the point of execution
// {{$imie := .}} - assignment to variable in a template
// {{range .}} {{.}} {{end}}
// {{.NazwaPolaWStrukturze}}
// {{funkcja .}} - wola func z argumentem
// {{funkcja}} - wola func bezz argumentu
// {{. | foo1 | foo2}} - pipeline - wez wartosc, zawolaj foo1, wynik do foo2
// {{.mojaMetoda}} metody tez mozna wolac
// {{.mojaMetoda | .metodaPrzyjmujacaArg}} 

func main() {
	parseFile("basicTemplate.gohtml", map[string]string {"userName": "Jacek"})
	parseFile("letterTemplate", "Asd")
	
	processAdHocTemplate([]string{"Gandi","Gates","Kacyznski"},
`===================
Moje ziomki:
{{range .}}
* {{.}} 
{{end}}`)
	processAdHocTemplate(map[string]int{
		"first":1,
		"second":2,
		"third":3,
	},
`=================
takie typy:
{{range $key, $value := .}}
{{$key}} -> {{$value}} 
{{end}}`)

	myStruct := struct {
		Imie string
		Nazwisko string
		Wiek int
	}{ "Jan", "Kowalski", 15 }
	processAdHocTemplate(myStruct,
`=============
Dane: {{.Imie}} {{.Nazwisko}}
Wiek: {{.Wiek}}
`)
	

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
	tmpl.ExecuteTemplate(os.Stdout, "", "my input data")

	// using predefined funs
	processAdHocTemplate([]int{1,2},
`============
passed 2 values, are they equal? {{index . 0}} and {{index . 1}} ?
{{$first := index . 0}} {{$second := index . 1}}-> {{eq $first $second}}
{{$res := eq $first $second}}
{{if $res}} yes, equal!
{{else }} no, not equal
{{end}}
`)

	// nested template

	processAdHocTemplate([]string{"Adam", "Pawel"},
`=============
{{define "myTemplate"}} Hello {{.}} from custom template {{end}}
{{define "myEmptyTemplate"}} This is empty template {{end}}

Nested template usage:
{{range .}}
{{template "myTemplate" .}} {{end}}

{{template "myEmptyTemplate"}}
`)

	type Address struct {
		City string
		Street string
	}
	type Asd struct{
		Review string
	}
	hotelData := []struct {
		Name string
		Age int
		Address
		Note Asd
	} {
		{"First hotel", 10, Address{"Zgierz", "asd"}, Asd{"Smierdzi"}},
		{"Second hotel", 20, Address{"Krakow", "bar"}, Asd{"Piekny"}},
		{"Third hotel", 30, Address{"Pcim", "foo"}, Asd{"Stary"}},
	}
	hotelTemplate := `============
List of hotels:
{{range .}}
-----------
{{.Name}}, age: {{.Age}}
{{.City}}, {{.Street}}
Customer notes: {{.Note.Review}}
----------
{{end}}
`
	processAdHocTemplate(hotelData, hotelTemplate)

	log.Printf("done")
}
