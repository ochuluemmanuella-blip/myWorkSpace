package main

import (
	"fmt"
	"strings"
)

func main() {

	line := "/$$\\@"
	cleaned := strings.TrimRight(line, "@")
	cleaned = strings.ReplaceAll(cleaned, "$", " ")
	fmt.Println(cleaned) // what do you think this prints?
}
