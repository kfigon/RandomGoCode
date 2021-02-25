package landing

import (
	"io"
	"html/template"
)

var BASE_PATH string = ""

func Render(w io.Writer) {
	tpl := template.Must(template.ParseFiles(BASE_PATH + "templates/base.html", BASE_PATH + "templates/landingPage.html"))
	tpl.Execute(w, "ziomx")
}