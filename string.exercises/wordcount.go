package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)
	for _,w := range words {
		wordCount[w]++
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}