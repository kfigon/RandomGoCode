package main

import (
	"strings"
	"io"
	"net/http"
)

func helloFunc(response http.ResponseWriter, req* http.Request) {
	io.WriteString(response, "Hi!")
}

func cat(response http.ResponseWriter, req* http.Request) {
	io.WriteString(response, "Hello Cat!")
}

func queried(response http.ResponseWriter, req* http.Request) {
	req.ParseForm()
	responseString := "Got "
	i := 0
	for k,v := range req.Form {
		responseString += k+"="+v[0]
		if i < len(req.Form) -1 {
			responseString += " and "
		}
		i++
	}
	io.WriteString(response, responseString)
}

// need to parse manually :(
func parametered(response http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/foo/")
	if len(id) == 0 {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	io.WriteString(response, "Id: " + id)
}

// mux == multiplexer == "server" (dispatch servlet z javy)
func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloFunc)
	mux.HandleFunc("/cat", cat) // /cat/ - match everything after cat
	mux.HandleFunc("/queried", queried)
	mux.HandleFunc("/foo/", parametered)

	return mux
}

func main() {
	// if nil in ListenAndServe(port, nil) - we can use:
	// http.Handle("/", cat). This will add to DefaultServeMux
	http.ListenAndServe(":8080", createMux())
}