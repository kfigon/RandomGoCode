package main

import (
	"io"
	"net/http"
)

// cookie - small file that server can write to clients machine!
// used for adding state to http (session)
// cookies are per domain
// browser checks if we got cookie for that domain and just sends that every time, on every request

	// Name  string
	// Value string
	// Path       string    // optional
	// Domain     string    // optional
	// Expires    time.Time // optional
	// RawExpires string    // for reading cookies only
	// 	// MaxAge=0 means no 'Max-Age' attribute specified.
	// 	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// 	// MaxAge>0 means Max-Age attribute present and given in seconds
	// MaxAge   int
	// Secure   bool
	// HttpOnly bool
	// SameSite SameSite
	// Raw      string
	// Unparsed []string // Raw text of unparsed attribute-value pairs

func login(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:"ziomCookie",
		Value:"asd",
	})
	io.WriteString(w, "Hi!")
}

func check(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("ziomCookie")
	if err != nil {
		http.Error(w, "", http.StatusForbidden)
		return
	}
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/", check)
	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}