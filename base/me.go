package main

import (
	"fmt"
	"strings"
)

func lastN(text []string) []string {
	for i := 0; i < len(text); i++ {
		if text[i] == "(low)" {
			text[i-1] = strings.ToUpper(text[i-1])
		}
		text = append(text[:i], text[i+1:]...)
	}
	return text
}
func main() {
	fmt.Println(lastN([]string{"go", "is", "fun", "(up)"}))
}
