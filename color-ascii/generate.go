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
func GenerateArtColor(input string, banner map[rune][]string, substr string, colorCode string) string {
	const reset = "\033[0m"
	lines := SplitInput(input)
	var sb strings.Builder
	start := 0
	if len(lines) > 0 && lines[0] == "" {
		start = 1
	}
	colorAll := substr == ""
	for _, line := range lines[start:] {
		if line == "" {
			sb.WriteString("\n")
			continue
		}
		runes := []rune(line)
		positions := findAllOccurrences(line, substr) // start indexes of substr in this line

		for row := 0; row < 8; row++ {
			for i, r := range runes {
				colored := colorAll || inColoredRange(i, positions, len(substr))
				if colored {
					sb.WriteString(colorCode)
					sb.WriteString(banner[r][row])
					sb.WriteString(reset)
				} else {
					sb.WriteString(banner[r][row])
				}
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
func findAllOccurrences(line, substr string) []int {
	var positions []int
	if substr == "" {
		return positions
	}
	start := 0
	for {
		idx := strings.Index(line[start:], substr)
		if idx == -1 {
			break
		}
		pos := start + idx
		positions = append(positions, pos)
		start = pos + 1
	}
	return positions
}
func inColoredRange(i int, positions []int, length int) bool {
	for _, pos := range positions {
		if i >= pos && i <= pos+length-1 {
			return true
		}
	}
	return false
}
