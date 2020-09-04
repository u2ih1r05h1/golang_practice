package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := map[string]int{}
	for _, word := range strings.Fields(s) {
		result[word]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
