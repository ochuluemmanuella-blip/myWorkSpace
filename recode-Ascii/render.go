package main

import (
	"fmt"
	"strings"
)

func RenderLine(input string, m map[rune][]string) {
	for row := 0; row < 8; row++ {
		for _, r := range input {
			fmt.Println(m[r][row])
		}
		fmt.Println()
	}

}
func GenerateArt(input string, m map[rune][]string) string {
	lines := SplitInput(input)
	var sb strings.Builder
	for _, line := range lines {
		if line == "" {
			sb.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, r := range line {
				sb.WriteString(m[r][row])
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
