package main

import (
	"io"
	"sync"
	"fmt"
	"encoding/json"
	"strings"
	"time"
	"log"
	"net/http"
)

func main() {
	service := newService()

	mux := http.NewServeMux()
	mux.HandleFunc("/save", logMiddleware(service.handleSave))
	mux.HandleFunc("/", logMiddleware(service.handleRedirect))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Request start", r.URL)
		next(w,r)
		log.Println("Request end", r.URL, "took", time.Since(start))
	}
}

type Request struct {
	URL string `json:"url"`
}
func (r *Request) fromJson(reader io.Reader) bool {
	err := json.NewDecoder(reader).Decode(r)
	return err == nil && r.URL != ""
}


type Response struct {
	URL string `json:"url"`
}
func (r *Response) toJson(writer io.Writer) {
	json.NewEncoder(writer).Encode(r)
}


type service struct {
	urls map[string]string
	mutex sync.Mutex
}

func newService() *service {
	return &service {
		urls: make(map[string]string),
		mutex: sync.Mutex{},
	}
}

func (s *service) save(url string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for id := range s.urls {
		if s.urls[id] == url {
			log.Println(url,"already present")
			return id
		}
	}

	redirectId := fmt.Sprint(len(s.urls)+1)
	s.urls[redirectId] = url
	log.Println(url,"saved")
	return redirectId
}

func (s *service) read(id string) (string,bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	v,ok := s.urls[id]
	return v,ok
}

// curl -XPOST -d '{"url":"https://www.wykop.pl/"}' localhost:8080/save -i
func (s *service) handleSave(w http.ResponseWriter, r *http.Request) {
	body := &Request{}
	if !body.fromJson(r.Body) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	url := fmt.Sprint("http://localhost:8080/", s.save(body.URL))
	(&Response{url}).toJson(w)
}

// curl localhost:8080/1 -i
func (s *service) handleRedirect(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.RequestURI, "/")
	v, ok := s.read(path)
	if !ok {
		log.Println("Invalid input:", path)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	log.Println("Redirect to:", v)
	http.Redirect(w,r, v, http.StatusSeeOther)
}
