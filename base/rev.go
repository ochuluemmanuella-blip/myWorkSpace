package main

import (
	"fmt"
	"strings"
)

func upS(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			if i > 0 {
				s[i-1] = strings.ToUpper(s[i-1])
			}
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}
func mai3() {
	fmt.Println(upS([]string{"This", "is", "the", "goal", "(up)"}))
}
