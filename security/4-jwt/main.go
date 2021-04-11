package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"net/http"
)

// JWT - praktyczne wykorzystanie HMACa do autentykacji
// {standard fields}.{custom fields}.{signature (hmac)}
//these parts are base64 encoded. It's not encrypted!!!1
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/secret", handleSecret)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// basic auth
// curl -XPOST -u ziom:asd localhost:8080/login -v
func handleLogin(w http.ResponseWriter, r *http.Request) {
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
	io.WriteString(w, createToken(`someSessionId`))
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

func createToken(data string) string {
	j := &jwtHelper{}
	return j.createToken(data)
}

func authenticated(r *http.Request) bool {
	j := &jwtHelper{}
	return j.authenticated(r)
}

type jwtHelper struct{}

func (j *jwtHelper) key() []byte                 { return []byte("mySecretKey") }
func (j *jwtHelper) alg() *jwt.SigningMethodHMAC { return jwt.SigningMethodHS512 }
func (j *jwtHelper) headerKey() string           { return "MYTOKEN" }
func (j *jwtHelper) tokenKey() string            { return "foo" }

func (j *jwtHelper) createToken(data string) string {
	token := jwt.New(j.alg())
	token.Claims = jwt.MapClaims{
		"foo": data,
	}

	tokenString, err := token.SignedString(j.key())
	if err != nil {
		log.Println("Error during encoding", err)
		return ""
	}
	log.Println("Token created", tokenString)
	return tokenString
}

func (j *jwtHelper) authenticated(r *http.Request) bool {
	receivedToken := r.Header.Get(j.headerKey())
	log.Println("Got token", receivedToken)

	token, err := jwt.Parse(receivedToken, j.validate)
	if err != nil {
		log.Println("Got validation error:", err)
		return false
	}
	valid := token.Valid
	log.Println("Token valid:", valid)
	return valid
}

func (j *jwtHelper) validate(token *jwt.Token) (interface{}, error) {
	log.Println("Got Header:", token.Header)
	claims := token.Claims.(jwt.MapClaims)
	log.Println("Got Payload:", claims)

	if token.Header["alg"] != j.alg().Alg() {
		log.Println("Got invalid algorithm, error!")
		return nil, fmt.Errorf("Invalid algorithm")
	}

	val, ok := claims[j.tokenKey()]
	if !ok || val != "someSessionId" {
		log.Println("Got invalid claims, error!")
		return nil, fmt.Errorf("Invalid token")
	}
	return j.key(), nil
}