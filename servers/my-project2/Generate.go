package main

import "strings"

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
	input = strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(input, "\n")
}
