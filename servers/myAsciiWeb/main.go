package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
	Text string
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
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Only post and Get methods allowed", http.StatusMethodNotAllowed)
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
	fmt.Println("Text: ", text)
	fmt.Println("Banner: ", banner)
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
	tmpl.Execute(w, PageData{Result: result, Text: text})
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainHandler)
	mux.HandleFunc("/ascii", AsciiHandler)
	fmt.Println("server is running on http://localhost:9000")
	http.ListenAndServe(":9000", mux)
}
