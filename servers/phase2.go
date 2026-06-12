package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}

func mai1n() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", mux)
}