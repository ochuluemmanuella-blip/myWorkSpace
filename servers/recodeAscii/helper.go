package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Result string
	Text   string
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

func main()
