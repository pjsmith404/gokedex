package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	splitText := strings.Fields(text)

	var loweredText []string
	for _, word := range splitText {
		loweredText = append(loweredText, strings.ToLower(word))
	}

	return loweredText
}

func main() {
	fmt.Println("Hello, World!")
}
