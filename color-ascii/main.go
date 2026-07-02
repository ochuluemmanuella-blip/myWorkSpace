package main

import (
	"fmt"
	"os"
	"strings"
)

var validBanners = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
}

var ansiColors = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

const usageMsg = `Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring to be colored> "something"`

func main() {
	var colorFlag string
	var banner string = "standard"
	var plainArgs []string

	for _, arg := range os.Args[1:] {
		switch {
		case strings.HasPrefix(arg, "--color="):
			colorFlag = strings.TrimPrefix(arg, "--color=")
			if colorFlag == "" {
				fmt.Println(usageMsg)
				os.Exit(1)
			}
		case strings.HasPrefix(arg, "--"):
			// looks like a flag but isn't a recognized one
			fmt.Println(usageMsg)
			os.Exit(1)
		case validBanners[arg]:
			banner = arg
		default:
			plainArgs = append(plainArgs, arg)
		}
	}

	var substr, input string
	switch len(plainArgs) {
	case 1:
		input = plainArgs[0]
	case 2:
		if strings.Contains(plainArgs[1], plainArgs[0]) {
			substr = plainArgs[0]
			input = plainArgs[1]
		} else if strings.Contains(plainArgs[0], plainArgs[1]) {
			substr = plainArgs[1]
			input = plainArgs[0]
		} else {
			// Neither contains the other — assume standard order: substring first, string second.
			substr = plainArgs[0]
			input = plainArgs[1]
		}
	default:
		fmt.Println(usageMsg)
		os.Exit(1)
	}

	font, err := LoadBanners(banner + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	if colorFlag == "" {
		fmt.Print(GenerateArt(input, font))
		return
	}

	code, ok := ansiColors[colorFlag]
	if !ok {
		fmt.Printf("Error: unknown color %q\n", colorFlag)
		os.Exit(1)
	}

	fmt.Print(GenerateArtColor(input, font, substr, code))
}
