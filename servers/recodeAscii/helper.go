package main

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Result string
	Text   string
	Banner string
}

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.ReplaceAll(string(data), "\r\n", "\n")
	content := strings.Split(lines, "\n")

	banner := make(map[rune][]string)

	start := 1

	for i := 32; i <= 126; i++ {
		art := content[start : start+8]
		banner[rune(i)] = art
		start += 9
	}
	return banner, nil
}
func GenerateArt(input string, banner map[rune][]string) string {
	lines := SplitInput(input)
	var sb strings.Builder

	start := 0
	if len(lines) > 0 && lines[0] == "" {
		start = 1
	}
	for _, line := range lines[start:] {
		if line == "" {
			sb.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, r := range line {
				sb.WriteString(banner[r][row])
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
func ValidateInput(input string) (rune, error) {
	for _, r := range input {
		if r == '\n' || r == '\r' {
			continue
		}
		if r < 32 || r > 126 {
			return r, errors.New("Invalid character")
		}
	}
	return 0, nil
}
func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, `\n`, "\n")
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")
	return strings.Split(input, "\n")
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
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
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post methods allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "no input confirmed", http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	if banner == "" {
		http.Error(w, "no banner choice confirmed", http.StatusBadRequest)
		return
	}

	filename := "banners/" + banner + ".txt"

	Map, err := LoadBanner(filename)
	if err != nil {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	_, err = ValidateInput(text)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	result := GenerateArt(text, Map)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, PageData{Result: result, Text: text})
}
func SwitchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-switch" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "only Get methods allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "no input confirmed", http.StatusBadRequest)
		return
	}
	banner := r.URL.Query().Get("banner")
	if banner == "" {
		http.Error(w, "no banner choice confirmed", http.StatusBadRequest)
		return
	}

	filename := "banners/" + banner + ".txt"

	Map, err := LoadBanner(filename)
	if err != nil {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	_, err = ValidateInput(text)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	result := GenerateArt(text, Map)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, PageData{Result: result, Text: text, Banner: banner})

}
