package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/dgrijalva/jwt-go"
)

// JWT - praktyczne wykorzystanie HMACa do autentykacji
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
	w.Write([]byte(createToken(`someSessionId`)))
}

type MyClaims struct {
	jwt.Claims
	MySessionId string
}

func createToken(data string) string {
	c := MyClaims{
		MySessionId: data,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	// todo: we should periodically rotate keys
	token, err := t.SignedString("mySecretKey")
	if err != nil {
		log.Println("Got error during signing", err)
	}
	return token
}

// curl -XGET localhost:8080/secret -I -H "MYTOKEN: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.VFb0qJ1LRg_4ujbZoRMXnVkUgiuKq5KxWqNdbKq_G9Vvz-S1zZa9LPxtHWKa64zDl2ofkT8F6jBt_K4riU-fPg"
func handleSecret(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusNotFound)
		return
	}
	if !authenticated(r) {
		http.Error(w, "auth error", http.StatusForbidden)
		return
	}

	w.Write([]byte(`hello mr president`))
}

func authenticated(r *http.Request) bool {
	receivedToken := r.Header.Get("MYTOKEN")
	var claims *MyClaims
	token, err := jwt.ParseWithClaims(receivedToken, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid sign method")
		}
		return t, nil
	})
	if err != nil {
		log.Println("Error during parseToken", err)
		return false
	}

	log.Println("got token", token)
	log.Println("claims", claims)
	return token.Valid
}