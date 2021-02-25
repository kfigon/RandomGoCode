package getlist

import (
	"io"
	"html/template"
)

type GetListView struct {
	controller *GetController
}

func CreateView(c *GetController) *GetListView {
	return &GetListView{c}
}

func (v *GetListView) Render(w io.Writer) {
	tpl := template.Must(template.ParseFiles("templates/base.html", "templates/list.html"))
	tpl.Execute(w, v.controller.getList())
}