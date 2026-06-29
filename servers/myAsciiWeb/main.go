package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template not found", 404)
		return
	}
	tmpl.Execute(w, PageData{})
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post methods allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "no input confirmed", http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	if banner == "" {
		http.Error(w, "Banner is empty", http.StatusBadRequest)
		return
	}
	filename := "banners/" + banner + ".txt"

	bannerMap, err := LoadBanner(filename)
	if err != nil {
		http.Error(w, "Banner file not found", http.StatusNotFound)
		return
	}

	_, err = ValidateInput(text)
	if err != nil {
		http.Error(w, "Invalid Input!", http.StatusBadRequest)
		return

	}

	result := GenerateArt(text, bannerMap)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, PageData{Result: result})
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainHandler)
	mux.HandleFunc("/ascii", AsciiHandler)
	fmt.Println("server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
