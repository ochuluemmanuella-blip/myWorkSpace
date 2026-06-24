package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <input> [banner]optional")
		os.Exit(1)
	}
	input := os.Args[1]
	//content := strings.ReplaceAll(input, "\r\n", "\n")
	//lines := strings.Split(content, "\n")

	_, err := ValidateInput(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bannerNAME := "standard"
	if len(os.Args) == 3 {
		bannerNAME = os.Args[2]
	}
	validBanners := map[string]bool{
		"shadow":     true,
		"standard":   true,
		"thinkertoy": true,
	}
	if !validBanners[bannerNAME] {
		fmt.Printf("error: unknown banner %q. choose: standard, shadow, thinkertoy\n", bannerNAME)
		return
	}
	bannerfile, err := LoadBanner(bannerNAME + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result := GenerateArt(input, bannerfile)
	fmt.Print(result)
}
