package main

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"log"
	"crypto/hmac"
	"crypto/sha256"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleAuth)
	mux.HandleFunc("/resource", handleSecure)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type Auth struct {
	User string 	`json:"user"`
	Password string `json:"password"`
}

var secretKey = []byte("secretKey")
var passwords = make(map[string]string)

// curl -i -d '{ "user": "ziomx", "password": "123" }' localhost:8080
func handleAuth(w http.ResponseWriter, r *http.Request) {
	log.Println("got request")
	var data Auth
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "error in decoding json", http.StatusBadRequest)
		return
	}

	log.Println("Json decoded", data)

	hash := hmac.New(sha256.New, secretKey)
	_, err = hash.Write([]byte(data.Password))
	if err != nil {
		http.Error(w, "error in writing hash", http.StatusBadRequest)
		return
	}

	signature := hex.EncodeToString(hash.Sum(nil))
	passwords[data.User] = signature
	log.Println("Got hash for", data.User, signature)
	w.Write([]byte(signature))
}

// curl -i -H "MYTOKEN: de1e6ed9340473135c23247a0dc6f1f349157c48e14835d17969106ae3c98fd5" localhost:8080/resource
func handleSecure(w http.ResponseWriter, r *http.Request) {
	if !authorised(r) {
		http.Error(w, "Unauthenticated", http.StatusForbidden)
		return
	}

	w.Write([]byte(`hello mr president`))
}

func authorised(r *http.Request) bool {
	header := r.Header.Get("MYTOKEN")
	if header == "" {
		log.Println("No token provided")
		return false
	}

	for userName := range passwords {
		storedPass := passwords[userName]
		ok := hmac.Equal([]byte(header), []byte(storedPass))
		if ok {
			log.Printf("%v authorized\n", userName)
			return true
		}
	}

	log.Println("Unauthorized")
	return false
}