package main

import (
	"fmt"
)

func commandCatch(conf *config, subCommand string) error {
	if subCommand == "" {
		return fmt.Errorf("Pick a pokemon to catch! eg: catch pikachu")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", subCommand)

	return nil
}
