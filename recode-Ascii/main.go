package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage:  go run . <string>")
		return
	}
	input := os.Args[1]
	_, err := validateInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}
	validBannners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !validBannners[bannerName] {
		fmt.Printf("error: unknown banner %q. choose: standard, shadow, thinkertoy\n", bannerName)
		return
	}
	banner, err := LoadBanner(bannerName + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	result := GenerateArt(input, banner)
	fmt.Print(result)

}
