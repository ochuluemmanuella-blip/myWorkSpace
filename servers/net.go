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
		`<DOCTYPE html>
		<html>
			<head>
			<title>PRACTICE</title>
			</head>
			<body>
			HI, SAMUEL!
			<br>
			</body>
			</html>
				`,
	)
	fmt.Fprintln(w, "Hello, world!, it's so annoying")

}

func main() {
	http.HandleFunc("/", ciceroHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server running on http://localhost:8080")
}
