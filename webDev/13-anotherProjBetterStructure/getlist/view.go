package getlist

import (
	"io"
	"html/template"
	"mywebapp/model"
)

func Render(w io.Writer, elements []model.Element) {
	tpl := template.Must(template.ParseFiles("templates/base.html", "templates/list.html"))
	tpl.Execute(w, elements)
}