package main

import "strings"

func RenderLine(input string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for _, r := range input {
			if art, ok := banner[r]; ok {
				sb.WriteString(art[row])
			} else if space, ok := banner[' ']; ok {
				sb.WriteString(space[row])
			}
		}
		result[row] = sb.String()
	}
	return result
}
