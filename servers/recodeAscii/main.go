package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainHandler)
	mux.HandleFunc("/ascii-art", AsciiHandler)
	mux.HandleFunc("/ascii-art-switch", SwitchHandler)
	fmt.Println("server is running on http://localhost:3030")
	http.ListenAndServe(":3030", mux)
}
