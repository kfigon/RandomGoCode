package main

// JWT - praktyczne wykorzystanie HMACa do autentykacji

import (
	"net/http"
	"log"
	"github.com/dgrijalva/jwt-go"
)

// JWT
// {standard fields}.{custom fields}.{signature (hmac)}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleIdx)
	mux.HandleFunc("/secret", handleSecret)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// basic auth
// curl -XPOST -u ziom:asd localhost:8080/login -I
func handleIdx(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusNotFound)
		return
	}
	user, pass, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "basic auth not ok", http.StatusForbidden)
		return
	}

	if user != "ziom" || pass != "asd" {
		http.Error(w, "invalid credentials", http.StatusForbidden)
		return
	}

	log.Println("Login ok")
}

func handleSecret(w http.ResponseWriter, r *http.Request) {

}