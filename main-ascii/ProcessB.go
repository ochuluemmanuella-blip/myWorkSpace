package main

import (
	"os"
	"strings"
)

func LoadBanners(fileN string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileN)
	if err != nil {
		return nil, err
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	line := strings.Split(content, "\n")
	font := make(map[rune][]string)
	for i := 0; i < len(line)/9; i++ {
		ch := rune(32 + i)
		start := i * 9
		font[ch] = line[start : start+8]
	}
	return font, nil
}
func PrintAsc(input string, font map[rune][]string) string {
	var sb strings.Builder
	lines := strings.Split(input, "\\n")
	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				sb.WriteString("\n")
			}
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if clean, ok := font[ch]; ok {
					sb.WriteString(clean[row])
				}
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
