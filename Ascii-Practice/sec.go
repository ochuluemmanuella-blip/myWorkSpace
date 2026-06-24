package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	word := os.Args[1]
	//fmt.Println(word)
	content := `flf2a$ 5 4 8 -1 0
	$$__$@
	$/$$\@
	/----\@
	|$$$$|@
	|$$$$|@@
	|---\$@
	|$$/$@
	|---\$@
	|$$$\@
	|---/@@
	$___$@
	/$$$$@
	|$$$$@
	\$$$$@
	$___/@@`

	lines := strings.Split(content, "\n")
	font := make(map[rune][]string)

	currentChar := 'A'        // start from A
	currentRows := []string{} // collect rows for current letter

	for i, line := range lines {
		if i == 0 {
			continue // skip the header line for now
		}

		isLastRow := strings.HasSuffix(line, "@@")
		cleaned := strings.TrimRight(line, "@")
		cleaned = strings.ReplaceAll(cleaned, "$", " ")

		currentRows = append(currentRows, cleaned)

		if isLastRow {
			font[currentChar] = currentRows // save the letter
			currentChar++                   // move to next letter (A→B→C)
			currentRows = []string{}        // reset for next letter
		}
	}

	// // render a word
	// word := "ABC"
	fmt.Print("\033[31m")
	for row := 0; row < 5; row++ {
		for _, ch := range word {
			fmt.Print(font[ch][row] + "  ")
		}
		fmt.Println()
	}

	//switch to red
	fmt.Println(word)
	fmt.Print("\033[0m") // reset back to normal
}
