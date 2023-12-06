package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	Id int
	Desc string
	Done bool
}

// todo: return templates on errors
// todo: embedd templates
func main() {
	todoList := &Protected[[]Todo]{}
	todoList.Access(func(todos *[]Todo) {
		*todos = append(*todos, Todo{Id: 0, Desc: "buy things", Done: false})
		*todos = append(*todos, Todo{Id: 1, Desc: "play games", Done: true})
		*todos = append(*todos, Todo{Id: 2, Desc: "code", Done: false})
	})
	
	port := 3000
	http.HandleFunc("/", get(index))
	http.HandleFunc("/todos", get(todoHandler(todoList)))
	http.HandleFunc("/todo", post(addItem(todoList)))
	http.HandleFunc("/todo/", put(flipDone(todoList)))

	fmt.Println("starting on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err)
	}
}

func todoHandler(todoList *Protected[[]Todo]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := []Todo{}
		todoList.Access(func(list *[]Todo) {
			todos = append(todos, *list...)
		})

		t := template.Must(template.ParseFiles("templates/todos.html", "templates/single_element.html"))
		if err := t.Execute(w, todos); err != nil {
			fmt.Println(err)
		}	
	}
}

func addItem(todoList *Protected[[]Todo]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "cant access request data", http.StatusBadRequest)
			return
		}

		v := strings.TrimSpace(r.FormValue("description"))
		if v == "" {
			http.Error(w, "no description proved", http.StatusBadRequest)
			return
		}

		var newItem Todo

		todoList.Access(func(list *[]Todo) {
			newItem = Todo{
				Desc: v,
				Id: len(*list),
			}
			*list = append(*list, newItem)
		})

		t := template.Must(template.ParseFiles("templates/single_element.html"))
		if err := t.Execute(w, newItem); err != nil {
			fmt.Println(err)
		}
	}
}

func flipDone(todoList *Protected[[]Todo]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.RequestURI, "/todo/"))
		if err != nil {
			http.Error(w, "no id provided", http.StatusBadRequest)	
			return
		}
		
		todoList.Access(func(list *[]Todo) {
			if id < 0 || id >= len(*list) {
				err = fmt.Errorf("invalid id provided")
				return
			}
			el := &(*list)[id]
			el.Done = !el.Done
		})

		if err != nil {
			http.Error(w, "invalid id provided", http.StatusBadRequest)	
			return
		}
	}
}

