package main

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"log"
	"crypto/hmac"
	"crypto/sha512"
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

	hash := hmac.New(sha512.New, secretKey)
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

// curl -i -H "MYTOKEN: 1958d869e16437892f9a8b0366889da9559b6327a880825bb18cb8401d2fc4b47ff80eeef93b7860530b22c76c4fdd258d741c9db9658f1bb37e0d08311af171" localhost:8080/resource
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