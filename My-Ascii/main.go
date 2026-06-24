package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("wrong format of arguments! Usage: go run . [STRING]")
		os.Exit(1)
	}
	input := os.Args[1]

	content := loadBanner("standard.txt")

	// fmt.Println(content[:50])
	banner := parseBanner(content)
	// fmt.Println(banner[23])
	// fmt.Println(len(banner))
	bannerMap := buildMap(banner)
	// fmt.Println(bannerMap['H'])
	// fmt.Println(bannerMap['!'])
	validateInput(input, bannerMap)
	lines := splitLines(input)
	//fmt.Println(lines)
	//fmt.Print(input)
	for _, r := range lines {
		if r == "" {
			fmt.Println()
		} else {
			renderLine(r, bannerMap)
		}

	}
	fmt.Printf("%q\n", bannerMap[' '][0]) // row 0 of space
	fmt.Printf("%q\n", bannerMap[' '][1]) // row 1 of space
	fmt.Printf("%q\n", bannerMap[' '][2]) // row 2 of space
}
func loadBanner(filename string) string {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		os.Exit(1)
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return content
}
func parseBanner(content string) [][]string {
	lines := strings.Split(content, "\n\n")
	con := [][]string{}
	for _, r := range lines {
		res := strings.Split(r, "\n")
		con = append(con, res)

	}

	return con
}
func buildMap(banner [][]string) map[rune][]string {
	store := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(' ') + rune(i)
		store[char] = banner[i]
	}
	return store
}
func validateInput(input string, bannerMap map[rune][]string) {
	for i, char := range input {
		if char == '\\' && i+1 < len(input) && input[i+1] == 'n' {
			continue
		}
		if char == 'n' && i > 0 && input[i-1] == '\\' {
			continue
		}
		_, ok := bannerMap[char]
		if !ok {
			fmt.Printf("Invalid character: %c\n", char)
			os.Exit(1)
		}
	}
}
func splitLines(input string) []string {
	return strings.Split(input, "\\n")
}
func renderLine(word string, bannerMap map[rune][]string) {
	for i := 0; i < 8; i++ {
		for _, char := range word {
			fmt.Print(bannerMap[char][i])
		}
		fmt.Println()
	}
}
