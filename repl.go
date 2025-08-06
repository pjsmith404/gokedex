package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
		"map": {
			name:        "map",
			description: "Get the next page of maps",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of maps",
			callback:    commandMapBack,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		next:     "",
		previous: "",
	}

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

		err := command.callback(&conf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
