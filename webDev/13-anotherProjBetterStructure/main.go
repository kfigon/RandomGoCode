package main

import (
	"mywebapp/model"
	"net/http"
	"log"
	"mywebapp/getlist"
	"mywebapp/insert"
	"mywebapp/landing"
)

func createMux(c *getlist.GetController, i *insert.InsertController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req* http.Request) {
		landing.Render(w)
	})

	mux.HandleFunc("/list", func(w http.ResponseWriter, req* http.Request) {
		getlist.Render(w, c.GetList())
	})

	mux.HandleFunc("/addNew", func(w http.ResponseWriter, req* http.Request) {
		defer http.Redirect(w,req,"/list",http.StatusSeeOther)

		// todo error handling on both methods?
		model,_ := createModelFromRequest(req)
		i.Insert(model)
	})
	return mux
}

func createModelFromRequest(req *http.Request) (model.Element, error) {
	err := req.ParseForm()
	if err != nil {
		return model.Element{}, err
	}

	newTodoEntry := model.Element{
		Name: req.FormValue("title"),
		Date: req.FormValue("date"),
	}
	return newTodoEntry, nil
}

func main() {
	getListRepo := getlist.MakeDb()
	insertRepo := &insertDb{getListRepo}

	getListController := getlist.CreateGetListController(getListRepo)
	insertController := insert.CreateInsertController(insertRepo)

	log.Fatal(http.ListenAndServe(":8080", createMux(getListController, insertController)))
}


type insertDb struct {
	db *getlist.MyDb
}

func (i *insertDb) Insert(el model.Element) {
	i.db.Elems = append(i.db.Elems, el)
}