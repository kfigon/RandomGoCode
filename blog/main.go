package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello")
	repo,err := newRepo()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer repo.conn.Close()

	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.RequestURI, "/blog/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		blog, err := repo.queryBlog(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		template.Must(template.New("blog").Parse(`{{ . }}`)).Execute(w, blog)
	})

	http.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		blog, err := repo.queryAllBlogs()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		template.Must(template.New("blogs").Parse(`{{ . }}`)).Execute(w, blog)
	})

	http.ListenAndServe(":8080", nil)
}