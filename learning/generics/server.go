package main

import (
	"encoding/json"
	"net/http"
)

func healthcheckMap(w http.ResponseWriter, r *http.Request) {
	writeJson(w, map[string]string{
		"status": "ok",
	})
}

func healthcheckTyped(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	writeJson(w, response{"ok"})
}

func writeJson[T any](w http.ResponseWriter, content T) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&content)
}