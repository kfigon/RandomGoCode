package main

import (
	"sync"
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
)
type void struct{}
var emptyVal void

type login struct {
	lock *sync.Mutex
	savedTokens map[string]void
}

type LoginResponse struct {
	Token string `json:"name"`
}

func newLogin() *login {
	return &login{
		savedTokens: make(map[string]void),
		lock: &sync.Mutex{},
	}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || !l.checkPass(user, pass) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := l.createToken()
	response := LoginResponse{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
}

func (l *login) auth(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("MY_TOKEN")
	l.lock.Lock()
	_, ok := l.savedTokens[token]
	l.lock.Unlock()
	if  !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (l *login) checkPass(user string, pass string) bool {
	return user == "John" && pass == "secret"
}

func (l *login) createToken() string {
	token := uuid.New().String()

	l.lock.Lock()
	l.savedTokens[token] = emptyVal
	l.lock.Unlock()

	return token
}