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
	http.HandleFunc("/todo/", modify(flipDone(todoList), deleteItem(todoList)))

	fmt.Println("starting on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.ExecuteTemplate(w, "index", nil); err != nil {
		fmt.Println(err)
	}
}

func todoHandler(todoList *Protected[[]Todo]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := []Todo{}
		todoList.Access(func(list *[]Todo) {
			todos = append(todos, *list...)
		})

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "todos", todos); err != nil {
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
			lastId := 0
			for _,i := range *list {
				lastId = i.Id
			}
			newItem = Todo{
				Desc: v,
				Id: lastId+1,
			}
			*list = append(*list, newItem)
		})

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "single_element", newItem); err != nil {
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
		} else if id < 0 {
			err = fmt.Errorf("invalid id provided")
			return
		}

		todoList.Access(func(list *[]Todo) {
			for i := 0; i < len(*list); i++ {
				if (*list)[i].Id == id {
					(*list)[i].Done = !(*list)[i].Done
					break
				}
			}
			err = fmt.Errorf("invalid Id")
		})

		if err != nil {
			http.Error(w, "invalid id provided", http.StatusBadRequest)	
			return
		}
	}
}

func modify(flipFn http.HandlerFunc, delteFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			flipFn(w,r)
			return
		} else if r.Method == http.MethodDelete {
			delteFn(w,r)
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func deleteItem(todoList *Protected[[]Todo]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.RequestURI, "/todo/"))
		if err != nil {
			http.Error(w, "no id provided", http.StatusBadRequest)	
			return
		} else if id < 0 {
			http.Error(w, "invalid id provided", http.StatusBadRequest)
			return
		}
		
		var todos []Todo
		todoList.Access(func(list *[]Todo) {
			todos = *list
			splitEl := -1
			for i, el := range *list {
				if el.Id == id {
					el.Done = !el.Done
					splitEl = i
					break
				}
			}
			if splitEl == -1 {
				return
			}

			*list = append((*list)[:splitEl], (*list)[splitEl+1:]...)
			todos = *list
		})

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "todos", todos); err != nil {
			fmt.Println(err)
		}	
	}
}


