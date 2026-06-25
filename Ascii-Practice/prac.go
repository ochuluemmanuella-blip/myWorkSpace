package main

import (
	"fmt"
	"strings"
)

func ma1n() {
	content := `A
 __ 
/  \
|__|
|  |
|  |
B
___ 
|  \
|__/
|  \
|__/
C
 ___
/   
|   
\   
 \___`

	lines := strings.Split(content, "\n")
	font := make(map[rune][]string)
	var currentChar rune
	for _, line := range lines {

		if len(line) == 1 {

			//var rowCount int

			// when you see a 1-char line:
			currentChar = rune(line[0])
			font[currentChar] = []string{}
		} else {

			// when you see an art line:
			font[currentChar] = append(font[currentChar], line)
			//rowCount++
		}

	}
	// TODO: loop over lines and build the font map
	// Hint: if len(line) == 1 → it's a letter key
	//       otherwise         → it's an art row

	// Once font is built, render a word:
	word := "ABC"
	for row := 0; row < 5; row++ {
		for _, ch := range word {
			fmt.Print(font[ch][row] + " ")
		}
		fmt.Println()
	}
}
