package main

import "strings"

func WordFrequency(s string) map[string]int {
	freq := make(map[string]int)

	cleaned := strings.ToLower(s)
	cleaned = strings.NewReplacer(",", "", ".", "", "!", "").Replace(cleaned)

	words := strings.Fields(cleaned)

	for _, word := range words {
		freq[word]++
	}
	return freq
}
