package main

import (
	"fmt"
)

func mai3() {
	font := map[rune][]string{
		'A': {
			" __ ",
			"/  \\",
			"|__|",
			"|  |",
			"|  |",
		},
		'B': {
			"|__ ",
			"|  \\",
			"|__/",
			"|  \\",
			"|__/",
		},
		'C': {
			"___",
			"/   ",
			"|   ",
			"\\   ",
			" \\___",
		},
	}
	word := "ABC"

	// outer loop: row 0 to 4
	for row := 0; row < 5; row++ {
		// inner loop: each character in the word
		for _, ch := range word {
			fmt.Print(font[ch][row] + "  ")
		}
		fmt.Println()
	}

}

// func main() {
// for i := 0; i < 5; i++ {
// for j := 0; j < (4 - i); j++ {
// fmt.Print(" ")
// }
// for k := 0; k < (2i + 1); k++ {
// fmt.Print("")
// }
// fmt.Println()
// }
// for i := 3; i >= 0; i-- {
// for j := 0; j < (4 - i); j++ {
// fmt.Print(" ")
// }
// for k := 0; k < (2i + 1); k++ {
// fmt.Print("")
// }
// fmt.Println()
// }
// }
// func main() {
// for i := 0; i < 5; i++ {
// for j := 0; j < 5; j++ {
// if i == 0 || i == 4 || j == 0 || j == 4 {
// fmt.Print("*")
// } else {
// fmt.Print(" ")
// }
// }
// fmt.Println()
// }
// }

// func main() {
// for i := 0; i < 5; i++ {
// for j := 0; j < (4 - i); j++ {
// fmt.Print(" ")
// }
// for k := 0; k < (2i + 1); k++ {
// fmt.Print("")
// }
// fmt.Println()
// }
// }

// func main() {
// for i := 0; i < 5; i++ {
// for j := i + 1; j <= 5; j++ {
// fmt.Print("*")
// }
// fmt.Println()
// }
// }

// func main() {
// for i := 0; i < 5; i++ {
// for j := 0; j <= i; j++ {
// fmt.Print("*")
// }
// fmt.Println()
// }
// }

// func main() {
// for i := 0; i < 5; i++ {
// for j := 0; j < 5; j++ {
// fmt.Print("*")
// }
// fmt.Println()
// }
// }

// func main() {
// for i := 0; i < 10; i++ {
// fmt.Print("*")
// }
// fmt.Println()
// }
