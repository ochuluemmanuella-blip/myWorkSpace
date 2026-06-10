package main

import (
	"os"
	"strings"
)

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
