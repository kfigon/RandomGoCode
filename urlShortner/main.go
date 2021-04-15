package main

import (
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
type Response struct {
	URL string `json:"url"`
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
			return id
		}
	}

	redirectId := fmt.Sprint(len(s.urls)+1)
	s.urls[redirectId] = url
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
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil || body.URL == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(Response{
		URL: fmt.Sprint("http://localhost:8080/", s.save(body.URL)),
	})
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
