package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func ciceroHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)

}
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
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
	fmt.Println("Text:", text)
	fmt.Println("Banner:", banner)

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

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, result)

}
func main() {
	http.HandleFunc("/", ciceroHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
