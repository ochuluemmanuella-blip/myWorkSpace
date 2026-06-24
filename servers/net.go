package main

import (
	"fmt"
	"io"
	"net/http"
)

func ciceroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<!DOCTYPE html>
		<html>
			<head>
			<title>PRACTICE</title>
			</head>
			<body>
			<h1>ASCII Art Generator</h1>
			<p>Welcome! This tool converts text into ASCII art.</p>
			<br>
			</body>
			</html>
				`,
	)

}

func main() {
	http.HandleFunc("/", ciceroHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
