package main

import (
	"errors"
)

func ValidateInput(input string) (rune, error) {
	for _, r := range input {
		if r == '\n' {
			continue
		}
		if r < 32 || r > 126 {
			return r, errors.New("Invalid character!")
		}
	}
	return 0, nil
}

// func SplitInput(input string) []string {
// 	input = strings.ReplaceAll(input, "\\n", "\n")
// 	return strings.Split(input, "\n")
// }
