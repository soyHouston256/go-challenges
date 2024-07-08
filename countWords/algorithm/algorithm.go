package algorithm

import (
	"fmt"
	"strings"
)

func Count(text string) {
	delimiter := " "
	mapWords := make(map[string]int)
	stringLower := strings.ToLower(text)

	arrayWord := strings.Split(stringLower, delimiter)

	for _, word := range arrayWord {
		mapWords[word]++
	}

	for word, count := range mapWords {
		fmt.Println("word: ", word, "\t count: ", count)
	}
}
