package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.Replace("hello world", "world", "Go", 1))
	fmt.Println(strings.Split("hello world", " "))
}
