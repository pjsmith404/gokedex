package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	splitText := strings.Fields(loweredText)

	return splitText
}

func getSupportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return fmt.Errorf("Failed to exit program")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, command := range getSupportedCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) <= 0 {
			continue
		}

		supportedCommands := getSupportedCommands()
		command := supportedCommands[input[0]]

		if command.name == "" {
			fmt.Fprintln(os.Stderr, "Unknown command")
			continue
		}

		command.callback()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
