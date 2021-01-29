package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/satori/go.uuid"
)
// sessions are just cookies with unique values to have state
// cookie value is stored in DB/memory and stores user context

// expire session - on every request set cookies MaxAge, cookie will be invalid (but will be present in our map)
// periodically go through sessions and clean old. Need lastActivity field
func greet(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		http.Error(w, "Not authorized", http.StatusForbidden)
		return
	}
	session.numberOfEntrance++
	io.WriteString(w, fmt.Sprintf("Hi there my user. Your entry number: %v", session.numberOfEntrance))
}

func getSession(r *http.Request) (*userContext, error) {
	c, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}
	sessionData, ok := sessions[c.Value]
	if !ok || sessionData == nil {
		return nil, fmt.Errorf("Invalid session")
	}
	return sessionData, nil
}

type userContext struct {
	numberOfEntrance int
}
// usually this should be separate space:
// dbSessions map[sessionId]userId
// users map[userId]userData
var sessions = map[string]*userContext{}

func login(w http.ResponseWriter, r *http.Request) {
	sessionID := uuid.Must(uuid.NewV4()).String()
	sessions[sessionID] = &userContext{0}

	http.SetCookie(w, &http.Cookie{
		Name:"session",
		Value:sessionID,
		// Secure: true, // only when https
		HttpOnly: true, // not accessed by JS, only http
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}