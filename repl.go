package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"github.com/pjsmith404/gokedex/internal/pokeapi"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
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

func getSupportedCommands() map[string]CliCommand {
	return map[string]CliCommand{
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

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return fmt.Errorf("Failed to exit program")
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, command := range getSupportedCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *Config) error {
	var url string
	if config.next == "" {
		url = baseUrl + "location-area/"
	} else {
		url = config.next
	}

	res := pokeapi.Get(url)
	locationArea := LocationArea{}
	err := json.Unmarshal(res, &locationArea)
	if err != nil {
		fmt.Println(err)
	}

	for _,location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	config.next = locationArea.Next
	config.previous = locationArea.Previous

	return nil
}

func commandMapBack(config *Config) error {
	var url string
	if config.previous == "" {
		url = baseUrl + "location-area/"
	} else {
		url = config.previous
	}

	res := pokeapi.Get(url)
	locationArea := LocationArea{}
	err := json.Unmarshal(res, &locationArea)
	if err != nil {
		fmt.Println(err)
	}

	for _,location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	config.next = locationArea.Next
	config.previous = locationArea.Previous

	return nil
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	config := Config{
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

		command.callback(&config)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
