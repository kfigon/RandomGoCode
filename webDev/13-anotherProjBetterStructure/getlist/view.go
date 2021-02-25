package getlist

import (
	"io"
	"html/template"
	"mywebapp/model"
)

var BASE_PATH string = ""

func Render(w io.Writer, elements []model.Element) {
	tpl := template.Must(template.ParseFiles(BASE_PATH + "templates/base.html", BASE_PATH + "templates/list.html"))
	tpl.Execute(w, elements)
}