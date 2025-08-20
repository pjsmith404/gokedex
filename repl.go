package main

import (
	"bufio"
	"fmt"
	"github.com/pjsmith404/gokedex/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
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
		"explore": {
			name: "explore",
			description: "Explore a given area",
			callback: commandExplore,
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
		pokeapiClient: pokeapi.NewClient(),
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

		subCommand := ""

		if len(input) > 1 {
			subCommand = input[1]
		}

		if command.name == "" {
			fmt.Fprintln(os.Stderr, "Unknown command")
			continue
		}

		err := command.callback(&conf, subCommand)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
