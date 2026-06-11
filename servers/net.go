package main

import (
	"fmt"
	"net/http"
)

func ciceroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", ciceroHandler)
	http.ListenAndServe(":8080", nil)
}
