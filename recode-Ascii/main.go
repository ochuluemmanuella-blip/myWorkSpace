package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage:  go run . <string>")
		return
	}
	input := os.Args[1]
	_, err := validateInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := GenerateArt(input, banner)
	fmt.Print(result)

}
