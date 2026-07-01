package main

import (
	"errors"
)

func ValidateInput(input string) (rune, error) {
	for _, r := range input {
		if r == '\n' || r == '\r' {
			continue
		}
		if r < 32 || r > 126 {
			return r, errors.New("Invalid character!")
		}
	}
	return 0, nil
}
