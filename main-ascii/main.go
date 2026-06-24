package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [string] [banner]")
		os.Exit(1)

	}
	input := os.Args[1]
	banner := "standard"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !validBanners[banner] {
		fmt.Printf("Error: unknown banner %q. Choose: standard, shadow, thinkertoy\n", banner)
		os.Exit(1)
	}
	font, err := LoadBanners(banner + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(PrintAsc(input, font))

}
