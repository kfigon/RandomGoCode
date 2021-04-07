package main

import (
	"encoding/base64"
	"net/http"
	"log"
	"strings"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleReq)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

// curl -u myuser:secretPassword localhost:8080 -v
func handleReq(w http.ResponseWriter, req *http.Request) {
	if !authenticated(req) {
		w.WriteHeader(http.StatusForbidden)
		return
	} 
	
	w.Write([]byte(`hello mr president`))
}

func authenticated(req *http.Request) bool {
	header := req.Header.Get("Authorization")
	if header == "" {
		return false
	}
	stripped := strings.Replace(header, "Basic ", "", -1)
	decoded, _ := base64.StdEncoding.DecodeString(stripped)
	log.Println("Got request: ", string(decoded))

	userData := strings.Split(string(decoded), ":")
	if len(userData) != 2 ||
		userData[0] != "myuser" ||
		userData[1] != "secretPassword" {
		
		log.Println("Invalid pass, reject")
		return false	
	}

	log.Println("Logged in")
	return true
}