package main

import (
	"net/http"
)

func allowedMethod(method string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "invalid method", http.StatusMethodNotAllowed)
			return
		}
		fn(w,r)
	}
}

func get(fn http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(http.MethodGet, fn)
}

func post(fn http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(http.MethodPost, fn)
}

func put(fn http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(http.MethodPut, fn)
}

func delete(fn http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(http.MethodDelete, fn)
}