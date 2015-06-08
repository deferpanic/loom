package main

import (
	"fmt"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func hi2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi2")
}

func main() {
	http.HandleFunc("/hi", hiHandler)
	http.HandleFunc("/hi2", hi2Handler)

	http.ListenAndServe(":3000", nil)
}
