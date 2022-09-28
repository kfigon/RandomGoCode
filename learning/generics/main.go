package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/map", healthcheckMap)
	http.HandleFunc("/type", healthcheckTyped)

	fmt.Println("running 8080")
	http.ListenAndServe(":8080", nil)
}
