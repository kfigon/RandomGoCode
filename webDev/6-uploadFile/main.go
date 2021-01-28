package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"io"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil || len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func form(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	io.WriteString(w, 
		`<form method="POST" enctype="multipart/form-data">
		<input type="file" name="usersFile">
		<input type="submit">
		</form>`)

	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("usersFile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		io.WriteString(w, `<br>`+string(fileData))
	}
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/form", form)

	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", createMux()))
}