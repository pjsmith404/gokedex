package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			fmt.Printf("Your command was: %v\n", input[0])
		}	
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
